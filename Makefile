.PHONY: proto-gen
proto-gen:
	@if [ -z "$(PROTO_PATH)" ]; then echo "Usage: proto-gen PROTO_PATH=<path_to_proto_file>"; exit 1;fi
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	$$PROTO_PATH
jaeger:
	@docker run \
		--rm \
		-d \
		--name jaeger \
		-p 16686:16686 \
		-p 4317:4317 \
		-p 4318:4318 \
		-p 5778:5778 \
		-p 9411:9411 \
		cr.jaegertracing.io/jaegertracing/jaeger:2.13.0