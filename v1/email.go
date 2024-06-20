package email

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ErrUnimplemented = errors.New("Not implemented")

var verbose = os.Getenv("EMAIL_VERBOSE") != ""

type Address struct {
	Email string
	Name  string
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
	Recipients []Address
	Variables  Variables
	Subject    string
}

func (p Personalization) With(conf Config) Personalization {
	if !conf.OverrideRecipient.IsZero() {
		p.Recipients = []Address{conf.OverrideRecipient}
	}
	return p
}

type Attachment struct {
	Type        string
	Filename    string
	Disposition string
	ContentId   string
	Content     []byte
}

type Template struct {
	From             Address
	ReplyTo          Address
	Personalizations []Personalization
	Attachments      []Attachment
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
	Id        string
	Email     string
	FirstName string
	LastName  string
	ListsIds  []string
	Fields    Fields
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
