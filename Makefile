gofmt:
	@gofmt -w -s main.go && goimports -w main.go && go vet main.go
gobuild-all: gobuild-mac gobuild-linux gobuild-win
gobuild-linux:
	@CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o bypassurl main.go 
gobuild-mac:
	@CGO_ENABLED=0 GOOS=darwin go build -ldflags "-s -w" -o bypassurl main.go
gobuild-win:
	@CGO_ENABLED=0 GOOS=windows go build -ldflags "-s -w" -o bypassurl main.go