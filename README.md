# keystore

### Prerequisites

Set Azure Environment Variables

```bash
export AZURE_ENVIRONMENT=xxxxxx
export AZURE_CLIENT_ID=xxxxxx
export AZURE_CLIENT_SECRET=xxxxxx
export AZURE_TENANT_ID=xxxxxx
export KVAULT=xxxxxx
```

### Sample

```
cli, err := client.NewAzureClient()
if err != nil {
    panic(err)
}
key := "test-only-pwd"
val, err := cli.GetValue(key)
if err != nil {
    panic(err)
}
fmt.Println(val)
```
