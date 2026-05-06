package test

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
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
	// ruleid: http-error-missing-return
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
		// ruleid: http-error-missing-return
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
