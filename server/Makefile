.PHONY: all clean install uninstall

all: server

server:
	export CGO_ENABLED=0; \
	if [ "$(uname)" = "Linux" ]; then \
		go get -t -d .; \
		go build; \
	else \
		export GOOS=linux GOARCH=amd64; \
		go get -t -d .; \
		go build; \
	fi

server.exe:
	export CGO_ENABLED=0 GOOS=windows GOARCH=386; \
	go get -t -d .; \
	go build; \

clean:
	rm -rf server server.exe
