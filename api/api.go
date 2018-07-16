package api

import (
	"net/http"
	"fmt"
	"awesomeProject/model"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"awesomeProject/utils"
)

func CreateGame()(string){
	var bufer bytes.Buffer
	req := make(map[string]string)
	req["Map"] = "RectSmall"
	data ,_ := json.Marshal(req)
	bufer.WriteString(string(data))
	fmt.Printf(string(data))
	url := fmt.Sprintf(model.CreateGame,model.ServerHost)
	fmt.Printf("%s\n",url)
	//bytes.NewBufferString("{'Map':'RectSmall'}")
	//strings.NewReader("Map=RectSamll")
	resq,err := http.Post(url,"application/json",&bufer)
	if resq.StatusCode != 200{
		fmt.Println("statuscode is not 200\n")
		return ""
	}

	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return ""
	}
	defer resq.Body.Close()
	fmt.Println(resq.StatusCode)
	body, err := ioutil.ReadAll(resq.Body)   //请求数据进行读取
	//fmt.Print(body)
	if err != nil {
		fmt.Printf("err:%v\n",err)
		return ""
		// handle error

	}
	//fmt.Printf("1")
	return string(body)

}


func CreatePlayer()(responseData model.CreatePlayerResponse,err error){
	name := utils.RandomString(10, "Aa0")
	requestDataString := fmt.Sprintf("{'Name':'%s','Color':6}",name)
	fmt.Println("create player data is ",requestDataString)
	data := bytes.NewBufferString(requestDataString)
	url := fmt.Sprintf(model.CreatePlayer,model.ServerHost)
	fmt.Printf("%s\n",url)
	client := &http.Client{}
	request, _ := http.NewRequest("POST",url,data)
	request.Header.Set("Cache-Control","no-cache")
	request.Header.Set("Content-Type","application/json")
	resq, _ := client.Do(request)
	if resq.StatusCode != 200{
		fmt.Println("statuscode is not 200\n")
		fmt.Println("create player info is",resq)
		err = fmt.Errorf("创建用户失败")
		return
	}

	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resq.Body.Close()
	fmt.Println(resq.Body)
	fmt.Println(resq.StatusCode)
	body, err := ioutil.ReadAll(resq.Body)   //请求数据进行读取
	if err != nil {
		fmt.Printf("err:%v\n",err)
		return
		// handle error

	}

	jsonErr := json.Unmarshal(body,&responseData)
	if jsonErr != nil {
		fmt.Println("json failed ",jsonErr)
		err = jsonErr
		return
	}
	return
}

func GameList()(responseData []model.GameListResponse ,err error){
	//var responseData []model.GameListResponse
	//dataStr := fmt.Sprintf("{'id':'%s','map':'RectSmall','state':0}",id_str)
	//data :=  bytes.NewBufferString(dataStr)
	client := &http.Client{}
	url := fmt.Sprintf(model.GameList,model.ServerHost)
	request, _ := http.NewRequest("GET",url,nil)
	request.Header.Set("Cache-Control","no-cache")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
		jsonerr := json.Unmarshal(body,&responseData)
		if jsonerr != nil {
			fmt.Println("json failed ",jsonerr)
			err = jsonerr
			return
		}
		fmt.Println(len(responseData))
	}
	return
}

func GetGame(id_str string)(responsedata model.GetGameResponse,err error){
	//var responsedata model.GetGameResponse
	client := &http.Client{}
	url := fmt.Sprintf(model.GetGame,model.ServerHost,id_str)
	fmt.Println(url)
	request, _ := http.NewRequest("GET",url,nil)
	response, client_err := client.Do(request)
	if response.StatusCode == 200 {
		body, _:= ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
		jsonerr := json.Unmarshal(body,&responsedata)
		if jsonerr != nil {
			fmt.Println("json failed ",jsonerr)
			err = jsonerr
			return
		}

		return
	} else {
		fmt.Print("获取失败")
		err = client_err
		return
	}
}

func GetMap() (responsedata model.GetMapResponse,err error){
	//var responsedata model.GetMapResponse
	mapName := "RectSmall"
	client := &http.Client{}
	url := fmt.Sprintf(model.GetMap,model.ServerHost,mapName)
	request, _ := http.NewRequest("GET",url,nil)
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body,_ := ioutil.ReadAll(response.Body)

		json_err := json.Unmarshal(body,&responsedata)
		if json_err != nil {
			fmt.Println("json unmarshal failed ",json_err)
			err = json_err
			return
		}
		fmt.Println("result is ",responsedata)
		return
	} else {
		print(response.StatusCode)
		err = fmt.Errorf("statuscode is not 200 ,now is %d",int(response.StatusCode))
		return
	}

}


func GetPlayer(playerId string)(responseData model.CreatePlayerResponse,err error){
	//var responsedata model.CreatePlayerResponse
	client := &http.Client{}
	url := fmt.Sprintf(model.GetPlayer,model.ServerHost,playerId)
	fmt.Println("get player url is ",url)
	request,_ := http.NewRequest("GET",url,nil)
	request.Header.Set("Cache-Control","no-cache")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		fmt.Println("getplayer response.body ",response.Body)
		body,_ := ioutil.ReadAll(response.Body)
		jsonerr := json.Unmarshal(body,&responseData)
		if jsonerr != nil {
			fmt.Println("json unmarshal failed ",jsonerr)
			err = jsonerr
			return
		}
		fmt.Println("当前player状态",responseData)
		return
	} else {
		err = fmt.Errorf("request statscode is not 200,now is %d",int(response.StatusCode))
		return
	}
}

func JoinGame(Game string,Player string)(status bool,err error){
	var requestData model.JoinGameRequest
	requestData.Game = Game
	requestData.Player = Player
	client := &http.Client{}
	data ,_:= json.Marshal(requestData)
	url := fmt.Sprintf(model.JoinGame,model.ServerHost)
	fmt.Println("请求的地址是",url)
	fmt.Println(requestData)
	request,_ := http.NewRequest("PATCH",url,bytes.NewBuffer([]byte(data)))
	request.Header.Set("Content-Type","application/json")
	request.Header.Set("Cache-Control","no-cache")
	response, _ := client.Do(request)

	if response.StatusCode == 200 {
			status = true
			return
	}
	fmt.Println("join game error statuscode is",response.StatusCode)
	fmt.Println(response)
	status = false
	return
}


func StartGame(Game string)(status bool,err error){
	client := &http.Client{}
	url := fmt.Sprintf(model.StartGame,model.ServerHost,Game)
	request,_ := http.NewRequest("PUT",url,nil)
	request.Header.Set("Cache-Control","no-cache")
	response, _ := client.Do(request)
	fmt.Println(response.StatusCode)
	if response.StatusCode == 200 {
		fmt.Println("开始游戏")
		status = true
		return
	}
	fmt.Println("开始失败")
	status = false
	return
}


func Attack(gameId string,playerId string,attachInfo []int)(err error) {
	var requestData model.AttachRequest
	requestData.Game = gameId
	requestData.Player = playerId
	requestData.X = attachInfo[0]

	requestData.Y = attachInfo[1]
	data,_ := json.Marshal(requestData)

	client := &http.Client{}
	url := fmt.Sprintf(model.AttachUrl,model.ServerHost)
	request,_ := http.NewRequest("PUT",url,bytes.NewBuffer(data))
	request.Header.Set("Content-Type","application/json")
	response,_ := client.Do(request)
	fmt.Println("attack请求发出")
	if response.StatusCode != 200 {
		fmt.Println("attack 发送失败")
		err = fmt.Errorf("攻击请求失败")
		return
	}
	fmt.Println(response.StatusCode,response.Body)
	fmt.Printf("attach请求成功 ",attachInfo)
	return
}