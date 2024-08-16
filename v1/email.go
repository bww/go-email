package email

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bww/go-util/v1/text/template"
)

var ErrUnimplemented = errors.New("Not implemented")

var verbose = os.Getenv("EMAIL_VERBOSE") != ""

type Address struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

func (a Address) IsZero() bool {
	return a.Email == ""
}

func (a Address) String() string {
	if a.Name != "" {
		return fmt.Sprintf("%s <%s>", a.Name, a.Email)
	} else {
		return a.Email
	}
}

type Variables map[string]string

type Personalization struct {
	Recipients []Address `json:"recipients"`
	Variables  Variables `json:"variables"`
	Subject    string    `json:"subject"`
}

// Interpolate considers each field in a Personalization as a template string
// which is evaluated with the provided data as context. The result is returned
// as a Personalization.
func (p Personalization) Interpolate(data interface{}) (Personalization, error) {
	s, err := interpolate(p.Subject, data)
	if err != nil {
		return Personalization{}, err
	}

	vars := make(map[string]string)
	for k, v := range p.Variables {
		vars[k], err = interpolate(v, data)
		if err != nil {
			return Personalization{}, err
		}
	}

	rcps := make([]Address, len(p.Recipients))
	for i, e := range p.Recipients {
		n, err := interpolate(e.Name, data)
		if err != nil {
			return Personalization{}, err
		}
		a, err := interpolate(e.Email, data)
		if err != nil {
			return Personalization{}, err
		}
		rcps[i] = Address{
			Name:  n,
			Email: a,
		}
	}

	return Personalization{
		Subject:    s,
		Variables:  vars,
		Recipients: rcps,
	}, nil
}

func (p Personalization) With(conf Config) Personalization {
	if !conf.OverrideRecipient.IsZero() {
		p.Recipients = []Address{conf.OverrideRecipient}
	}
	return p
}

type Attachment struct {
	Type        string `json:"type"`
	Filename    string `json:"filename"`
	Disposition string `json:"disposition,omitempty"`
	ContentId   string `json:"content_id,omitempty"`
	Content     []byte `json:"content"`
}

type Template struct {
	From             Address           `json:"from"`
	ReplyTo          Address           `json:"reply_to"`
	Personalizations []Personalization `json:"personalizations"`
	Attachments      []Attachment      `json:"attachments,omitempty"`
}

func (t Template) With(conf Config) Template {
	if t.From.IsZero() {
		t.From = conf.DefaultSender
	}
	if !conf.OverrideRecipient.IsZero() {
		for i, e := range t.Personalizations {
			t.Personalizations[i] = e.With(conf)
		}
	}
	return t
}

type Fields map[string]interface{}

type Contact struct {
	Id        string   `json:"id"`
	Email     string   `json:"email"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	ListsIds  []string `json:"list_ids,omitempty"`
	Fields    Fields   `json:"fields,omitempty"`
}

type Error struct {
	Message string `json:"message"`
	Indices []int  `json:"error_indices,omitempty"`
}

func (e Error) Error() string {
	var s strings.Builder
	s.WriteString(e.Message)
	if verbose && len(e.Indices) > 0 {
		s.WriteString(" (input indices: ")
		for i, e := range e.Indices {
			if i > 0 {
				s.WriteString(", ")
			}
			s.WriteString(strconv.Itoa(e))
		}
		s.WriteString(")")
	}
	return s.String()
}

func interpolate(f string, v interface{}) (string, error) {
	t, err := template.Parse(f)
	if err != nil {
		return "", err
	}
	r, err := t.Exec(v)
	if err != nil {
		return "", err
	}
	return string(r), err
}
