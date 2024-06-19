package email

type Config struct {
	DefaultSender     string
	OverrideRecipient string
}

func (c Config) With(opts []Option) Config {
	for _, opt := range opts {
		c = opt(c)
	}
	return c
}

type Option func(Config) Config

func WithDefaultSender(s string) Option {
	return func(c Config) Config {
		c.DefaultSender = s
		return c
	}
}

func WithOverrideRecipient(s string) Option {
	return func(c Config) Config {
		c.OverrideRecipient = s
		return c
	}
}
