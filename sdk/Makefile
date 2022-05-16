
install:
	# go
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		golang.org/x/tools/cmd/goimports \
		mvdan.cc/gofumpt \
		github.com/tmthrgd/go-bindata/go-bindata \
		golang.org/x/tools/cmd/stringer \
		github.com/vektra/mockery/v2 \
		github.com/go-swagger/go-swagger/cmd/swagger \
		github.com/sebdah/markdown-toc \
		github.com/golang/mock/mockgen \
		github.com/99designs/gqlgen \
		github.com/cucumber/godog/cmd/godog
	go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	go install github.com/kyleconroy/sqlc/cmd/sqlc@$(SQLC_VERSION)

FORMAT_FILES := grep -L -R "^\/\/ Code generated .* DO NOT EDIT\.$$" --exclude-dir=.git --exclude-dir=vendor --include="*.go" .

format:
	go install mvdan.cc/gofumpt golang.org/x/tools/cmd/goimports
	$(FORMAT_FILES) | xargs gofumpt -w
	$(FORMAT_FILES) | xargs goimports -w -local github.com/utilitywarehouse/partner-mono
	buf format proto -w

lint-imports:
	go install github.com/fatih/faillint github.com/pavius/impi/cmd/impi
	faillint -paths \
	    "log=github.com/utilitywarehouse/partner-mono/pkg/log/v1,errors=github.com/pkg/errors,github.com/kelseyhightower/envconfig=github.com/utilitywarehouse/partner-mono/pkg/app/flags,github.com/urfave/cli=github.com/urfave/cli/v2" \
	    ./...
	impi --local github.com/utilitywarehouse/partner-mono --scheme stdThirdPartyLocal --ignore-generated=true ./...

download-protos:
	./proto/download.sh

generate:
	markdown-toc --skip-headers=2 --replace --inline README.md
	./proto/generate.sh
	go generate -x ./...

# *-changed commands
# Run specific commands only on changed files, allowing quick checks on them. Can be used with git hooks
# Use `check` to run all *-changed commands

generate-protos-changed:
	@echo :: Running proto generator
	@$(eval GEN = $(shell git diff --dirstat=files,0 master | grep proto))
	@$(if $(strip $(GEN)),make generate,echo Skipping generators)

format-changed:
	@echo :: Running gofumports
	@git diff --dirstat=files,0 master | sed -E 's/^[ 0-9.]+% //g' | xargs -I {} gofumports -w -local github.com/utilitywarehouse/partner-mono {}

lint-changed:
	@echo :: Running linter
	@git diff --dirstat=files,0 master | sed -E 's/^[ 0-9.]+% //g' | xargs -I {} golangci-lint run ./{}

test-changed:
	@echo :: Running tests
	@git diff --dirstat=files,0 master | sed -E 's/^[ 0-9.]+% //g' | xargs -I {} go test ./{}...

check: generate-protos-changed format-changed lint-changed test-changed
