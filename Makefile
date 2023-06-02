#!make

vet:
	go vet ./...

test:
	go test -v ./...

# should make this a variable
# should check that that the variable is not empty
# to avoid deleting the whole system
removefolder:
	rm -rf /Users/berrio/xyz.memorystore