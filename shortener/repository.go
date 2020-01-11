package shortener

// RedirectRepository implement same function signatures as
// the service interface but this will be used to abstract
// database access
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}

/*
We need an interface to connect our business logic to our
repository (storage)
*/
