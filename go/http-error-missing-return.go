package test

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)


var errSentinel = errors.New("sentinel")

func checkAuth(r *http.Request) error                  { return errors.New("x") }
func getUser(r *http.Request) (string, error)               { return "", errors.New("x") }
func doSomething(r *http.Request) error                     { return errors.New("x") }
func processRequest(w http.ResponseWriter, r *http.Request) {}

func handlerAuthCheckBug(w http.ResponseWriter, r *http.Request) {
	if err := checkAuth(r); err != nil {
		// ruleid: http-error-missing-return
		http.Error(w, fmt.Sprintf("auth check failed: %s", err.Error()), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func handlerSimpleBug(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if err != nil {
		// ruleid: http-error-missing-return
		http.Error(w, "user not found", http.StatusNotFound)
	}

	w.Write([]byte("user: " + user))
}

func handlerLogThenContinue(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Print("bad method")
		// ruleid: http-error-missing-return
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	processRequest(w, r)
}

// http.Error followed by another statement in the same block — Case B.
func handlerErrorFollowedByWrite(w http.ResponseWriter, r *http.Request) {
	// ruleid: http-error-missing-return-followed-by-stmt
	http.Error(w, "boom", http.StatusInternalServerError)
	w.Write([]byte("leaked body"))
}

func handlerCorrectReturn(w http.ResponseWriter, r *http.Request) {
	if err := checkAuth(r); err != nil {
		// ok: http-error-missing-return
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handlerCorrectLogThenReturn(w http.ResponseWriter, r *http.Request) {
	if err := doSomething(r); err != nil {
		log.Print("error: ", err)
		// ok: http-error-missing-return
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("ok"))
}

func helperReturningError(w http.ResponseWriter, r *http.Request) error {
	if err := doSomething(r); err != nil {
		// ok: http-error-missing-return
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func handlerLastStatement(w http.ResponseWriter, r *http.Request) {
	if err := doSomething(r); err != nil {
		// ok: http-error-missing-return
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlerIfElseBothReturn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// ok: http-error-missing-return
		http.Error(w, "GET not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		w.Write([]byte("ok"))
		return
	}
}

func handlerNestedReturn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.ContentLength == 0 {
			// ok: http-error-missing-return
			http.Error(w, "empty body", http.StatusBadRequest)
			return
		}
	}

	w.Write([]byte("ok"))
}

// The if-block is the LAST statement of the function — there is no
// `$NEXT` after it, so Case A's pattern-inside requirement (a trailing
// statement after the if-chain) is not met and the rule does not fire.
func handlerEndOfFunction(w http.ResponseWriter, r *http.Request) {
	if err := doSomething(r); err != nil {
		// ok: http-error-missing-return
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// log.Fatal terminates execution; treat as a control-flow terminator.
func handlerLogFatal(w http.ResponseWriter, r *http.Request) {
	if err := doSomething(r); err != nil {
		// ok: http-error-missing-return
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("unrecoverable: ", err)
	}

	w.Write([]byte("ok"))
}

// log.Fatalln terminator.
func handlerLogFatalln(w http.ResponseWriter, r *http.Request) {
	if err := doSomething(r); err != nil {
		// ok: http-error-missing-return
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("unrecoverable")
	}

	w.Write([]byte("ok"))
}

// log.Fatalf terminator.
func handlerLogFatalf(w http.ResponseWriter, r *http.Request) {
	if err := doSomething(r); err != nil {
		// ok: http-error-missing-return
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("unrecoverable: %v", err)
	}

	w.Write([]byte("ok"))
}

// os.Exit terminator.
func handlerOsExit(w http.ResponseWriter, r *http.Request) {
	if err := doSomething(r); err != nil {
		// ok: http-error-missing-return
		http.Error(w, err.Error(), http.StatusInternalServerError)
		os.Exit(1)
	}

	w.Write([]byte("ok"))
}

// Terminator follows the if-block at the enclosing scope (Case A's
// "next-stmt is terminator" exclusion).
func handlerNextStmtLogFatal(w http.ResponseWriter, r *http.Request) {
	if err := doSomething(r); err != nil {
		// ok: http-error-missing-return
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Fatal("downstream did not run")
}

// Case B: log.Fatal immediately after http.Error — not flagged.
func handlerCaseBLogFatal(w http.ResponseWriter, r *http.Request) {
	// ok: http-error-missing-return
	http.Error(w, "boom", http.StatusInternalServerError)
	log.Fatalf("boom %d", 1)
}

// Case B: os.Exit immediately after http.Error — not flagged.
func handlerCaseBOsExit(w http.ResponseWriter, r *http.Request) {
	// ok: http-error-missing-return
	http.Error(w, "boom", http.StatusInternalServerError)
	os.Exit(2)
}

// Sanity: an identifier whose name *starts with* "return" must not be
// mistaken for a return statement by the terminator regex.
func handlerReturnPrefixIsNotTerminator(w http.ResponseWriter, r *http.Request) {
	if err := doSomething(r); err != nil {
		// ruleid: http-error-missing-return-followed-by-stmt
		http.Error(w, err.Error(), http.StatusInternalServerError)
		returnedSomething()
	}

	w.Write([]byte("ok"))
}

func returnedSomething() {}

// http.Error in the middle branch of an `else if` chain — handler
// keeps running past the if-else and reaches the trailing write.
func handlerElseIfMissingReturn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("get"))
	} else if r.Method == "POST" {
		// ruleid: http-error-missing-return
		http.Error(w, "no post", http.StatusMethodNotAllowed)
	}

	w.Write([]byte("done"))
}

// `else if` middle branch with a `return` is fine.
func handlerElseIfWithReturn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("get"))
	} else if r.Method == "POST" {
		// ok: http-error-missing-return
		http.Error(w, "no post", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("done"))
}

// `else if` middle branch where a terminator follows the entire
// if-else chain — no fall-through, so no bug.
func handlerElseIfNextTerm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("get"))
	} else if r.Method == "POST" {
		// ok: http-error-missing-return
		http.Error(w, "no post", http.StatusMethodNotAllowed)
	}

	log.Fatal("downstream did not run")
}

// Trailing `return` after the inner if-else satisfies Case A's
// terminator subtraction at the outer scope, so neither inner
// `http.Error` is flagged.
func handlerNestedNoInnerReturn(w http.ResponseWriter, r *http.Request) {
	if err := doSomething(r); err != nil {
		if errors.Is(err, errSentinel) {
			// ok: http-error-missing-return
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			// ok: http-error-missing-return
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	w.Write([]byte("ok"))
}

// `continue` legitimately ends an iteration; the response was written for
// this request item, so the loop just moves on.
func handlerForLoopContinue(w http.ResponseWriter, r *http.Request, items []string) {
	for range items {
		if err := doSomething(r); err != nil {
			// ok: http-error-missing-return
			http.Error(w, err.Error(), http.StatusInternalServerError)
			continue
		}
	}
}

// `break` is treated as a terminator by the rule even though it exits
// the loop, not the handler. Conservative choice: legitimate retry-style
// loops commonly write a single error and `break`, and flagging them
// produces noise. Real fall-through bugs after `break` (loop is followed
// by code that double-writes) are accepted as a known FN.
func handlerForLoopBreak(w http.ResponseWriter, r *http.Request) {
	for {
		if r.Method != "POST" {
			// ok: http-error-missing-return
			http.Error(w, "no", http.StatusMethodNotAllowed)
			break
		}
	}
}

// `goto` to a cleanup label is also a control-flow terminator for the
// purpose of this rule.
func handlerGoto(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// ok: http-error-missing-return
		http.Error(w, "no", http.StatusMethodNotAllowed)
		goto cleanup
	}
	w.Write([]byte("ok"))
cleanup:
}

// `runtime.Goexit` terminates the current goroutine — also a terminator.
func handlerRuntimeGoexit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// ok: http-error-missing-return
		http.Error(w, "no", http.StatusMethodNotAllowed)
		runtime.Goexit()
	}
	w.Write([]byte("ok"))
}

// 3-way else-if chain: http.Error in the deepest branch, no return, and
// code follows the entire chain. Should be flagged. (Previously a FN.)
func handlerThreeWayElseIf(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("g"))
	} else if r.Method == "POST" {
		w.Write([]byte("p"))
	} else if r.Method == "DELETE" {
		// ruleid: http-error-missing-return
		http.Error(w, "no del", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("done"))
}

// 4-way else-if: still a bug, exercises arbitrary depth.
func handlerFourWayElseIf(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("g"))
	} else if r.Method == "POST" {
		w.Write([]byte("p"))
	} else if r.Method == "PUT" {
		w.Write([]byte("u"))
	} else if r.Method == "DELETE" {
		// ruleid: http-error-missing-return
		http.Error(w, "no del", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("done"))
}

// 4-way else-if where the trailing chain is followed by a terminator —
// no fall-through, so no bug.
func handlerFourWayElseIfNextTerm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("g"))
	} else if r.Method == "POST" {
		w.Write([]byte("p"))
	} else if r.Method == "PUT" {
		w.Write([]byte("u"))
	} else if r.Method == "DELETE" {
		// ok: http-error-missing-return
		http.Error(w, "no del", http.StatusMethodNotAllowed)
	}
	log.Fatal("downstream did not run")
}

// 4-way else-if where the deepest branch terminates with `return` — no
// bug. Verifies the inner-block-tail check correctly excludes calls that
// have a terminator after them in their own block.
func handlerFourWayElseIfWithReturn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("g"))
	} else if r.Method == "POST" {
		w.Write([]byte("p"))
	} else if r.Method == "PUT" {
		w.Write([]byte("u"))
	} else if r.Method == "DELETE" {
		// ok: http-error-missing-return
		http.Error(w, "no del", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("done"))
}

// Switch case where http.Error is the tail of a case body and the
// switch is followed by non-terminating code. Switch cases don't
// fall through in Go, but the function continues past the switch.
func handlerSwitchCase(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// ruleid: http-error-missing-return
		http.Error(w, "no posts", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("leaked"))
}

// Switch case with `return` after http.Error — fine.
func handlerSwitchCaseReturn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// ok: http-error-missing-return
		http.Error(w, "no posts", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("ok"))
}

// Switch is the LAST stmt of the function — handler returns naturally,
// so missing terminator inside a case is harmless.
func handlerSwitchCaseEndOfFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// ok: http-error-missing-return
		http.Error(w, "no posts", http.StatusMethodNotAllowed)
	}
}

// Switch is followed by a terminator — no fall-through past the switch.
func handlerSwitchCaseNextTerm(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// ok: http-error-missing-return
		http.Error(w, "no posts", http.StatusMethodNotAllowed)
	}
	log.Fatal("downstream did not run")
}

// Switch with default clause — same fall-through bug.
func handlerSwitchDefault(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("g"))
	default:
		// ruleid: http-error-missing-return
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("leaked"))
}

// Multi-case switch — fall-through bug in a middle case body, with
// other case bodies present alongside.
func handlerSwitchMultiCase(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("g"))
	case "POST":
		// ruleid: http-error-missing-return
		http.Error(w, "no posts", http.StatusMethodNotAllowed)
	case "PUT":
		w.Write([]byte("u"))
	}
	w.Write([]byte("done"))
}

// Switch with init clause: `switch $INIT; $X { ... }` followed by
// trailing code — exercises the init-form `pattern-inside` branches.
func handlerSwitchWithInit(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case "POST":
		// ruleid: http-error-missing-return
		http.Error(w, "no", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("done"))
}

type server struct{}

// Method-receiver handler — same fall-through bug, different signature.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := doSomething(r); err != nil {
		// ruleid: http-error-missing-return
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte("ok"))
}

// Method-receiver handler with `return` — fine.
func (s *server) Get(w http.ResponseWriter, r *http.Request) {
	if err := doSomething(r); err != nil {
		// ok: http-error-missing-return
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("ok"))
}

// Anonymous handler passed to http.HandleFunc — bug fires inside the closure.
func registerHandlerBug() {
	http.HandleFunc("/x", func(w http.ResponseWriter, r *http.Request) {
		if err := doSomething(r); err != nil {
			// ruleid: http-error-missing-return
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte("ok"))
	})
}

// Anonymous handler with `return` inside the closure — fine.
func registerHandlerOK() {
	http.HandleFunc("/x", func(w http.ResponseWriter, r *http.Request) {
		if err := doSomething(r); err != nil {
			// ok: http-error-missing-return
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("ok"))
	})
}

// `select` cases are caught by the existing switch-shaped patterns
// (Semgrep's Go AST treats `select` similarly to `switch`), so a
// fall-through bug inside a `select` case is flagged the same way.
func handlerSelectFallThrough(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-ctx.Done():
		// ruleid: http-error-missing-return
		http.Error(w, "timeout", http.StatusGatewayTimeout)
	}
	w.Write([]byte("done"))
}
