# Go Proxy

A lightweight HTTP/HTTPS proxy server written in Go that forwards requests to an upstream server while modifying request bodies.

## Table of Contents

- [Description](#description)
- [Requirements](#requirements)
- [Local Installation](#local-installation)
- [Environment Variables](#environment-variables)

## Description

This project consists of two main components:

1. **Proxy Server** (`cmd/proxy/main.go`): A TLS-enabled proxy server that:

   - Listens on HTTPS (default port 443)
   - Forwards requests to an upstream server
   - Injects custom key-value pair into request bodies
   - Provides a `/ping` endpoint for health checks
   - Handles request/response modification

2. **Upstream Server** (`cmd/upstreamserver/main.go`): A mock HTTP server that:
   - Serves as the target for proxied requests
   - Provides a `/test` endpoint for testing
   - Runs on HTTP (default port 8081)

The proxy server acts as an intermediary that can modify requests before forwarding them to the upstream server, making it useful for testing, debugging, and request transformation scenarios.

## Requirements

### System Requirements

- Go 1.24.4 or higher (as specified in `go.mod`)

### TLS and Self-Signed Certificates

The proxy server requires TLS certificates to run on HTTPS. Self-signed certificates are used for development purposes. See the [Local Installation](#local-installation) section for certificate generation instructions.

## Local Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/RonFelsenfeld/go-proxy.git
   cd go-proxy
   ```

2. **Generate TLS certificates:**

   ```bash
   # Create the certs directory
   mkdir -p certs

   # Generate self-signed certificate and private key
   openssl req -x509 -newkey rsa:4096 -keyout certs/key.pem -out certs/cert.pem -days 365 -nodes
   ```

   **Note:** When using self-signed certificates, your browser will show a security warning. You can safely proceed by accepting the risk (this is normal for development environments).

3. **Install dependencies:**

   ```bash
   make install-dependencies
   ```

4. **Run the servers:**

   **Option 1: Using Makefile (Recommended)**

   ```bash
   # Run both servers simultaneously
   make run-all-servers

   # Or run them separately
   make run-upstream-server  # In one terminal
   make run-proxy-server     # In another terminal
   ```

   **Option 2: Direct Go commands**

   ```bash
   # Terminal 1 - Start upstream server
   go run cmd/upstreamserver/main.go

   # Terminal 2 - Start proxy server
   go run cmd/proxy/main.go
   ```

5. **Test the setup:**

   ```bash
   # Test the upstream server directly
   curl -X POST http://localhost:<UPSTREAM_PORT>/test \
     -H "Content-Type: application/json" \
     -d '{"test": "data"}'

   # Test the proxy ping endpoint
   curl -k https://localhost:<PROXY_PORT>/ping

   # Test proxying a request - notice the difference in response
   curl -k -X POST https://localhost:<PROXY_PORT>/proxy \
     -H "Content-Type: application/json" \
     -d '{"test": "data"}'
   ```

   **Notes:**

   - The `-k` flag tells curl to skip SSL certificate verification, which is necessary when using self-signed certificates.
   - Replace `<UPSTREAM_PORT>` and `<PROXY_PORT>` with your actual ports. Default values are `8081` for upstream and `443` for proxy.
   - The difference between direct upstream access and proxied requests is that the proxy injects additional key-value pairs into the request body before forwarding it to the upstream server.

## Environment Variables

The application uses environment variables for configuration. You can set these in a `.env` file or as system environment variables.

| Variable        | Default Value           | Description                         |
| --------------- | ----------------------- | ----------------------------------- |
| `UPSTREAM_URL`  | `http://localhost:8081` | URL of the upstream server          |
| `UPSTREAM_PORT` | `8081`                  | Port for the upstream server        |
| `PROXY_PORT`    | `443`                   | Port for the proxy server           |
| `TLS_CERT_PATH` | `certs/cert.pem`        | Path to TLS certificate file        |
| `TLS_KEY_PATH`  | `certs/key.pem`         | Path to TLS private key file        |
| `INJECT_KEY`    | `injected_key`          | Key to inject into request bodies   |
| `INJECT_VALUE`  | `injected_value`        | Value to inject into request bodies |

### Example `.env` file:

```env
UPSTREAM_URL=http://localhost:8081
UPSTREAM_PORT=8081
PROXY_PORT=443
TLS_CERT_PATH=certs/cert.pem
TLS_KEY_PATH=certs/key.pem
INJECT_KEY=custom_header
INJECT_VALUE=custom_value
```

### Notes:

- If no `.env` file is found, the application will fall back to system environment variables
- If neither `.env` nor system environment variables are set, the default values will be used
- The proxy server requires valid TLS certificates to start successfully
