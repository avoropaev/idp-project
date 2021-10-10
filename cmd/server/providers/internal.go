package providers

import (
	"database/sql/driver"
	codeModule "github.com/avoropaev/idp-project/internal/app/code"
	"github.com/avoropaev/idp-project/sdk/s1sdk"
	"github.com/avoropaev/idp-project/sdk/s2sdk"
)

func ProvideCodeRepository(_ driver.Connector) *codeModule.Repository {
	rep := codeModule.NewRepository()

	return &rep
}

func ProvideCodeService(s1Client s1sdk.S1Client, s2Client s2sdk.S2Client, rep *codeModule.Repository) *codeModule.Service {
	ser := codeModule.NewService(s1Client, s2Client, *rep)

	return &ser
}
