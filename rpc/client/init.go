package client

func InitRpcClient() {
	initFeedClient()     //8880
	initUserClient()     //8881
	initFavoriteClient() //8882
	initCommentClient()  //8883
	initPublishClient()  //8884
	initRelationClient() //8885
	initMessageClient()  //8886
}
