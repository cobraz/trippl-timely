package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

type TimelyConfig struct {
	ClientID     string `survey:"client_id"`
	ClientSecret string `survey:"client_secret"`
	Token        *oauth2.Token
}

type TripletexConfig struct {
	ConsumerToken string `survey:"consumer_token"`
	EmployeeToken string `survey:"employee_token"`
	ActivityCode  int32
}

type Config struct {
	Timely    TimelyConfig
	Tripletex TripletexConfig
}

func Init() (*viper.Viper, error) {
	userDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := fmt.Sprintf("%s/.trippl-timely-auth", userDir)
	_ = os.Mkdir(configPath, os.ModePerm)

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(configPath)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file does not exists. Creating!")
		} else {
			return nil, err
		}
	}
	return v, nil
}

func SetConfig(config Config) error {
	v, err := Init()
	if err != nil {
		return err
	}

	userDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configPath := fmt.Sprintf("%s/.trippl-timely-auth/config.json", userDir)

	v.Set("timely", config.Timely)
	v.Set("tripletex", config.Tripletex)

	err = v.WriteConfigAs(configPath)
	if err != nil {
		return err
	}
	return nil
}

func values() (Config, error) {
	v, err := Init()
	if err != nil {
		return Config{}, err
	}
	var c Config
	err = v.Unmarshal(&c)

	// TODO: Fix ugly hotfix for unmarshalling time.Time
	c.Timely.Token = &oauth2.Token{
		AccessToken:  v.GetString("timely.Token.access_token"),
		TokenType:    v.GetString("timely.Token.token_type"),
		RefreshToken: v.GetString("timely.Token.refresh_token"),
		Expiry:       v.GetTime("timely.Token.expiry"),
	}

	return c, nil
}

func GetConfig() (Config, error) {
	return values()
}

func Print() {
	v, err := Init()
	if err != nil {
		log.Fatal("Something happened when parsing config")
	}

	c := v.AllSettings()
	bs, err := yaml.Marshal(c)

	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}

	fmt.Println(string(bs))
}
