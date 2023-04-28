package main

import (
	"net"
    "net/http"
    "net/url"
	"fmt"
)

var players map[string]net.IP = make(map[string]net.IP)
var start bool = false

func createGame(username string) {
	players[username] = GetOutboundIP();
	waitForPlayers();
}

func joinGame(ip string, username string){
    data := url.Values{
        "ip": {GetOutboundIP().String()},
        "username": {username},
    }

    http.PostForm("http://"+ip+":9000/join", data);
}

func waitForPlayers(){
    go listenForPlayers();
    loop:=true
    for(loop) {
        resp := ask();
        switch(lobbyCommands(resp)){
        case 1:
            fmt.Println("start");
            start = true
            loop = false
            break;
        case -1:
            loop = false
            break;
        };
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

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        fmt.Printf("error");
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}