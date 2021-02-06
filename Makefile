
all: test

# Run tests
test: verify-generate
	go test ./... -mod vendor

# Run the apiserver locally
run-apiserver:
	go run ./cmd/apiserver/main.go --secure-port=9443 --bind-address=127.0.0.1 --standalone-debug-mode

vendor:
	go mod vendor
	go mod tidy

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate: pkg/apis/core/v1alpha1/manifest_types.go hack/update-codegen.sh
	hack/update-codegen.sh

verify-generate:
	hack/verify-codegen.sh
