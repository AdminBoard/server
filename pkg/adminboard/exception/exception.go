package exception

import "github.com/pkg/errors"

func Wrap(err error, msg string) error {
	e := errors.Cause(err)
	if e != nil {
		if unwrap := errors.Unwrap(e); unwrap != nil {
			e = errors.Wrap(unwrap, msg)
		} else {
			e = errors.Wrap(e, msg)
		}
	} else {
		e = errors.New(msg)
	}
	return e
}
