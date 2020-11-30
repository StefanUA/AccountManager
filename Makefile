build:
	go build;

clean:
	go clean;

test:
	go test ./...;

run:
	./AccountManager --input=input.txt --output=results.txt;