package model

const (
	//远程服务器地址
	ServerHost = "http://106.75.33.221:6000"
	AttachUrl = "%s/api/cell"
	CreateGame = "%s/api/game"
	CreatePlayer = "%s/api/player"
	GameList = "%s/api/game"
	//{GameId}: 欲获取信息的游戏Id
	GetGame = "%s/api/game/%s"
	//{name}: 欲获取地图的名称
	GetMap = "%s/api/map/%s"
	//{PlayerId}: 欲获取信息的玩家Id
	GetPlayer = "%s/api/player/%s"
	JoinGame = "%s/api/game"
	//{GameId}: 欲开始的游戏Id
	StartGame = "%s/api/game/%s"
)
