package main

import (
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
	"net/url"
	"strconv"
	"strings"
)


type boat struct{
	position [2]int
	length int
	isVertical bool
}

var boats []boat;
var playedCases [10][10]bool;

func startGame(){
	boats = generateBoats();

	for i:= 0; i < 10; i++{
		for j:= 0; j < 10; j++{
			playedCases[i][j] = false;
		}
	}

	go allowRoutes();
	for {
		resp := ask();
		gameCommands(resp, boats);
	}
}

func attack (username string, playCase string){
	data := url.Values{
		"playCase": {playCase},
	}

	http.PostForm("http://"+players[username].String()+"/attack", data);
}

func allowRoutes(){
	http.HandleFunc("/gameSet", gameSetHandler);
	http.HandleFunc("/attack", attackHandler);
	http.ListenAndServe(":9000", nil)
}

func attackHandler(w http.ResponseWriter, req *http.Request){
	if(req.Method == http.MethodPost){
		if err := req.ParseForm(); err != nil { // Parsing des paramètres envoyés
            fmt.Println("Something went bad"); // par le client et gestion
            // d’erreurs
            fmt.Fprintln(w, "Something went bad");
            return
            }
		attackCase := req.PostForm["attack"];

		x, y := parseAttack(attackCase[0]);
		if(!playedCases[x][y]){
			if(isBoatPos([2]int{x, y}, boats)){
				if(checkBoatAlive([2]int{x,y}, boats)){
					fmt.Fprintf(w, "true");
				} else {
					fmt.Fprintf(w, "flood");
				}
				playedCases[x][y] = true;
			} else {
				fmt.Fprintf(w, "false");
			}
		}
	}
}

func parseAttack(attackCase string)(int ,int){
	parts := strings.Split(attackCase, "");
	firstPartMap := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
		"H": 7,
		"I": 8,
		"J": 9,
	}

	partTwo, _ := strconv.Atoi(parts[1])

	return firstPartMap[parts[0]], partTwo;
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
		for j:= 0; j < 10; j++{
			gameSet +=" | ";
			if(isBoatPos([2]int{i,j}, boats)) {
				if(private){
					gameSet +="1";
				} else {
					if(playedCases[i][j]){
						gameSet += "X";
					}else {
						gameSet +=" ";
					}
				}
			} else {
				if(playedCases[i][j]){
					gameSet += "."
				}else {
					gameSet +=" ";
				}
			}
		}
		gameSet +="\n";
	}
	return gameSet;
}

func checkBoatAlive(pos [2]int, boats []boat) bool{
	boatIndex := getBoat(pos, boats);
	for i:= 0; i < boats[boatIndex].length; i++{
		if(boats[boatIndex].isVertical){
			if(!playedCases[boats[boatIndex].position[0]][boats[boatIndex].position[1] + i]){
				return true
			}
		}
	}
	return false;
}

func getBoat(pos [2]int, boats []boat)int {
	for i:= 0; i< len(boats); i++{
		x := boats[i].position[0];
		y:= boats[i].position[1];
		for j:= 0; j < boats[i].length; j++{
			if(x == pos[0] && y == pos[1]){
				return i;
			}
			if(boats[i].isVertical){
				y++;
			}else{
				x++;
			}
		}
	} 
	return -1;
}