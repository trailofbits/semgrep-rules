package testdata

import (
	"fmt"
	"log"
	"os"
	"runtime"

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
// TP3: branchy reassignment — `err` is reassigned only inside an inner
// `if branch { ... return }`, so along the wrap path `err` is still
// proven-nil from the outer guard. The reassignment lives in a sibling
// block, not the wrap's enclosing block.
func tp3(branch bool) error {
	err := anotherCall()
	if err != nil {
		return err
	}

	if someCondition() {
		if branch {
			_, err = mayFail()
			return err
		}

		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "branchy reassign")
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

// TP17: chained calls — each call uses `:=` (re)binding `err` and is
// followed by its own `if err != nil { return ... }` guard. After both
// guards, `err` is provably nil, so the wrap inside the trailing
// non-err `if $COND { ... }` is the bug.
func tp17(threshold int) error {
	id, err := someCall()
	if err != nil {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrapf(err, "lookup %q", id)
	}

	count, err := mayFail()
	if err != nil {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrapf(err, "count %q", id)
	}

	if count < threshold {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrapf(err, "below threshold %d", threshold)
	}
	return nil
}

// TN11: classic FP shape — `err` is reassigned via tuple-assign between
// the guard and the wrap, so the wrap site is on a freshly-assigned
// (potentially non-nil) err, not the proven-nil one.
func tn11(cond bool) error {
	res, err := mayFail()
	if err != nil {
		return errors.Wrap(err, "first call")
	}

	if res > 0 {
		return nil
	}

	if cond {
		_, err = mayFail()
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "second call")
	}
	return nil
}

// TN12: same as TN11 but the reassignment lives in the surrounding
// function scope (not inside the `if cond` block). The wrap is the last
// statement of an enclosing `if cond { ... }` block.
func tn12(cond bool) error {
	res, err := mayFail()
	if err != nil {
		return errors.Wrap(err, "first call")
	}

	if cond {
		if res > 0 {
			return nil
		}

		_, err = mayFail()

		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "second call")
	}
	return nil
}

// TN13: single-LHS reassignment between guard and wrap.
func tn13(cond bool) error {
	err := anotherCall()
	if err != nil {
		return err
	}

	if cond {
		err = anotherCall()
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrapf(err, "single reassign")
	}
	return nil
}

// TN14: reassignment happens inside an else branch before the wrap —
// Case B FP shape.
func tn14() error {
	err := anotherCall()
	if err != nil {
		return err
	} else {
		_, err = mayFail()
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "reassigned in else")
	}
}

// TN15: alt-named err reassigned between guard and wrap.
func tn15(cond bool) error {
	parseErr := anotherCall()
	if parseErr != nil {
		return parseErr
	}

	if cond {
		parseErr = anotherCall()
		// ok: pkg-errors-wrap-nil-err
		return errors.WithMessage(parseErr, "after reassign")
	}
	return nil
}

// TN16: multiple reassignments before the wrap — still a TN because the
// last value of err is not provably nil.
func tn16(cond bool) error {
	_, err := mayFail()
	if err != nil {
		return err
	}

	if cond {
		_, err = mayFail()
		_, err = mayFail()
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "after multiple reassigns")
	}
	return nil
}

func threeVals() (int, string, error) { return 0, "", nil }

// TN17: 3-value tuple reassignment between guard and wrap (err is last,
// per Go convention).
func tn17(cond bool) error {
	_, _, err := threeVals()
	if err != nil {
		return err
	}

	if cond {
		_, _, err = threeVals()
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "three-tuple reassign")
	}
	return nil
}

// TN18: 3-value tuple reassignment inside an else branch.
func tn18() error {
	_, _, err := threeVals()
	if err != nil {
		return err
	} else {
		_, _, err = threeVals()
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "three-tuple else reassign")
	}
}

var ErrSentinelA = errors.New("sentinel a")
var ErrSentinelB = errors.New("sentinel b")

// TN19: compound `if $ERR != nil && ...` — `&&` guarantees err is
// non-nil inside the body, so the wrap is intentional.
func tn19() error {
	_, err := mayFail()
	if err != nil {
		return err
	}

	if err != nil && !errors.Is(err, ErrSentinelA) {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "compound err first")
	}
	return nil
}

// TN20: compound condition with `$ERR != nil` as the LAST conjunct.
func tn20(extra bool) error {
	_, err := mayFail()
	if err != nil {
		return err
	}

	if extra && err != nil {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "compound err last")
	}
	return nil
}

// TN21: compound condition with `$ERR != nil` in the MIDDLE position
// of a 3-conjunct `&&` chain.
func tn21() error {
	_, err := mayFail()
	if err != nil {
		return err
	}

	if err != nil && !errors.Is(err, ErrSentinelA) && !errors.Is(err, ErrSentinelB) {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrapf(err, "compound err middle")
	}
	return nil
}

// TN22: reverse guard `if $ERR == nil { return ... }` followed by a
// wrap on the err-non-nil path.
func tn22(cond bool) error {
	err := anotherCall()
	if err != nil {
		return err
	}

	if err == nil {
		return nil
	}

	if cond {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "after err==nil reverse guard")
	}
	return nil
}

// TP18: guard body has a logging stmt before `return` — `err` is still
// provably nil after the guard, so the wrap inside the trailing
// non-err `if` is the bug.
func tp18() error {
	err := anotherCall()
	if err != nil {
		log.Print("init failed: ", err)
		return err
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "after multi-stmt guard")
	}
	return nil
}

// TP19: assignment-form wrap inside the else of `if err != nil { ... } else { ... }`.
// Wrap returns nil for nil input, so `err` stays nil.
func tp19() error {
	err := anotherCall()
	if err != nil {
		return err
	} else {
		// ruleid: pkg-errors-wrap-nil-err
		err = errors.Wrap(err, "assignment-form")
		return err
	}
}

// TN24: assignment-form wrap in else preceded by a single-LHS reassign
// — `err` is no longer provably nil, so the wrap is intentional.
func tn24() error {
	err := anotherCall()
	if err != nil {
		return err
	} else {
		err = anotherCall()
		// ok: pkg-errors-wrap-nil-err
		err = errors.Wrap(err, "single reassign in else")
		return err
	}
}

// TN25: assignment-form wrap in else preceded by a tuple reassign.
func tn25() error {
	_, err := mayFail()
	if err != nil {
		return err
	} else {
		_, err = mayFail()
		// ok: pkg-errors-wrap-nil-err
		err = errors.Wrap(err, "tuple reassign in else")
		return err
	}
}

// TN23: reverse guard inside a deferred closure — the wrap runs only
// when the captured err is non-nil.
func tn23() (err error) {
	err = anotherCall()
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			return
		}

		if cleanupErr := anotherCall(); cleanupErr != nil {
			if err == nil {
				err = cleanupErr
				return
			}
			// ok: pkg-errors-wrap-nil-err
			err = errors.Wrap(err, "with cleanup failure")
		}
	}()

	return nil
}

// TP20: panic-based guard. After `if err != nil { panic(err) }`, `err`
// is provably nil, so wrapping it inside the trailing non-err `if`
// swallows the failure path.
func tp20() error {
	err := anotherCall()
	if err != nil {
		panic(err)
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "after panic guard")
	}
	return nil
}

// TP21: log.Fatal-based guard.
func tp21() error {
	err := anotherCall()
	if err != nil {
		log.Fatal("init: ", err)
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "after log.Fatal guard")
	}
	return nil
}

// TP22: log.Fatalf-based guard.
func tp22() error {
	err := anotherCall()
	if err != nil {
		log.Fatalf("init: %v", err)
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrapf(err, "after log.Fatalf guard")
	}
	return nil
}

// TP23: log.Fatalln-based guard.
func tp23() error {
	err := anotherCall()
	if err != nil {
		log.Fatalln("init failed", err)
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.WithMessage(err, "after log.Fatalln guard")
	}
	return nil
}

// TP24: os.Exit-based guard.
func tp24() error {
	err := anotherCall()
	if err != nil {
		os.Exit(1)
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.WithStack(err)
	}
	return nil
}

// TP25: runtime.Goexit-based guard.
func tp25() error {
	err := anotherCall()
	if err != nil {
		runtime.Goexit()
	}

	if someCondition() {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.WithMessagef(err, "after runtime.Goexit guard")
	}
	return nil
}

// TN26: reverse guard with panic — `err` is non-nil after.
func tn26(cond bool) error {
	err := anotherCall()
	if err != nil {
		return err
	}

	if err == nil {
		panic("unreachable: err already non-nil")
	}

	if cond {
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "after err==nil panic guard")
	}
	return nil
}

// TN27: reverse guard with log.Fatal — `err` is non-nil after.
func tn27(cond bool) error {
	err := anotherCall()
	if err != nil {
		return err
	}

	if err == nil {
		log.Fatal("unreachable")
	}

	if cond {
		// ok: pkg-errors-wrap-nil-err
		return errors.WithStack(err)
	}
	return nil
}

// TP26: wrap inside `else if` arm — Go parses this as
// `if err != nil { return err } else { if cond { return errors.Wrap(...) } }`,
// and inside the inner if (which runs only when err is nil) the wrap
// swallows the error path.
func tp26(cond bool) error {
	err := anotherCall()
	if err != nil {
		return err
	} else if cond {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "else-if arm")
	}
	return nil
}

// TP27: wrap inside a nested `else if` (depth 2).
func tp27(cond1, cond2 bool) error {
	err := anotherCall()
	if err != nil {
		return err
	} else if cond1 {
		return nil
	} else if cond2 {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.WithStack(err)
	}
	return nil
}

// TP28: assignment-form wrap inside an `else if` arm.
func tp28(cond bool) error {
	err := anotherCall()
	if err != nil {
		return err
	} else if cond {
		// ruleid: pkg-errors-wrap-nil-err
		err = errors.Wrap(err, "assignment-form in else-if")
		return err
	}
	return nil
}

// TP29: `else if` arm with init clause.
func tp29() error {
	err := anotherCall()
	if err != nil {
		return err
	} else if x := 1; x > 0 {
		// ruleid: pkg-errors-wrap-nil-err
		return errors.Wrapf(err, "init else-if: %d", x)
	}
	return nil
}

// TN28: `else if` arm where err is reassigned before the wrap — wrap
// is on a freshly-assigned (potentially non-nil) err.
func tn28(cond bool) error {
	err := anotherCall()
	if err != nil {
		return err
	} else if cond {
		_, err = mayFail()
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "after reassign in else-if")
	}
	return nil
}

// TN29: `:=` shadow inside the wrap branch. `_, err := mayFail()` in a
// nested scope declares a fresh `err` distinct from the proven-nil
// outer one, so the wrap is on the fresh, possibly non-nil shadow.
func tn29(cond bool) error {
	err := anotherCall()
	if err != nil {
		return err
	}

	if cond {
		_, err := mayFail()
		// ok: pkg-errors-wrap-nil-err
		return errors.Wrap(err, "shadowed via :=")
	}
	return nil
}

