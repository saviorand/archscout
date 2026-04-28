package typestructfixture

import (
	"io"

	"example.com/typestructfixture/inner"
)

type User struct {
	inner.Base

	Name      string `json:"name"`
	Email     string `json:"email,omitempty"`
	Age, Year int
}

type Service interface {
	io.Reader
	inner.Pinger

	Run() error
	Stop(force bool) error
}
