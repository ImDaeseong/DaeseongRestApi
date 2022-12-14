// main
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Res struct {
	Res     string `json:"response"`
	Message string `json:"message"`
}

type Game struct {
	ID          string    `json:"id"`
	PackageName string    `json:"packagename"`
	GameTitle   string    `json:"gametitle"`
	GameDesc    *GameDesc `json:"gamedesc"`
}

type GameDesc struct {
	Details1 string `json:"details1"`
	Details2 string `json:"details2"`
}

var gamedata []Game

//AllList
func getGames(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	json.NewEncoder(response).Encode(gamedata)
}

//item
func getGame(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("COntent-Type", "application/json")

	//get
	id := request.URL.Query().Get("id")

	for _, items := range gamedata {
		if items.ID == id {
			json.NewEncoder(response).Encode(items)
			return
		}
	}

	//Not Found
	res := Res{Res: "err", Message: "Not found"}
	json.NewEncoder(response).Encode(res)
}

//del
func deleteGame(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("COntent-Type", "application/json")

	id := request.URL.Query().Get("id")
	for index, items := range gamedata {

		if items.ID == id {
			//찾은 데이터 제외하고 전체 데이터(gamedata[:index]:id보다 작은 데이터, gamedata[index+1:]:id보다 큰 데이터)
			gamedata = append(gamedata[:index], gamedata[index+1:]...)
			break
		}
	}

	//전체 내용 보여줌
	json.NewEncoder(response).Encode(gamedata)
}

//update
func updateGame(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("COntent-Type", "application/json")

	//post
	id := request.FormValue("id")
	name := request.FormValue("name")
	title := request.FormValue("title")
	detail1 := request.FormValue("detail1")
	detail2 := request.FormValue("detail2")

	for index, items := range gamedata {

		if items.ID == id {

			var j1 = ""
			var j2 = ""
			var j3 = ""
			var j4 = ""

			if gamedata[index].PackageName != name && len(name) > 0 {
				j1 = name
			} else {
				j1 = gamedata[index].PackageName
			}

			if gamedata[index].GameTitle != title && len(title) > 0 {
				j2 = title
			} else {
				j2 = gamedata[index].GameTitle
			}

			if gamedata[index].GameDesc.Details1 != detail1 && len(detail1) > 0 {
				j3 = detail1
			} else {
				j3 = gamedata[index].GameDesc.Details1
			}

			if gamedata[index].GameDesc.Details2 != detail2 && len(detail2) > 0 {
				j4 = detail2
			} else {
				j4 = gamedata[index].GameDesc.Details2
			}

			//찾은 데이터 제외하고 전체 데이터(gamedata[:index]:id보다 작은 데이터, gamedata[index+1:]:id보다 큰 데이터)
			gamedata = append(gamedata[:index], gamedata[index+1:]...)

			var update = Game{ID: id, PackageName: j1, GameTitle: j2, GameDesc: &GameDesc{Details1: j3, Details2: j4}}
			gamedata = append(gamedata, update)

			//수정 내용만 보여줌
			//json.NewEncoder(response).Encode(update)

			break
		}
	}

	//전체 내용 보여줌
	json.NewEncoder(response).Encode(gamedata)
}

//add
func addGame(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("COntent-Type", "application/json")

	//post
	name := request.FormValue("name")
	title := request.FormValue("title")
	detail1 := request.FormValue("detail1")
	detail2 := request.FormValue("detail2")

	if len(name) == 0 || len(title) == 0 || len(detail1) == 0 || len(detail2) == 0 {

		res := Res{Res: "err", Message: "데이터를 입력하세요"}
		json.NewEncoder(response).Encode(res)
		return
	}

	id := strconv.Itoa(len(gamedata) + 1)
	//fmt.Println("ID:", id)

	var add = Game{ID: id, PackageName: name, GameTitle: title, GameDesc: &GameDesc{Details1: detail1, Details2: detail2}}
	gamedata = append(gamedata, add)

	//추가 내용만 보여줌
	//json.NewEncoder(response).Encode(add)

	//전체 내용 보여줌
	json.NewEncoder(response).Encode(gamedata)
}

func main() {

	gamedata = append(gamedata, Game{ID: "1", PackageName: "com.pearlabyss.blackdesertm", GameTitle: "검은사막 모바일", GameDesc: &GameDesc{Details1: "당신이 진짜로 원했던 모험의 시작", Details2: "월드클래스 MMORPG “검은사막 모바일”"}})
	gamedata = append(gamedata, Game{ID: "2", PackageName: "com.kakaogames.moonlight", GameTitle: "달빛조각사", GameDesc: &GameDesc{Details1: "500만 구독자의 게임 판타지 대작 '달빛조각사'", Details2: "- 5레벨만 달성해도 달빛조각사 이모티콘 100% 지급!"}})
	gamedata = append(gamedata, Game{ID: "3", PackageName: "com.ncsoft.lineagem19", GameTitle: "리니지M", GameDesc: &GameDesc{Details1: "PC의 향수! 리니지 본질 그대로 리니지M", Details2: "PC리니지와 동일한 아덴월드의 오픈 필드"}})
	gamedata = append(gamedata, Game{ID: "4", PackageName: "com.netmarble.bnsmkr", GameTitle: "블레이드&소울 레볼루션", GameDesc: &GameDesc{Details1: "원작 감성의 방대한 세계관과 복수 중심의 흥미진진한 스토리", Details2: "MMORPG의 필드를 제대로 즐길 수 있는 경공"}})
	gamedata = append(gamedata, Game{ID: "5", PackageName: "com.cjenm.sknights", GameTitle: "세븐나이츠", GameDesc: &GameDesc{Details1: "Netmarble롤플레잉", Details2: "세나의 재탄생, 세븐나이츠: 리부트"}})
	gamedata = append(gamedata, Game{ID: "6", PackageName: "com.google.android.youtube", GameTitle: "YouTube", GameDesc: &GameDesc{Details1: "Google LLC동영상 플레이어/편집기", Details2: "좋아하는 동영상 빠르게 검색하기"}})

	http.HandleFunc("/api/AllList", getGames)
	http.HandleFunc("/api/item", getGame)
	http.HandleFunc("/api/del", deleteGame)
	http.HandleFunc("/api/update", updateGame)
	http.HandleFunc("/api/add", addGame)

	server := http.Server{
		Addr: "127.0.0.1:3333",
	}
	server.ListenAndServe()

	fmt.Println("http://localhost:3333/api/AllList")
	fmt.Println("http://localhost:3333/api/item?id=1")
	fmt.Println("http://localhost:3333/api/del?id=1")
	fmt.Println("http://localhost:3333/api/update")
	fmt.Println("http://localhost:3333/api/add")
}
