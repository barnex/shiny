all:
	goimports -w *.go
	go build -race

run:
	go build
	./shiny
