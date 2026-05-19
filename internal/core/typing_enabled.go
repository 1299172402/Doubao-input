//go:build typing
// +build typing

package core

import "github.com/go-vgo/robotgo"

func typeText(msg string) {
	robotgo.Type(msg)
}
