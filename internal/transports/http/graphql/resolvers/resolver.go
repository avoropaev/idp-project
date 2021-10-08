package resolvers

import (
	codeModule "github.com/avoropaev/idp-project/internal/app/code"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CodeService codeModule.Service
}
