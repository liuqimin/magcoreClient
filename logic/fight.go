package logic

import (
	"fmt"
	"reflect"
	"strings"
	"strconv"
	"awesomeProject/api"
	"awesomeProject/model"
	"sync"
)

func AttachStart(){
	game_id,palyer_id,_,_,mapRows,_ := Gameinit()
	gameinfo ,_:= api.GetGame(game_id)
	fmt.Println("game all info iss",gameinfo)

	api.GetMap()
	fmt.Println(mapRows)
	//api.Attack(game_id,palyer_id)
	//getMapInfo(game_id)
	allAttackLocation := powerAllAttach(mapRows)
	attachAllLocation(game_id,palyer_id,allAttackLocation)
}

func attachAllLocation(game_id string,palyer_id string,allAttackLocation [][]int){
	ggameState := getMapInfo(game_id)
	if (ggameState.State != 2 && ggameState.State != 3 && ggameState.State != 0){
		var waitGroup sync.WaitGroup
		for _,location := range allAttackLocation {
			fmt.Println("开始攻击",location)
			waitGroup.Add(1)
			go threadAttach(game_id,palyer_id,location,&waitGroup)
		}
		waitGroup.Wait()
		fmt.Println("开始查看状态")
		attachAllLocation(game_id,palyer_id,allAttackLocation)
	} else {
		fmt.Println("over")
		//attachAllLocation(game_id,palyer_id,allAttackLocation)
		return
	}

}

func threadAttach(game_id string,palyer_id string,location []int,waitGroup *sync.WaitGroup){
	api.Attack(game_id,palyer_id,location)
	waitGroup.Done()
}

func testAttach(BasesIp []string,ownIndex int, game_id string,palyer_id string){
	var attchState string
	attchState = BasesIp[0]
	XYSplice := strings.Split(attchState,",")
	//XYSplice[0],_ =  strconv.Atoi(XYSplice[0])
	//XYSplice[1] = int(XYSplice[1]) + 1
	//attchState[1] = BasesIp[1]
	x,_ := strconv.Atoi(XYSplice[0])
	y,_ := strconv.Atoi(XYSplice[1])
	//x = x
	y = y + 1
	location := []int{x, y}
	fmt.Println("baseip",BasesIp,reflect.TypeOf(attchState[0]),reflect.TypeOf(XYSplice[0]),XYSplice[1])
	api.Attack(game_id,palyer_id,location)
	fmt.Println(attchState[0])
	fmt.Println(ownIndex)
	fmt.Println(game_id,palyer_id)
	x = x +1
	y = y - 1
	location1 := []int{x, y}
	api.Attack(game_id,palyer_id,location1)
	x = 0
	y = 0
	location2 := []int{x, y}
	api.Attack(game_id,palyer_id,location2)
	getOwnBase(game_id,ownIndex)
	fmt.Println("\n")
}

func powerAllAttach(mapRows []string)(allAttackLocation [][]int){
	//得到所有可以攻击的地点
	fmt.Println(reflect.TypeOf(mapRows))
	//allAttackLocation := [][2]int{}
	//allAttackLocation = allAttackLocation{}
	for yindex,rows := range mapRows {
		rowByte := []byte(rows)

		for xIndex,valueS := range rowByte {
			value,_ := strconv.Atoi(string(valueS))
			if value == 1 {
				tmpLocation :=[]int{xIndex,yindex}
				//fmt.Println("ssaa ",tmpLocation)
				allAttackLocation = append(allAttackLocation, tmpLocation)
			}
			//fmt.Println("id is %d,%s,%d",yindex,xIndex,value)
		}
	}
	fmt.Println("所有可以攻击的地点",allAttackLocation)
	return
}

func getOwnBase(game_id string,ownIndex int)(result [][]int){
	//获取自己的基地
	var gameinfo model.GetGameResponse
	gameinfo ,_= api.GetGame(game_id)
	for _,v := range gameinfo.Cells {
		fmt.Println("kaishi")
		fmt.Print(v)
		result := [][2]int{}
		for _, vv := range v {
			if vv.Owner == ownIndex {
				ownBase :=[2]int{vv.X,vv.Y}
				result = append(result, ownBase)
			}

		}
	}

	return
}

func getMapInfo(game_id string)(gameinfo model.GetGameResponse){
	gameinfo ,_= api.GetGame(game_id)
	fmt.Println("now state is ",gameinfo.State)
	fmt.Println("game all info iss",gameinfo)
	return
}

func GetCanAttachState(mapRows []string){

}