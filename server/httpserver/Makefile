start:
	go run databaseparams.go urlHandler.go fetchInterface.go databaseop.go cacheInterface.go

build:
	go build

install:
	sudo apt install golang-go
	make libs

libs:
	go get github.com/lib/pq

old:
	curl -O https://dl.google.com/go/go1.16.3.linux-amd64.tar.gz
	sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz
	export PATH=${PATH}:/usr/local/go/bin
	go version
