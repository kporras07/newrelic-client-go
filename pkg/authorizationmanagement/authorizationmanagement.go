package authorizationmanagement

// Package authorizationmanagement provides a programmatic API for interacting with the New Relic authorization management
import (
	"github.com/newrelic/newrelic-client-go/v2/internal/http"
	"github.com/newrelic/newrelic-client-go/v2/pkg/config"
	"github.com/newrelic/newrelic-client-go/v2/pkg/logging"
)

type Authorizationmanagement struct {
	client http.Client
	logger logging.Logger
}

// New is used to create a new Account Management.
func New(config config.Config) Authorizationmanagement {
	client := http.NewClient(config)

	pkg := Authorizationmanagement{
		client: client,
		logger: config.GetLogger(),
	}

	return pkg
}
