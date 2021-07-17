/*
Copyright 2021 The Pixiu Authors.

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

package pixiu

// PullPolicy describes a policy for if/when to pull a container image
type PullPolicy string

const (
	Pull   = "pull"
	Remove = "remove"

	// PullAlways means that imageSet always attempts to pull the latest image. Container will fail If the pull fails.
	PullAlways PullPolicy = "Always"
	// PullNever means that imageSet never pulls an image, but only uses a local image. Container will fail if the image isn't present
	PullNever PullPolicy = "Never"
	// PullIfNotPresent means that imageSet pulls if the image isn't present on disk. Container will fail if the image isn't present and the pull fails.
	PullIfNotPresent PullPolicy = "IfNotPresent"
)

var (
	AvailableActions = map[string]bool{
		Pull:   true,
		Remove: true,
	}

	AvailableImagePullPolicy = map[PullPolicy]bool{
		PullAlways:       true,
		PullIfNotPresent: true,
		PullNever:        true,
	}
)
