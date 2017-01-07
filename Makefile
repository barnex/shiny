all:
	goimports -w *.go
	go build

run: all
	./shiny
