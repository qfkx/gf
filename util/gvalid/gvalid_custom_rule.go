// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gvalid

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
)

// RuleFunc is the custom function for data validation.
type RuleFunc func(ctx context.Context, in RuleFuncInput) error

// RuleFuncInput holds the input parameters that passed to custom rule function RuleFunc.
type RuleFuncInput struct {
	// Rule specifies the validation rule string, like "required", "between:1,100", etc.
	Rule string

	// Message specifies the custom error message or configured i18n message for this rule.
	Message string

	// Value specifies the value for this rule to validate.
	Value *gvar.Var

	// Data specifies the `data` which is passed to the Validator. It might be a type of map/struct or a nil value.
	// You can ignore the parameter `Data` if you do not really need it in your custom validation rule.
	Data *gvar.Var
}

var (
	// customRuleFuncMap stores the custom rule functions.
	// map[Rule]RuleFunc
	customRuleFuncMap = make(map[string]RuleFunc)
)

// RegisterRule registers custom validation rule and function for package.
func RegisterRule(rule string, f RuleFunc) error {
	customRuleFuncMap[rule] = f
	return nil
}

// RegisterRuleByMap registers custom validation rules using map for package.
func RegisterRuleByMap(m map[string]RuleFunc) error {
	for k, v := range m {
		customRuleFuncMap[k] = v
	}
	return nil
}

// DeleteRule deletes custom defined validation one or more rules and associated functions from global package.
func DeleteRule(rules ...string) {
	for _, rule := range rules {
		delete(customRuleFuncMap, rule)
	}
}
