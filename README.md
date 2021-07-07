# Trail of Bits public Semgrep rules

This repository contains Semgrep rules developed by Trail of Bits and made available to the public. They are part of our ongoing development efforts and are used in our security audits, vulnerability reseach, and internal projects. They will evolve over time as we identify new techniques.

## Using Semgrep

The easiest way to run the rules is to run them from the [Semgrep registry](https://semgrep.dev/p/trailofbits). To do so, navigate to the root folder of your project and run the following:

```shell
$ semgrep --config "p/trailofbits"
```

Alternatively, you can clone this repository, navigate to the root folder of your project, and run individual rules using the command below :

```shell
$ semgrep --config /path/to/semgrep-rules/semgreprule.yml
```

To run all rules from the cloned repository:

```shell
$ semgrep --config /path/to/semgrep-rules/ .
```

## Useful flags

Semgrep will run against all supported code files except for those in your `.gitignore` file. If you want to run the rules against all files and directories, including those in your `.gitignore`, add the `--no-git-ignore` flag.

```shell
$ semgrep --config /path/to/semgrep-rules/ . --no-git-ignore
```

You can also tell Semgrep to ignore files and directories that match any pattern. For instance, if you want to tell Semgrep to ignore all Go test files you can run the following:


```shell
$ $ semgrep --config /path/to/semgrep-rules/ . --exclude='*_test.go'
```

Use `-o` to output results to a file:

```shell
$ semgrep --config /path/to/semgrep-rules/hanging-goroutine.yml -o leaks.txt'
```

## Rules

Rule ID | Language | What it Finds
--- | --- | ---
[anonymous-race-condition](go/anonymous-race-condition.yml) | Go | Race conditions within anonymous goroutines
[hanging-goroutine](go/hanging-goroutine.yml) | Go | Goroutine leaks
[iterate-over-empty-collection](go/iterate-over-empty-collection.yml) | Go | Iterations over empty collection
[nil-check-after-call](go/nil-check-after-call.yml) | Go | Possible nil dereferences
[questionable-assignment](go/questionable-assignment.yml) | Go | Possible unintentional assignment when an error occurs
[nondeterministic-select](go/nondeterministic-select.yml) | Go | Nondeterministic `select` logic. 
[servercodec-readrequestbody-unhandled-nil](go/servercodec-readrequestbody-unhandled-nil.yml) | Go | Possible incorrect `ServerCodec` interface implementation
[sleep-used-for-synchronizations](go/sleep-used-for-synchronizations.yml) | Go | Uses `time.Sleep` for goroutine synchronizations
[string-to-int-signedness-cast](go/string-to-int-signedness-cast.yml) | Go | Integer underflows
[sync-mutex-value-copied](go/sync-mutex-value-copied.yml) | Go | Copying of `sync.Mutex` via value receivers
[waitgroup-add-called-inside-goroutine](go/waitgroup-add-called-inside-goroutine.yml) | Go | Calls to `sync.WaitGroup.Add` inside of anonymous goroutines
[waitgroup-wait-inside-loop](go/waitgroup-wait-inside-loop.yml) | Go | Calls to `sync.WaitGroup.Wait` inside a loop
[tarfile-extractall-traversal](python/tarfile-extractall-traversal.yml) | Python | Potential path traversal in call to `extractall` for a `tarfile`
[unchecked-type-assertion](go/unchecked-type-assertion.yml) | Go | Unchecked type assertion