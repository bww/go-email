package sendgrid

import (
	"net/url"
)

const Scheme = "sendgrid"

type Provider struct{}

func New(dsn *url.URL) (Provider, error) {
	return &Provider{}, nil
}

func (p *Provider) Send() error {
	return nil
}
