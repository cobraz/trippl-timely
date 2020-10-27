package tripletex

import (
	apiclient "github.com/bjerkio/tripletex-go/client"
	"github.com/bjerkio/tripletex-go/helpers"
	"github.com/cobraz/trippl-timely/internal/pkg/config"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type TripletexClient struct {
	client   *apiclient.Tripletex
	authInfo runtime.ClientAuthInfoWriter
	config   config.Config
}

// New returns a authenticated Tripletex client
func New(config config.Config) (*TripletexClient, error) {

	token, err := helpers.CreateToken(config.Tripletex.ConsumerToken, config.Tripletex.EmployeeToken)
	if err != nil {
		return nil, err
	}

	r := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	r.DefaultAuthentication = httptransport.BasicAuth("0", token)

	// Fix "application/json; charset=utf-8" issue
	r.Producers["application/json; charset=utf-8"] = runtime.JSONProducer()

	return &TripletexClient{
		client:   apiclient.New(r, strfmt.Default),
		authInfo: r.DefaultAuthentication,
		config:   config,
	}, nil
}
