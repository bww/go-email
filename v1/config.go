package email

type Config struct {
	DefaultSender     Address
	OverrideRecipient Address
}

func (c Config) With(opts []Option) Config {
	for _, opt := range opts {
		c = opt(c)
	}
	return c
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
