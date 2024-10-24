.PHONY: goa-gen
goa-gen:
	go run goa.design/goa/v3/cmd/goa@v3.19.1 gen github.com/SatoKeiju/shiharai-kun/design;
