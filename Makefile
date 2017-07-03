BINARY=miner

.DEFAULT_GOAL := build

build:
	go build -i -o ${BINARY}

test:
	go test ./...

install:
	go install

clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi

linux:
	env GOOS=linux GOARCH=arm go build -o ${BINARY}

windows:
	env GOOS=windows GOARCH=386 go build -o ${BINARY}.exe

.PHONY: clean build test linux
