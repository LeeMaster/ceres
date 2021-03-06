package account

import (
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/account"
	"ceres/pkg/utility/auth"
	"strconv"
)

// ListAccounts list all accounts of the Comer
func ListAccounts(ctx *router.Context) {
	uin, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	response, err := service.GetComerAccounts(uin)
	if err != nil {
		ctx.ERROR(router.ErrBuisnessError, err.Error())
		return
	}

	ctx.OK(response)
}

// UnlinkAccount unlink accounts for the Comer
func UnlinkAccount(ctx *router.Context) {
	uin, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	identifier, err := strconv.ParseInt(ctx.Query("identifier"), 10, 64)
	if err != nil {
		ctx.ERROR(router.ErrParametersInvaild, err.Error())
		return
	}
	err = service.UnlinkComerAccount(uin, uint64(identifier))
	if err != nil {
		ctx.ERROR(router.ErrBuisnessError, err.Error())
		return
	}

	ctx.OK(nil)
}

// LinkWithGithub link current account with github
func LinkWithGithub(ctx *router.Context) {
	uin, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	requestToken := ctx.Query("request_token")
	if requestToken == "" {
		ctx.ERROR(router.ErrParametersInvaild, "request_token missed")
		return
	}
	client := auth.NewGithubOauthClient(requestToken)
	err := service.LinkOauthAccountToComer(uin, client, model.GithubOauth)
	if err != nil {
		ctx.ERROR(router.ErrBuisnessError, err.Error())
		return

	}
	ctx.OK(nil)
}

// LinkWithGithub link current account with github
// FIXME: should eliminate the duplicate code in the login api
func LinkWithMetamask(ctx *router.Context) {
	uin, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	signature := &model.EthSignatureObject{}
	err := ctx.BindJSON(signature)
	if err != nil {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"wrong metamask login parameter",
		)
		return
	}

	err = service.LinkEthAccountToComer(
		uin,
		signature.Address,
		signature.MessageHash,
		signature.Signature,
		model.MetamaskEth,
	)

	if err != nil {
		ctx.ERROR(
			router.ErrBuisnessError,
			err.Error(),
		)
		return
	}

	ctx.OK(nil)
}

// CheckComerExists
func CheckComerExists(ctx *router.Context) {
	oin := ctx.Query("oin")
	result, err := service.CheckComerExists(oin)
	if err != nil {
		ctx.ERROR(
			router.ErrBuisnessError,
			err.Error(),
		)
		return
	}

	ctx.OK(result)
}
