// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: b19c730e3a5c747d9055c95884b5c0310f7f2f16

package appservicecertificateorders

import original "github.com/Azure/azure-sdk-for-go/service/web/management/2015-08-01/appServiceCertificateOrders"

type GroupClient = original.GroupClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type CertificateOrderActionType = original.CertificateOrderActionType

const (
	CertificateExpirationWarning	CertificateOrderActionType	= original.CertificateExpirationWarning
	CertificateExpired		CertificateOrderActionType	= original.CertificateExpired
	CertificateIssued		CertificateOrderActionType	= original.CertificateIssued
	CertificateOrderCanceled	CertificateOrderActionType	= original.CertificateOrderCanceled
	CertificateOrderCreated		CertificateOrderActionType	= original.CertificateOrderCreated
	CertificateRevoked		CertificateOrderActionType	= original.CertificateRevoked
	DomainValidationComplete	CertificateOrderActionType	= original.DomainValidationComplete
	FraudCleared			CertificateOrderActionType	= original.FraudCleared
	FraudDetected			CertificateOrderActionType	= original.FraudDetected
	FraudDocumentationRequired	CertificateOrderActionType	= original.FraudDocumentationRequired
	OrgNameChange			CertificateOrderActionType	= original.OrgNameChange
	OrgValidationComplete		CertificateOrderActionType	= original.OrgValidationComplete
	SanDrop				CertificateOrderActionType	= original.SanDrop
	Unknown				CertificateOrderActionType	= original.Unknown
)

type CertificateOrderStatus = original.CertificateOrderStatus

const (
	Canceled		CertificateOrderStatus	= original.Canceled
	Denied			CertificateOrderStatus	= original.Denied
	Expired			CertificateOrderStatus	= original.Expired
	Issued			CertificateOrderStatus	= original.Issued
	NotSubmitted		CertificateOrderStatus	= original.NotSubmitted
	Pendingissuance		CertificateOrderStatus	= original.Pendingissuance
	PendingRekey		CertificateOrderStatus	= original.PendingRekey
	Pendingrevocation	CertificateOrderStatus	= original.Pendingrevocation
	Revoked			CertificateOrderStatus	= original.Revoked
	Unused			CertificateOrderStatus	= original.Unused
)

type CertificateProductType = original.CertificateProductType

const (
	StandardDomainValidatedSsl		CertificateProductType	= original.StandardDomainValidatedSsl
	StandardDomainValidatedWildCardSsl	CertificateProductType	= original.StandardDomainValidatedWildCardSsl
)

type KeyVaultSecretStatus = original.KeyVaultSecretStatus

const (
	KeyVaultSecretStatusAzureServiceUnauthorizedToAccessKeyVault	KeyVaultSecretStatus	= original.KeyVaultSecretStatusAzureServiceUnauthorizedToAccessKeyVault
	KeyVaultSecretStatusCertificateOrderFailed			KeyVaultSecretStatus	= original.KeyVaultSecretStatusCertificateOrderFailed
	KeyVaultSecretStatusExternalPrivateKey				KeyVaultSecretStatus	= original.KeyVaultSecretStatusExternalPrivateKey
	KeyVaultSecretStatusInitialized					KeyVaultSecretStatus	= original.KeyVaultSecretStatusInitialized
	KeyVaultSecretStatusKeyVaultDoesNotExist			KeyVaultSecretStatus	= original.KeyVaultSecretStatusKeyVaultDoesNotExist
	KeyVaultSecretStatusKeyVaultSecretDoesNotExist			KeyVaultSecretStatus	= original.KeyVaultSecretStatusKeyVaultSecretDoesNotExist
	KeyVaultSecretStatusOperationNotPermittedOnKeyVault		KeyVaultSecretStatus	= original.KeyVaultSecretStatusOperationNotPermittedOnKeyVault
	KeyVaultSecretStatusSucceeded					KeyVaultSecretStatus	= original.KeyVaultSecretStatusSucceeded
	KeyVaultSecretStatusUnknown					KeyVaultSecretStatus	= original.KeyVaultSecretStatusUnknown
	KeyVaultSecretStatusUnknownError				KeyVaultSecretStatus	= original.KeyVaultSecretStatusUnknownError
	KeyVaultSecretStatusWaitingOnCertificateOrder			KeyVaultSecretStatus	= original.KeyVaultSecretStatusWaitingOnCertificateOrder
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled	ProvisioningState	= original.ProvisioningStateCanceled
	ProvisioningStateDeleting	ProvisioningState	= original.ProvisioningStateDeleting
	ProvisioningStateFailed		ProvisioningState	= original.ProvisioningStateFailed
	ProvisioningStateInProgress	ProvisioningState	= original.ProvisioningStateInProgress
	ProvisioningStateSucceeded	ProvisioningState	= original.ProvisioningStateSucceeded
)

type AppServiceCertificate = original.AppServiceCertificate
type AppServiceCertificateCollection = original.AppServiceCertificateCollection
type AppServiceCertificateOrder = original.AppServiceCertificateOrder
type AppServiceCertificateOrderProperties = original.AppServiceCertificateOrderProperties
type AppServiceCertificateResource = original.AppServiceCertificateResource
type CertificateDetails = original.CertificateDetails
type CertificateEmail = original.CertificateEmail
type CertificateEmailProperties = original.CertificateEmailProperties
type CertificateOrderAction = original.CertificateOrderAction
type CertificateOrderActionProperties = original.CertificateOrderActionProperties
type Collection = original.Collection
type ListCertificateEmail = original.ListCertificateEmail
type ListCertificateOrderAction = original.ListCertificateOrderAction
type NameIdentifier = original.NameIdentifier
type ReissueCertificateOrderRequest = original.ReissueCertificateOrderRequest
type ReissueCertificateOrderRequestProperties = original.ReissueCertificateOrderRequestProperties
type RenewCertificateOrderRequest = original.RenewCertificateOrderRequest
type RenewCertificateOrderRequestProperties = original.RenewCertificateOrderRequestProperties
type Resource = original.Resource
type SiteSeal = original.SiteSeal
type SiteSealRequest = original.SiteSealRequest

func NewGroupClient(subscriptionID string) GroupClient {
	return original.NewGroupClient(subscriptionID)
}
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}