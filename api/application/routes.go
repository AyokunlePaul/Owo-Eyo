package application

func mapOwoEyoRoutes() {
	blockChainGroup := owoEyoRouter.Group("/v1")
	blockChainGroup.GET("/mine")
}
