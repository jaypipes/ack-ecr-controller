// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package repository

import (
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"

	svcapitypes "github.com/aws-controllers-k8s/ecr-controller/apis/v1alpha1"
)

var (
	_ = svcapitypes.GroupVersion
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.EncryptionConfiguration, b.ko.Spec.EncryptionConfiguration) {
		delta.Add("Spec.EncryptionConfiguration", a.ko.Spec.EncryptionConfiguration, b.ko.Spec.EncryptionConfiguration)
	} else if a.ko.Spec.EncryptionConfiguration != nil && b.ko.Spec.EncryptionConfiguration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.EncryptionConfiguration.EncryptionType, b.ko.Spec.EncryptionConfiguration.EncryptionType) {
			delta.Add("Spec.EncryptionConfiguration.EncryptionType", a.ko.Spec.EncryptionConfiguration.EncryptionType, b.ko.Spec.EncryptionConfiguration.EncryptionType)
		} else if a.ko.Spec.EncryptionConfiguration.EncryptionType != nil && b.ko.Spec.EncryptionConfiguration.EncryptionType != nil {
			if *a.ko.Spec.EncryptionConfiguration.EncryptionType != *b.ko.Spec.EncryptionConfiguration.EncryptionType {
				delta.Add("Spec.EncryptionConfiguration.EncryptionType", a.ko.Spec.EncryptionConfiguration.EncryptionType, b.ko.Spec.EncryptionConfiguration.EncryptionType)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.EncryptionConfiguration.KMSKey, b.ko.Spec.EncryptionConfiguration.KMSKey) {
			delta.Add("Spec.EncryptionConfiguration.KMSKey", a.ko.Spec.EncryptionConfiguration.KMSKey, b.ko.Spec.EncryptionConfiguration.KMSKey)
		} else if a.ko.Spec.EncryptionConfiguration.KMSKey != nil && b.ko.Spec.EncryptionConfiguration.KMSKey != nil {
			if *a.ko.Spec.EncryptionConfiguration.KMSKey != *b.ko.Spec.EncryptionConfiguration.KMSKey {
				delta.Add("Spec.EncryptionConfiguration.KMSKey", a.ko.Spec.EncryptionConfiguration.KMSKey, b.ko.Spec.EncryptionConfiguration.KMSKey)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ImageScanningConfiguration, b.ko.Spec.ImageScanningConfiguration) {
		delta.Add("Spec.ImageScanningConfiguration", a.ko.Spec.ImageScanningConfiguration, b.ko.Spec.ImageScanningConfiguration)
	} else if a.ko.Spec.ImageScanningConfiguration != nil && b.ko.Spec.ImageScanningConfiguration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ImageScanningConfiguration.ScanOnPush, b.ko.Spec.ImageScanningConfiguration.ScanOnPush) {
			delta.Add("Spec.ImageScanningConfiguration.ScanOnPush", a.ko.Spec.ImageScanningConfiguration.ScanOnPush, b.ko.Spec.ImageScanningConfiguration.ScanOnPush)
		} else if a.ko.Spec.ImageScanningConfiguration.ScanOnPush != nil && b.ko.Spec.ImageScanningConfiguration.ScanOnPush != nil {
			if *a.ko.Spec.ImageScanningConfiguration.ScanOnPush != *b.ko.Spec.ImageScanningConfiguration.ScanOnPush {
				delta.Add("Spec.ImageScanningConfiguration.ScanOnPush", a.ko.Spec.ImageScanningConfiguration.ScanOnPush, b.ko.Spec.ImageScanningConfiguration.ScanOnPush)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ImageTagMutability, b.ko.Spec.ImageTagMutability) {
		delta.Add("Spec.ImageTagMutability", a.ko.Spec.ImageTagMutability, b.ko.Spec.ImageTagMutability)
	} else if a.ko.Spec.ImageTagMutability != nil && b.ko.Spec.ImageTagMutability != nil {
		if *a.ko.Spec.ImageTagMutability != *b.ko.Spec.ImageTagMutability {
			delta.Add("Spec.ImageTagMutability", a.ko.Spec.ImageTagMutability, b.ko.Spec.ImageTagMutability)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RepositoryName, b.ko.Spec.RepositoryName) {
		delta.Add("Spec.RepositoryName", a.ko.Spec.RepositoryName, b.ko.Spec.RepositoryName)
	} else if a.ko.Spec.RepositoryName != nil && b.ko.Spec.RepositoryName != nil {
		if *a.ko.Spec.RepositoryName != *b.ko.Spec.RepositoryName {
			delta.Add("Spec.RepositoryName", a.ko.Spec.RepositoryName, b.ko.Spec.RepositoryName)
		}
	}

	if !equalTagList(a.ko.Spec.Tags, b.ko.Spec.Tags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}

func equalTagList(a, b []*svcapitypes.Tag) bool {
	//TODO(a-hilaly) implement this function
	return true
}