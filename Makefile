fmt:
	golangci-lint run && treefmt

build:
	go build -o kube-detox .
