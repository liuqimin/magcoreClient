package logic

import (
	"awesomeProject/api"
	"fmt"
)

func Gameinit()(game_id string,palyer_id string,BasesIp []string,ownIndex int,mapRows []string ,err error ) {
	//var game_id string
	game_list,_ := api.GameList()
	fmt.Println("当前的gamelist是",game_list)
	if len(game_list) == 0 {
		game_id = api.CreateGame()
	} else {
		var flag int8
		for _,list := range game_list {
			if list.State == 0 {
				game_id = list.Id
				flag = 1
			}
		}
		if flag != 1 {
			game_id = api.CreateGame()
		}

	}

	playerInfo,createplayeerr := api.CreatePlayer()
	if createplayeerr != nil {
		fmt.Println(createplayeerr)
		return
	}
	fmt.Println(playerInfo)
	fmt.Println("createplayer over")
	status,_ := api.JoinGame(game_id,playerInfo.Id)
	palyer_id = playerInfo.Id
	fmt.Println(palyer_id)
	if status {
		fmt.Println("已经加入游戏")
		startgameStatus,_ := api.StartGame(game_id)
		if startgameStatus {

			now_info ,getplayererr := api.GetPlayer(playerInfo.Id)
			if getplayererr != nil  {
				fmt.Println(getplayererr)
				err = getplayererr
				return
			}
			fmt.Println(now_info)
			BasesIp = now_info.Bases
			ownIndex = now_info.Index
			mapInfo, _ := api.GetMap()
			mapRows = mapInfo.Rows
			fmt.Println("地图信息是 ",mapRows)
			return
		}
	}
	err = fmt.Errorf("加入游戏失败")
	return
}
