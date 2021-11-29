.PHONY: build

build:
	@go build -o build/sensors-informer
	@cp config.yml.dist build/config.yml

clean:
	@rm -rf ./build

install:
	@cp build/sensors-informer /usr/local/bin/sensors-informer
	@mkdir -p /etc/sensors-informer/
	@cp build/config.yml /etc/sensors-informer/config.yml

uninstall:
	@rm -f /usr/local/bin/sensors-informer
	@rm -rf /etc/sensors-informer/
