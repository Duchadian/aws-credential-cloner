build:
	go build .

install: build
	sudo cp aws-credentials-cloner /usr/local/bin/.
