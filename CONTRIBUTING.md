# Contributing to Telnyx CLI

We love your input! We want to make contributing to Telnyx CLI as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## Development Process

We use GitHub to host code, to track issues and feature requests, as well as accept pull requests.

1. Fork the repo and create your branch from `main`
2. If you've added code that should be tested, add tests
3. If you've changed APIs, update the documentation
4. Ensure the test suite passes
5. Make sure your code follows the existing style
6. Issue that pull request!

## Development Setup

1. Install Go 1.21.5 or later (we recommend using `asdf` as described in the README)
2. Configure GitHub credentials for private repository access
3. Clone your fork of the repository
4. Install dependencies: `go mod download`
5. Build the project: `go build`

## Code Style

- Follow standard Go formatting guidelines (`go fmt`)
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions focused and concise
- Use proper error handling
- Follow the project's existing patterns for Cobra commands

## Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally after the first line

## Pull Request Process

1. Update the README.md with details of changes to the interface, if applicable
2. Update the documentation with any new commands or flags
3. The PR must be reviewed and approved by at least one maintainer
4. Ensure all checks pass (tests, linting, etc.)

## Testing

- Write unit tests for new functionality
- Ensure all tests pass before submitting: `go test ./...`
- Include both positive and negative test cases
- Mock external dependencies (e.g., Consul) in tests

## Documentation

When adding or modifying features:

1. Update command help text (`Use`, `Short`, `Long`, `Example`)
2. Update README.md if adding new commands
3. Add godoc comments to exported functions
4. Include example usage in command documentation

## Adding New Commands

When adding new commands:

1. Follow the existing command structure in `cmd/` directory
2. Use Cobra's command template
3. Implement proper flag handling
4. Add appropriate validation
5. Include progress indicators for long-running operations
6. Add user confirmation for destructive operations
7. Implement shell completion where appropriate

## Reporting Bugs

Report bugs using GitHub's [issue tracker](https://github.com/team-telnyx/telnyx-cli/issues)

**Great Bug Reports** tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

## License

By contributing, you agree that your contributions will be licensed under the same license that covers the project.

## Questions?

Don't hesitate to ask questions by creating an issue or contacting the maintainers directly.
