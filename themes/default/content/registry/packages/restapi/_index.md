---
# WARNING: this file was fetched from https://djoiyj6oj2oxz.cloudfront.net/docs/registry.opentofu.org/mastercard/restapi/2.0.1/index.md
# Do not edit by hand unless you're certain you know what you are doing!
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Restapi Provider
meta_desc: Provides an overview on how to configure the Pulumi Restapi provider.
layout: package
---

## Generate Provider

The Restapi provider must be installed as a Local Package by following the [instructions for Any Terraform Provider](https://www.pulumi.com/registry/packages/terraform-provider/):

```bash
pulumi package add terraform-provider mastercard/restapi
```
## Example Usage

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    restapi:createMethod:
        value: PUT
    restapi:debug:
        value: true
    restapi:destroyMethod:
        value: PUT
    restapi:headers:
        value: |-
            TODO: {
            "X-Auth-Token"=var.AUTH_TOKEN,
            "Content-Type"="application/json"
            }
    restapi:updateMethod:
        value: PUT
    restapi:uri:
        value: https://api.url.com
    restapi:writeReturnsObject:
        value: true

```
## Configuration Reference
### Required

- `uri` (String) URI of the REST API endpoint. This serves as the base of all requests.

- `certFile` (String) When set with the keyFile parameter, the provider will load a client certificate as a file for mTLS authentication.
- `certString` (String) When set with the keyString parameter, the provider will load a client certificate as a string for mTLS authentication.
- `copyKeys` (List of String) When set, any PUT to the API for an object will copy these keys from the data the provider has gathered about the object. This is useful if internal API information must also be provided with updates, such as the revision of the object.
- `createMethod` (String) Defaults to `POST`. The HTTP method used to CREATE objects of this type on the API server.
- `createReturnsObject` (Boolean) Set this when the API returns the object created only on creation operations (POST). This is used by the provider to refresh internal data structures.
- `debug` (Boolean) Enabling this will cause lots of debug information to be printed to STDOUT by the API client.
- `destroyMethod` (String) Defaults to `DELETE`. The HTTP method used to DELETE objects of this type on the API server.
- `headers` (Map of String) A map of header names and values to set on all outbound requests. This is useful if you want to use a script via the 'external' provider or provide a pre-approved token or change Content-Type from `application/json`. If `username` and `password` are set and Authorization is one of the headers defined here, the BASIC auth credentials take precedence.
- `idAttribute` (String) When set, this key will be used to operate on REST objects. For example, if the ID is set to 'name', changes to the API object will be to <http://foo.com/bar/VALUE_OF_NAME>. This value may also be a '/'-delimeted path to the id attribute if it is multple levels deep in the data (such as `attributes/id` in the case of an object `{ "attributes": { "id": 1234 }, "config": { "name": "foo", "something": "bar"}}`
- `insecure` (Boolean) When using https, this disables TLS verification of the host.
- `keyFile` (String) When set with the certFile parameter, the provider will load a client certificate as a file for mTLS authentication. Note that this mechanism simply delegates to golang's tls.LoadX509KeyPair which does not support passphrase protected private keys. The most robust security protections available to the keyFile are simple file system permissions.
- `keyString` (String) When set with the certString parameter, the provider will load a client certificate as a string for mTLS authentication. Note that this mechanism simply delegates to golang's tls.LoadX509KeyPair which does not support passphrase protected private keys. The most robust security protections available to the keyFile are simple file system permissions.
- `oauthClientCredentials` (Block List, Max: 1) Configuration for oauth client credential flow using the <https://pkg.go.dev/golang.org/x/oauth2> implementation (see below for nested schema)
- `password` (String) When set, will use this password for BASIC auth to the API.
- `rateLimit` (Number) Set this to limit the number of requests per second made to the API.
- `readMethod` (String) Defaults to `GET`. The HTTP method used to READ objects of this type on the API server.
- `rootCaFile` (String) When set, the provider will load a root CA certificate as a file for mTLS authentication. This is useful when the API server is using a self-signed certificate and the client needs to trust it.
- `rootCaString` (String) When set, the provider will load a root CA certificate as a string for mTLS authentication. This is useful when the API server is using a self-signed certificate and the client needs to trust it.
- `testPath` (String) If set, the provider will issue a readMethod request to this path after instantiation requiring a 200 OK response before proceeding. This is useful if your API provides a no-op endpoint that can signal if this provider is configured correctly. Response data will be ignored.
- `timeout` (Number) When set, will cause requests taking longer than this time (in seconds) to be aborted.
- `updateMethod` (String) Defaults to `PUT`. The HTTP method used to UPDATE objects of this type on the API server.
- `useCookies` (Boolean) Enable cookie jar to persist session.
- `username` (String) When set, will use this username for BASIC auth to the API.
- `writeReturnsObject` (Boolean) Set this when the API returns the object created on all write operations (POST, PUT). This is used by the provider to refresh internal data structures.
- `xssiPrefix` (String) Trim the xssi prefix from response string, if present, before parsing.

<a id="nestedblock--oauth_client_credentials"></a>
### Nested Schema for `oauthClientCredentials`

Required:

- `oauthClientId` (String) client id
- `oauthClientSecret` (String) client secret
- `oauthTokenEndpoint` (String) oauth token endpoint

Optional:

- `endpointParams` (Map of String) Additional key/values to pass to the underlying Oauth client library (as EndpointParams)
- `oauthScopes` (List of String) scopes