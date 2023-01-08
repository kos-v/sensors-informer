.PHONY: build

build: clean
	@go build -o build/sensinf cmd/informer/main.go
	@cp etc/conf/config.yml build/config.yml

clean:
	@rm -rf ./build

install:
	@cp build/sensinf /usr/local/bin/sensinf
	@mkdir -p /etc/sensors-informer/
	@cp build/config.yml /etc/sensors-informer/config.yml

uninstall:
	@rm -f /usr/local/bin/sensinf
	@rm -rf /etc/sensors-informer/

dev-fmt:
	@go fmt ./...