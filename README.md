# Transitive dependencies in Go

This repository was an experiment to learn better how transitive dependencies work in Go. It contains two Go modules:

- A library module that houses two packages. One of these packages `cgo` has a dependency on the Confluent Kafka client, which requires `CGO_ENABLED` to build. The other package `clean` is standard Go.
- An application module that has a dependency on the library module. This application requires the library as a whole but only imports the standard Go package in its source code.

Notice that the application does not state any transitive dependency on the Confluent Kafka client in its `go.mod`. You can build the application successfully without `CGO_ENABLED=1`
using the Makefile target.

Use `make cgo` to switch the application code so that it now imports the `cgo` package from the library. Then run `go mod tidy`. You will now be able to see the Confluent client in the `go.mod`
as an indirect dependency. You will also be unable to build `main.go` unless you have `CGO_ENBALED=1`.
