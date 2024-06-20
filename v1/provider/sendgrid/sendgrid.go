package sendgrid

import (
	"context"
	"fmt"
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
	client, err := api.New(api.WithAuthorizer(
		api.NewBearerAuthorizer(dsn.Host)),
		api.WithBaseURL("https://api.sendgrid.com/v3/"),
		api.WithHeader("Content-Type", "application/json"),
	)
	if err != nil {
		return nil, fmt.Errorf("Could not create API client: %w", err)
	}
	return &Provider{
		Config: conf,
		client: client,
		log:    slog.Default().With("provider", "sendgrid"),
	}, nil
}

func (p *Provider) Send(cxt context.Context, tmplName string, msg email.Template) error {
	msg = msg.With(p.Config)

	var psn []personalization
	for _, e := range msg.Personalizations {
		psn = append(psn, personalization{
			Recipients: newAddresses(e.Recipients),
			Variables:  e.Variables,
			Subject:    e.Subject,
		})
	}

	var att []attachment
	for _, e := range msg.Attachments {
		att = append(att, attachment{
			Type:        e.Type,
			Filename:    e.Filename,
			Disposition: e.Disposition,
			Content:     e.Content,
			ContentId:   e.ContentId,
		})
	}

	var rply *address
	if !msg.ReplyTo.IsZero() {
		v := newAddress(msg.ReplyTo)
		rply = &v
	}

	tmpl := template{
		TemplateId:       tmplName,
		From:             newAddress(msg.From),
		ReplyTo:          rply,
		Personalizations: psn,
		Attachments:      att,
	}

	_, err := p.client.Post(cxt, "mail/send", &tmpl, nil)
	if err != nil {
		return fmt.Errorf("Could not send email: %w", err)
	}

	return nil
}

func (p *Provider) String() string {
	return "Sendgrid"
}
