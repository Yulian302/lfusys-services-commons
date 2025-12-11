gen:
	@protoc --go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	api/say_hello.proto
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