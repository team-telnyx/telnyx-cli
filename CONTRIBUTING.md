# Contributing to Telnyx CLI

Thanks for your interest in contributing! This document outlines the process for contributing to this project.

## Development Setup

1. **Fork and clone the repository**
   ```bash
   git clone https://github.com/YOUR_USERNAME/telnyx-cli.git
   cd telnyx-cli
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Run tests**
   ```bash
   npm test
   ```

4. **Run the CLI locally**
   ```bash
   node bin/telnyx --help
   ```

## Making Changes

1. Create a new branch for your feature or fix:
   ```bash
   git checkout -b feature/my-new-feature
   ```

2. Make your changes and add tests if applicable

3. Ensure tests pass:
   ```bash
   npm test
   ```

4. Commit your changes with a descriptive message:
   ```bash
   git commit -m "feat: add support for XYZ"
   ```

## Commit Message Guidelines

We follow [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` — New feature
- `fix:` — Bug fix
- `docs:` — Documentation changes
- `test:` — Adding or updating tests
- `refactor:` — Code changes that neither fix bugs nor add features
- `chore:` — Maintenance tasks

## Pull Request Process

1. Push your branch to your fork
2. Open a Pull Request against `main`
3. Ensure CI checks pass
4. Request review from maintainers

## Code Style

- Use meaningful variable and function names
- Add JSDoc comments for public functions
- Keep functions focused and small
- Handle errors gracefully with user-friendly messages

## Adding New Commands

1. Create a new file in `src/commands/`
2. Follow the pattern of existing commands
3. Add any new API methods to `src/api/client.js`
4. Update the main CLI entry point in `bin/telnyx`
5. Add tests in `src/__tests__/`
6. Update README.md with usage examples

## Reporting Issues

- Check existing issues first
- Include steps to reproduce
- Include your Node.js version and OS
- Include relevant error messages

## Questions?

Open an issue or reach out to the Telnyx team.

Thank you for contributing! 🚀
