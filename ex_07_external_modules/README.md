# EXTERNAL MODULES IN GOLANG

## Creating a CLI with COBRA module

First, we must initialize the Go project:

```bash
go mod init example.com/external_modules
```

After that, we need to add the required external modules (in this case, COBRA), we can run:

```bash
go get github.com/spf13/cobra/cobra
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
```

These commands should update the `go.mod` and `go.sum` files with the necessary information about the specified dependencies.

Then, we can create the folder for the CLI and initialize it with:

```bash
mkdir santicli
cd santicli
go mod init github.com/spf13/santicli
```

Then, we can use COBRA commands to create the CLI's project structure:

```bash
# Initialize the Go project with the help of cobra command
~/go/bin/cobra-cli init
```

Here, we can modify the main logic (I added a cool hello message).

To build the the CLI, we can run:

```bash
go build .
```

Now we can use the generated binary (`./santicli`) as this:

The first usage is the binary without any flags:

```bash
# Simple usage
./santicli
```

And it shows:

```txt
THE COOL MESSAGE IS: KEEP WORKING HARD MY FRIEND!
```

The second usage is the binary with the `--help` or `-h` flag:

```bash
# Help usage
./santicli --help
```

This command shows:

```txt
For now, this CLI only shows a cool message.

        It was developed using COBRA by san99tiago.

Usage:
  santicli [flags]

Flags:
  -h, --help   help for santicli
THE COOL MESSAGE IS: KEEP WORKING HARD MY FRIEND!
```
