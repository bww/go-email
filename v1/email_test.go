package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpolatePersonalization(t *testing.T) {
	tests := []struct {
		In     Personalization
		Vars   interface{}
		Expect Personalization
		Err    func(error) error
	}{
		{
			In: Personalization{
				Recipients: []Address{},
				Variables:  Variables{},
				Subject:    "Oh, hello, {{ .Name }}",
			},
			Vars: map[string]string{
				"Name": "Joseph Steel",
			},
			Expect: Personalization{
				Recipients: []Address{},
				Variables:  Variables{},
				Subject:    "Oh, hello, Joseph Steel",
			},
		},
	}
	for _, e := range tests {
		res, err := e.In.Interpolate(e.Vars)
		if e.Err != nil {
			assert.NoError(t, e.Err(err))
		} else if assert.NoError(t, err) {
			assert.Equal(t, e.Expect, res)
		}
	}
}
