package providers

import (
	codeModule "github.com/avoropaev/idp-project/internal/app/code"
	"github.com/avoropaev/idp-project/sdk/s1sdk"
	"github.com/avoropaev/idp-project/sdk/s2sdk"
)

func ProvideCodeService(s1Client s1sdk.S1Client, s2Client s2sdk.S2Client) *codeModule.Service {
	ser := codeModule.NewService(s1Client, s2Client)

	return &ser
}
