.PHONY: build

build:
	@go build -o build/sensors-informer
	@cp config.yml.dist build/config.yml