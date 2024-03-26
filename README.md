# Gitorchk

Gitorchk is a command-line tool written in Go to help you stay up-to-date with changes in your Git repository. It checks if your local main branch is behind the remote main branch and prompts you to take action if updates are available.

## Installation

To install Gitorchk, you can use `go get`:

```bash
go get github.com/danny-molnar/gitorchk
```

Ensure your Go environment is set up properly and your `$GOPATH/bin` directory is in your system's PATH.

## Usage

Run Gitorchk in your Git repository directory:

```bash
gitorchk
```

If your local main branch is behind the remote main branch, Gitorchk will prompt you to take action to update your repository.

## Features

- Check if local main branch is behind remote main branch.
- Prompt user to take action if updates are available.

## Contributing

Contributions are welcome! If you find a bug or have an idea for a new feature, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
