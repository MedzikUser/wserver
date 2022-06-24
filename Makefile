GO := go
BIN_OUT := bin/server
PLUGINS := plugins

all: build build-plugins

build:
	$(GO) build \
		-v \
		-o "${BIN_OUT}"

build-plugins:
	mkdir -p plugins
	for plugin in ./plugin_*; do \
		$(GO) build \
		-v \
		-o plugins/ \
		-buildmode=plugin \
		"$${plugin}" ; \
	done

lint:
	$(GO) fmt
	$(GO) mod tidy

clean:
	rm -rf bin
	rm -rf plugins
	$(GO) clean
