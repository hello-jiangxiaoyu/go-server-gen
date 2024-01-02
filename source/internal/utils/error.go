package utils

import "errors"

func WrapError(err1 error, err2 error) error {
	if err1 == nil && err2 == nil {
		return nil
	} else if err1 == nil {
		return err2
	} else if err2 == nil {
		return err1
	}

	return errors.New(err2.Error() + ": " + err1.Error())
}

func WithMessage(err error, msg string) error {
	if err == nil && msg == "" {
		return nil
	} else if err == nil {
		return errors.New(msg)
	} else if msg == "" {
		return err
	}

	return errors.New(msg + ": " + err.Error())
}

// DeferErr handle defer function err
func DeferErr(errFunc func() error) {
	if err := errFunc(); err != nil {
		println("### Defer err: ", err.Error())
	}
}
