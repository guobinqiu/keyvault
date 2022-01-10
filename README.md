# keystore

### Prerequisites

Set environment variables

```bash
export AZURE_ENVIRONMENT=xxxxxx
export AZURE_CLIENT_ID=xxxxxx
export AZURE_CLIENT_SECRET=xxxxxx
export AZURE_TENANT_ID=xxxxxx
```

### Sample

```
package main

import (
	"fmt"
	"github.com/guobinqiu/keyvault/client"
)

func main() {
	vaultBaseURL := "https://cn-apac-kv-np-akstesting.vault.azure.cn/"
	cli, err := client.NewAzureClient(vaultBaseURL)
	if err != nil {
		panic(err)
	}
	key := "test-only-pwd"
	val, err := cli.GetValue(key)
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

```
