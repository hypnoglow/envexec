# envexec

[![Codefresh build status]( https://g.codefresh.io/api/badges/pipeline/hypnoglow/hypnoglow%2Fenvexec%2Fenvexec?type=cf-1)]( https://g.codefresh.io/public/accounts/hypnoglow/pipelines/hypnoglow/envexec/envexec)
[![GolangCI](https://golangci.com/badges/github.com/hypnoglow/envexec.svg)](https://golangci.com/r/github.com/hypnoglow/envexec)

**envexec** helps to provision an application by taking values from
sources like Vault and bringing them as environment variables.

## Features

- Out of process (no code dependency)
- Works with any app written in any language
- One small static binary (💙 Golang)
- Familiar configuration format, with versions
- No supervising, just replaces the process with `exec`
- Simple Docker integration

## Usage

### Docker

The easiest way to embed **envexec** into your Docker image is to just
copy the binary from the prebuilt image:

```docker
FROM alpine:3.8

COPY --from=hypnoglow/envexec:latest-scratch /envexec /usr/local/bin/envexec

ENTRYPOINT ["envexec", "--"]
CMD ["echo", "Hello from envexec!"]
```

An alternative approach is to build your image with **envexec** image
as a base:

```
FROM hypnoglow/envexec:latest-alpine

ENTRYPOINT ["envexec", "--"]
CMD ["echo", "Hello from envexec!"]
```

*NOTE: Using "latest" tags is not recommended. Prefer tagged versions.*

See [examples](_examples/docker/) for more info.

## Providers

### Vault

To fetch secrets from Vault and export values as environment
variables, you need to prepare a spec. Example:

```yaml
apiVersion: envexec/v1alpha1
kind: VaultSecrets
secrets:
  - path: secret/namespace/service/some
    key: api_key
    env: SOME_API_KEY
  - path: secret/namespace/service/db
    key: password
    env: DB_PASSWORD
```

Store this spec in the file `vaultsecrets.yaml`.

Next you need to prepare environment variables to authenticate
in Vault. This depends on the Vault Auth Method. Lets consider the
simplest token authentication method:

```bash
export VAULT_ADDR="https://vault.company.tld"
export VAULT_METHOD="token"
export VAULT_TOKEN="put-vault-token-here"
```

Now you just run your app through **envexec**:

```bash
envexec --spec-file vaultsecrets.yaml -- /usr/bin/env
```

#### Auth Methods

##### Token

See: https://www.vaultproject.io/docs/auth/token.html

```bash
export VAULT_ADDR="https://vault.company.tld"
export VAULT_METHOD="token"
export VAULT_TOKEN="put-vault-token-here"

envexec --spec-file vaultsecrets.yaml -- /usr/bin/env
```

##### Kubernetes

See: https://www.vaultproject.io/docs/auth/kubernetes.html

```bash
export VAULT_ADDR="https://vault.company.tld"
export VAULT_AUTH_METHOD="kubernetes"
export VAULT_AUTH_KUBERNETES_ROLE="foo-app"

envexec --spec-file vaultsecrets.yaml /usr/bin/env
```

## Acknowledgements

Inspired by:

- https://github.com/hashicorp/envconsul
- https://github.com/channable/vaultenv
