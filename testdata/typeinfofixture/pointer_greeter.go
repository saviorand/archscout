package main

// PointerGreeter implements api.Greeter via a pointer receiver. This makes
// it a member of the pointer-method set only, exercising the alternate
// branch of typeImplements.
type PointerGreeter struct {
	prefix string
}

// Greet satisfies api.Greeter for *PointerGreeter.
func (p *PointerGreeter) Greet(name string) string {
	return p.prefix + " " + name
}
