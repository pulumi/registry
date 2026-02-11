---
# WARNING: this file was fetched from https://raw.githubusercontent.com/kotaicode/pulumi-dex/v0.7.2/docs/installation-configuration.md
# Do not edit by hand unless you're certain you know what you are doing!
edit_url: https://github.com/kotaicode/pulumi-dex/blob/v0.7.2/docs/installation-configuration.md
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

Install the provider using `go get`:

```bash
go get github.com/kotaicode/pulumi-provider-dex/sdk/go/dex
```

### Python

Install the provider using pip:

```bash
pip install pulumi-dex
```

### Provider Binary

The provider binary is automatically downloaded by Pulumi when you use the provider. If you need to install it manually:

```bash
pulumi plugin install resource dex v0.1.0
```

Or download from [GitHub Releases](https://github.com/kotaicode/pulumi-provider-dex/releases).

## Configuration

The provider requires configuration to connect to your Dex gRPC API.

### Required Configuration

- **`host`** (string): Dex gRPC host and port (e.g., `dex.example.com:5557`)

### Optional Configuration

- **`caCert`** (string, secret): PEM-encoded CA certificate for validating Dex's TLS certificate
- **`clientCert`** (string, secret): PEM-encoded client certificate for mTLS authentication
- **`clientKey`** (string, secret): PEM-encoded private key for the client certificate
- **`insecureSkipVerify`** (boolean): Skip TLS verification (development only, default: `false`)
- **`timeoutSeconds`** (number): Per-RPC timeout in seconds (default: `5`)

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

// Development configuration (insecure)
const devProvider = new dex.Provider("dex-dev", {
    host: "localhost:5557",
    insecureSkipVerify: true,
});
```

#### Go

```go
import (
    "github.com/kotaicode/pulumi-provider-dex/sdk/go/dex"
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
    insecure_skip_verify=True
)
```

### Environment Variables

You can also configure the provider using environment variables, though you'll still need to create a provider instance in your Pulumi program:

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

## Security Best Practices

1. **Use TLS/mTLS in Production**: Never use `insecureSkipVerify: true` in production
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

**Error: `certificate verify failed`**
- Ensure CA certificate matches Dex's certificate chain
- Check certificate expiration dates
- For development, use `insecureSkipVerify: true` (not recommended for production)

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

- See [Resource Documentation](../README.md#resources) for available resources
- Check [Examples](../examples/) for complete working examples
- Review [Dex Documentation](https://dexidp.io/docs/) for Dex-specific configuration

