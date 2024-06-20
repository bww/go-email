package email

import (
	"net/mail"
	"net/url"
)

type Config struct {
	DefaultSender     Address
	OverrideRecipient Address
}

func (c Config) WithOptions(opts []Option) Config {
	for _, opt := range opts {
		c = opt(c)
	}
	return c
}

func (c Config) WithParams(query url.Values) (Config, error) {
	if s := query.Get("sender"); s != "" {
		a, err := mail.ParseAddress(s)
		if err != nil {
			return c, err
		}
		c.DefaultSender = Address{
			Name:  a.Name,
			Email: a.Address,
		}
	}
	return c, nil
}

type Option func(Config) Config

func DefaultSender(s Address) Option {
	return func(c Config) Config {
		c.DefaultSender = s
		return c
	}
}

func OverrideRecipient(s Address) Option {
	return func(c Config) Config {
		c.OverrideRecipient = s
		return c
	}
}
