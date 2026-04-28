package callerfixture

import (
	"errors"
	"strings"
)

// ErrSentinel is initialized at package level — calls in this initializer
// have no enclosing FuncDecl.
var ErrSentinel = errors.New("sentinel")

// Service exercises pointer-receiver method dispatch.
type Service struct {
	name string
}

// Validate calls errors.New from the method body and strings.ToUpper from a
// closure inside the same method. Both call sites should report Validate as
// their enclosing caller.
func (s *Service) Validate(input string) error {
	if input == "" {
		return errors.New("empty input")
	}

	normalize := func(in string) string {
		return strings.ToUpper(in)
	}
	if normalize(input) == "BAD" {
		return errors.New("bad input")
	}
	return nil
}
