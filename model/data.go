package model


func getRectSamll()(string){
	return "RectSmall"
}

type GameListResponse struct {
	Id string `json:"id"`
	Map string `json:"map"`
	State int `json:"state"`
}



type GetGameResponse struct {
	Id string `json:"Id"`
	Map string `json:"Map"`
	State int `json:"State"`
	Players []GetplayerMapResponse `json:"Players"`
	Cells [][]CellsResponse `json:"Cells"`
}

type GetplayerMapResponse struct {
	//getgame 返回数据的 players 数据
	Index int `json:"Index"`
	Color int `json:"Color"`
	Name string `json:"Name"`
	State int `json:"State"`
}

type PlayerResponse struct {
	Index string `json:"Index"`
	Color int `json:"Color"`
	Name string `json:"Name"`
	//Cells []CellsResponse `json:"Cells"`
}

type CellsResponse struct {
	X int `json:"X"`
	Y int `json:"Y"`
	Type int `json:"Type"`
	State int `json:"State"`
	Owner int `json:"Owner"`
}


type GetMapResponse struct {
	Edge int `json:"Edge"`
	Shift int `json:"Shift"`
	Direction int `json:"Direction"`
	Rows []string
}


type CreatePlayerResponse struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	Token string `json:"Token"`
	Energy int `json:"Energy"`
	Color int `json:"Color"`
	State int `json:"State"`
	Index int `json:"Index"`
	Bases []string
}


 type JoinGameRequest struct {
 	Game string `json:"Game"`
 	Player string `json:"Player"`
 }

 type AttachRequest struct {
 	JoinGameRequest
 	X int `json:"X"`
 	Y int `json:"Y"`
 }