// Copyright 2021 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	gw "github.com/chaos-mesh/chaos-mesh/api/v1alpha1/genericwebhook"
)

type ScheduleItem struct {
	EmbedChaos `json:",inline"`
	// +optional
	Workflow *WorkflowSpec `json:"workflow,omitempty"`
}

func (in EmbedChaos) Validate(chaosType string) field.ErrorList {
	gw.Default(&in)
	return gw.Validate(&in)
}