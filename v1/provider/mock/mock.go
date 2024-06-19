package mock

import (
	"fmt"
	"log/slog"
	"net/url"
	"strings"

	"github.com/bww/go-email/v1"
	"github.com/bww/go-util/v1/slices"
)

const Scheme = "mock"

type Provider struct {
	log *slog.Logger
}

func New(dsn *url.URL, conf email.Config) (*Provider, error) {
	return &Provider{
		log: slog.Default().With("email", "mock"),
	}, nil
}

func (p *Provider) Send(tmplName string, msg *email.Template) error {
	p.log.With(
		"template", tmplName,
		"sender", msg.From,
		"recipients", summaryOf(slices.Flatten(slices.Map(msg.Personalizations, func(p email.Personalization) []email.Address { return p.Recipients })), 3),
	).Info("Send email")
	return nil
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
