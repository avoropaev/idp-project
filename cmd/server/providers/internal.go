package providers

import (
	codeModule "github.com/avoropaev/idp-project/internal/app/code"
)

func ProvideCodeService() *codeModule.Service {
	ser := codeModule.NewService()

	return &ser
}
