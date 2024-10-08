LOCAL_BIN:=$(CURDIR)/bin

# линтер производит статический анализ кода и на основе правил, устанавливаю его локально
install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

# если не указать конфиг, линтер применит базовые настройки, но я указываю
lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml