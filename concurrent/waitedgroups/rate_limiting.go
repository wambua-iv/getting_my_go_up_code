package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

func createRateLimit (){
	var wg sync.WaitGroup
	total, max := 24,6

	for i := 0; i <= total; i+= max {
		limit := max

		if i + max > total {
			limit = total - i
		}

		wg.Add(limit)
		for j := 0; j< limit; j++ {
			go func (j int) {
				defer wg.Done()
				conn, err := net.Dial("tcp", ":8080")
				if err != nil {
					log.Fatal("Connection failed")
				}

				stringed, err := io.ReadAll(conn)
				if err != nil {
					log.Fatal("Connection failed")
				}
				fmt.Printf("%v %v %v \n",i +j + 1, j, string(stringed))
			}(j)
		} 
		wg.Wait()

	}
}

func server () {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Connection creation failed")
	}

	var connects int32

	for {
		connect, err := listener.Accept()

		if err != nil{
			continue
		}

		connects++ 

		go func () {
			defer func(){
				_ = connect.Close()
				atomic.AddInt32(&connects, -1)
			}()

			if atomic.LoadInt32(&connects) > 6 {
				return
			}

			time.Sleep(time.Second)
			_, err = connect.Write([]byte("Made it Out of the Block"))

			if err != nil {
				log.Fatal("Connection write failed")
			}
		}()
	}


}