.PHONY: gen
schema:
	hack/ensure-schema.sh

.PHONY: gen
gen: schema
	go generate ./...

build:
	go build -o .out/ovn-dia .

.PHONY: test
test: build 
	hack/test.sh
