GO := go

all: build build-plugins

build:
	$(GO) build \
		-v \
		-o bin/example \
		.

build-plugins:
	mkdir -p plugins
	for plugin in ./plugin_*; do \
		$(GO) build \
		-v \
		-o plugins/ \
		-buildmode=plugin \
		"$${plugin}" ; \
	done
