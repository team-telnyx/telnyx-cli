# Telnyx CLI

Telnyx CLI is a toolkit with useful commands for operating services in the Telnyx infrastructure. It is built on top of the powerful [Cobra](https://github.com/spf13/cobra)
and [Viper](https://github.com/spf13/viper) libraries, which provide a foundation for building configurable CLIs.

## Getting started

**Option 1: Install from GitHub Releases**

1. Visit the [Releases](https://github.com/team-telnyx/telnyx-cli/releases) page of the repository
2. Download the appropriate version for your platform:
   - For Mac M1/M2 (Apple Silicon): `telnyx-cli_Darwin_arm64.tar.gz`
   - For Mac Intel: `telnyx-cli_Darwin_x86_64.tar.gz`
   - For Linux: `telnyx-cli_Linux_x86_64.tar.gz`
3. Extract the archive:
   ```bash
   tar -xzf telnyx-cli_<OS>_<ARCH>.tar.gz
   ```
4. Move the binary to a location in your PATH:
   ```bash
   sudo mv telnyx-cli /usr/local/bin/
   ```

**Option 2: Install using Go**

Before installing the CLI, you need to set up Go in your system:

1. Install Go:
   - Using asdf (recommended):
     ```bash
     # Install asdf if you haven't already
     git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.13.1
     
     # Add to your shell (for bash)
     echo '. "$HOME/.asdf/asdf.sh"' >> ~/.bashrc
     echo '. "$HOME/.asdf/completions/asdf.bash"' >> ~/.bashrc
     
     # Install Go plugin
     asdf plugin add golang https://github.com/kennyp/asdf-golang.git
     
     # Install Go 1.21.5
     asdf install golang 1.21.5
     asdf global golang 1.21.5
     ```
   
   - Alternative methods:
     - macOS: `brew install go`
     - Linux: Use your distribution's package manager
     - Windows: Download from [Go's official website](https://golang.org/dl/)

2. Configure GitHub credentials (required for private repository access):
   - Create a GitHub Personal Access Token following [these instructions](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
   - Add the following line to `~/.netrc` (create the file if it doesn't exist):
     ```
     machine github.com login your_github_username password your_github_access_token
     ```

3. Install the CLI:
   ```bash
   GOPRIVATE=github.com/team-telnyx go install github.com/team-telnyx/telnyx-cli@latest
   ```

**Initialize the CLI**

The CLI depends on some local configuration files to work properly. You can auto-generate them by running `telnyx-cli init`. This will create the folder `$HOME/.telnyx-cli/`
and add the initial configuration files there.

**(Optional) Configure GitHub Token for Service Diffs**

The `get service-diff` command compares service versions between environments using GitHub. To use this feature:

1. Create a GitHub Personal Access Token:
   - Go to https://github.com/settings/tokens
   - Click "Generate new token (classic)"
   - Select scope: `repo` (Full control of private repositories)
   - Generate and copy the token

2. Add token to configuration:
   ```bash
   # Edit ~/.telnyx-cli/config.yaml and add:
   github_token: ghp_your_token_here
   ```

3. Test the configuration:
   ```bash
   # Compare prod vs dev versions
   telnyx-cli get service-diff billing

   # Plain format (parsable with awk/cut/grep)
   telnyx-cli get service-diff billing --format plain

   # JSON format (parsable with jq)
   telnyx-cli get service-diff billing --format json
   ```

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
