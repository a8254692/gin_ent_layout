package auth

import (
	"back-end/common/libs/jwtHelper"
	"back-end/merchant/internal/config"
	"back-end/merchant/pkg/ginx"
)

// GetJwtUserId ...
func GetJwtUserId(ctx *ginx.Context) uint64 {
	jwtToken, ok := ctx.Get(config.AUTHUSERKEY)
	if !ok {
		return 0
	}

	jwtTokenInfo := jwtToken.(*jwtHelper.MerchantClaims)
	if jwtTokenInfo.Id <= 0 {
		return 0
	}

	return jwtTokenInfo.Id
}
