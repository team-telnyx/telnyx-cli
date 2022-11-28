# Telnyx CLI

Telnyx CLI is a toolkit with useful commands for operating services in the Telnyx infrastructure. It is built on top of the powerful [Cobra](https://github.com/spf13/cobra)
and [Viper](https://github.com/spf13/viper) libraries, which provide a foundation for building configurable CLIs.

## Installation

**Install Go in your system**

I suggest using a tool like [asdf](https://asdf-vm.com/) to manage the Go versions in your system. This [plugin](https://github.com/kennyp/asdf-golang) should
get the job done.

**Build the project**

For now, there is no precompiled binary for the CLI yet. So you need to clone this repository in your machine and build it locally:

1) Clone this repository
2) Run `go install`. This will build the project and copy the CLI executable into your `$GOPATH/bin` or `$HOME/go/bin` folders, depending on how Go is configured.
3) You should now have access to the `telnyx-cli` command in your shell. If not, restart the shell.

**Initialize the CLI**

The CLI depends on some local configuration files to work properly. You can auto-generate them by running `telnyx-cli init`. This will create the folder `$HOME/.telnyx-cli/`
and add the initial configuration files there.

**(Optional) Configure shell autocompletion**

You can generate shell autocompletion scripts by running the command `telnyx-cli completion (SHELL)`. This will print the autocompletion script, that should be placed in
the correct place expected by your shell.


## Using the CLI

You can see the available commands by running `telnyx-cli help`.
