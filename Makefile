.PHONY: all test race cover fmt fmt-check vet tidy staticcheck ci

# Subdirectories containing their own go.mod
MODULES := . gin fiber echo iris beego fasthttp

all: fmt vet test

ci: fmt-check vet race

test:
	@for dir in $(MODULES); do \
		echo "Testing $$dir..."; \
		(cd $$dir && go test -cover ./...) || exit 1; \
	done

race:
	@for dir in $(MODULES); do \
		echo "Race testing $$dir..."; \
		(cd $$dir && go test -race -cover ./...) || exit 1; \
	done

cover:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out | tail -1

fmt:
	@for dir in $(MODULES); do \
		echo "Formatting $$dir..."; \
		(cd $$dir && go fmt ./...) || exit 1; \
	done

fmt-check:
	@out=$$(gofmt -l $(MODULES)); \
	if [ -n "$$out" ]; then \
		echo "Not gofmt'd:"; echo "$$out"; exit 1; \
	fi

vet:
	@for dir in $(MODULES); do \
		echo "Vetting $$dir..."; \
		(cd $$dir && go vet ./...) || exit 1; \
	done

staticcheck:
	@command -v staticcheck >/dev/null || go install honnef.co/go/tools/cmd/staticcheck@latest
	@for dir in $(MODULES); do \
		echo "Staticcheck $$dir..."; \
		(cd $$dir && staticcheck ./...) || exit 1; \
	done

tidy:
	@for dir in $(MODULES); do \
		echo "Tidying $$dir..."; \
		(cd $$dir && go mod tidy) || exit 1; \
	done
