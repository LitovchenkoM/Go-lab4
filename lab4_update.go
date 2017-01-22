package main

import (
    "fmt"
    "time"
)

type Token struct {
	data string
	recipient int
}

 
func main() {
	var count = 10  //кол-во потоков
	var token Token;
	token.recipient = 3
	token.data = "token"
	if  count < token.recipient {
		count = token.recipient
	}
	var channels = make([]chan Token, count+1)
	recipient:= make(chan Token)
	for i:=0; i<= count; i++ {
		channels[i] = make(chan Token)
	}
	for i := 0; i < count; i++ {
		go transfer(channels[i], channels[i+1], i, recipient) 
	}
	channels[0] <- token
	fmt.Println(<-recipient)
	time.Sleep(1 * 1e9)   //для вывода всех ожидающих потоков 
}

func transfer (sender chan  Token,  recipient chan Token, n int, last chan Token){
	fmt.Println(" Waiting data from channel ", n+1)
	token := <-sender
	if n+1 == token.recipient {
		fmt.Println("Goroutine ",n+1," is a recipient, get token")
		last <- token
	} else {
		fmt.Println("Goroutine ",n+1,"send next, because recipient: ", token.recipient)
		recipient <- token
	}
}


// https://play.golang.org/p/YAoKcBU-dP