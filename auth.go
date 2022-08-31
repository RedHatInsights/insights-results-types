// Copyright 2020 Red Hat, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

// ContextKey is a type for user authentication token in request
type ContextKey string

const (
	// ContextKeyUser is a constant for user authentication token in request
	ContextKeyUser = ContextKey("user")
)

// Internal contains information about organization ID
type Internal struct {
	OrgID OrgID `json:"org_id,string"`
}

// Identity contains internal user info
type Identity struct {
	AccountNumber UserID   `json:"account_number"`
	Internal      Internal `json:"internal"`
	User          User     `json:"user"`
}

// User struct contains single user data passed in the auth token
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserID   UserID `json:"user_id"`
}

// IdentityV2 contains user info after July 22 modifications (org_id moved to top level)
// will be replacing Identity after testing
type IdentityV2 struct {
	// AccountNumber to be removed after EBS/org_id migration is finalized
	AccountNumber UserID `json:"account_number"`
	OrgID         OrgID  `json:"org_id,string"`
	User          User   `json:"user"`
}

// Token is x-rh-identity struct
type Token struct {
	Identity Identity `json:"identity"`
}

// TokenV2 is x-rh-identity struct after July 22 modifications (org_id moved to top level)
// will be replacing Token after testing
type TokenV2 struct {
	IdentityV2 IdentityV2 `json:"identity"`
}

// JWTPayload is structure that contain data from parsed JWT token
type JWTPayload struct {
	AccountNumber UserID `json:"account_number"`
	OrgID         OrgID  `json:"org_id,string"`
	UserID        UserID `json:"user_id"`
}
