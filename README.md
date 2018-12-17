# envexec

envexec helps to provision an application by taking values from
sources like Vault and bringing them as environment variables.

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

Now you just run your app through envexec:

```bash
envexec --spec-file vaultsecrets.yaml -- /usr/bin/env
```

#### Auth Methods

##### Token

```bash
export VAULT_ADDR="https://vault.company.tld"
export VAULT_METHOD="token"
export VAULT_TOKEN="put-vault-token-here"

envault --spec-file vaultsecrets.yaml -- /usr/bin/env
```

##### Kubernetes

```bash
export VAULT_ADDR="https://vault.company.tld"
export VAULT_AUTH_METHOD="kubernetes"
export VAULT_AUTH_KUBERNETES_ROLE="foo-app"

envault --spec-file vaultsecrets.yaml /usr/bin/env
```

## Acknowledgements

Inspired by:

- https://github.com/hashicorp/envconsul
- https://github.com/channable/vaultenv
