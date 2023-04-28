package main
    // "net/http"
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
    // http.HandleFunc("/join", joinHandler);
    for {
        resp := ask();
		if(!homeCommands(resp)){
			break;
		}
	}
}