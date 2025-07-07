.PHONY: build
build:
	templ generate internal/view
	go build -o ./bin/vv-tickets cmd/main.go

.PHONY: run
run: build
	./bin/vv-tickets