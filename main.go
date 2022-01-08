package main

import (
	"fmt"
	"keyvault/client"
)

func main() {
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
}
