all: test cover
test:
	go test ./... -v --coverprofile=cover.txt
cover:
	go tool cover -html=cover.txt -o cover.html
