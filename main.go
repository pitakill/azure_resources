// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-09-02 21:19:42.711677 -0500 CDT m=+0.000309641
package azure_resources

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

type resource interface {
	GetProperties() ([]byte, error)
}

func GetAllByGroupName(subscriptionID, groupName string) []resource {
	client := resources.NewClient(subscriptionID)

	authorizer, err := auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
	if err == nil {
		client.Authorizer = authorizer
	}

	results, err := client.ListByResourceGroup(context.Background(), groupName, "", "", nil)
	if err != nil {
		log.Fatalln(err)
	}

	var resources []resource

	for _, resource := range results.Values() {
		switch *resource.Type {
		case "Microsoft.Compute/virtualMachines":
			r := &ComputeVirtualMachines{
				subscriptionID: subscriptionID,
				groupName:      groupName,
				resourceName:   resource.Name,
				authorizer:     authorizer,
			}
			resources = append(resources, r)

		case "Microsoft.Compute/virtualMachines/Extensions":
			r := &ComputeVirtualMachineExtensions{
				subscriptionID: subscriptionID,
				groupName:      groupName,
				resourceName:   resource.Name,
				authorizer:     authorizer,
			}
			resources = append(resources, r)

		case "Microsoft.Network/publicIPAddresses":
			r := &NetworkPublicIPAddresses{
				subscriptionID: subscriptionID,
				groupName:      groupName,
				resourceName:   resource.Name,
				authorizer:     authorizer,
			}
			resources = append(resources, r)

		case "Microsoft.Compute/disks":
			r := &ComputeDisks{
				subscriptionID: subscriptionID,
				groupName:      groupName,
				resourceName:   resource.Name,
				authorizer:     authorizer,
			}
			resources = append(resources, r)

		case "Microsoft.Network/virtualNetworks":
			r := &NetworkVirtualNetworks{
				subscriptionID: subscriptionID,
				groupName:      groupName,
				resourceName:   resource.Name,
				authorizer:     authorizer,
			}
			resources = append(resources, r)
		}
	}

	return resources
}
