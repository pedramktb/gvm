# gvm (go version manager)

## Installation
Note: This package utilizes `~/.bachrc` for version management. Windows is subsequently not supported.

To install gvm run `go install github.com/pedramktb/gvm@latest`.

## Requirement
You need to have `$GOROOT/bin` in your `$PATH` in order to use `gvm` in all sessions. You can do this by adding `PATH=$PATH:$GOROOT/bin` or `PATH=$PATH:$HOME/go/bin` (if you don't use `$GOROOT`) in your `~/.bashrc`.

## Usage
You can install a custom go version with `go install golang.org/dl/{Version}@latest` followed by `{Version} download`.
Now you can use `gvm {Version}` to set it as the default go runtime for future terminal sessions.

Example: `gvm go1.19.13`

In order to use the default go installation again, run `gvm go`.
