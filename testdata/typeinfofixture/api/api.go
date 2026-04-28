package api

// Greeter is a single-method interface used to verify that interface
// method calls resolve to the interface's defining package.
type Greeter interface {
	Greet(name string) string
}

// Service exercises both pointer and value receiver methods.
type Service struct{}

// Run is called via a pointer receiver.
func (s *Service) Run() string {
	return "running"
}

// Describe is called via a value receiver.
func (s Service) Describe() string {
	return "describing"
}
