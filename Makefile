out_dir = "./build"
out_bin = "dumb-launcher"

default:
	@go run *.go

pkg:
	@go build -o=${out_dir}/${out_bin}
