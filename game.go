package main

import (
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
	"strconv"
)


type boat struct{
	position [2]int
	length int
	isVertical bool
}

var boats []boat;

func startGame(){
	boats = generateBoats();

	go allowRoutes();
	for {
		resp := ask();
		gameCommands(resp, boats);
	}
}

func allowRoutes(){
	http.HandleFunc("/gameSet", gameSetHandler);
	http.ListenAndServe(":9000", nil)
}

func gameSetHandler(w http.ResponseWriter, req *http.Request){
	if(req.Method == http.MethodGet){
		fmt.Fprintf(w, getGameSet(false, boats));
	}
}

func askForGameSet(player string){
	fmt.Println("ok");
	if value, ok := players[player]; ok {

		resp, err := http.Get("http://" + value.String() + ":9000/gameSet");

		if err != nil {
			log.Fatal(err)
		}
	
		defer resp.Body.Close()
	
		body, err := ioutil.ReadAll(resp.Body)
	
		if err != nil {
			log.Fatal(err)
		}
	
		fmt.Println(string(body))
	}
}

func getGameSet(private bool, boats []boat) string{
	gameSet := "";
	gameSet +="   ";
	for i:= 0; i < 10; i++{
		gameSet +=" ";
			if(i != 0){
				gameSet +="| ";
			}
			gameSet += strconv.Itoa(i) ;
	}
	gameSet +="\n";
	letters := [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"};
	for i:= 0; i < 10; i++{
		gameSet += letters[i];
		for j:= 0; j < 11; j++{
			gameSet +=" | ";
			if(isBoatPos([2]int{i,j}, boats)) {
				if(private){
					gameSet +="1";
				} else {
					gameSet +=" ";
				}
			} else {
				gameSet +=" ";
			}
		}
		gameSet +="\n";
	}
	return gameSet;
}