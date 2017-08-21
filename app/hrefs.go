// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "cellar": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=project/goa-tutorial/design
// --out=$(GOPATH)/src/project/goa-tutorial
// --version=v1.2.0

package app

import (
	"fmt"
	"strings"
)

// AccountHref returns the resource href.
func AccountHref(accountID interface{}) string {
	paramaccountID := strings.TrimLeftFunc(fmt.Sprintf("%v", accountID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/cellar/accounts/%v", paramaccountID)
}

// BottleHref returns the resource href.
func BottleHref(accountID, bottleID interface{}) string {
	paramaccountID := strings.TrimLeftFunc(fmt.Sprintf("%v", accountID), func(r rune) bool { return r == '/' })
	parambottleID := strings.TrimLeftFunc(fmt.Sprintf("%v", bottleID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/cellar/accounts/%v/bottles/%v", paramaccountID, parambottleID)
}
