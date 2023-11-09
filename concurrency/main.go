package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen() {

running:
	for {
		select {
		case msg := <-s.msgch:
			fmt.Printf("Message from %s: %s\n", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Print("The server is shutting down")
			break running
		default:

		}
	}

	fmt.Print("The server is shutting down")
}

func sendMessageToServer(msgch chan Message, payload string) {

	msg := Message{
		From:    "Justin",
		Payload: "Hello",
	}

	msgch <- msg

}

func gracefullyShutdown(quitch chan struct{}) {
	close(quitch)
}

func main() {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	go s.StartAndListen()
	go func() {
		time.Sleep(500 * time.Millisecond)
		sendMessageToServer(s.msgch, "Hello World")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		gracefullyShutdown(s.quitch)
	}()

	select {}
}

// func main() {
// 	now := time.Now()
// 	userID := 10
// 	respch := make(chan string, 3)
// 	wg := &sync.WaitGroup{}

// 	go fetchUserData(userID, respch, wg)
// 	wg.Add(1)
// 	go fetchuserRecommendations(userID, respch, wg)
// 	wg.Add(1)
// 	go fetchuserLikes(userID, respch, wg)
// 	wg.Add(1)

// 	wg.Wait()
// 	close(respch)
// 	for resp := range respch {
// 		fmt.Println(resp)
// 	}

// 	fmt.Println(time.Since(now))

// }

// func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(80 * time.Millisecond)
// 	respch <- "user data"
// 	wg.Done()
// }
// func fetchuserRecommendations(userID int, respch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(120 * time.Millisecond)
// 	respch <- "user recommendations"
// 	wg.Done()
// }
// func fetchuserLikes(userID int, respch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(50 * time.Millisecond)
// 	respch <- "user likes"
// 	wg.Done()
// }
