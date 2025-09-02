GO=		go

GO_PACKAGE=	github.com/fumiyas/qrc/cmd/qrc
CROSS_TARGETS=	linux/amd64 linux/386 darwin/amd64 windows/386

default: build

get:
	$(GO) mod tidy

vendor:
	$(GO) mod vendor

build:
	$(GO) build -mod=vendor ./cmd/qrc

cross:
	@set -e; \
	for target in $(CROSS_TARGETS); do \
		OS=$${target%/*}; ARCH=$${target#*/}; \
		echo "Building $$OS/$$ARCH"; \
		GOOS=$$OS GOARCH=$$ARCH $(GO) build -mod=vendor -o dist/qrc-$$OS-$$ARCH ./cmd/qrc; \
	done

test:
	$(GO) test -mod=vendor ./...

