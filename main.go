package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

var (
	port = flag.Int("port", 50050, "The server port")
)

func main() {
	quotes := []string{
		"Don't communicate by sharing memory, share memory by communicating.",
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
		"Make the zero value useful.",
		"interface{} says nothing.",
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		"A little copying is better than a little dependency.",
		"Syscall must always be guarded with build tags.",
		"Cgo must always be guarded with build tags.",
		"Cgo is not Go.",
		"With the unsafe package there are no guarantees.",
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Don't just check errors, handle them gracefully.",
		"Design the architecture, name the components, document the details.",
		"Documentation is for users.",
		"Don't panic.",
	}
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	for {
		rand.Seed(time.Now().UnixNano())
		conn, err := lis.Accept()
		fmt.Println(lis.Addr())

		go func() {

			if err != nil {
				log.Fatal(err)
				fmt.Println("!!!!!!!!!!!!")
			}
			defer conn.Close()
			for {
				randomNumber := rand.Intn(len(quotes))

				conn.Write([]byte(quotes[randomNumber]))
				conn.Write([]byte("\n"))
				time.Sleep(time.Second * 3)
			}
		}()
	}
}
