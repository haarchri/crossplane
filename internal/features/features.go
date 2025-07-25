/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package features defines Crossplane feature flags.
package features

import "github.com/crossplane/crossplane-runtime/pkg/feature"

// Alpha Feature flags.
const (
	// EnableAlphaDependencyVersionUpgrades enables alpha support for
	// upgrading the version of a package's dependencies when needed.
	EnableAlphaDependencyVersionUpgrades feature.Flag = "EnableAlphaDependencyVersionUpgrades"

	// EnableAlphaDependencyVersionDowngrades enables alpha support for
	// downgrading the version of a package's dependencies when needed.
	// It builds on EnableAlphaDependencyVersionUpgrades; both flags must
	// be set to enable downgrades.
	EnableAlphaDependencyVersionDowngrades feature.Flag = "EnableAlphaDependencyVersionDowngrades"

	// EnableAlphaSignatureVerification enables alpha support for verifying
	// the package signatures via ImageConfig API.
	EnableAlphaSignatureVerification feature.Flag = "EnableAlphaSignatureVerification"

	// EnableAlphaFunctionResponseCache enables alpha support for caching
	// composition function responses.
	EnableAlphaFunctionResponseCache feature.Flag = "EnableAlphaFunctionResponseCache"

	// EnableAlphaOperations enables alpha support for Operations, including
	// CronOperations and WatchOperations.
	EnableAlphaOperations feature.Flag = "EnableAlphaOperations"
)

// Beta Feature Flags.
const (
	// EnableBetaDeploymentRuntimeConfigs enables beta support for deployment
	// runtime configs. See the below design for more details.
	// https://github.com/crossplane/crossplane/blob/c2e206/design/one-pager-package-runtime-config.md
	EnableBetaDeploymentRuntimeConfigs feature.Flag = "EnableBetaDeploymentRuntimeConfigs"

	// EnableBetaUsages enables beta support for deletion ordering and
	// protection with Usage resource. See the below design for more details.
	// https://github.com/crossplane/crossplane/blob/19ea23/design/one-pager-generic-usage-type.md
	EnableBetaUsages feature.Flag = "EnableBetaUsages"

	// EnableBetaClaimSSA enables beta support for using server-side apply in
	// the claim controller. See the below issue for more details:
	// https://github.com/crossplane/crossplane/issues/4581
	EnableBetaClaimSSA feature.Flag = "EnableBetaClaimSSA"

	// EnableBetaRealtimeCompositions enables beta support for realtime
	// compositions, i.e. watching composed resources and reconciling
	// compositions immediately when any composed resource is updated.
	EnableBetaRealtimeCompositions feature.Flag = "EnableBetaRealtimeCompositions"
)
