binary_name = wiki
destination = /usr/local/bin

.PHONY: all
all: build install

.PHONY: build
build:
	go build -o $(binary_name)

.PHONY: install
install:
	sudo mv $(binary_name) $(destination)
