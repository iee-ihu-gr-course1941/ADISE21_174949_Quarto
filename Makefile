.PHONY: placeholder

placeholder:
	echo "Makefile doesn't do anything yet"

static:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-extldflags "-f no-PIC -static"' -tags 'osusergo netgo static_build'
