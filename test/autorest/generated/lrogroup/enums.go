// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

// OperationResultStatus - The status of the request
type OperationResultStatus string

const (
	OperationResultStatusAccepted  OperationResultStatus = "Accepted"
	OperationResultStatusCanceled  OperationResultStatus = "canceled"
	OperationResultStatusCreated   OperationResultStatus = "Created"
	OperationResultStatusCreating  OperationResultStatus = "Creating"
	OperationResultStatusDeleted   OperationResultStatus = "Deleted"
	OperationResultStatusDeleting  OperationResultStatus = "Deleting"
	OperationResultStatusFailed    OperationResultStatus = "Failed"
	OperationResultStatusOk        OperationResultStatus = "OK"
	OperationResultStatusSucceeded OperationResultStatus = "Succeeded"
	OperationResultStatusUpdated   OperationResultStatus = "Updated"
	OperationResultStatusUpdating  OperationResultStatus = "Updating"
)

func PossibleOperationResultStatusValues() []OperationResultStatus {
	return []OperationResultStatus{OperationResultStatusAccepted, OperationResultStatusCanceled, OperationResultStatusCreated, OperationResultStatusCreating, OperationResultStatusDeleted, OperationResultStatusDeleting, OperationResultStatusFailed, OperationResultStatusOk, OperationResultStatusSucceeded, OperationResultStatusUpdated, OperationResultStatusUpdating}
}

func (c OperationResultStatus) ToPtr() *OperationResultStatus {
	return &c
}

type ProductPropertiesProvisioningStateValues string

const (
	ProductPropertiesProvisioningStateValuesAccepted  ProductPropertiesProvisioningStateValues = "Accepted"
	ProductPropertiesProvisioningStateValuesCanceled  ProductPropertiesProvisioningStateValues = "canceled"
	ProductPropertiesProvisioningStateValuesCreated   ProductPropertiesProvisioningStateValues = "Created"
	ProductPropertiesProvisioningStateValuesCreating  ProductPropertiesProvisioningStateValues = "Creating"
	ProductPropertiesProvisioningStateValuesDeleted   ProductPropertiesProvisioningStateValues = "Deleted"
	ProductPropertiesProvisioningStateValuesDeleting  ProductPropertiesProvisioningStateValues = "Deleting"
	ProductPropertiesProvisioningStateValuesFailed    ProductPropertiesProvisioningStateValues = "Failed"
	ProductPropertiesProvisioningStateValuesOk        ProductPropertiesProvisioningStateValues = "OK"
	ProductPropertiesProvisioningStateValuesSucceeded ProductPropertiesProvisioningStateValues = "Succeeded"
	ProductPropertiesProvisioningStateValuesUpdated   ProductPropertiesProvisioningStateValues = "Updated"
	ProductPropertiesProvisioningStateValuesUpdating  ProductPropertiesProvisioningStateValues = "Updating"
)

func PossibleProductPropertiesProvisioningStateValuesValues() []ProductPropertiesProvisioningStateValues {
	return []ProductPropertiesProvisioningStateValues{ProductPropertiesProvisioningStateValuesAccepted, ProductPropertiesProvisioningStateValuesCanceled, ProductPropertiesProvisioningStateValuesCreated, ProductPropertiesProvisioningStateValuesCreating, ProductPropertiesProvisioningStateValuesDeleted, ProductPropertiesProvisioningStateValuesDeleting, ProductPropertiesProvisioningStateValuesFailed, ProductPropertiesProvisioningStateValuesOk, ProductPropertiesProvisioningStateValuesSucceeded, ProductPropertiesProvisioningStateValuesUpdated, ProductPropertiesProvisioningStateValuesUpdating}
}

func (c ProductPropertiesProvisioningStateValues) ToPtr() *ProductPropertiesProvisioningStateValues {
	return &c
}

type SubProductPropertiesProvisioningStateValues string

const (
	SubProductPropertiesProvisioningStateValuesAccepted  SubProductPropertiesProvisioningStateValues = "Accepted"
	SubProductPropertiesProvisioningStateValuesCanceled  SubProductPropertiesProvisioningStateValues = "canceled"
	SubProductPropertiesProvisioningStateValuesCreated   SubProductPropertiesProvisioningStateValues = "Created"
	SubProductPropertiesProvisioningStateValuesCreating  SubProductPropertiesProvisioningStateValues = "Creating"
	SubProductPropertiesProvisioningStateValuesDeleted   SubProductPropertiesProvisioningStateValues = "Deleted"
	SubProductPropertiesProvisioningStateValuesDeleting  SubProductPropertiesProvisioningStateValues = "Deleting"
	SubProductPropertiesProvisioningStateValuesFailed    SubProductPropertiesProvisioningStateValues = "Failed"
	SubProductPropertiesProvisioningStateValuesOk        SubProductPropertiesProvisioningStateValues = "OK"
	SubProductPropertiesProvisioningStateValuesSucceeded SubProductPropertiesProvisioningStateValues = "Succeeded"
	SubProductPropertiesProvisioningStateValuesUpdated   SubProductPropertiesProvisioningStateValues = "Updated"
	SubProductPropertiesProvisioningStateValuesUpdating  SubProductPropertiesProvisioningStateValues = "Updating"
)

func PossibleSubProductPropertiesProvisioningStateValuesValues() []SubProductPropertiesProvisioningStateValues {
	return []SubProductPropertiesProvisioningStateValues{SubProductPropertiesProvisioningStateValuesAccepted, SubProductPropertiesProvisioningStateValuesCanceled, SubProductPropertiesProvisioningStateValuesCreated, SubProductPropertiesProvisioningStateValuesCreating, SubProductPropertiesProvisioningStateValuesDeleted, SubProductPropertiesProvisioningStateValuesDeleting, SubProductPropertiesProvisioningStateValuesFailed, SubProductPropertiesProvisioningStateValuesOk, SubProductPropertiesProvisioningStateValuesSucceeded, SubProductPropertiesProvisioningStateValuesUpdated, SubProductPropertiesProvisioningStateValuesUpdating}
}

func (c SubProductPropertiesProvisioningStateValues) ToPtr() *SubProductPropertiesProvisioningStateValues {
	return &c
}