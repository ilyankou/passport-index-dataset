IMAGE := passport-parser

.PHONY: build upgrade

build:
	docker build -t $(IMAGE) ./parser

upgrade: build
	docker run --rm -v "$(CURDIR):/data" $(IMAGE)
