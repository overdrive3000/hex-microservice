package shortener

// RedirectService is a service interface
// defines all the action implemented in this microserviceS
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}

/*
This file is where business logic should be implemented (I guess)
this way of work is compatible with go-kit model
*/
