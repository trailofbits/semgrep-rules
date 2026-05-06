package fixtures

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

func someFunc() error                         { return nil }
func multiReturn() (int, error)               { return 0, nil }
func tripleReturn() (int, string, error)      { return 0, "", nil }
func xFunc(_ context.Context) error     { return nil }
func compute() int                            { return 0 }
func httpServer() *http.Server                { return &http.Server{} }
func newCtx() context.Context                 { return context.Background() }
func errorsJoin(errs ...error) error          { return errors.Join(errs...) }
func wrap(s string, e error) error            { return fmt.Errorf(s+": %w", e) }

func bugSimple(err error) error {
	// ruleid: shadowed-err-check
	if xErr := xFunc(newCtx()); err != nil {
		return wrap("x", xErr)
	}
	return nil
}

func bugListenAndServe() (err error) {
	srv := httpServer()
	// ruleid: shadowed-err-check
	if svrErr := srv.ListenAndServe(); err != nil {
		err = errorsJoin(err, wrap("listen", svrErr))
	}
	return
}

func bugMultiAssign() (err error) {
	// ruleid: shadowed-err-check
	if v, xErr := multiReturn(); err != nil {
		_ = v
		err = errorsJoin(err, xErr)
	}
	return
}

func bugTripleAssign() (err error) {
	// ruleid: shadowed-err-check
	if a, b, innerErr := tripleReturn(); err != nil {
		_ = a
		_ = b
		err = errorsJoin(err, innerErr)
	}
	return
}

func bugEqualsNil() (err error) {
	// ruleid: shadowed-err-check
	if xErr := someFunc(); err == nil {
		_ = xErr
	}
	return
}


func okSimple() error {
	// ok: shadowed-err-check
	if err := someFunc(); err != nil {
		return err
	}
	return nil
}

func okSimpleEqualsNil() {
	// ok: shadowed-err-check
	if err := someFunc(); err == nil {
		fmt.Println("success")
	}
}

func okMultiAssign() error {
	// ok: shadowed-err-check
	if v, err := multiReturn(); err != nil {
		_ = v
		return err
	}
	return nil
}

func okTripleAssign() error {
	// ok: shadowed-err-check
	if a, b, err := tripleReturn(); err != nil {
		_ = a
		_ = b
		return err
	}
	return nil
}

func okCustomNamedErr() error {
	// ok: shadowed-err-check
	if xErr := xFunc(newCtx()); xErr != nil {
		return xErr
	}
	return nil
}

func okCustomNamedErrEqualsNil() {
	// ok: shadowed-err-check
	if svrErr := someFunc(); svrErr == nil {
		fmt.Println("ok")
	}
}

func okNonErrorVariable() {
	// ok: shadowed-err-check
	if err := compute(); err > 0 {
		fmt.Println(err)
	}
}

func okNonErrorMultiAssign() {
	// ok: shadowed-err-check
	if v, ok := map[string]int{"a": 1}["a"]; ok {
		_ = v
	}
}

func okConditionUsesOtherBoolean() {
	flag := true
	// ok: shadowed-err-check
	if err := someFunc(); flag {
		_ = err
	}
}

// Variable names that merely contain `err` as a substring (e.g.
// `terraform`, `cherry`) must not be treated as error variables.
func okSubstringNotErrLike(cherry int) {
	// ok: shadowed-err-check
	if terraform := compute(); cherry > 0 {
		_ = terraform
	}
}
