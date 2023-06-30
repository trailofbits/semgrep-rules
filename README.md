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
$ semgrep --config /path/to/semgrep-rules/ . --exclude='*_test.go'
```

Use `-o` to output results to a file:

```shell
$ semgrep --config /path/to/semgrep-rules/hanging-goroutine.yml -o leaks.txt'
```

## Rules

### go

| ID | Playground | Impact | Confidence | Description |
| -- | :--------: | :----: | :--------: | ----------- |
| [anonymous-race-condition](go/anonymous-race-condition.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.anonymous-race-condition.anonymous-race-condition) | 🟧 | 🌗 | Race conditions within anonymous goroutines |
| [hanging-goroutine](go/hanging-goroutine.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.hanging-goroutine.hanging-goroutine) | 🟩 | 🌗 | Goroutine leaks |
| [invalid-usage-of-modified-variable](go/invalid-usage-of-modified-variable.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.invalid-usage-of-modified-variable.invalid-usage-of-modified-variable) | 🟧 | 🌘 | Possible unintentional assignment when an error occurs |
| [iterate-over-empty-map](go/iterate-over-empty-map.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.iterate-over-empty-map.iterate-over-empty-map) | 🟩 | 🌗 | Probably redundant iteration over an empty map |
| [missing-runlock-on-rwmutex](go/missing-runlock-on-rwmutex.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.missing-runlock-on-rwmutex.missing-runlock-on-rwmutex) | 🟧 | 🌗 | Missing `RUnlock` on an `RWMutex` lock before returning from a function |
| [missing-unlock-before-return](go/missing-unlock-before-return.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.missing-unlock-before-return.missing-unlock-before-return) | 🟧 | 🌗 | Missing `mutex` unlock before returning from a function |
| [nil-check-after-call](go/nil-check-after-call.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.nil-check-after-call.nil-check-after-call) | 🟧 | 🌗 | Possible nil dereferences |
| [racy-append-to-slice](go/racy-append-to-slice.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.racy-append-to-slice.racy-append-to-slice) | 🟧 | 🌗 | Concurrent calls to `append` from multiple goroutines |
| [racy-write-to-map](go/racy-write-to-map.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.racy-write-to-map.racy-write-to-map) | 🟧 | 🌗 | Concurrent writes to the same map in multiple goroutines |
| [servercodec-readrequestbody-unhandled-nil](go/servercodec-readrequestbody-unhandled-nil.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.servercodec-readrequestbody-unhandled-nil.servercodec-readrequestbody-unhandled-nil) | 🟩 | 🌘 | Possible incorrect `ServerCodec` interface implementation |
| [string-to-int-signedness-cast](go/string-to-int-signedness-cast.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.string-to-int-signedness-cast.string-to-int-signedness-cast) | 🟧 | 🌘 | Integer underflows |
| [sync-mutex-value-copied](go/sync-mutex-value-copied.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.sync-mutex-value-copied.sync-mutex-value-copied) | 🟩 | 🌘 | Copying of `sync.Mutex` via value receivers |
| [unsafe-dll-loading](go/unsafe-dll-loading.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.unsafe-dll-loading.unsafe-dll-loading) | 🟥 | 🌘 | Use of function vulnerable to DLL hijacking attacks |
| [waitgroup-add-called-inside-goroutine](go/waitgroup-add-called-inside-goroutine.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.waitgroup-add-called-inside-goroutine.waitgroup-add-called-inside-goroutine) | 🟧 | 🌗 | Calls to `sync.WaitGroup.Add` inside of anonymous goroutines |
| [waitgroup-wait-inside-loop](go/waitgroup-wait-inside-loop.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.waitgroup-wait-inside-loop.waitgroup-wait-inside-loop) | 🟧 | 🌗 | Calls to `sync.WaitGroup.Wait` inside a loop |


### python

| ID | Playground | Impact | Confidence | Description |
| -- | :--------: | :----: | :--------: | ----------- |
| [automatic-memory-pinning](python/automatic-memory-pinning.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.automatic-memory-pinning.automatic-memory-pinning) | 🟩 | 🌘 | `PyTorch` memory not automatically pinned |
| [lxml-in-pandas](python/lxml-in-pandas.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.lxml-in-pandas.lxml-in-pandas) | 🟧 | 🌘 | Potential XXE attacks from loading `lxml` in pandas |
| [numpy-distutils](python/numpy-distutils.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.numpy-distutils.numpy-distutils) | 🟩 | 🌘 | Use of deprecated `numpy.distutils` |
| [numpy-f2py-compile](python/numpy-f2py-compile.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.numpy-f2py-compile.numpy-f2py-compile) | 🟥 | 🌗 | Potential arbitrary code execution from `NumPy` `f2py` compilation |
| [numpy-in-pytorch-datasets](python/numpy-in-pytorch-datasets.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.numpy-in-pytorch-datasets.numpy-in-pytorch-datasets) | 🟩 | 🌘 | Calls to the `NumPy` RNG inside of a `Torch` dataset |
| [numpy-in-pytorch-modules](python/numpy-in-pytorch-modules.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.numpy-in-pytorch-modules.numpy-in-pytorch-modules) | 🌫️ | 🌗 | Uses of `NumPy` functions inside `PyTorch` modules |
| [numpy-load-library](python/numpy-load-library.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.numpy-load-library.numpy-load-library) | 🟥 | 🌗 | Potential arbitrary code execution from `NumPy` library loading |
| [onnx-session-options](python/onnx-session-options.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.onnx-session-options.onnx-session-options) | 🟥 | 🌗 | Potential arbitrary code execution from `ONNX` library loading |
| [pickles-in-numpy](python/pickles-in-numpy.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pickles-in-numpy.pickles-in-numpy) | 🟥 | 🌗 | Potential arbitrary code execution from `NumPy` functions reliant on pickling |
| [pickles-in-pandas](python/pickles-in-pandas.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pickles-in-pandas.pickles-in-pandas) | 🟥 | 🌗 | Potential arbitrary code execution from `Pandas` functions reliant on pickling |
| [pickles-in-pytorch-distributed](python/pickles-in-pytorch-distributed.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pickles-in-pytorch-distributed.pickles-in-pytorch-distributed) | 🟥 | 🌗 | Potential arbitrary code execution from `PyTorch.Distributed` functions reliant on pickling |
| [pickles-in-pytorch](python/pickles-in-pytorch.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pickles-in-pytorch.pickles-in-pytorch) | 🟥 | 🌗 | Potential arbitrary code execution from `PyTorch` functions reliant on pickling |
| [pytorch-classes-load-library](python/pytorch-classes-load-library.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pytorch-classes-load-library.pytorch-classes-load-library) | 🟥 | 🌗 | Potential arbitrary code execution from `PyTorch` library loading |
| [pytorch-package](python/pytorch-package.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pytorch-package.pytorch-package) | 🟥 | 🌕 | Potential arbitrary code execution from `torch.package` |
| [pytorch-tensor](python/pytorch-tensor.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pytorch-tensor.pytorch-tensor) | 🌫️ | 🌘 | Possible parsing issues and inefficiency from improper tensor creation |
| [scikit-joblib-load](python/scikit-joblib-load.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.scikit-joblib-load.scikit-joblib-load) | 🟥 | 🌗 | Potential arbitrary code execution from `SciKit.Joblib` functions reliant on pickling |
| [tarfile-extractall-traversal](python/tarfile-extractall-traversal.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.tarfile-extractall-traversal.tarfile-extractall-traversal) | 🟧 | 🌗 | Potential path traversal in call to `extractall` for a `tarfile` |
| [tensorflow-load-library](python/tensorflow-load-library.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.tensorflow-load-library.tensorflow-load-library) | 🟥 | 🌗 | Potential arbitrary code execution from `TensorFlow` library loading |
| [waiting-with-pytorch-distributed](python/waiting-with-pytorch-distributed.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.waiting-with-pytorch-distributed.waiting-with-pytorch-distributed) | 🟩 | 🌗 | Possible `PyTorch` undefined behavior when not waiting for requests |


### rs

| ID | Playground | Impact | Confidence | Description |
| -- | :--------: | :----: | :--------: | ----------- |
| [panic-in-function-returning-result](rs/panic-in-function-returning-result.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.rs.panic-in-function-returning-result.panic-in-function-returning-result) | 🟩 | 🌘 | Calling `unwrap` or `expect` in a function returning a `Result` |


### javascript

| ID | Playground | Impact | Confidence | Description |
| -- | :--------: | :----: | :--------: | ----------- |
| [schema-directives](javascript/apollo-graphql/schema-directives.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.javascript.apollo-graphql.schema-directives.schema-directives) | 🟥 | 🌗 | Use of outdated ApolloServer option 'schemaDirectives' |
| [use-of-graphql-upload](javascript/apollo-graphql/use-of-graphql-upload.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.javascript.apollo-graphql.use-of-graphql-upload.use-of-graphql-upload) | 🟧 | 🌕 | Use of the graphql-upload library |
| [v3-potentially-bad-cors](javascript/apollo-graphql/v3-cors-audit.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.javascript.apollo-graphql.v3-cors-audit.v3-potentially-bad-cors) | 🟧 | 🌕 | Potentially bad CORS policy |
| [v3-express-bad-cors](javascript/apollo-graphql/v3-cors-express.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.javascript.apollo-graphql.v3-cors-express.v3-express-bad-cors) | 🟥 | 🌗 | Bad CORS policy |
| [v3-express-no-cors](javascript/apollo-graphql/v3-cors-express.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.javascript.apollo-graphql.v3-cors-express.v3-express-no-cors) | 🟩 | 🌘 | Lack of CORS policy |
| [v3-bad-cors](javascript/apollo-graphql/v3-cors.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.javascript.apollo-graphql.v3-cors.v3-bad-cors) | 🟥 | 🌗 | Bad CORS policy |
| [v3-no-cors](javascript/apollo-graphql/v3-cors.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.javascript.apollo-graphql.v3-cors.v3-no-cors) | 🟩 | 🌘 | Lack of CORS policy |
| [v3-csrf-prevention](javascript/apollo-graphql/v3-csrf-prevention.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.javascript.apollo-graphql.v3-csrf-prevention.v3-csrf-prevention) | 🟧 | 🌘 | Lack of CSRF prevention |
| [v4-csrf-prevention](javascript/apollo-graphql/v4-csrf-prevention.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.javascript.apollo-graphql.v4-csrf-prevention.v4-csrf-prevention) | 🟧 | 🌘 | CSRF protection disabled |


## Contributing

Pull Requests and issues are welcomed!

See [CONTRIBUTING.md](CONTRIBUTING.md) for more information.

## Licensing

The rules defined in this repository are licensed under AGPLv3.

The sidecar examples *may* be derived from other works, and retain their
original licenses where required.
