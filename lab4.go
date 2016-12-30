

package main

import (
    "fmt"
    "time"
)

type Token struct {
	data string
	recipient int
}

var count = 7 //����� ��������
func main() {
	var token Token;
	token.recipient = count
	token.data = "token"
	first := make(chan Token) 
	last := make(chan Token)
	go send(first, token) //�����: 1 ����� �������� token
	last = channels(first,count-1) // ��������� �� ������� ������ ����������
	go get(last, count-1)  //�����: ��������� ����� �������� �� �������������� token
	time.Sleep(1 * 1e9) // main ����� ����������� ������ goroutines
}

func send (channel chan Token, t Token){
	fmt.Println("Send token to 1 channel")
	channel <- t
}

func transfer (sender <-chan  Token,  recipient chan<- Token, n int){
	fmt.Println("get from channel ", n+1)
	fmt.Println("send into channel ", n+2)
	for {
		token := <-sender
		recipient <- token
	}
}

func get(last <-chan Token, n int){
	for {
		fmt.Println("The recipient get token=",<-last, " from the last channel ", n+1)
	}
}

func channels(in chan Token,n int) (out chan Token){
	out = make(chan Token)
	previous := in
	for i := 0; i < n; i++ {   
		next := make(chan Token)
		go transfer(previous, next, i) //� ����� ��������� count-1 �������, ������� ���������� token
		previous = next  
	}
	out = previous
	return out
}
