package errorx

import "fmt"

// Functions to used in various specific error handling scenarios

// PanicOnErr panics if any of the given errors is not nil.
func PanicOnErr(err ...error) {
	for _, e := range err {
		if e != nil {
			panic(e)
		}
	}
}

// Chain runs func one by one until an error occurred.
func Chain(fn ...func() error) error {
	for _, f := range fn {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

// Wrap returns err with given msg
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", message, err)
}

func Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
}
