all:
	goimports -w *.go
	go build -race

run: all
	./shiny
