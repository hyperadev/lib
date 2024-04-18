# Welcome!

Welcome to [`hypera.dev/lib`](https://github.com/HyperaDev/lib), first off, thank you for taking the time to consider
contributing!

All contributions to `hypera.dev/lib` are extremely helpful and greatly appreciated! We are trying our best
to make this project as good as possible, but we're still improving things.

This document contains a set of guides for contributing to this project.

<details>
<summary>Table of Contents</summary>

<!-- TOC -->
* [Welcome!](#welcome)
* [Code of Conduct](#code-of-conduct)
* [Questions](#questions)
* [Contributing](#contributing)
  * [Bug reports](#bug-reports)
    * [Security vulnerabilities](#security-vulnerabilities)
  * [Suggesting features](#suggesting-features)
  * [Code contributions](#code-contributions)
    * [Testing](#testing)
    * [Commit messages](#commit-messages)
      * [Allowed types](#allowed-types)
      * [Allowed scopes](#allowed-scopes)
    * [Code review](#code-review)
  * [Supporting the Authors](#supporting-the-authors)
<!-- TOC -->
</details>

# Code of Conduct

Please help keep this project open and inclusive for all.  
Read and follow the [Code of Conduct](https://github.com/HyperaDev/.github/blob/main/CODE_OF_CONDUCT.md) before
contributing to this repository.

If you have encountered someone who is not following the Code of Conduct, please report them
to [oss@hypera.dev](mailto:oss@hypera.dev).

# Questions

> **Please do not use GitHub issues to ask questions.** You will get a faster response if you ask on Discord!

If you wish to ask a question, please contact us using Discord by joining
the [Hypera Development Discord server](https://discord.hypera.dev/), and you will get a response as soon as
someone is next available.

# Contributing

There are many ways to contribute to `hypera.dev/lib`, and they all help!  
Here are the most common types of contributions:

* [Bug reports](#bug-reports)
  * [Security vulnerabilities](#security-vulnerabilities)
* [Suggesting features](#suggesting-features)
* [Code contributions](#code-contributions)
* [Supporting the authors](#supporting-the-authors)

## Bug reports

If you have discovered a bug in `hypera.dev/lib`, you can help us
by [creating an issue](https://github.com/HyperaDev/lib/issues/new), or if you have the time
and required knowledge, and really want to help this project, you
can [create a Pull Request](https://github.com/HyperaDev/lib/compare) with a fix.

### Security vulnerabilities

We take the security of `hypera.dev/lib` and our users very seriously. As such, we encourage responsible
disclosure of security vulnerabilities in `hypera.dev/lib`.

If you have discovered a security vulnerability in `hypera.dev/lib`, please report it in accordance with
our [Security Policy](SECURITY.md#reporting-a-vulnerability).<br/>
**Never use GitHub issues to report a security vulnerability.**

## Suggesting features

If you have an idea for something that could be added to `hypera.dev/lib`, you can suggest it
by [creating an issue](https://github.com/HyperaDev/lib/issues/new)!  
Before submitting a feature request, please be sure to check that it hasn't already been suggested.

## Code contributions

Code contributions are often the most helpful way to contribute to this project, and all code contributions will be
greatly appreciated!

You can contribute code changes that you have written for `hypera.dev/lib`
by [creating a Pull Request](https://github.com/HyperaDev/lib/compare).

### Testing

Adding test coverage is extremely helpful and highly recommended for any major changes you make.  
Testing helps us catch problems early before they have the change to cause big issues in production.

### Commit messages

Whilst **not required for commits in pull requests**, all commits made in the `main` branch **must**
follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).
This allows for the Git history to be more readable and helps us generate changelogs automatically.

#### Allowed types

- `fix`, when the commit fixes a bug or other issue.
- `feat`, when adding a new feature.
- `refactor`, when refactoring or improving existing code.
- `build`, when modifying a build file.
- `ci`, when modifying a GitHub Actions workflow.
- `docs`, when changing documentation.
- `style`, when correcting a code-style issue.
- `perf`, when improving the performance of a feature.
- `test`, when adding or improving tests.
- `chore`, when doing something that does not fit into the types above.

#### Allowed scopes

- `slog/pretty`, when making changes in the `slog/pretty` package.
- `util/retry`, when making changes in the `util/retry` package.
- `deps`, when adding, updating, or removing dependencies.

When using the type `ci`, workflow names (`.github/workflows/` files, excluding extensions) may be
used as the scope.

### Code review

We have strict guidelines for merging pull requests to maintain code quality, and to prevent future
problems:

1. Pull requests must successfully build, pass all tests, and adhere to style guidelines. Any pull
   requests that fail to do so, will not be merged.
2. **Code contributions must be licensed under the [MIT License](LICENSE)**.
3. Modifications must be reviewed by the [code owners](.github/CODEOWNERS) for the respective files.

We value high code quality, which is why our pull request reviews are rigorous. This approach
ensures that problems or mistakes are caught before being merged into the `main` branch.

If you receive numerous comments pointing out mistakes in your pull request, please don't take
offense. Mistakes are normal, and you should not be ashamed of them.

If you spot issues in someone else's pull request, kindly leave a polite comment to make others
aware of the problem before it gets merged. Your help in maintaining code quality is greatly
appreciated.

## Supporting the Authors

If you wish to support this project in another way, the authors accept donations!
These donations go towards enabling the authors to spend more time working on this project, paying
for infrastructure/domains, etc. All donations are extremely appreciated! :D

- [Joshua (joshuasing)](https://github.com/sponsors/joshuasing)
- [Luis (LooFifteen)](https://ko-fi.com/loofifteen)

Thank you to everyone who has donated or otherwise contributed to `hypera.dev/lib`!
