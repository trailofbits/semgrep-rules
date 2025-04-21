# Trail of Bits public Semgrep rules

This repository contains Semgrep rules developed by Trail of Bits and made available to the public. They are part of our ongoing development efforts and are used in our security audits, vulnerability reseach, and internal projects. They will evolve over time as we identify new techniques.

Visit [Testing Handbook](https://appsec.guide/docs/static-analysis/semgrep/) for Semgrep guidance.

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
| [eth-rpc-tracetransaction](go/eth-rpc-tracetransaction.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.eth-rpc-tracetransaction.eth-rpc-tracetransaction) | 🟥 | 🌕 | Detects attempts to extract trace information from an EVM transaction or block. In exchange or bridge applications, extra logic must be implemented encapsulating these endpoints to prevent the values transferred during reverted call frames from being counted. |
| [eth-txreceipt-status](go/eth-txreceipt-status.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.eth-txreceipt-status.eth-txreceipt-status) | 🟥 | 🌕 | Detects when a transaction receipt's status is read |
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
| [unmarshal-tag-is-dash](go/unmarshal_tag_is_dash.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.unmarshal_tag_is_dash.unmarshal-tag-is-dash) | 🟧 | 🌘 |  |
| [unmarshal-tag-is-omitempty](go/unmarshal_tag_is_omitempty.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.unmarshal_tag_is_omitempty.unmarshal-tag-is-omitempty) | 🟩 | 🌘 |  |
| [unsafe-dll-loading](go/unsafe-dll-loading.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.unsafe-dll-loading.unsafe-dll-loading) | 🟥 | 🌘 | Use of function vulnerable to DLL hijacking attacks |
| [waitgroup-add-called-inside-goroutine](go/waitgroup-add-called-inside-goroutine.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.waitgroup-add-called-inside-goroutine.waitgroup-add-called-inside-goroutine) | 🟧 | 🌗 | Calls to `sync.WaitGroup.Add` inside of anonymous goroutines |
| [waitgroup-wait-inside-loop](go/waitgroup-wait-inside-loop.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.go.waitgroup-wait-inside-loop.waitgroup-wait-inside-loop) | 🟧 | 🌗 | Calls to `sync.WaitGroup.Wait` inside a loop |


### python

| ID | Playground | Impact | Confidence | Description |
| -- | :--------: | :----: | :--------: | ----------- |
| [automatic-memory-pinning](python/automatic-memory-pinning.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.automatic-memory-pinning.automatic-memory-pinning) | 🟩 | 🌘 | `PyTorch` memory not automatically pinned |
| [lxml-in-pandas](python/lxml-in-pandas.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.lxml-in-pandas.lxml-in-pandas) | 🟧 | 🌘 | Potential XXE attacks from loading `lxml` in pandas |
| [msgpack-numpy](python/msgpack-numpy.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.msgpack-numpy.msgpack-numpy) | 🟥 | 🌗 | Potential arbitrary code execution from functions reliant on pickling |
| [numpy-distutils](python/numpy-distutils.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.numpy-distutils.numpy-distutils) | 🟩 | 🌘 | Use of deprecated `numpy.distutils` |
| [numpy-f2py-compile](python/numpy-f2py-compile.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.numpy-f2py-compile.numpy-f2py-compile) | 🟥 | 🌗 | Potential arbitrary code execution from `NumPy` `f2py` compilation |
| [numpy-in-pytorch-datasets](python/numpy-in-pytorch-datasets.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.numpy-in-pytorch-datasets.numpy-in-pytorch-datasets) | 🟩 | 🌘 | Calls to the `NumPy` RNG inside of a `Torch` dataset |
| [numpy-in-pytorch-modules](python/numpy-in-pytorch-modules.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.numpy-in-pytorch-modules.numpy-in-pytorch-modules) | 🌫️ | 🌗 | Uses of `NumPy` functions inside `PyTorch` modules |
| [numpy-load-library](python/numpy-load-library.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.numpy-load-library.numpy-load-library) | 🟥 | 🌗 | Potential arbitrary code execution from `NumPy` library loading |
| [onnx-session-options](python/onnx-session-options.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.onnx-session-options.onnx-session-options) | 🟥 | 🌗 | Potential arbitrary code execution from `ONNX` library loading |
| [pandas-eval](python/pandas-eval.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pandas-eval.pandas-eval) | 🟥 | 🌕 | Potential arbitrary code execution from `pandas` functions that evaluate user-provided expressions |
| [pickles-in-keras-deprecation](python/pickles-in-keras-deprecation.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pickles-in-keras-deprecation.pickles-in-keras-deprecation) | 🟥 | 🌗 | Potential arbitrary code execution from Keras' load_model function |
| [pickles-in-keras](python/pickles-in-keras.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pickles-in-keras.pickles-in-keras) | 🟥 | 🌗 | Potential arbitrary code execution from Keras' load_model function |
| [pickles-in-numpy](python/pickles-in-numpy.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pickles-in-numpy.pickles-in-numpy) | 🟥 | 🌗 | Potential arbitrary code execution from `NumPy` functions reliant on pickling |
| [pickles-in-pandas](python/pickles-in-pandas.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pickles-in-pandas.pickles-in-pandas) | 🟥 | 🌗 | Potential arbitrary code execution from `Pandas` functions reliant on pickling |
| [pickles-in-pytorch-distributed](python/pickles-in-pytorch-distributed.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pickles-in-pytorch-distributed.pickles-in-pytorch-distributed) | 🟥 | 🌗 | Potential arbitrary code execution from `PyTorch.Distributed` functions reliant on pickling |
| [pickles-in-pytorch](python/pickles-in-pytorch.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pickles-in-pytorch.pickles-in-pytorch) | 🟥 | 🌗 | Potential arbitrary code execution from `PyTorch` functions reliant on pickling |
| [pickles-in-tensorflow](python/pickles-in-tensorflow.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.python.pickles-in-tensorflow.pickles-in-tensorflow) | 🟥 | 🌗 | Potential arbitrary code execution from tensorflow's load function |
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


### ruby

| ID | Playground | Impact | Confidence | Description |
| -- | :--------: | :----: | :--------: | ----------- |
| [action-dispatch-insecure-ssl](ruby/action-dispatch-insecure-ssl.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.action-dispatch-insecure-ssl.action-dispatch-insecure-ssl) | 🟥 | 🌘 |  |
| [action-mailer-insecure-tls](ruby/action-mailer-insecure-tls.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.action-mailer-insecure-tls.action-mailer-insecure-tls) | 🟥 | 🌘 |  |
| [active-record-encrypts-misorder](ruby/active-record-encrypts-misorder.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.active-record-encrypts-misorder.active-record-encrypts-misorder) | 🟥 | 🌘 |  |
| [active-record-hardcoded-encryption-key](ruby/active-record-hardcoded-encryption-key.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.active-record-hardcoded-encryption-key.active-record-hardcoded-encryption-key) | 🟥 | 🌘 |  |
| [faraday-disable-verification](ruby/faraday-disable-verification.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.faraday-disable-verification.faraday-disable-verification) | 🟥 | 🌘 |  |
| [global-timeout](ruby/global-timeout.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.global-timeout.global-timeout) | 🟩 | 🌘 |  |
| [insecure-rails-cookie-session-store](ruby/insecure-rails-cookie-session-store.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.insecure-rails-cookie-session-store.insecure-rails-cookie-session-store) | 🟩 | 🌘 |  |
| [json-create-deserialization](ruby/json-create-deserialization.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.json-create-deserialization.json-create-deserialization) | 🟥 | 🌕 |  |
| [rails-cache-store-marshal](ruby/rails-cache-store-marshal.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.rails-cache-store-marshal.rails-cache-store-marshal) | 🟩 | 🌗 |  |
| [rails-cookie-attributes](ruby/rails-cookie-attributes.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.rails-cookie-attributes.rails-cookie-attributes) | 🟩 | 🌘 |  |
| [rails-params-json](ruby/rails-params-json.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.rails-params-json.rails-params-json) | 🟥 | 🌕 |  |
| [rest-client-disable-verification](ruby/rest-client-disable-verification.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.rest-client-disable-verification.rest-client-disable-verification) | 🟥 | 🌘 |  |
| [ruby-saml-skip-validation](ruby/ruby-saml-skip-validation.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.ruby-saml-skip-validation.ruby-saml-skip-validation) | 🟧 | 🌘 |  |
| [yaml-unsafe-load](ruby/yaml-unsafe-load.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.ruby.yaml-unsafe-load.yaml-unsafe-load) | 🟥 | 🌘 |  |


### hcl

| ID | Playground | Impact | Confidence | Description |
| -- | :--------: | :----: | :--------: | ----------- |
| [docker-hardcoded-password](hcl/nomad/docker-hardcoded-password.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.hcl.nomad.docker-hardcoded-password.docker-hardcoded-password) | 🟥 | 🌘 |  |
| [docker-privileged-mode](hcl/nomad/docker-privileged-mode.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.hcl.nomad.docker-privileged-mode.docker-privileged-mode) | 🟩 | 🌘 |  |
| [podman-tls-verify-disabled](hcl/nomad/podman-tls-verify-disabled.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.hcl.nomad.podman-tls-verify-disabled.podman-tls-verify-disabled) | 🟩 | 🌘 |  |
| [root-user](hcl/nomad/root-user.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.hcl.nomad.root-user.root-user) | 🟩 | 🌘 |  |
| [tls-hostname-verification-disabled](hcl/nomad/tls-hostname-verification-disabled.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.hcl.nomad.tls-hostname-verification-disabled.tls-hostname-verification-disabled) | 🟥 | 🌘 |  |
| [aws-oidc-role-policy-duplicate-condition](hcl/terraform/aws-oidc-role-policy-duplicate-condition.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.hcl.terraform.aws-oidc-role-policy-duplicate-condition.aws-oidc-role-policy-duplicate-condition) | 🟥 | 🌘 |  |
| [aws-oidc-role-policy-missing-sub](hcl/terraform/aws-oidc-role-policy-missing-sub.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.hcl.terraform.aws-oidc-role-policy-missing-sub.aws-oidc-role-policy-missing-sub) | 🟥 | 🌘 |  |
| [vault-hardcoded-token](hcl/terraform/vault-hardcoded-token.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.hcl.terraform.vault-hardcoded-token.vault-hardcoded-token) | 🟥 | 🌘 |  |
| [vault-skip-tls-verify](hcl/terraform/vault-skip-tls-verify.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.hcl.terraform.vault-skip-tls-verify.vault-skip-tls-verify) | 🟥 | 🌘 |  |


### jvm

| ID | Playground | Impact | Confidence | Description |
| -- | :--------: | :----: | :--------: | ----------- |
| [gc-call](jvm/gc-call.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.jvm.gc-call.gc-call) | 🟩 | 🌘 |  |
| [mongo-hostname-verification-disabled](jvm/mongo-hostname-verification-disabled.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.jvm.mongo-hostname-verification-disabled.mongo-hostname-verification-disabled) | 🟥 | 🌘 |  |


### yaml

| ID | Playground | Impact | Confidence | Description |
| -- | :--------: | :----: | :--------: | ----------- |
| [apt-key-unencrypted-url](yaml/ansible/apt-key-unencrypted-url.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.apt-key-unencrypted-url.apt-key-unencrypted-url) | 🟥 | 🌘 |  |
| [apt-key-validate-certs-disabled](yaml/ansible/apt-key-validate-certs-disabled.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.apt-key-validate-certs-disabled.apt-key-validate-certs-disabled) | 🟥 | 🌘 |  |
| [apt-unencrypted-url](yaml/ansible/apt-unencrypted-url.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.apt-unencrypted-url.apt-unencrypted-url) | 🟥 | 🌘 |  |
| [dnf-unencrypted-url](yaml/ansible/dnf-unencrypted-url.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.dnf-unencrypted-url.dnf-unencrypted-url) | 🟥 | 🌘 |  |
| [dnf-validate-certs-disabled](yaml/ansible/dnf-validate-certs-disabled.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.dnf-validate-certs-disabled.dnf-validate-certs-disabled) | 🟥 | 🌘 |  |
| [get-url-unencrypted-url](yaml/ansible/get-url-unencrypted-url.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.get-url-unencrypted-url.get-url-unencrypted-url) | 🟥 | 🌘 |  |
| [get-url-validate-certs-disabled](yaml/ansible/get-url-validate-certs-disabled.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.get-url-validate-certs-disabled.get-url-validate-certs-disabled) | 🟥 | 🌘 |  |
| [rpm-key-unencrypted-url](yaml/ansible/rpm-key-unencrypted-url.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.rpm-key-unencrypted-url.rpm-key-unencrypted-url) | 🟥 | 🌘 |  |
| [rpm-key-validate-certs-disabled](yaml/ansible/rpm-key-validate-certs-disabled.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.rpm-key-validate-certs-disabled.rpm-key-validate-certs-disabled) | 🟥 | 🌘 |  |
| [unarchive-unencrypted-url](yaml/ansible/unarchive-unencrypted-url.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.unarchive-unencrypted-url.unarchive-unencrypted-url) | 🟥 | 🌘 |  |
| [unarchive-validate-certs-disabled](yaml/ansible/unarchive-validate-certs-disabled.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.unarchive-validate-certs-disabled.unarchive-validate-certs-disabled) | 🟥 | 🌘 |  |
| [wrm-cert-validation-ignore](yaml/ansible/wrm-cert-validation-ignore.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.wrm-cert-validation-ignore.wrm-cert-validation-ignore) | 🟥 | 🌘 |  |
| [yum-unencrypted-url](yaml/ansible/yum-unencrypted-url.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.yum-unencrypted-url.yum-unencrypted-url) | 🟥 | 🌘 |  |
| [yum-validate-certs-disabled](yaml/ansible/yum-validate-certs-disabled.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.yum-validate-certs-disabled.yum-validate-certs-disabled) | 🟥 | 🌘 |  |
| [zypper-repository-unencrypted-url](yaml/ansible/zypper-repository-unencrypted-url.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.zypper-repository-unencrypted-url.zypper-repository-unencrypted-url) | 🟥 | 🌘 |  |
| [zypper-unencrypted-url](yaml/ansible/zypper-unencrypted-url.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.ansible.zypper-unencrypted-url.zypper-unencrypted-url) | 🟥 | 🌘 |  |
| [port-all-interfaces](yaml/docker-compose/port-all-interfaces.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.docker-compose.port-all-interfaces.port-all-interfaces) | 🟩 | 🌕 |  |
| [aws-secret-key](yaml/github-actions/aws-secret-key.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.github-actions.aws-secret-key.aws-secret-key) | 🟧 | 🌘 |  |
| [azure-principal-secret](yaml/github-actions/azure-principal-secret.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.github-actions.azure-principal-secret.azure-principal-secret) | 🟧 | 🌘 |  |
| [gcp-credentials-json](yaml/github-actions/gcp-credentials-json.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.github-actions.gcp-credentials-json.gcp-credentials-json) | 🟧 | 🌘 |  |
| [jfrog-hardcoded-credential](yaml/github-actions/jfrog-hardcoded-credential.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.github-actions.jfrog-hardcoded-credential.jfrog-hardcoded-credential) | 🟧 | 🌘 |  |
| [pypi-publish-password](yaml/github-actions/pypi-publish-password.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.github-actions.pypi-publish-password.pypi-publish-password) | 🟧 | 🌘 |  |
| [rubygems-publish-key](yaml/github-actions/rubygems-publish-key.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.github-actions.rubygems-publish-key.rubygems-publish-key) | 🟧 | 🌘 |  |
| [vault-token](yaml/github-actions/vault-token.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.yaml.github-actions.vault-token.vault-token) | 🟧 | 🌘 |  |


### generic

| ID | Playground | Impact | Confidence | Description |
| -- | :--------: | :----: | :--------: | ----------- |
| [amqp-unencrypted-transport](generic/amqp-unencrypted-transport.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.amqp-unencrypted-transport.amqp-unencrypted-transport) | 🟥 | 🌘 |  |
| [container-privileged](generic/container-privileged.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.container-privileged.container-privileged) | 🟥 | 🌗 |  |
| [container-user-root](generic/container-user-root.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.container-user-root.container-user-root) | 🟥 | 🌗 |  |
| [curl-insecure](generic/curl-insecure.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.curl-insecure.curl-insecure) | 🟥 | 🌗 |  |
| [curl-unencrypted-url](generic/curl-unencrypted-url.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.curl-unencrypted-url.curl-unencrypted-url) | 🟥 | 🌗 |  |
| [gpg-insecure-flags](generic/gpg-insecure-flags.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.gpg-insecure-flags.gpg-insecure-flags) | 🟥 | 🌗 |  |
| [installer-allow-untrusted](generic/installer-allow-untrusted.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.installer-allow-untrusted.installer-allow-untrusted) | 🟥 | 🌘 |  |
| [mongodb-insecure-transport](generic/mongodb-insecure-transport.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.mongodb-insecure-transport.mongodb-insecure-transport) | 🟥 | 🌘 |  |
| [mysql-insecure-sslmode](generic/mysql-insecure-sslmode.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.mysql-insecure-sslmode.mysql-insecure-sslmode) | 🟥 | 🌗 |  |
| [node-disable-certificate-validation](generic/node-disable-certificate-validation.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.node-disable-certificate-validation.node-disable-certificate-validation) | 🟥 | 🌘 |  |
| [openssl-insecure-flags](generic/openssl-insecure-flags.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.openssl-insecure-flags.openssl-insecure-flags) | 🟥 | 🌗 |  |
| [postgres-insecure-sslmode](generic/postgres-insecure-sslmode.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.postgres-insecure-sslmode.postgres-insecure-sslmode) | 🟥 | 🌘 |  |
| [redis-unencrypted-transport](generic/redis-unencrypted-transport.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.redis-unencrypted-transport.redis-unencrypted-transport) | 🟥 | 🌘 |  |
| [ssh-disable-host-key-checking](generic/ssh-disable-host-key-checking.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.ssh-disable-host-key-checking.ssh-disable-host-key-checking) | 🟥 | 🌗 |  |
| [tar-insecure-flags](generic/tar-insecure-flags.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.tar-insecure-flags.tar-insecure-flags) | 🟥 | 🌗 |  |
| [wget-no-check-certificate](generic/wget-no-check-certificate.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.wget-no-check-certificate.wget-no-check-certificate) | 🟥 | 🌗 |  |
| [wget-unencrypted-url](generic/wget-unencrypted-url.yaml) | [🛝🔗](https://semgrep.dev/playground/r/trailofbits.generic.wget-unencrypted-url.wget-unencrypted-url) | 🟥 | 🌗 |  |


## Contributing

Pull Requests and issues are welcomed!

See [CONTRIBUTING.md](CONTRIBUTING.md) for more information.

## Licensing

The rules defined in this repository are licensed under AGPLv3.

The sidecar examples *may* be derived from other works, and retain their
original licenses where required.
