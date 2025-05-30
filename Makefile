.PHONY: build install clean

BINARY_NAME=tangerine-vault
INSTALL_DIR=/usr/local/bin

build:
	go build -o $(BINARY_NAME) .

install: build
	sudo mv $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME) 