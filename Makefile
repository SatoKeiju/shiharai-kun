.PHONY: goa-gen
goa-gen:
	go run goa.design/goa/v3/cmd/goa@v3.19.1 gen ./design;
