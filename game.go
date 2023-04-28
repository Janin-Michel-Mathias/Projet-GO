package main

import (
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
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
		gameCommands(resp);
	}
}

func allowRoutes(){
	http.HandleFunc("/gameSet", gameSetHandler);
	http.ListenAndServe(":9000", nil)
}

func gameSetHandler(w http.ResponseWriter, req *http.Request){
	if(req.Method == http.MethodGet){
		fmt.Fprintf(w, " 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 ");
	}
}

func askForGameSet(player string){
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