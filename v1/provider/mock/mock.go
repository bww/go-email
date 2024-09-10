package mock

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"strings"

	"github.com/bww/go-email/v1"
	"github.com/bww/go-util/v1/slices"
	"github.com/bww/go-util/v1/text"
)

const Scheme = "mock"

type Provider struct {
	email.Config
	log *slog.Logger
}

func New(dsn *url.URL, conf email.Config) (*Provider, error) {
	return &Provider{
		Config: conf,
		log:    slog.Default().With("provider", "mock"),
	}, nil
}

func (p *Provider) Send(cxt context.Context, tmplName string, msg email.Template) error {
	msg = msg.With(p.Config)
	log := p.log.With(
		"template", tmplName,
		"sender", msg.From,
		"recipients", summaryOf(slices.Flatten(slices.Map(msg.Personalizations, func(p email.Personalization) []email.Address { return p.Recipients })), 3),
	)
	log.Info("Send email")
	if p.Verbose {
		data, err := json.MarshalIndent(msg, "", "  ")
		if err != nil {
			log.Error("Could not marshal email template", "error", err)
		} else {
			fmt.Println(text.Indent(string(data), " > "))
		}
	}
	return nil
}

func (p *Provider) String() string {
	if p.DefaultSender.IsZero() {
		return "mock sender"
	} else {
		return fmt.Sprintf("mock sender: %s", p.DefaultSender)
	}
}

func summaryOf[E fmt.Stringer](e []E, n int) string {
	sb := &strings.Builder{}
	l := len(e)
	for i := 0; i < min(l, n); i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(e[i].String())
	}
	if l > n {
		sb.WriteString(fmt.Sprintf(", ...and %d more", l-n))
	}
	return sb.String()
}
