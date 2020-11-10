package config

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cobraz/trippl-timely/internal/pkg/config"
	tx "github.com/cobraz/trippl-timely/internal/pkg/tripletex"
	"github.com/urfave/cli/v2"
)

func GetConfig(c *cli.Context) error {
	config.Print()
	return nil
}

var timelyQs = []*survey.Question{
	{
		Name:     "client_id",
		Prompt:   &survey.Input{Message: "Client ID?"},
		Validate: survey.Required,
	},
	{
		Name:     "client_secret",
		Prompt:   &survey.Input{Message: "Client Secret?"},
		Validate: survey.Required,
	},
}

var tripletexQs = []*survey.Question{
	{
		Name:     "consumer_token",
		Prompt:   &survey.Input{Message: "Consumer Token?"},
		Validate: survey.Required,
	},
	{
		Name:     "employee_token",
		Prompt:   &survey.Input{Message: "Employee Token?"},
		Validate: survey.Required,
	},
}

func askAboutActivityCode(tx *tx.TripletexClient) (*int32, error) {
	activities, err := tx.GetActivities()

	if err != nil {
		return nil, err
	}

	var aOpts []string
	for _, a := range activities {
		aOpts = append(aOpts, a.Name)
	}

	prompt := &survey.Select{
		Message: "Choose an activity to track on:",
		Options: aOpts,
	}

	var selectedOpt int

	survey.AskOne(prompt, &selectedOpt, survey.WithValidator(survey.Required))

	var res *int32 = &activities[selectedOpt].ID

	return res, nil
}

func askAboutEmployeeId(tx *tx.TripletexClient) (*int32, error) {
	employees, err := tx.GetEmployees()

	if err != nil {
		return nil, err
	}

	var aOpts []string
	for _, a := range employees {
		aOpts = append(aOpts, fmt.Sprintf("%s %s", *a.FirstName, *a.LastName))
	}

	prompt := &survey.Select{
		Message: "Choose which employee you want to copy to:",
		Options: aOpts,
	}

	var selectedOpt int

	survey.AskOne(prompt, &selectedOpt, survey.WithValidator(survey.Required))

	var res *int32 = &employees[selectedOpt].ID

	return res, nil
}

func SetConfig(c *cli.Context) error {
	var timley config.TimelyConfig

	fmt.Println("Let's start with Timely configuration.")
	err := survey.Ask(timelyQs, &timley)
	if err != nil {
		return err
	}

	// ActivityCode

	var tripletex config.TripletexConfig

	fmt.Println("Now, let's add Tripletex configuration.")

	err = survey.Ask(tripletexQs, &tripletex)
	if err != nil {
		return err
	}

	client, err := tx.New(config.Config{
		Tripletex: tripletex,
	})

	if err != nil {
		return err
	}

	actCode, err := askAboutActivityCode(client)

	if err != nil {
		return err
	}

	tripletex.ActivityCode = *actCode

	empID, err := askAboutEmployeeId(client)

	if err != nil {
		return err
	}

	tripletex.EmployeeId = fmt.Sprint(*empID)

	cnf := config.Config{
		Timely:    timley,
		Tripletex: tripletex,
	}

	err = config.SetConfig(cnf)
	if err != nil {
		return err
	}

	fmt.Println("\n\nGratz! You are all set. This is your config:")
	config.Print()
	return nil
}
