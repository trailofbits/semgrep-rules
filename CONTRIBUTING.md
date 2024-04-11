Contributing to Trail of Bits Semgrep Rules
=========================

Thank you for your interest in contributing to ToB `semgrep-rules`!

The information below will help you set up a local development environment,
as well as performing common development tasks.

## Requirements

`semgrep-rules`'s only development environment requirement *should* be Python 3.7
or newer. Development and testing is actively performed on macOS and Linux,
but Windows and other supported platforms that are supported by Python
should also work.

## Development steps

First, clone this repository:

```bash
git clone https://github.com/trailofbits/semgrep-rules
cd semgrep-rules
```

Then [install semgrep CLI](https://semgrep.dev/docs/getting-started/), and you are good to start development.

### Linting

First, [install `prettier`](https://prettier.io/docs/en/install), or [use `brew`](https://formulae.brew.sh/formula/prettier) to do so.

Use the following command to check rule files for formatting errors:

```bash
prettier --check '**/*.{yaml,yml}'
```

Any issues can be automatically fixed with the following command:

```bash
prettier --write '**/*.{yaml,yml}'
```

### Testing

You can run tests locally with:

```bash
semgrep --test --test-ignore-todo --metrics=off
```

To test a specific file:

```bash
semgrep --test --test-ignore-todo --metrics=off --config ./go/iterate-over-empty-map.yaml ./go/iterate-over-empty-map.go
```

### Development practices

Before publishing a new rule, or updating an existing one, make sure to review the checklist below:

- [ ] Check if the rule does not already exists. Review this repository and [Semgrep registry](https://semgrep.dev/r). If there already is a rule that finds the vulnerability your new rule is targeting, consider making updates to this rule instead of creating a new one.

- [ ] Add metadata. Semgrep [defines which metadata fields are required](https://semgrep.dev/docs/contributing/contributing-to-semgrep-rules-repository/#writing-a-rule-for-semgrep-registry)
    - [ ] Add a non-standard `metadata.description` field. It will be used as a description in the `semgrep-rules` README table.
    - For `metadata.references` provide a link to official documentation, Trail of Bits blogpost, GitHub issue, or some reputable website. Avoid linking to websites that may disappear in the future.

- [ ] Validate metadata against the official schema
    - Download python validation script `wget https://raw.githubusercontent.com/returntocorp/semgrep-rules/develop/.github/scripts/validate-metadata.py`
    - Download rules schema `wget https://raw.githubusercontent.com/returntocorp/semgrep-rules/develop/metadata-schema.yaml.schm`
    - Run `python ./validate-metadata.py -s ./metadata-schema.yaml.schm -f .`

- [ ] Add tests
  - [ ] At least one true positive (`ruleid: ` comment)
  - [ ] At least one true negative (`ok: ` comment)
  - Tests are allowed to crash when running them directly or to be meaningless
  - However, try writing tests that can be compiled or parsed by the language interpreter
  - The first few test cases should be easy to understand, the later should be more complex or check for edge-cases
  - [ ] Make sure all tests pass, run `semgrep --test --test-ignore-todo --metrics=off`

- [ ] Run official semgrep lints with `semgrep --validate --metrics=off --config ./<new-rule>.yaml`

- [ ] Review style of the rules
    - [ ] Use 2 spaces for indentation
    - [ ] Use `>-` for multiline messages
    - [ ] Use backticks in messages e.g., `$VAR`, `$FUNC`, `some.method()`
    - The `languages` field in `[go, java]` format are preferable (not `- go \n -java`)
    - [ ] Run prettier (see [Linting](#linting))

- [ ] Check amount of false-positives on some large public repositories

- [ ] Check performance - take a look at [r2c methodology](https://github.com/returntocorp/semgrep-rules/blob/main/tests/performance/test_public_repos.py)

- [ ] Add the new rules to the README
    - Run `python ./rules_table_generator.py` to re-generate the table
    - Manually check if the table was correctly generated

### Documentation

We don't provide any documentation for the rules. All information that you need to understand a rule is inside it. Semgrep documentation can be found [here](https://semgrep.dev/docs/).

### Releasing

**NOTE**: If you're a non-maintaining contributor, you don't need the steps
here! They're documented for completeness and for onboarding future maintainers.

We don't have a release cycle yet.

All changes to the repository's `main` branch are automatically pushed to the semgrep registry (with a GitHub action).

Modifying rule's filename, path, or ID will result in duplication of the rule in the registry.
This is a known issue, r2c team still works on resolving it.
