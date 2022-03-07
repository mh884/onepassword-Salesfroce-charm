package main

import (
	"fmt"

	"github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
)

type Config struct {
	Username      string `opfield:"username"`
	Password      string `opfield:"password"`
	SecurityToken string `opfield:"security-token" opsection:"details"`
	InstnaceUrl   string `opfield:"instance-url" opsection:"details"`
}

func getAllCredentials(vault string) []onepassword.Item {
	var OP_CONNECT_TOKEN = viperEnvVariable("OP_CONNECT_TOKEN")
	var OP_CONNECT_HOST = viperEnvVariable("OP_CONNECT_HOST")
	client := connect.NewClient(OP_CONNECT_HOST, OP_CONNECT_TOKEN)
	Items, err := client.GetItems(vault)
	if err != nil {
		fmt.Println(err)
	}
	return Items
}

func getCredentials(itemUUID string, vault string) Config {
	var OP_CONNECT_TOKEN = viperEnvVariable("OP_CONNECT_TOKEN")
	var OP_CONNECT_HOST = viperEnvVariable("OP_CONNECT_HOST")
	item := Config{}
	client := connect.NewClient(OP_CONNECT_HOST, OP_CONNECT_TOKEN)
	err := client.LoadStructFromItem(&item, itemUUID, vault)
	if err != nil {
		fmt.Println(err)
	}
	return item
}
