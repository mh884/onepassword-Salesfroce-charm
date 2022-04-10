package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
)

// login to devlops
// select subscribe
// get List team projects.  az devops project list
// get current libery Veriable Groups.  az pipelines variable-group list --project=
// get selected username from DevOps
// get username detils from 1pass
// update username detials in DevOps

func auth() {

	tenantID := viperEnvVariable("AZURE_TENANT_ID")
	clientID := viperEnvVariable("AZURE_CLIENT_ID")
	clientSecret := viperEnvVariable("AZURE_CLIENT_SECRET")
	//cred, err := azidentity.NewDefaultAzureCredential(nil)
	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		fmt.Println("Error running program:", err)
	}
	//AZURE_CLIENT_SECRET=iTT7Q~H0TZtuzhowW_w3e4NH9ddA-q0bIMCao
	var AZURE_SUBSCRIPTION_ID = viperEnvVariable("AZURE_SUBSCRIPTION_ID")
	client := armdevops.NewPipelinesClient(AZURE_SUBSCRIPTION_ID, cred, nil)
	ctx := context.Background()
	pager := client.ListByResourceGroup("VisualStudioOnline-AE3C29BC36C640C583D4CC179E75430A", nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}

	// res, err := client.Get(ctx,
	// 	"VisualStudioOnline-AE3C29BC36C640C583D4CC179E75430A",
	// 	"SFApp Forms - Regression - Build",
	// 	nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Response result: %#v\n", res.PipelinesClientGetResult)
}

func main() {
	auth()
}
