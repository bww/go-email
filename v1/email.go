package email

import (
	"os"
	"strconv"
	"strings"
)

var verbose = os.Getenv("EMAIL_VERBOSE") != ""

type Address struct {
	Email string
	Name  string
}

func (a Address) IsZero() bool {
	return a.Email == ""
}

type Variables map[string]string

type Personalization struct {
	Recipients []Address
	Variables  Variables
	Subject    string
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
	Attachments      []*Attachment
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
