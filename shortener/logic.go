package shortener

import (
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
)

var (
	// ErrRedirectNotFound custom error
	ErrRedirectNotFound = errors.New("Redirect Not Found")

	// ErrRedirectInvalid custom error
	ErrRedirectInvalid = errors.New("Redirect invalid")
)

// Redirect service structure which implements the redirect service interface
type redirectService struct {
	redirectRepo RedirectRepository
}

// NewRedirectService instantiate a new redirect service object
func NewRedirectService(redirectRepo RedirectRepository) RedirectService {
	return &redirectService{
		redirectRepo,
	}
}

// Abstract the find method in order to don't take care which
// kind of repository is being used.
func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().Unix()
	return r.redirectRepo.Store(redirect)
}
