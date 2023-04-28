package main

import (
	"net"
    "net/http"
    "net/url"
	"fmt"
)

var players map[string]net.IP = make(map[string]net.IP)
var start bool = false

// PART For game creation

func createGame(username string) {
	players[username] = GetOutboundIP();
	waitForPlayers();
}

func waitForPlayers(){
    go listenForPlayers();
    loop:=true
    for(loop) {
        resp := ask();
        switch(lobbyCommands(resp)){
        case 1:
            go sendStartToPlayers();
            startGame();
            start = true
            loop = false
            break;
        case -1:
            loop = false
            break;
        };
    }
    
}

func sendStartToPlayers(){
    me := true
    data := url.Values{}
    for username, ip := range(players){
        data.Add(username, ip.String());
    }

    for _, value := range(players){
        if(!me){
            http.PostForm("http://" + value.String() + ":9000/start", data);
        }
        me = false
    }
}

func listenForPlayers() {
    http.HandleFunc("/join", listenPlayersHandler);
    http.ListenAndServe(":9000", nil);
}

func listenPlayersHandler(w http.ResponseWriter, req *http.Request){
    if(req.Method == http.MethodPost){
        if err := req.ParseForm(); err != nil { // Parsing des paramètres envoyés
            fmt.Println("Something went bad"); // par le client et gestion
            // d’erreurs
            fmt.Fprintln(w, "Something went bad");
            return
            }
            username, ok := req.PostForm["username"];
            if ok{
                ip, ok := req.PostForm["ip"];
                if ok{
                    _ , ok := players[username[0]];
                    if !ok{
                        players[username[0]] = net.ParseIP(ip[0]);
                    }
                }
            }

            fmt.Fprintf(w, "Information received: %v\n", req.PostForm);
               
    }
}


// PART For a player to join a game


func joinGame(ip string, username string){
    data := url.Values{
        "ip": {GetOutboundIP().String()},
        "username": {username},
    }

    http.PostForm("http://"+ip+":9000/join", data);
    waitForGameStart();
}

func waitForGameStart(){
    http.HandleFunc("/start", startHandler);
    http.ListenAndServe(":9000", nil)
}

func startHandler(w http.ResponseWriter, req *http.Request){
    if(req.Method == http.MethodPost){
        if err := req.ParseForm(); err != nil { // Parsing des paramètres envoyés
            fmt.Println("Something went bad"); // par le client et gestion
            // d’erreurs
            fmt.Fprintln(w, "Something went bad");
            return
        }
        for username, ip := range(req.PostForm){
            players[username] = net.ParseIP(ip[0]);
        }

        for username, ip := range(players){
            fmt.Println(username + " => " + ip.String());
        }
        startGame();
    }
}

// Get IP func

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        fmt.Printf("error");
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}