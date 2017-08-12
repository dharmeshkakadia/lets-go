package main

import "fmt"
import "sync"
import "math/rand"
import "sync/atomic"
import "time"

func main() {
	map1 := make(map[int]int)
	mutex := &sync.Mutex{}
	var readOps uint64 = 0
	var writeOps uint64 = 0
	for i := 0; i < 100; i++ {
		go func() {
			for {
				total := 0
				key := rand.Intn(5)
				mutex.Lock()
				total += map1[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				mutex.Lock()
				map1[key] = rand.Intn(100)
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("#read ops", atomic.LoadUint64(&readOps))
	fmt.Println("#write ops", atomic.LoadUint64(&writeOps))

	mutex.Lock()
	fmt.Println(map1)
	mutex.Unlock()
}
