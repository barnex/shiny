all:
	goimports -w *.go
	go build

run:
	go build
	./shiny
