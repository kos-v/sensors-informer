.PHONY: build

DC_DEV_CONF=docker-compose.dev.yml

build: clean
	@go build -o build/sensinf cmd/informer/main.go
	@cp etc/conf/config.yml build/config.yml
	@go build -o build/sensinf-monitor cmd/monitor/main.go
	@mkdir -p build/ui/monitor && \
		cp -r ./ui/monitor/dist build/ui/monitor/dist && \
		cp -r ./ui/monitor/templates build/ui/monitor/templates

clean:
	@rm -rf ./build

install:
	@cp build/sensinf /usr/local/bin/sensinf
	@mkdir -p /etc/sensors-informer/
	@cp build/config.yml /etc/sensors-informer/config.yml

uninstall:
	@rm -f /usr/local/bin/sensinf
	@rm -rf /etc/sensors-informer/

fmt:
	@go fmt ./...

build-dev-env:
	@docker-compose -f ${DC_DEV_CONF} up --build -d
	@docker-compose -f ${DC_DEV_CONF} exec node sh -c "npm install -g @vue/cli"