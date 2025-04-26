out_dir = "./build"
out_bin = "dumb-launcher"

default:
	@go run *.go

sync:
	@go mod tidy && go mod vendor

pkg:
	@go build -o=${out_dir}/${out_bin}
