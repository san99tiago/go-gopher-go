# INSTALLING GO LIBRARIES

In Golang, we usually have "Standard Libraries" and "External Libraries". Both of them are important and we use them based on our needs and requirements. The standard ones are the "built-in" modules that are really common to use in a daily basis. The external ones are the ones developed by any maintainer and are not directly supported by Golang community.

## Standard Libraries

To search for standard libraries, we can go to:

- [Go Standard Library](https://pkg.go.dev/std)

## External Libraries

We could search for them at:

- [Go Packages](https://pkg.go.dev)

However, we could also find them somewhere else (as they are customized by any maintainer).

### Install dependency

```bash
go get <MODULE_PATH>
```

### Uninstall dependencies

```bash
go mod tidy
```

### Examples

```bash
go get example.com/theirmodule@latest
```
