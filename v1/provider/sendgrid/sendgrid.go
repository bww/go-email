package sendgrid

import (
	"log/slog"
	"net/url"

	api "github.com/bww/go-apiclient/v1"
	"github.com/bww/go-email/v1"
)

const Scheme = "sendgrid"

type Provider struct {
	email.Config
	client *api.Client
	log    *slog.Logger
}

func New(dsn *url.URL, conf email.Config) (*Provider, error) {
	return &Provider{
		Config: conf,
		client: api.New(api.WithAuthorizer(
			api.NewBearerAuthorizer(dsn.Host)),
			api.WithBaseURL("https://api.sendgrid.com/v3"),
			api.WithHeader("Content-Type", "application/json"),
		),
		log: slog.Default().With("provider", "sendgrid"),
	}, nil
}

func (p *Provider) Send(tmplName string, msg email.Template) error {
	msg = msg.With(p.Config)
	return email.ErrUnimplemented
}

func (p *Provider) String() string {
	return "Sendgrid"
}
