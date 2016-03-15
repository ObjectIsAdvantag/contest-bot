PROGRAM=webhook
GOFLAGS= -tags netgo

default: all

.PHONY: all
all: build run

.PHONY: goget
goget:
	go get github.com/spf13/viper

.PHONY: clean
clean:
	rm -f $(PROGRAM) $(PROGRAM).exe $(PROGRAM).zip $(PROGRAM).debug

.PHONY: build
build: clean goget
	go build $(GOFLAGS) *.go

.PHONY: dev
dev:
	go run $(GOFLAGS) *.go

.PHONY: run
run:
	(go run $(GOFLAGS) *.go &)
	(lt -p 8081 -s contestbot &)
