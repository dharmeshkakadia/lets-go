package main

import "fmt"
import "time"
import "math/rand"
import "sync/atomic"

type readop struct {
	key     int
	results chan int
}

type writeop struct {
	key    int
	val    int
	result chan bool
}

func main() {
	var totalReads uint64 = 0
	var totalWrites uint64 = 0

	reads := make(chan readop)
	writes := make(chan writeop)
	go func() {
		m := make(map[int]int)
		for {
			select {
			case r := <-reads:
				r.results <- m[r.key]
			case w := <-writes:
				m[w.key] = w.val
				w.result <- true
			}
		}
	}()

	// send read requests
	for i := 0; i < 100; i++ { // 100 parrellel reuesters
		go func() {
			for { // every requester generating read requests
				r := readop{key: rand.Intn(5), results: make(chan int)}
				reads <- r  // send the read requst
				<-r.results // discard the result so it can process next
				atomic.AddUint64(&totalReads, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// send write reuests
	for i := 0; i < 10; i++ {
		go func() {
			for {
				w := writeop{key: rand.Intn(5), val: rand.Intn(100), result: make(chan bool)}
				writes <- w // send the write requet
				<-w.result
				atomic.AddUint64(&totalWrites, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("#total reads:", atomic.LoadUint64(&totalReads))
	fmt.Println("#total writes:", atomic.LoadUint64(&totalWrites))
}
