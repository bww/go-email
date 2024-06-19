package provider

import (
	"fmt"
	"net/url"
)

var (
	ErrUnsupported = fmt.Errorf("Provider is not supported")
)

type Provider interface {
	Send() error
}

func New(dsn string) (Provider, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return fmt.Errorf("Malformed spec: %w", err)
	}
	switch u.Scheme {
	case sendgrid.Scheme:
		return sendgrid.New(u)
	case mock.Scheme:
		return mock.New(u)
	default:
		return fmt.Errorf("%w: %s", ErrUnsupported, u.Scheme)
	}
}
