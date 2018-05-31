PKG="github.com/bramz/systemgo"

build: | test
    go build $(PKG)

fmt:
    go fmt $(PKG)

test: | fmt
    go test $(PKG)

