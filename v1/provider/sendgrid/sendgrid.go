package sendgrid

import (
	"net/url"

	"github.com/bww/go-email/v1"
)

const Scheme = "sendgrid"

type Provider struct{}

func New(dsn *url.URL, conf email.Config) (*Provider, error) {
	return &Provider{}, nil
}

func (p *Provider) Send(tmplName string, msg *email.Template) error {
	return email.ErrUnimplemented
}
