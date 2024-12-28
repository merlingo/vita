build-all:
    GOOS=linux GOARCH=amd64 go build -o ./build/myproject-linux-amd64 ./cmd/myprojectd/main.go
    GOOS=linux GOARCH=arm64 go build -o ./build/myproject-linux-arm64 ./cmd/myprojectd/main.go
    GOOS=darwin GOARCH=amd64 go build -o ./build/myproject-darwin-amd64 ./cmd/myprojectd/main.go

do-checksum:
    cd build && sha256sum myproject-linux-amd64 myproject-linux-arm64 myproject-darwin-amd64 > myproject_checksum

build-with-checksum: build-all do-checksum
