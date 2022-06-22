# MUTABLE/IMMUTABLE DATA TYPES IN GOLANG

To create a simple GoLang structure for a project, we can execute in the terminal:

```bash
go mod init example.com/functions_for
```

This command will create a file called `go.mod`, which has the initial GoLang setup.

The file for this example is `main.go`. This is the "main" package and it imports a couple of dependencies.

This example is really interesting, because it shows us how the underlying memory is used for referencing values when creating multiple variables (numbers, arrays, slices, maps, ...).

To run the script, we can execute:

```bash
go run main.go
```
