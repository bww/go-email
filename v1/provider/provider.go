package provider

import (
	"fmt"
	"net/url"

	"github.com/bww/go-email/v1"
	"github.com/bww/go-email/v1/provider/mock"
	"github.com/bww/go-email/v1/provider/sendgrid"
)

var ErrUnsupported = fmt.Errorf("Provider is not supported")

type Provider interface {
	Send(tmplName string, msg *email.Template) error
}

func New(dsn string) (Provider, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("Malformed spec: %w", err)
	}
	switch u.Scheme {
	case sendgrid.Scheme:
		return sendgrid.New(u)
	case mock.Scheme:
		return mock.New(u)
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupported, u.Scheme)
	}
}
