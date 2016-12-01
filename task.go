/* 
 * IronFunctions
 *
 * The open source serverless platform.
 *
 * OpenAPI spec version: 0.1.23
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package functions

import (
	"time"
)

type Task struct {

	// Name of Docker image to use. This is optional and can be used to override the image defined at the group level.
	Image string `json:"image,omitempty"`

	// Payload for the task. This is what you pass into each task to make it do something.
	Payload string `json:"payload,omitempty"`

	// Group this task belongs to.
	GroupName string `json:"group_name,omitempty"`

	// The error message, if status is 'error'. This is errors due to things outside the task itself. Errors from user code will be found in the log.
	Error_ string `json:"error,omitempty"`

	// Machine usable reason for task being in this state. Valid values for error status are `timeout | killed | bad_exit`. Valid values for cancelled status are `client_request`. For everything else, this is undefined. 
	Reason string `json:"reason,omitempty"`

	// Time when task was submitted. Always in UTC.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Time when task started execution. Always in UTC.
	StartedAt time.Time `json:"started_at,omitempty"`

	// Time when task completed, whether it was successul or failed. Always in UTC.
	CompletedAt time.Time `json:"completed_at,omitempty"`

	// If this field is set, then this task is a retry of the ID in this field.
	RetryOf string `json:"retry_of,omitempty"`

	// If this field is set, then this task was retried by the task referenced in this field.
	RetryAt string `json:"retry_at,omitempty"`

	// Env vars for the task. Comes from the ones set on the Group.
	EnvVars map[string]string `json:"env_vars,omitempty"`
}
