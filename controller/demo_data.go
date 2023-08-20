package controller


var DemoVideos = []Video{
	//{
	//	Id:            1,
	//	Author:        DemoUser,
	//	//PlayUrl:       "https://web-tlias-amireux.oss-cn-hangzhou.aliyuncs.com/0e148e93-1c67-44ce-aa81-7835e3d62f63.mp4",
	//	//PlayUrl:       "https://web-tlias-amireux.oss-cn-hangzhou.aliyuncs.com/6aff7637-d31a-4f7f-ba7a-b96183886391.mp4",
	//	PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
	//	CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
	//	FavoriteCount: 0,
	//	CommentCount:  0,
	//	IsFavorite:    false,
	//},
	{
		Id:            2,
		Author:        DemoUser,
		PlayUrl:       "https://web-tlias-amireux.oss-cn-hangzhou.aliyuncs.com/2ef4e4b2-2a8b-4793-b745-6c0f9db5b593.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},

}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
