---
# WARNING: this file was fetched from https://raw.githubusercontent.com/kotaicode/pulumi-dex/v0.12.0/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/kotaicode/pulumi-dex/blob/v0.12.0/docs/installation-configuration.md
title: Installation and Configuration
meta_desc: Learn how to install and configure the Dex provider for Pulumi.
layout: package
---

# Installation and Configuration

## Installation

### TypeScript/JavaScript

Install the provider using npm:

```bash
npm install @kotaicode/pulumi-dex
```

### Go

Install the provider using `go get` (the Go module lives in `sdk/`; the package import path is `.../sdk/go/dex`):

```bash
go get github.com/kotaicode/pulumi-dex/sdk
```

### Python

Install the provider using pip:

```bash
pip install pulumi-dex
```

### Provider Binary

The provider binary is automatically downloaded by Pulumi when you use the provider. If you need to install it manually:

```bash
pulumi plugin install resource dex v<version> --server github://api.github.com/kotaicode/pulumi-dex
```

Use the same version as your SDK package (for npm: `npm ls @kotaicode/pulumi-dex`).

Or download from [GitHub Releases](https://github.com/kotaicode/pulumi-dex/releases).

## Configuration

The provider requires configuration to connect to your Dex gRPC API.

### Required Configuration

- **`host`** (string): Dex gRPC host and port (e.g., `dex.example.com:5557`)

### Optional Configuration

- **`caCert`** (string, secret): PEM-encoded CA certificate that issued Dex's TLS certificate. When set, only this CA is trusted instead of the system's root CAs
- **`clientCert`** (string, secret): PEM-encoded client certificate for mTLS; requires `clientKey`
- **`clientKey`** (string, secret): PEM-encoded private key for the client certificate
- **`insecureSkipVerify`** (boolean): TLS without certificate verification (development only, default: `false`); the connection is still encrypted
- **`plaintext`** (boolean): Connect without TLS (default: `false`). Only for a Dex whose gRPC API has no TLS and is reachable only on a trusted network, such as localhost or inside a cluster. Cannot be combined with the TLS options
- **`timeoutSeconds`** (number): Per-RPC timeout in seconds (default: `5`)

Without TLS options the provider connects with TLS 1.2 or later and verifies Dex's certificate and host name against the system's root CAs. It never falls back to plaintext.

### Configuration Examples

#### TypeScript

```typescript
import * as dex from "@kotaicode/pulumi-dex";
import * as fs from "fs";

// Production configuration with mTLS
const provider = new dex.Provider("dex", {
    host: "dex.example.com:5557",
    caCert: fs.readFileSync("certs/ca.crt", "utf-8"),
    clientCert: fs.readFileSync("certs/client.crt", "utf-8"),
    clientKey: fs.readFileSync("certs/client.key", "utf-8"),
    timeoutSeconds: 10,
});

// Local Dex without TLS (for example the docker-compose setup)
const devProvider = new dex.Provider("dex-dev", {
    host: "localhost:5557",
    plaintext: true,
});
```

#### Go

```go
import (
    "github.com/kotaicode/pulumi-dex/sdk/go/dex"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "os"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        caCert, _ := os.ReadFile("certs/ca.crt")
        clientCert, _ := os.ReadFile("certs/client.crt")
        clientKey, _ := os.ReadFile("certs/client.key")
        
        provider, err := dex.NewProvider(ctx, "dex", &dex.ProviderArgs{
            Host:       pulumi.String("dex.example.com:5557"),
            CaCert:     pulumi.String(string(caCert)),
            ClientCert: pulumi.String(string(clientCert)),
            ClientKey:  pulumi.String(string(clientKey)),
            TimeoutSeconds: pulumi.Int(10),
        })
        if err != nil {
            return err
        }
        
        // Use provider...
        return nil
    })
}
```

#### Python

```python
import pulumi_dex as dex

# Production configuration
provider = dex.Provider("dex",
    host="dex.example.com:5557",
    ca_cert=open("certs/ca.crt").read(),
    client_cert=open("certs/client.crt").read(),
    client_key=open("certs/client.key").read(),
    timeout_seconds=10
)

# Development configuration
dev_provider = dex.Provider("dex-dev",
    host="localhost:5557",
    plaintext=True
)
```

### Environment Variables

The provider does not read environment variables itself. To take settings from the environment, read them in your Pulumi program:

```bash
export DEX_HOST="dex.example.com:5557"
export DEX_CA_CERT="$(cat certs/ca.crt)"
export DEX_CLIENT_CERT="$(cat certs/client.crt)"
export DEX_CLIENT_KEY="$(cat certs/client.key)"
export DEX_TIMEOUT_SECONDS="10"
```

Then reference them in your code:

```typescript
const provider = new dex.Provider("dex", {
    host: process.env.DEX_HOST!,
    caCert: process.env.DEX_CA_CERT,
    clientCert: process.env.DEX_CLIENT_CERT,
    clientKey: process.env.DEX_CLIENT_KEY,
    timeoutSeconds: parseInt(process.env.DEX_TIMEOUT_SECONDS || "5"),
});
```

## Dex Setup Requirements

### Enable gRPC API

Ensure your Dex instance has the gRPC API enabled. In your Dex configuration:

```yaml
grpc:
  addr: 0.0.0.0:5557
  tlsCert: /etc/dex/tls/grpc-server.crt
  tlsKey: /etc/dex/tls/grpc-server.key
  tlsClientCA: /etc/dex/tls/grpc-client-ca.crt
```

### Enable Connector CRUD (for connector management)

For managing connectors via the API, Dex must be started with:

```bash
DEX_API_CONNECTORS_CRUD=true dex serve config.yaml
```

Or in your deployment:

```yaml
env:
  - name: DEX_API_CONNECTORS_CRUD
    value: "true"
```

### Network Access

Ensure your Pulumi program can reach Dex's gRPC endpoint:
- Default port: `5557`
- Protocol: gRPC (HTTP/2)
- For production: Use TLS/mTLS

## Local development and testing

To build, generate SDKs, and test against a local Dex instance:

1. **Build the provider and generate schema/SDKs** (from repo root):
   ```bash
   make build
   make generate-sdks     # builds, writes schema.json from the binary, then generates the TypeScript, Go and Python SDKs
   ```
   Generation needs the Pulumi CLI version in `.pulumi.version`; with another version `make` prints a warning, and CI fails because the generated code may differ. Any failed step makes `make` fail. CI regenerates schema and SDKs and fails if the committed files differ, so commit the regenerated `schema.json` and `sdk/` together with provider changes.

2. **Install the provider so Pulumi can find it**:
   ```bash
   make install
   ```
   Or manually: `pulumi plugin install resource dex v<version> --file bin/pulumi-resource-dex`, where `<version>` is the `VERSION` set in the Makefile

3. **Start local Dex** (Docker required):
   ```bash
   make dex-up
   ```

4. **Run the TypeScript example**:
   ```bash
   cd examples/typescript
   npm install
   pulumi stack init dev   # or pulumi stack select dev
   pulumi preview
   pulumi up               # confirm with yes
   ```
   See `examples/typescript/TESTING.md` for more detail and troubleshooting.

5. **Test the Go SDK** (optional). In a separate Go project:
   ```bash
   go mod init myapp
   go get github.com/kotaicode/pulumi-dex/sdk
   ```
   Then use `import "github.com/kotaicode/pulumi-dex/sdk/go/dex"` in your Pulumi program. To test against the provider repo before publishing, use a `replace` in your `go.mod`:
   ```go
   replace github.com/kotaicode/pulumi-dex/sdk => /path/to/pulumi-dex/sdk
   ```

6. **Stop Dex when done**: `make dex-down` (from repo root).

## Security Best Practices

1. **Use TLS/mTLS in Production**: Never use `insecureSkipVerify: true` or `plaintext: true` for a Dex outside your machine or cluster
2. **Protect Secrets**: Use Pulumi secrets for sensitive configuration:
   ```typescript
   const provider = new dex.Provider("dex", {
       host: "dex.example.com:5557",
       clientKey: pulumi.secret(fs.readFileSync("certs/client.key", "utf-8")),
   });
   ```
3. **Restrict Network Access**: Limit access to Dex's gRPC API to trusted networks/IPs
4. **Rotate Certificates**: Regularly rotate mTLS certificates
5. **Monitor Access**: Monitor Dex logs for unauthorized access attempts

## Troubleshooting

### Connection Errors

**Error: `failed to connect to Dex`**
- Verify Dex is running and gRPC API is enabled
- Check network connectivity to Dex host:port
- Verify firewall rules allow gRPC traffic

**Error: `did not answer with TLS`**
- This Dex serves its gRPC API without TLS. Set `plaintext: true` if it is only reachable on a trusted network, or enable TLS in Dex (`grpc.tlsCert`, `grpc.tlsKey`)

**Error: `is not trusted by this system` or `is not issued by the CA in caCert`**
- Set `caCert` to the CA certificate that issued Dex's gRPC certificate
- Check certificate expiration dates
- For development only, `insecureSkipVerify: true` skips the check

**Error: `rejected the TLS connection`**
- Dex requires client certificates (`grpc.tlsClientCA`): set `clientCert` and `clientKey`, signed by a CA Dex trusts

### Authentication Errors

**Error: `authentication failed`**
- Verify client certificate and key are correct
- Ensure client certificate is signed by Dex's trusted CA
- Check that mTLS is properly configured on both sides

### Timeout Errors

**Error: `context deadline exceeded`**
- Increase `timeoutSeconds` if operations are slow
- Check network latency to Dex
- Verify Dex is not overloaded

## Next Steps

- See [Resource Documentation](../api-docs) for available resources
- Check [Examples](https://github.com/kotaicode/pulumi-dex/tree/main/examples) for complete working examples
- Review [Dex Documentation](https://dexidp.io/docs/) for Dex-specific configuration

