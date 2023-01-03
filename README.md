# Telnyx CLI

Telnyx CLI is a toolkit with useful commands for operating services in the Telnyx infrastructure. It is built on top of the powerful [Cobra](https://github.com/spf13/cobra)
and [Viper](https://github.com/spf13/viper) libraries, which provide a foundation for building configurable CLIs.

## Getting started

**Install Go in your system**

I suggest using a tool like [asdf](https://asdf-vm.com/) to manage the Go versions in your system. This [plugin](https://github.com/kennyp/asdf-golang) should
get the job done.

**Download the CLI**

You will first need to configure your Github credentials so Go can download the CLI from the Telnyx-owned private repository:

1) Create a Github Personal Access Token following [these instructions](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token).
2) Add the following line `machine github.com login your_github_username password your_github_access_token` to `~/.netrc` (create the file if it doesn't exist).

This will allow you to install the CLI binary directly:

```
GOPRIVATE=github.com/team-telnyx go install github.com/team-telnyx/telnyx-cli@latest
```

**Initialize the CLI**

The CLI depends on some local configuration files to work properly. You can auto-generate them by running `telnyx-cli init`. This will create the folder `$HOME/.telnyx-cli/`
and add the initial configuration files there.

**(Optional) Configure shell autocompletion**

You can generate shell autocompletion scripts by running the command `telnyx-cli completion (SHELL)`. The currently supported shells are:

* Bash
* Zsh
* fish
* PowerShell

Running the command will print the autocompletion script, that should be placed in the correct place expected by your shell. You can see more details about how Cobra handles
shell completions [here](https://github.com/spf13/cobra/blob/main/shell_completions.md).

**Using the CLI**

You can see the available commands (with instructions) by running `telnyx-cli help`.

## Contributing

**Install Go in your system**

Follow the same steps as listed above.

**Build the project**

You need to clone this repository in your machine and build it locally:

1) Clone this repository
2) Run `go build`. This will build the project and create the CLI executable into the local folder
3) You can run the binary via the `./telnyx-cli` command in your shell.

Note: Running `go install` will compile the local project and overwrite the current binary installed in `$GOPATH/bin` or `$HOME/go/bin`. You will need to reinstall the
official release in case you want to revert to it (or revert your local changes and run `go install` again).
