package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/bww/go-email/v1"
	"github.com/bww/go-email/v1/provider/mock"
	"github.com/bww/go-email/v1/provider/sendgrid"
)

var ErrUnsupported = fmt.Errorf("Provider is not supported")

type Provider interface {
	fmt.Stringer
	Send(cxt context.Context, tmplName string, msg email.Template) error
}

func New(dsn string, opts ...email.Option) (Provider, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("Malformed spec: %w", err)
	}
	conf := email.Config{}.WithOptions(opts)
	conf, err = conf.WithParams(u.Query())
	if err != nil {
		return nil, fmt.Errorf("Invalid paramters: %w", err)
	}
	switch u.Scheme {
	case sendgrid.Scheme:
		return sendgrid.New(u, conf)
	case mock.Scheme:
		return mock.New(u, conf)
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupported, u.Scheme)
	}
}
