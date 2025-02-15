//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/compute/armcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/stretchr/testify/suite"
)

type SampleTestSuite struct {
	suite.Suite

	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	armEndpoint       string
	fakeStepVar       string
	resourceName      string
	testPrefix        string
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *SampleTestSuite) SetupSuite() {
	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/compute/armcompute/testdata")

	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.armEndpoint = "https://management.azure.com"
	testsuite.fakeStepVar = "signalrswaggertest4"
	testsuite.resourceName = "signalrswaggertest4"
	testsuite.testPrefix = testutil.GenerateAlphaNumericID(testsuite.T(), "test", 6)
	testsuite.location = testutil.GetEnv("LOCATION", "westus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "00000000-00000000-00000000-00000000")
	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
	testsuite.Prepare()
}

func (testsuite *SampleTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestSampleTestSuite(t *testing.T) {
	suite.Run(t, new(SampleTestSuite))
}

func (testsuite *SampleTestSuite) Prepare() {
	var err error
	// From step Delete-proximity-placement-group
	fmt.Println("Call operation: ProximityPlacementGroups_Delete")
	proximityPlacementGroupsClient, err := armcompute.NewProximityPlacementGroupsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = proximityPlacementGroupsClient.Delete(testsuite.ctx, testsuite.resourceGroupName, testsuite.resourceName, nil)
	testsuite.Require().NoError(err)
}

// Microsoft.SignalRService/Basic_CRUD
func (testsuite *SampleTestSuite) TestScenario0() {
	fakeScenarioVar := "signalrswaggertest5"
	resourceName := testsuite.resourceName
	var err error
	// From step Generate_Unique_Name
	template := map[string]any{
		"$schema":        "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
		"contentVersion": "1.0.0.0",
		"outputs": map[string]any{
			"name": map[string]any{
				"type":  "string",
				"value": "[variables('name').value]",
			},
			"resourceName": map[string]any{
				"type":  "string",
				"value": "[variables('name').value]",
			},
		},
		"resources": []any{},
		"variables": map[string]any{
			"name": map[string]any{
				"type": "string",
				"metadata": map[string]any{
					"description": "Name of the SignalR service.",
				},
				"value": "[concat('sw',uniqueString(resourceGroup().id))]",
			},
		},
	}
	deployment := armresources.Deployment{
		Properties: &armresources.DeploymentProperties{
			Template: template,
			Mode:     to.Ptr(armresources.DeploymentModeIncremental),
		},
	}
	deploymentExtend, err := testutil.CreateDeployment(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName, "Generate_Unique_Name", &deployment)
	testsuite.Require().NoError(err)
	name = deploymentExtend.Properties.Outputs.(map[string]interface{})["name"].(map[string]interface{})["value"].(string)
	testsuite.resourceName = deploymentExtend.Properties.Outputs.(map[string]interface{})["resourceName"].(map[string]interface{})["value"].(string)

	// From step Create-or-Update-a-proximity-placement-group
	fmt.Println("Call operation: ProximityPlacementGroups_CreateOrUpdate")
	proximityPlacementGroupsClient, err := armcompute.NewProximityPlacementGroupsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	proximityPlacementGroupsClientCreateOrUpdateResponse, err := proximityPlacementGroupsClient.CreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.resourceName, armcompute.ProximityPlacementGroup{
		Location: to.Ptr(testsuite.location),
		Properties: &armcompute.ProximityPlacementGroupProperties{
			ProximityPlacementGroupType: to.Ptr(armcompute.ProximityPlacementGroupTypeStandard),
		},
	}, nil)
	testsuite.Require().NoError(err)
	fakeScenarioVar = *proximityPlacementGroupsClientCreateOrUpdateResponse.ID

	// From step Delete-proximity_placement_group
	fmt.Println("Call operation: ProximityPlacementGroups_Delete")
	_, err = proximityPlacementGroupsClient.Delete(testsuite.ctx, testsuite.resourceGroupName, testsuite.resourceName, nil)
	testsuite.Require().NoError(err)

	// From step Create_a_vm_with_Host_Encryption_using_encryptionAtHost_property
	fmt.Println("Call operation: VirtualMachines_CreateOrUpdate")
	virtualMachinesClient, err := armcompute.NewVirtualMachinesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	virtualMachinesClientCreateOrUpdateResponsePoller, err := virtualMachinesClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, "myVM", armcompute.VirtualMachine{
		Location: to.Ptr(testsuite.location),
		Plan: &armcompute.Plan{
			Name:      to.Ptr("signalrswaggertest6"),
			Product:   to.Ptr("windows-data-science-vm"),
			Publisher: to.Ptr("microsoft-ads"),
		},
		Properties: &armcompute.VirtualMachineProperties{
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardDS1V2),
			},
			NetworkProfile: &armcompute.NetworkProfile{
				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
					{
						ID: to.Ptr("/subscriptions/" + testsuite.subscriptionId + "/resourceGroups/" + testsuite.resourceGroupName + "/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
						Properties: &armcompute.NetworkInterfaceReferenceProperties{
							Primary: to.Ptr(true),
						},
					}},
			},
			OSProfile: &armcompute.OSProfile{
				AdminPassword: to.Ptr("{your-password}"),
				AdminUsername: to.Ptr("{your-username}"),
				ComputerName:  to.Ptr("myVM"),
			},
			SecurityProfile: &armcompute.SecurityProfile{
				EncryptionAtHost: to.Ptr(true),
			},
			StorageProfile: &armcompute.StorageProfile{
				ImageReference: &armcompute.ImageReference{
					Offer:     to.Ptr("windows-data-science-vm"),
					Publisher: to.Ptr(fakeScenarioVar),
					SKU:       to.Ptr("windows2016"),
					Version:   to.Ptr("latest"),
				},
				OSDisk: &armcompute.OSDisk{
					Name:         to.Ptr("myVMosdisk"),
					Caching:      to.Ptr(armcompute.CachingTypesReadOnly),
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
					},
				},
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, virtualMachinesClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)
}

// Microsoft.SignalRService/DeleteOnly
func (testsuite *SampleTestSuite) TestScenario1() {
	var err error
	// From step Delete_proximity_placement_group
	fmt.Println("Call operation: ProximityPlacementGroups_Delete")
	proximityPlacementGroupsClient, err := armcompute.NewProximityPlacementGroupsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = proximityPlacementGroupsClient.Delete(testsuite.ctx, testsuite.resourceGroupName, testsuite.resourceName, nil)
	testsuite.Require().NoError(err)
}
