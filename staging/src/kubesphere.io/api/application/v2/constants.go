/*
 * Copyright 2024 the KubeSphere Authors.
 * Please refer to the LICENSE file in the root directory of the project.
 * https://github.com/kubesphere/kubesphere/blob/master/LICENSE
 */

package v2

const (
	RepoIDLabelKey              = "application.kubesphere.io/repo-name"
	ReqUserAnnotationKey        = "application.kubesphere.io/req-user"
	AppIDLabelKey               = "application.kubesphere.io/app-id"
	AppOriginalNameLabelKey     = "application.kubesphere.io/app-originalName"
	AppVersionIDLabelKey        = "application.kubesphere.io/appversion-id"
	AppTypeLabelKey             = "application.kubesphere.io/app-type"
	HasCrdLabelKey              = "application.kubesphere.io/hascrd"
	AppStoreLabelKey            = "application.kubesphere.io/app-store"
	TimeoutRecheck              = "application.kubesphere.io/timeout-recheck"
	AppCategoryNameKey          = "application.kubesphere.io/app-category-name"
	LatestAppVersionKey         = "application.kubesphere.io/latest-app-version"
	AppMaintainersKey           = "application.kubesphere.io/app-maintainers"
	AppReleaseReferenceLabelKey = "application.kubesphere.io/app-release-name"
	UncategorizedCategoryID     = "kubesphere-app-uncategorized"
	StatusActive                = "active"
	StatusSuccessful            = "successful"
	StatusCreating              = "creating"
	StatusSyncing               = "syncing"
	StatusManualTrigger         = "manualTrigger"
	StatusDeleting              = "deleting"
	StatusUpgrading             = "upgrading"
	StatusClusterDeleted        = "clusterDeleted"
	StatusFailed                = "failed"
	StatusTimeout               = "timeout"
	StatusDeployFailed          = "deployFailed"
	StatusCreated               = "created"
	StatusUpgraded              = "upgraded"
	StatusNosync                = "nosync"
	AppTypeHelm                 = "helm"
	AppTypeYaml                 = "yaml"
	AppTypeEdge                 = "edge"
	BinaryKey                   = "BinaryKey"
	UploadRepoKey               = "upload"
	MaxNumOfVersions            = 10
	MaxImageWidth               = 128
	ApplicationNamespace        = "extension-openpitrix"
	StoreCleanFinalizer         = "storeCleanFinalizer.application.kubesphere.io"
	CleanupFinalizer            = "application.kubesphere.io/cleanup"
	SystemWorkspace             = "system-workspace"
	// App review status: draft, submitted, passed, rejected, suspended, active
	ReviewStatusDraft     = "draft"
	ReviewStatusSubmitted = "submitted"
	ReviewStatusPassed    = "passed"
	ReviewStatusRejected  = "rejected"
	ReviewStatusSuspended = "suspended"
	ReviewStatusActive    = "active"
)
