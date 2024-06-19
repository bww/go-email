package mock

import (
	"net/url"
)

const Scheme = "mock"

type Provider struct{}

func New(dsn *url.URL) (Provider, error) {
	return &Provider{}, nil
}

func (p *Provider) Send() error {
	return nil
}
