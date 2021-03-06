/*
Copyright 2019 The Kubernetes Authors.

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

package tags

const (
	TagOSFamilyRHEL   = "_rhel_family"
	TagOSFamilyDebian = "_debian_family"

	TagOSCentOS8 = "_centos8"
	TagOSRHEL8   = "_rhel8"

	TagSystemd = "_systemd"

	// Nodes with the "_automatic_upgrade" tag automatically update installed packages
	// during bootstrapping and daily for security updates (unless this update would require
	// a node reboot). To disable automatic node package updates, set:
	// `Cluster.Spec.UpdatePolicy = external`
	TagUpdatePolicyAuto = "_automatic_upgrades"
)

type HasTags interface {
	HasTag(tag string) bool
}
