package account

import (
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	service "ceres/pkg/service/account"
	"ceres/pkg/utility/auth"
)

// LoginWithGithub login with github oauth
func LoginWithGithub(ctx *router.Context) {
	requestToken := ctx.Query("request_token")
	if requestToken == "" {
		ctx.ERROR(router.ErrParametersInvaild, "request_token missed")
		return
	}
	client := auth.NewGithubOauthClient(requestToken)
	response, err := service.LoginWithOauth(client, model.GithubOauth)
	if err != nil {
		ctx.ERROR(router.ErrBuisnessError, err.Error())
		return
	}
	ctx.OK(response)
}

// LoginWithFacebook login with facebook oauth
func LoginWithFacebook(ctx *router.Context) {
	requestToken := ctx.Query("request_token")
	if requestToken == "" {
		ctx.ERROR(router.ErrParametersInvaild, "request_token missed")
		return
	}
	client := auth.NewFacebookClient(requestToken)
	response, err := service.LoginWithOauth(client, model.GithubOauth)
	if err != nil {
		ctx.ERROR(router.ErrBuisnessError, err.Error())
		return
	}
	ctx.OK(response)
}

// LoginWithTwitter login with twitter oauth
func LoginWithTwitter(_ *router.Context) {
	// TODO: should complete the twitter logic
}

// LoginWithLinkedIn login with linkedin oauth
func LoginWithLinkedIn(_ *router.Context) {
	// TODO: should complete the linkedin logic
}

// GetBlockchainLoginNonce get the blockchain login nonce.
func GetBlockchainLoginNonce(ctx *router.Context) {
	address := ctx.Query("address")
	if address == "" {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"no web3 public key",
		)
		return
	}
	nonce, err := service.GenerateWeb3LoginNonce(address)
	if err != nil {
		ctx.ERROR(
			router.ErrBuisnessError,
			err.Error(),
		)
		return
	}
	ctx.OK(nonce)
}

// LoginWithMetamask login with the metamask signature.
func LoginWithMetamask(ctx *router.Context) {
	signature := &model.EthSignatureObject{}
	err := ctx.BindJSON(signature)
	if err != nil {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"wrong metamask login parameter",
		)
		return
	}

	// Replace the 0x prefix

	resposne, err := service.VerifyEthLogin(
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
	ctx.OK(resposne)
}
