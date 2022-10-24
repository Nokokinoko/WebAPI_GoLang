package main

import (
	"webapi/handler"

	"github.com/labstack/echo/v4"
)

func router(echo *echo.Echo) {
	mapDir := map[string]handler.IController{
		"/dev": handler.DevelopController{},
	}

	for dir, ctrl := range mapDir {
		group := echo.Group(dir)

		nft := group.Group("/nft")
		nft.POST("/balance", ctrl.Balance)
		nft.POST("/list_collections", ctrl.ListCollections)
		nft.POST("/collection_detail", ctrl.CollectionDetail)
		nft.POST("/clone", ctrl.Clone)
		nft.POST("/mint", ctrl.Mint)
		nft.POST("/burn", ctrl.Burn)
		nft.POST("/burn_to_mint", ctrl.BurnToMint)
		nft.POST("/transfer", ctrl.Transfer)
		nft.POST("/list_tokens", ctrl.ListTokens)
		nft.POST("/list_tokens_by_collection", ctrl.ListTokensByCollection)
		nft.POST("/token_detail", ctrl.TokenDetail)
		nft.POST("/list_user_tokens", ctrl.ListUserTokens)
		nft.POST("/list_user_tokens_by_collection", ctrl.ListUserTokensByCollection)
		nft.POST("/user_token_detail", ctrl.UserTokenDetail)
		nft.POST("/list_users_by_token", ctrl.ListUsersByToken)

		group.POST("/ver", ctrl.Version)
		group.POST("/log", ctrl.Log)
	}
}
