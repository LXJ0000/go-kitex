.PHONY: api
# generate api proto
api:
	cd app/gateway && cwgo server --type HTTP --server_name gateway --module github.com/LXJ0000/go-kitex/app/gateway --I ../../idl --idl ../../idl/gateway/ping.proto
	cd app/gateway && cwgo server --type HTTP --server_name gateway --module github.com/LXJ0000/go-kitex/app/gateway --I ../../idl --idl ../../idl/gateway/auth2.proto