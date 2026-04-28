package inner

type Base struct {
	ID int
}

// Pinger is an interface embedded into the parent package's Service
// interface. Method-set resolution must surface "Pinger" (or its qname) in
// Service's Embeds slice, while Service.Methods only lists the directly
// declared methods.
type Pinger interface {
	Ping() error
}
