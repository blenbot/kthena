/*
Copyright The Volcano Authors.

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

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/volcano-sh/kthena/pkg/controller"
)

func TestParseControllers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]bool
	}{
		{
			name:  "wildcard_only",
			input: "*",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: true,
				controller.AutoscalerController:   true,
			},
		},
		{
			name:  "wildcard_with_other_controllers",
			input: "*,modelserving",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: true,
				controller.AutoscalerController:   true,
			},
		},
		{
			name:  "single_controller_modelserving",
			input: "modelserving",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: false,
				controller.AutoscalerController:   false,
			},
		},
		{
			name:  "single_controller_modelbooster",
			input: "modelbooster",
			expected: map[string]bool{
				controller.ModelServingController: false,
				controller.ModelBoosterController: true,
				controller.AutoscalerController:   false,
			},
		},
		{
			name:  "single_controller_autoscaler",
			input: "autoscaler",
			expected: map[string]bool{
				controller.ModelServingController: false,
				controller.ModelBoosterController: false,
				controller.AutoscalerController:   true,
			},
		},
		{
			name:  "multiple_controllers",
			input: "modelserving,modelbooster",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: true,
				controller.AutoscalerController:   false,
			},
		},
		{
			name:  "all_controllers_explicit",
			input: "modelserving,modelbooster,autoscaler",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: true,
				controller.AutoscalerController:   true,
			},
		},
		{
			name:  "controllers_with_spaces",
			input: " modelserving , modelbooster ",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: true,
				controller.AutoscalerController:   false,
			},
		},
		{
			name:  "invalid_controller_name",
			input: "invalid",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: true,
				controller.AutoscalerController:   true,
			},
		},
		{
			name:  "mixed_valid_and_invalid_controllers",
			input: "modelserving,invalid",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: false,
				controller.AutoscalerController:   false,
			},
		},
		{
			name:  "empty_string",
			input: "",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: true,
				controller.AutoscalerController:   true,
			},
		},
		{
			name:  "only_commas",
			input: ",,",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: true,
				controller.AutoscalerController:   true,
			},
		},
		{
			name:  "invalid_with_no_valid_controllers",
			input: "invalid1,invalid2",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: true,
				controller.AutoscalerController:   true,
			},
		},
		{
			name:  "duplicate_valid_controllers",
			input: "modelserving,modelserving",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: false,
				controller.AutoscalerController:   false,
			},
		},
		{
			name:  "invalid_controller",
			input: "*modelserving",
			expected: map[string]bool{
				controller.ModelServingController: true,
				controller.ModelBoosterController: true,
				controller.AutoscalerController:   true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseControllers(tt.input)
			assert.Equal(t, tt.expected, result, "parseControllers(%q) result mismatch", tt.input)
		})
	}
}
