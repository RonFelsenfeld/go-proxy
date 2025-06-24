.PHONY: run-proxy-server run-upstream-server run-all-servers

run-proxy-server:
	go run cmd/proxy/main.go

run-upstream-server:
	go run cmd/upstreamserver/main.go

run-all-servers:
	@echo "🟢 Starting upstream server..."
	go run cmd/upstreamserver/main.go & \
	echo "🟢 Starting proxy..." && \
	go run cmd/proxy/main.go

