package testdata

import (
	"fmt"
	"github.com/pkg/errors"
)

func someCall() (string, error) { return "", nil }
func anotherCall() error        { return nil }
func mayFail() (int, error)     { return 0, nil }
func someCondition() bool       { return true }

// TP1: simplest case — wrap inside non-err if-block after a guard.
func tp1() error {
	err := anotherCall()
	if err != nil {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "first")
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrapf(err, "second %s", "x")
	}

	return nil
}

// TP2: errors.Wrap (not Wrapf) variant.
func tp2() error {
	_, err := someCall()
	if err != nil {
		return err
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "boom")
	}

	return nil
}


// TP4: bug nested inside a for loop body.
func tp4(items []string) error {
	for range items {
		_, err := mayFail()
		if err != nil {
			// ok: pkg-errors-wrap-nil-err
			return errors.Wrap(err, "loop")
		}

		if someCondition() {
			// ruleid: pkg-errors-wrap-nil-err
			return errors.Wrapf(err, "post check")
		}
	}
	return nil
}

// TP5: wrap inside a complex non-err condition
func tp5(xs []int, xx int) error {
	_, err := mayFail()
	if err != nil {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "init")
	}

	if len(xs) < xx {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrapf(err, "too short, got %d want %d", len(xs), xx)
	}
	return nil
}

// TP6: guard returns a non-wrap value but err is still provably nil after.
func tp6() error {
	err := anotherCall()
	if err != nil {
		return fmt.Errorf("init: %w", err)
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrapf(err, "later")
	}
	return nil
}

// TP7: many statements between guard and wrap.
func tp7() error {
	_, err := someCall()
	if err != nil {
		return err
	}

	a := 1
	b := 2
	c := a + b
	_ = c

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "noisy")
	}
	return nil
}

// TP8: bug inside an `else` branch — still an `if $COND { ... }` body.
func tp8() error {
	err := anotherCall()
	if err != nil {
		return err
	}

	if someCondition() {
		return nil
	} else {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "else branch")
	}
}

// TP9: bug nested inside an inner if (deeper than the guard).
func tp9(n int) error {
	err := anotherCall()
	if err != nil {
		return err
	}

	if n > 0 {
		if n > 100 {
			// ruleid: pkg-errors-wrap-nil-err
			return errors.Wrapf(err, "way too big: %d", n)
		}
	}
	return nil
}

// TP10: errors.WithMessage on provably-nil err — same nil-passthrough bug.
func tp10() error {
	err := anotherCall()
	if err != nil {
		return err
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.WithMessage(err, "with-message")
	}
	return nil
}

// TP11: errors.WithMessagef on provably-nil err.
func tp11() error {
	err := anotherCall()
	if err != nil {
		return err
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.WithMessagef(err, "with-message %d", 42)
	}
	return nil
}

// TP12: errors.WithStack on provably-nil err — also returns nil for nil input.
func tp12() error {
	err := anotherCall()
	if err != nil {
		return err
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.WithStack(err)
	}
	return nil
}

// TP13: alternate variable name `e` — same bug shape.
func tp13() error {
	e := anotherCall()
	if e != nil {
		return e
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrap(e, "alt name e")
	}
	return nil
}

// TP14: alternate variable name `readErr` — same bug shape.
func tp14() error {
	readErr := anotherCall()
	if readErr != nil {
		return readErr
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrapf(readErr, "alt name readErr")
	}
	return nil
}

// TP15: PascalCase suffix — `parseErr` (still err-like).
func tp15() error {
	parseErr := anotherCall()
	if parseErr != nil {
		return parseErr
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.WithMessage(parseErr, "alt name parseErr")
	}
	return nil
}

// TP16: invalid else branch
func tp16() error {
	parseErr := anotherCall()
	if parseErr != nil {
		return parseErr
	} else {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.WithMessage(parseErr, "alt name parseErr")
	}
}

// TN1: wrap inside the `if err != nil` block — canonical correct usage.
func tn1() error {
	err := anotherCall()
	if err != nil {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "legit")
	}
	return nil
}

// TN2: wrap inside `if err == nil` (intentional success-side check).
func tn2() error {
	err := anotherCall()
	if err != nil {
		return err
	}

	if err == nil {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "intentional nil branch")
	}
	return nil
}

// TN3: legitimate "wrap result of trailing call" idiom — wrap is at function
// level, NOT inside a non-err if-block. Wrapping nil returns nil, so this
// pattern is intentional even when err might be nil.
func tn3() error {
	err := anotherCall()
	if err != nil {
		return err
	}

	err = anotherCall()
	// ok: pkg-errors-wrap-nil-err
	return errors.Wrap(err, "trailing call")
}

// TN4: a different error variable is being wrapped — the literal `err` token
// in the pattern doesn't bind to it.
func tn4() error {
	err := anotherCall()
	if err != nil {
		return err
	}

	otherErr := anotherCall()
	if someCondition() {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(otherErr, "other variable")
	}
	return nil
}

// TN5: no preceding `if err != nil { return }` guard at all — `err` may be
// non-nil at the wrap site.
func tn5() error {
	err := anotherCall()
	if someCondition() {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "no guard")
	}
	return nil
}

// TN6: wrap with a literal `nil` argument — not the `err` variable shape.
func tn6() error {
	err := anotherCall()
	if err != nil {
		return err
	}
	if someCondition() {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(nil, "explicitly nil")
	}
	return nil
}

// TN7: standard library `errors` package usage — irrelevant. Included to
// confirm we don't accidentally trip on it (no Wrap in std `errors`, but
// this stresses that the rule is scoped to Wrap/Wrapf calls).
func tn7() error {
	err := anotherCall()
	if err != nil {
		return err
	}
	if someCondition() {
		// ok: pkg-errors-wrap-nil-err
		return fmt.Errorf("plain %w", err)
	}
	return nil
}

// TN8: WithMessage / WithMessagef / WithStack inside `if err != nil` —
// canonical correct usage.
func tn8() error {
	err := anotherCall()
	if err != nil {
		// ok: pkg-errors-wrap-nil-err
		return errors.WithMessage(err, "ctx")
	}
	err = anotherCall()
	if err != nil {
		// ok: pkg-errors-wrap-nil-err
		return errors.WithMessagef(err, "ctx %d", 1)
	}
	err = anotherCall()
	if err != nil {
		// ok: pkg-errors-wrap-nil-err
		return errors.WithStack(err)
	}
	return nil
}

// TN9: variable name `result` doesn't match the err-like regex even after a
// preceding `if err != nil { return }` — we don't flag wraps of
// non-error-shaped names.
func tn9() error {
	err := anotherCall()
	if err != nil {
		return err
	}

	var result error = anotherCall()
	if someCondition() {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(result, "non-err-named")
	}
	return nil
}

// TN10: nil-check on a DIFFERENT err variable than the one being wrapped —
// the metavariable unification means $ERR must be the same name on both
// sides, so this is a TN.
func tn10() error {
	outerErr := anotherCall()
	if outerErr != nil {
		return outerErr
	}

	innerErr := anotherCall()
	if innerErr != nil {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(innerErr, "different err var than the guard")
	}
	return nil
}
