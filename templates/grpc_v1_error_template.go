package templates

const GrpcV1ErrorTemplate = `
package v1

import (
	err "github.com/d1zero/scratch/pkg/error"
	"google.golang.org/grpc/codes"
	"user-service/internal/entity"
)

var ErrCodeMap = map[err.ErrCode]codes.Code{
	entity.ErrCodeUnknown: codes.Unknown,
}
`
