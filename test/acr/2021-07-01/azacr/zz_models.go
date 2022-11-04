//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azacr

import "time"

type AccessToken struct {
	// The access token for performing authenticated requests
	AccessToken *string `json:"access_token,omitempty"`
}

// Annotations - Additional information provided through arbitrary metadata.
type Annotations struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// Contact details of the people or organization responsible for the image.
	Authors *string `json:"org.opencontainers.image.authors,omitempty"`

	// Date and time on which the image was built (string, date-time as defined by https://tools.ietf.org/html/rfc3339#section-5.6)
	Created *time.Time `json:"org.opencontainers.image.created,omitempty"`

	// Human-readable description of the software packaged in the image
	Description *string `json:"org.opencontainers.image.description,omitempty"`

	// URL to get documentation on the image.
	Documentation *string `json:"org.opencontainers.image.documentation,omitempty"`

	// License(s) under which contained software is distributed as an SPDX License Expression.
	Licenses *string `json:"org.opencontainers.image.licenses,omitempty"`

	// Name of the reference for a target.
	Name *string `json:"org.opencontainers.image.ref.name,omitempty"`

	// Source control revision identifier for the packaged software.
	Revision *string `json:"org.opencontainers.image.revision,omitempty"`

	// URL to get source code for building the image.
	Source *string `json:"org.opencontainers.image.source,omitempty"`

	// Human-readable title of the image
	Title *string `json:"org.opencontainers.image.title,omitempty"`

	// URL to find more information on the image.
	URL *string `json:"org.opencontainers.image.url,omitempty"`

	// Name of the distributing entity, organization or individual.
	Vendor *string `json:"org.opencontainers.image.vendor,omitempty"`

	// Version of the packaged software. The version MAY match a label or tag in the source code repository, may also be Semantic
	// versioning-compatible
	Version *string `json:"org.opencontainers.image.version,omitempty"`
}

// ArtifactManifestPlatform - The artifact's platform, consisting of operating system and architecture.
type ArtifactManifestPlatform struct {
	// READ-ONLY; Manifest digest
	Digest *string `json:"digest,omitempty" azure:"ro"`

	// READ-ONLY; CPU architecture
	Architecture *ArtifactArchitecture `json:"architecture,omitempty" azure:"ro"`

	// READ-ONLY; Operating system
	OperatingSystem *ArtifactOperatingSystem `json:"os,omitempty" azure:"ro"`
}

// ArtifactManifestProperties - Manifest attributes details
type ArtifactManifestProperties struct {
	// READ-ONLY; Manifest attributes
	Manifest *ManifestAttributesBase `json:"manifest,omitempty" azure:"ro"`

	// READ-ONLY; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string `json:"registry,omitempty" azure:"ro"`

	// READ-ONLY; Repository name
	RepositoryName *string `json:"imageName,omitempty" azure:"ro"`
}

// ArtifactTagProperties - Tag attributes
type ArtifactTagProperties struct {
	// READ-ONLY; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string `json:"registry,omitempty" azure:"ro"`

	// READ-ONLY; Image name
	RepositoryName *string `json:"imageName,omitempty" azure:"ro"`

	// READ-ONLY; List of tag attribute details
	Tag *TagAttributesBase `json:"tag,omitempty" azure:"ro"`
}

// AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenOptions contains the optional parameters for the authenticationClient.ExchangeAADAccessTokenForAcrRefreshToken
// method.
type AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenOptions struct {
	// AAD access token, mandatory when granttype is accesstokenrefreshtoken or access_token.
	AccessToken *string
	// AAD refresh token, mandatory when granttype is accesstokenrefreshtoken or refresh_token
	RefreshToken *string
	// AAD tenant associated to the AAD credentials.
	Tenant *string
}

// AuthenticationClientExchangeAcrRefreshTokenForAcrAccessTokenOptions contains the optional parameters for the authenticationClient.ExchangeAcrRefreshTokenForAcrAccessToken
// method.
type AuthenticationClientExchangeAcrRefreshTokenForAcrAccessTokenOptions struct {
	// Grant type is expected to be refresh_token
	GrantType *TokenGrantType
}

// AuthenticationClientGetAcrAccessTokenFromLoginOptions contains the optional parameters for the authenticationClient.GetAcrAccessTokenFromLogin
// method.
type AuthenticationClientGetAcrAccessTokenFromLoginOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientCancelUploadOptions contains the optional parameters for the containerRegistryBlobClient.CancelUpload
// method.
type ContainerRegistryBlobClientCancelUploadOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientCheckBlobExistsOptions contains the optional parameters for the containerRegistryBlobClient.CheckBlobExists
// method.
type ContainerRegistryBlobClientCheckBlobExistsOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientCheckChunkExistsOptions contains the optional parameters for the containerRegistryBlobClient.CheckChunkExists
// method.
type ContainerRegistryBlobClientCheckChunkExistsOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientCompleteUploadOptions contains the optional parameters for the containerRegistryBlobClient.CompleteUpload
// method.
type ContainerRegistryBlobClientCompleteUploadOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientDeleteBlobOptions contains the optional parameters for the containerRegistryBlobClient.DeleteBlob
// method.
type ContainerRegistryBlobClientDeleteBlobOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientGetBlobOptions contains the optional parameters for the containerRegistryBlobClient.GetBlob
// method.
type ContainerRegistryBlobClientGetBlobOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientGetChunkOptions contains the optional parameters for the containerRegistryBlobClient.GetChunk
// method.
type ContainerRegistryBlobClientGetChunkOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientGetUploadStatusOptions contains the optional parameters for the containerRegistryBlobClient.GetUploadStatus
// method.
type ContainerRegistryBlobClientGetUploadStatusOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientMountBlobOptions contains the optional parameters for the containerRegistryBlobClient.MountBlob
// method.
type ContainerRegistryBlobClientMountBlobOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientStartUploadOptions contains the optional parameters for the containerRegistryBlobClient.StartUpload
// method.
type ContainerRegistryBlobClientStartUploadOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientUploadChunkOptions contains the optional parameters for the containerRegistryBlobClient.UploadChunk
// method.
type ContainerRegistryBlobClientUploadChunkOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientCheckDockerV2SupportOptions contains the optional parameters for the containerRegistryClient.CheckDockerV2Support
// method.
type ContainerRegistryClientCheckDockerV2SupportOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientCreateManifestOptions contains the optional parameters for the containerRegistryClient.CreateManifest
// method.
type ContainerRegistryClientCreateManifestOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientDeleteManifestOptions contains the optional parameters for the containerRegistryClient.DeleteManifest
// method.
type ContainerRegistryClientDeleteManifestOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientDeleteRepositoryOptions contains the optional parameters for the containerRegistryClient.DeleteRepository
// method.
type ContainerRegistryClientDeleteRepositoryOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientDeleteTagOptions contains the optional parameters for the containerRegistryClient.DeleteTag method.
type ContainerRegistryClientDeleteTagOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientGetManifestOptions contains the optional parameters for the containerRegistryClient.GetManifest
// method.
type ContainerRegistryClientGetManifestOptions struct {
	// Accept header string delimited by comma. For example, application/vnd.docker.distribution.manifest.v2+json
	Accept *string
}

// ContainerRegistryClientGetManifestPropertiesOptions contains the optional parameters for the containerRegistryClient.GetManifestProperties
// method.
type ContainerRegistryClientGetManifestPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientGetManifestsOptions contains the optional parameters for the containerRegistryClient.GetManifests
// method.
type ContainerRegistryClientGetManifestsOptions struct {
	// Query parameter for the last item in previous query. Result set will include values lexically after last.
	Last *string
	// query parameter for max number of items
	N *int32
	// orderby query parameter
	Orderby *string
}

// ContainerRegistryClientGetPropertiesOptions contains the optional parameters for the containerRegistryClient.GetProperties
// method.
type ContainerRegistryClientGetPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientGetRepositoriesOptions contains the optional parameters for the containerRegistryClient.GetRepositories
// method.
type ContainerRegistryClientGetRepositoriesOptions struct {
	// Query parameter for the last item in previous query. Result set will include values lexically after last.
	Last *string
	// query parameter for max number of items
	N *int32
}

// ContainerRegistryClientGetTagPropertiesOptions contains the optional parameters for the containerRegistryClient.GetTagProperties
// method.
type ContainerRegistryClientGetTagPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientGetTagsOptions contains the optional parameters for the containerRegistryClient.GetTags method.
type ContainerRegistryClientGetTagsOptions struct {
	// filter by digest
	Digest *string
	// Query parameter for the last item in previous query. Result set will include values lexically after last.
	Last *string
	// query parameter for max number of items
	N *int32
	// orderby query parameter
	Orderby *string
}

// ContainerRegistryClientUpdateManifestPropertiesOptions contains the optional parameters for the containerRegistryClient.UpdateManifestProperties
// method.
type ContainerRegistryClientUpdateManifestPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientUpdatePropertiesOptions contains the optional parameters for the containerRegistryClient.UpdateProperties
// method.
type ContainerRegistryClientUpdatePropertiesOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientUpdateTagAttributesOptions contains the optional parameters for the containerRegistryClient.UpdateTagAttributes
// method.
type ContainerRegistryClientUpdateTagAttributesOptions struct {
	// placeholder for future optional parameters
}

// ContainerRepositoryProperties - Properties of this repository.
type ContainerRepositoryProperties struct {
	// REQUIRED; Writeable properties of the resource
	ChangeableAttributes *RepositoryWriteableProperties `json:"changeableAttributes,omitempty"`

	// READ-ONLY; Image created time
	CreatedOn *time.Time `json:"createdTime,omitempty" azure:"ro"`

	// READ-ONLY; Image last update time
	LastUpdatedOn *time.Time `json:"lastUpdateTime,omitempty" azure:"ro"`

	// READ-ONLY; Number of the manifests
	ManifestCount *int32 `json:"manifestCount,omitempty" azure:"ro"`

	// READ-ONLY; Image name
	Name *string `json:"imageName,omitempty" azure:"ro"`

	// READ-ONLY; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string `json:"registry,omitempty" azure:"ro"`

	// READ-ONLY; Number of the tags
	TagCount *int32 `json:"tagCount,omitempty" azure:"ro"`
}

// DeleteRepositoryResult - Deleted repository
type DeleteRepositoryResult struct {
	// READ-ONLY; SHA of the deleted image
	DeletedManifests []*string `json:"manifestsDeleted,omitempty" azure:"ro"`

	// READ-ONLY; Tag of the deleted image
	DeletedTags []*string `json:"tagsDeleted,omitempty" azure:"ro"`
}

// Descriptor - Docker V2 image layer descriptor including config and layers
type Descriptor struct {
	// Additional information provided through arbitrary metadata.
	Annotations *Annotations `json:"annotations,omitempty"`

	// Layer digest
	Digest *string `json:"digest,omitempty"`

	// Layer media type
	MediaType *string `json:"mediaType,omitempty"`

	// Layer size
	Size *int64 `json:"size,omitempty"`

	// Specifies a list of URIs from which this object may be downloaded.
	Urls []*string `json:"urls,omitempty"`
}

// ErrorInfo - Error information
type ErrorInfo struct {
	// Error code
	Code *string `json:"code,omitempty"`

	// Error details
	Detail any `json:"detail,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`
}

// Errors - Acr error response describing why the operation failed
type Errors struct {
	// Array of detailed error
	Errors []*ErrorInfo `json:"errors,omitempty"`
}

// FsLayer - Image layer information
type FsLayer struct {
	// SHA of an image layer
	BlobSum *string `json:"blobSum,omitempty"`
}

// History - A list of unstructured historical data for v1 compatibility
type History struct {
	// The raw v1 compatibility information
	V1Compatibility *string `json:"v1Compatibility,omitempty"`
}

// ImageSignature - Signature of a signed manifest
type ImageSignature struct {
	// A JSON web signature
	Header *JWK `json:"header,omitempty"`

	// The signed protected header
	Protected *string `json:"protected,omitempty"`

	// A signature for the image manifest, signed by a libtrust private key
	Signature *string `json:"signature,omitempty"`
}

// JWK - A JSON web signature
type JWK struct {
	// The algorithm used to sign or encrypt the JWT
	Alg *string `json:"alg,omitempty"`

	// JSON web key parameter
	Jwk *JWKHeader `json:"jwk,omitempty"`
}

// JWKHeader - JSON web key parameter
type JWKHeader struct {
	// crv value
	Crv *string `json:"crv,omitempty"`

	// kid value
	Kid *string `json:"kid,omitempty"`

	// kty value
	Kty *string `json:"kty,omitempty"`

	// x value
	X *string `json:"x,omitempty"`

	// y value
	Y *string `json:"y,omitempty"`
}

// Manifest - Returns the requested manifest file
type Manifest struct {
	// Schema version
	SchemaVersion *int32 `json:"schemaVersion,omitempty"`
}

// ManifestAttributesBase - Manifest details
type ManifestAttributesBase struct {
	// READ-ONLY; Created time
	CreatedOn *time.Time `json:"createdTime,omitempty" azure:"ro"`

	// READ-ONLY; Manifest
	Digest *string `json:"digest,omitempty" azure:"ro"`

	// READ-ONLY; Last update time
	LastUpdatedOn *time.Time `json:"lastUpdateTime,omitempty" azure:"ro"`

	// Writeable properties of the resource
	ChangeableAttributes *ManifestWriteableProperties `json:"changeableAttributes,omitempty"`

	// Config blob media type
	ConfigMediaType *string `json:"configMediaType,omitempty"`

	// READ-ONLY; CPU architecture
	Architecture *ArtifactArchitecture `json:"architecture,omitempty" azure:"ro"`

	// READ-ONLY; Operating system
	OperatingSystem *ArtifactOperatingSystem `json:"os,omitempty" azure:"ro"`

	// READ-ONLY; List of artifacts that are referenced by this manifest list, with information about the platform each supports.
	// This list will be empty if this is a leaf manifest and not a manifest list.
	RelatedArtifacts []*ArtifactManifestPlatform `json:"references,omitempty" azure:"ro"`

	// READ-ONLY; Image size
	Size *int64 `json:"imageSize,omitempty" azure:"ro"`

	// READ-ONLY; List of tags
	Tags []*string `json:"tags,omitempty" azure:"ro"`
}

// ManifestAttributesManifest - List of manifest attributes
type ManifestAttributesManifest struct {
	// List of manifest attributes details
	References []*ArtifactManifestPlatform `json:"references,omitempty"`
}

// ManifestList - Returns the requested Docker multi-arch-manifest file
type ManifestList struct {
	// List of V2 image layer information
	Manifests []*ManifestListAttributes `json:"manifests,omitempty"`

	// Media type for this Manifest
	MediaType *string `json:"mediaType,omitempty"`

	// Schema version
	SchemaVersion *int32 `json:"schemaVersion,omitempty"`
}

type ManifestListAttributes struct {
	// The digest of the content, as defined by the Registry V2 HTTP API Specification
	Digest *string `json:"digest,omitempty"`

	// The MIME type of the referenced object. This will generally be application/vnd.docker.image.manifest.v2+json, but it could
	// also be application/vnd.docker.image.manifest.v1+json
	MediaType *string `json:"mediaType,omitempty"`

	// The platform object describes the platform which the image in the manifest runs on. A full list of valid operating system
	// and architecture values are listed in the Go language documentation for $GOOS
	// and $GOARCH
	Platform *Platform `json:"platform,omitempty"`

	// The size in bytes of the object
	Size *int64 `json:"size,omitempty"`
}

// ManifestWrapper - Returns the requested manifest file
type ManifestWrapper struct {
	// (OCI, OCIIndex) Additional metadata
	Annotations *Annotations `json:"annotations,omitempty"`

	// (V1) CPU architecture
	Architecture *string `json:"architecture,omitempty"`

	// (V2, OCI) Image config descriptor
	Config *Descriptor `json:"config,omitempty"`

	// (V1) List of layer information
	FsLayers []*FsLayer `json:"fsLayers,omitempty"`

	// (V1) Image history
	History []*History `json:"history,omitempty"`

	// (V2, OCI) List of V2 image layer information
	Layers []*Descriptor `json:"layers,omitempty"`

	// (ManifestList, OCIIndex) List of V2 image layer information
	Manifests []*ManifestListAttributes `json:"manifests,omitempty"`

	// Media type for this Manifest
	MediaType *string `json:"mediaType,omitempty"`

	// (V1) Image name
	Name *string `json:"name,omitempty"`

	// Schema version
	SchemaVersion *int32 `json:"schemaVersion,omitempty"`

	// (V1) Image signature
	Signatures []*ImageSignature `json:"signatures,omitempty"`

	// (V1) Image tag
	Tag *string `json:"tag,omitempty"`
}

// ManifestWriteableProperties - Changeable attributes
type ManifestWriteableProperties struct {
	// Delete enabled
	CanDelete *bool `json:"deleteEnabled,omitempty"`

	// List enabled
	CanList *bool `json:"listEnabled,omitempty"`

	// Read enabled
	CanRead *bool `json:"readEnabled,omitempty"`

	// Write enabled
	CanWrite *bool `json:"writeEnabled,omitempty"`
}

// Manifests - Manifest attributes
type Manifests struct {
	Link *string `json:"link,omitempty"`

	// List of manifests
	Manifests []*ManifestAttributesBase `json:"manifests,omitempty"`

	// Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string `json:"registry,omitempty"`

	// Image name
	Repository *string `json:"imageName,omitempty"`
}

// OCIIndex - Returns the requested OCI index file
type OCIIndex struct {
	// Additional information provided through arbitrary metadata.
	Annotations *Annotations `json:"annotations,omitempty"`

	// List of OCI image layer information
	Manifests []*ManifestListAttributes `json:"manifests,omitempty"`

	// Schema version
	SchemaVersion *int32 `json:"schemaVersion,omitempty"`
}

// OCIManifest - Returns the requested OCI Manifest file
type OCIManifest struct {
	// Additional information provided through arbitrary metadata.
	Annotations *Annotations `json:"annotations,omitempty"`

	// V2 image config descriptor
	Config *Descriptor `json:"config,omitempty"`

	// List of V2 image layer information
	Layers []*Descriptor `json:"layers,omitempty"`

	// Schema version
	SchemaVersion *int32 `json:"schemaVersion,omitempty"`
}

type Paths108HwamOauth2ExchangePostRequestbodyContentApplicationXWwwFormUrlencodedSchema struct {
	// REQUIRED; Can take a value of accesstokenrefreshtoken, or accesstoken, or refresh_token
	GrantType *PostContentSchemaGrantType `json:"grant_type,omitempty"`

	// REQUIRED; Indicates the name of your Azure container registry.
	Service *string `json:"service,omitempty"`

	// AAD access token, mandatory when granttype is accesstokenrefreshtoken or access_token.
	AADAccessToken *string `json:"access_token,omitempty"`

	// AAD refresh token, mandatory when granttype is accesstokenrefreshtoken or refresh_token
	RefreshToken *string `json:"refresh_token,omitempty"`

	// AAD tenant associated to the AAD credentials.
	Tenant *string `json:"tenant,omitempty"`
}

type PathsV3R3RxOauth2TokenPostRequestbodyContentApplicationXWwwFormUrlencodedSchema struct {
	// REQUIRED; Must be a valid ACR refresh token
	AcrRefreshToken *string `json:"refresh_token,omitempty"`

	// REQUIRED; Grant type is expected to be refresh_token
	GrantType *TokenGrantType `json:"grant_type,omitempty"`

	// REQUIRED; Which is expected to be a valid scope, and can be specified more than once for multiple scope requests. You obtained
	// this from the Www-Authenticate response header from the challenge.
	Scope *string `json:"scope,omitempty"`

	// REQUIRED; Indicates the name of your Azure container registry.
	Service *string `json:"service,omitempty"`
}

// Platform - The platform object describes the platform which the image in the manifest runs on. A full list of valid operating
// system and architecture values are listed in the Go language documentation for $GOOS
// and $GOARCH
type Platform struct {
	// Specifies the CPU architecture, for example amd64 or ppc64le.
	Architecture *string `json:"architecture,omitempty"`

	// The optional features field specifies an array of strings, each listing a required CPU feature (for example sse4 or aes
	Features []*string `json:"features,omitempty"`

	// The os field specifies the operating system, for example linux or windows.
	OS *string `json:"os,omitempty"`

	// The optional os.features field specifies an array of strings, each listing a required OS feature (for example on Windows
	// win32k
	OSFeatures []*string `json:"os.features,omitempty"`

	// The optional os.version field specifies the operating system version, for example 10.0.10586.
	OSVersion *string `json:"os.version,omitempty"`

	// The optional variant field specifies a variant of the CPU, for example armv6l to specify a particular CPU variant of the
	// ARM CPU.
	Variant *string `json:"variant,omitempty"`
}

type RefreshToken struct {
	// The refresh token to be used for generating access tokens
	RefreshToken *string `json:"refresh_token,omitempty"`
}

// Repositories - List of repositories
type Repositories struct {
	Link *string `json:"link,omitempty"`

	// Repository names
	Repositories []*string `json:"repositories,omitempty"`
}

// RepositoryTags - Result of the request to list tags of the image
type RepositoryTags struct {
	// Name of the image
	Name *string `json:"name,omitempty"`

	// List of tags
	Tags []*string `json:"tags,omitempty"`
}

// RepositoryWriteableProperties - Changeable attributes for Repository
type RepositoryWriteableProperties struct {
	// Delete enabled
	CanDelete *bool `json:"deleteEnabled,omitempty"`

	// List enabled
	CanList *bool `json:"listEnabled,omitempty"`

	// Read enabled
	CanRead *bool `json:"readEnabled,omitempty"`

	// Write enabled
	CanWrite *bool `json:"writeEnabled,omitempty"`
}

// TagAttributesBase - Tag attribute details
type TagAttributesBase struct {
	// REQUIRED; Writeable properties of the resource
	ChangeableAttributes *TagWriteableProperties `json:"changeableAttributes,omitempty"`

	// Is signed
	Signed *bool `json:"signed,omitempty"`

	// READ-ONLY; Tag created time
	CreatedOn *time.Time `json:"createdTime,omitempty" azure:"ro"`

	// READ-ONLY; Tag digest
	Digest *string `json:"digest,omitempty" azure:"ro"`

	// READ-ONLY; Tag last update time
	LastUpdatedOn *time.Time `json:"lastUpdateTime,omitempty" azure:"ro"`

	// READ-ONLY; Tag name
	Name *string `json:"name,omitempty" azure:"ro"`
}

// TagAttributesTag - Tag
type TagAttributesTag struct {
	// SignatureRecord value
	SignatureRecord *string `json:"signatureRecord,omitempty"`
}

// TagList - List of tag details
type TagList struct {
	// REQUIRED; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string `json:"registry,omitempty"`

	// REQUIRED; Image name
	Repository *string `json:"imageName,omitempty"`

	// REQUIRED; List of tag attribute details
	TagAttributeBases []*TagAttributesBase `json:"tags,omitempty"`
	Link              *string              `json:"link,omitempty"`
}

// TagWriteableProperties - Changeable attributes
type TagWriteableProperties struct {
	// Delete enabled
	CanDelete *bool `json:"deleteEnabled,omitempty"`

	// List enabled
	CanList *bool `json:"listEnabled,omitempty"`

	// Read enabled
	CanRead *bool `json:"readEnabled,omitempty"`

	// Write enabled
	CanWrite *bool `json:"writeEnabled,omitempty"`
}

// V1Manifest - Returns the requested V1 manifest file
type V1Manifest struct {
	// CPU architecture
	Architecture *string `json:"architecture,omitempty"`

	// List of layer information
	FsLayers []*FsLayer `json:"fsLayers,omitempty"`

	// Image history
	History []*History `json:"history,omitempty"`

	// Image name
	Name *string `json:"name,omitempty"`

	// Schema version
	SchemaVersion *int32 `json:"schemaVersion,omitempty"`

	// Image signature
	Signatures []*ImageSignature `json:"signatures,omitempty"`

	// Image tag
	Tag *string `json:"tag,omitempty"`
}

// V2Manifest - Returns the requested Docker V2 Manifest file
type V2Manifest struct {
	// V2 image config descriptor
	Config *Descriptor `json:"config,omitempty"`

	// List of V2 image layer information
	Layers []*Descriptor `json:"layers,omitempty"`

	// Media type for this Manifest
	MediaType *string `json:"mediaType,omitempty"`

	// Schema version
	SchemaVersion *int32 `json:"schemaVersion,omitempty"`
}