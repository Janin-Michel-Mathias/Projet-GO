package main

import (
    "fmt"
)

 func homeCommands(resp []string) bool{
	switch(resp[0]){
	case "/create":
		if(len(resp) > 1){
			switch(resp[1]){
				case "new":
					if(len(resp) > 2 && resp[2] == "game"){
						if(len(resp) == 5 && resp[3] == "as"){
							createGame(resp[4]);
						}else{
							var username string
							fmt.Println("Type your username: ");
							fmt.Scanf("%s", &username)
							createGame(username);
						}
						fmt.Println("create");
					}
				case "-h":
					fmt.Println("/create: ");
		         	fmt.Println("new game: creates new game and wait for players");
				default:
					commandNotFound("create");
			}
		}else{
			commandNotFound("create");
		}
	case "/join":
		if(len(resp) > 1){
			if(len(resp) > 3 && resp[2] == "as"){
				joinGame(resp[1], resp[3]);
			}
			if(resp[1] == "-h"){
				fmt.Println("/join [addr]")
			}
		}else{
			commandNotFound("join");
		}
	case "/exit":
		return false;
	default:
		if(len(resp) > 0){
			fmt.Println("/!\\ '" + resp[0] + "' is not a command");
		}
	}
	return true;
}

func lobbyCommands(resp []string) int{
	switch(resp[0]){
	case "/players":
		for key, value := range(players){
			fmt.Println("[" + key + "] => " + value.String());
		}
	case "/start":
		return 1;
	case "/exit":
		return -1;
	}

	return 0;
}

func gameCommands(resp []string){
	if(len(resp) > 1){
		switch(resp[0]){
		case "/gameSet":
			if(len(resp) == 2 ){
				askForGameSet(resp[1]);
			}
	}
	}
}

func commandNotFound(command string){
	fmt.Println("/!\\ Command not found type '/"+ command +" -h' for help");
}