package templates

const InternalErrorTemplate = `
package entity

import (
	"fmt"
	err "github.com/d1zero/scratch/pkg/error"
)

var (
	ErrUnknown = fmt.Errorf("UNKNOWN_ERROR")
)

const (
	_ = err.ErrCode(iota)
	ErrCodeUnknown
)
`
