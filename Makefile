check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger: check_install
	# if this doesn't work, check that $$PATH includes the go 'bin' directory
	swagger generate spec -i ./swaggerseed.json -o ./swagger.json --scan-models