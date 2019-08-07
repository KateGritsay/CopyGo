
    all: install build test lint deps

    install:
	        go install $(GOPATH)/src/copygo.go
    build:
            go build -o copygo
    test:
            go test -v ./...
    lint:
    	    golangci-lint run

    deps:
            go get github.com/KateGritsay/CopyGo
            go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
