package typestructfixture

import (
	"io"

	"example.com/typestructfixture/inner"
)

// User exercises named, multi-name and embedded fields, plus struct tags.
type User struct {
	inner.Base // embedded cross-package struct

	Name      string `json:"name"`
	Email     string `json:"email,omitempty"`
	Age, Year int
}

// Service exercises directly declared interface methods plus embedded
// interfaces from both stdlib and a workspace package.
type Service interface {
	io.Reader
	inner.Pinger

	Run() error
	Stop(force bool) error
}
