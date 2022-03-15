# HELLO WORLD IN GOLANG

To create a simple GoLang structure for a project, we can execute in the terminal:

```bash
go mod init example.com/hello_world
```

This command will create a file called `go.mod`, which has the initial GoLang setup.

The file for this example is `main.go`. This is the "main" package and it imports a couple of dependencies that are used to print some info infinitely and wait "N" seconds to repeat the given message (a simple infinite loop that shows a message every "N" seconds).

To run the script, we can execute:

```bash
go run main.go
```
