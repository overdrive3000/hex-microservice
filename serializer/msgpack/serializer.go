package msgpack

import (
	"github.com/overdrive3000/hex-microservice/shortener"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

// Redirect object
type Redirect struct{}

// Decode a msgpack structure
func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode a msgpack structure
func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer,Redirect.Encode")
	}
	return rawMsg, nil
}
