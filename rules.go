// Copyright 2020, 2021  Red Hat, Inc
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

import (
	"time"
)

// RuleID represents type for rule id
type RuleID string

// RuleOnReport represents a single (hit) rule of the string encoded report
type RuleOnReport struct {
	Module          RuleID      `json:"component"`
	ErrorKey        ErrorKey    `json:"key"`
	UserVote        UserVote    `json:"user_vote"`
	Disabled        bool        `json:"disabled"`
	DisableFeedback string      `json:"disable_feedback"`
	DisabledAt      Timestamp   `json:"disabled_at"`
	TemplateData    interface{} `json:"details"`
}

// ReportRules is a helper struct for easy JSON unmarshalling of string encoded report
type ReportRules struct {
	HitRules     []RuleOnReport `json:"reports"`
	SkippedRules []RuleOnReport `json:"skips"`
	PassedRules  []RuleOnReport `json:"pass"`
	TotalCount   int
}

// RuleContentResponse represents a single rule in the response of /report endpoint
type RuleContentResponse struct {
	CreatedAt    string      `json:"created_at"`
	Description  string      `json:"description"`
	ErrorKey     string      `json:"-"`
	Generic      string      `json:"details"`
	Reason       string      `json:"reason"`
	Resolution   string      `json:"resolution"`
	TotalRisk    int         `json:"total_risk"`
	RiskOfChange int         `json:"risk_of_change"`
	RuleModule   RuleID      `json:"rule_id"`
	TemplateData interface{} `json:"extra_data"`
	Tags         []string    `json:"tags"`
	UserVote     UserVote    `json:"user_vote"`
	Disabled     bool        `json:"disabled"`
	Internal     bool        `json:"internal"`
}

// DisabledRuleResponse represents a single disabled rule displaying only identifying information
type DisabledRuleResponse struct {
	RuleModule  string `json:"rule_id"`
	Description string `json:"description"`
	Generic     string `json:"details"`
	DisabledAt  string `json:"disabled_at"`
}

// Rule represents the content of rule table
type Rule struct {
	Module     RuleID `json:"module"`
	Name       string `json:"name"`
	Summary    string `json:"summary"`
	Reason     string `json:"reason"`
	Resolution string `json:"resolution"`
	MoreInfo   string `json:"more_info"`
}

// RuleErrorKey represents the content of rule_error_key table
type RuleErrorKey struct {
	ErrorKey    ErrorKey  `json:"error_key"`
	RuleModule  RuleID    `json:"rule_module"`
	Condition   string    `json:"condition"`
	Description string    `json:"description"`
	Impact      int       `json:"impact"`
	Likelihood  int       `json:"likelihood"`
	PublishDate time.Time `json:"publish_date"`
	Active      bool      `json:"active"`
	Generic     string    `json:"generic"`
	Tags        []string  `json:"tags"`
}

// RuleWithContent represents a rule with content, basically the mix of rule and rule_error_key tables' content
type RuleWithContent struct {
	Module       RuleID    `json:"module"`
	Name         string    `json:"name"`
	Summary      string    `json:"summary"`
	Reason       string    `json:"reason"`
	Resolution   string    `json:"resolution"`
	MoreInfo     string    `json:"more_info"`
	ErrorKey     ErrorKey  `json:"error_key"`
	Condition    string    `json:"condition"`
	Description  string    `json:"description"`
	TotalRisk    int       `json:"total_risk"`
	RiskOfChange int       `json:"risk_of_change"`
	PublishDate  time.Time `json:"publish_date"`
	Active       bool      `json:"active"`
	Internal     bool      `json:"internal"`
	Generic      string    `json:"generic"`
	Tags         []string  `json:"tags"`
}
