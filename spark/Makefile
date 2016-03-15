PROGRAM=webhook
GOFLAGS= -tags netgo

default: all

.PHONY: all
all: build run

.PHONY: goget
goget:

.PHONY: clean
clean:
	rm -f $(PROGRAM) $(PROGRAM).exe $(PROGRAM).zip $(PROGRAM).debug

.PHONY: build
build: clean goget
	go build $(GOFLAGS) $(PROGRAM).go

.PHONY: run
run:
	($(PROGRAM) -port 80801 &)
	(lt -p 8081 -s contestbot &)
