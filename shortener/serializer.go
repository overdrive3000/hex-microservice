package shortener

// RedirectSerializer implements interface for endonde and decode messages
type RedirectSerializer interface {
	Decode(input []byte) (*Redirect, error)
	Encode(input *Redirect) ([]byte, error)
}
