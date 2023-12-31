cgo:
	sed -i '' -e 's/clean/cgo/g' application/main.go
	go mod tidy -C application

clean:
	sed -i '' -e 's/cgo/clean/g' application/main.go
	go mod tidy -C application

build/cgo_disabled:
	CGO_ENABLED=0 go build application/main.go

build/cgo_enabled:
	CGO_ENABLED=1 go build application/main.go
