.PHONY: build-ui
build-ui:
	cd ui && npm run build:go

build-and-run: build-ui
	cd be && go run .
