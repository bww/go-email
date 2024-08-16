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
				Recipients: []Address{
					{
						Name:  "{{ .Name }}",
						Email: "{{ .Email }}",
					},
				},
				Variables: Variables{
					"favorite_color": "{{ .FavoriteColor }}",
				},
				Subject: "Oh, hello, {{ .Name }}; hope you still like {{ .FavoriteColor }}!",
			},
			Vars: map[string]string{
				"Name":          "Joseph Steel",
				"Email":         "joe@ussr.ru",
				"FavoriteColor": "red",
			},
			Expect: Personalization{
				Recipients: []Address{
					{
						Name:  "Joseph Steel",
						Email: "joe@ussr.ru",
					},
				},
				Variables: Variables{
					"favorite_color": "red",
				},
				Subject: "Oh, hello, Joseph Steel; hope you still like red!",
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
