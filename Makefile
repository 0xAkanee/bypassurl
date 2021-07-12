gofmt:
	@gofmt -w -s main.go && goimports -w main.go && go vet main.go
gobuild-all: gobuild-mac gobuild-linux gobuild-win
gobuild-all-compress: gobuild-mac-compress gobuild-linux-compress gobuild-win-compress

mkout:
	@mkdir -p out/

gobuild-linux: mkout
	@rm -rf out/bypassurl-linux
	@CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o out/bypassurl-linux main.go

gobuild-linux-compress: gobuild-linux
	@cd out/ && upx -k --brute bypassurl-linux

gobuild-mac: mkout
	@rm -rf out/bypassurl-mac
	@CGO_ENABLED=0 GOOS=darwin go build -ldflags "-s -w" -o out/bypassurl-mac main.go

gobuild-mac-compress: gobuild-mac
	@cd out/ && upx -k --brute bypassurl-mac

gobuild-win: mkout
	@rm -rf bypassurl-win.exe
	@CGO_ENABLED=0 GOOS=windows go build -ldflags "-s -w" -o out/bypassurl-win.exe main.go

gobuild-win-compress: gobuild-win
	@cd out/ && upx -k --brute bypassurl-win.exe
