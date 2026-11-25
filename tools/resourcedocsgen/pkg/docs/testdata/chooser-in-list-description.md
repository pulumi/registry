### TLS Termination and Passthrough

- **TLS termination** — set `protocol = "https"` and attach one or more
  `certificates`.

  <!--Start PulumiCodeChooser -->
```typescript
const tlsTermination = new hcloud.LoadBalancerService("tls_termination", {
    protocol: "https",
});
```
<!--End PulumiCodeChooser -->

- **TLS passthrough** — set `protocol = "tcp"` and forward the TLS port.

  <!--Start PulumiCodeChooser -->
  ```python
  tls_passthrough = hcloud.LoadBalancerService("tls_passthrough",
      protocol="tcp")
  ```
  <!--End PulumiCodeChooser -->
