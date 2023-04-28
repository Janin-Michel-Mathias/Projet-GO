package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func ask() []string{
    scanner := bufio.NewScanner(os.Stdin);
    fmt.Print(">>>");
        scanner.Scan();
        resp := strings.Split(scanner.Text(), " ");
        return resp;
}

func main(){
    for {
        resp := ask();
		if(!homeCommands(resp)){
			break;
		}
	}
}