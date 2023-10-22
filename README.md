# `bashpack` - A `bash` Package Manager

`bashpack` is a tool for managing `bash` code for other `bash` programs.

(Note: Presently, `bashpack` can only use `git` sources as its package
identifiers)

## Uses

### Give `bash` a modern import system

You can use `bashpack` as a library manager for imports/`source`s:

```bash
#!/usr/bin/env bash

source "$(bashpack mainpath '<uri_to_pkg>')"
source "$(bashpack mainpath '<uri_to_other_pkg>')"

<... rest of your script ...>
```

For example, to use `ezlog`'s log functions:

```bash
# Load the ezlog library
source "$(bashpack mainpath 'https://github.com/opensourcecorp/ezlog')"

# Now you can use the logging functions from that package, like `log-info`
log-info 'Starting backup job...'
```

You likely notice that `bashpack` prefers to be very explicit in its URI
specification. This is by design and ***not*** a laziness hack -- you should
always know exactly where your dependencies come from `:)`

### Script runner

If you want to run a script from a package instead of use it as a library, then
you can use the `run` command:

```bash
bashpack run '<uri_to_pkg>'
```

This is useful in myriad ways. Say you have something like a generic database
backup/migration executable script in your org. `bashpack` lets you share it
across projects and teams!

## Making your code tree into a `bashpack` package

Getting `bashpack` to recognize your code tree as a package is very simple --
you just need a manifest file.

### Manifest file

A `bashpack` manifest file is named `manifest.bashpack`, found at the root of
the package tree. It is `ini`-formatted, with key-value pairs separated by
equals signs (`=`).

Currently, the only supported key is named `main`, and its value is the relative
path to your package's `main.sh` or equivalent. For example, the `ezlog`
package's `manifest.bashpack` looks like this:

```ini
main = src/main.sh
```

Thiss relative path is used by `bashpack` to know what your package's entrypoint
is.

## Installation

The easiest way to install the latest version of `bashpack` is to run the
following two lines:

```bash
git clone 'https://github.com/opensourcecorp/bashpack' "${HOME}/.local/share/bashpack"
ln -fs "${HOME}/.local/share/bashpack/src/main.sh" "${HOME}/.local/bin/bashpack"
```

This assumes that `${HOME}/.local/bin` is on your `$PATH`. If it is not, change
the destination on the first line to be something that is.

## Developing

`bashpack` requires `bash` (duh), `git`, and [the `bats` testing
framework](https://github.com/bats-core/bats-core).

Common development tasks are driven by the included `Makefile`.

`make test` runs tests.
