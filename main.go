package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan Book)
	dbCh := make(chan Book)
	for i:=0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryDatabase(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)
		//create one goroutine per query to handle response
		go func(cacheCh, dbCh <-chan Book, ) {
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				<- dbCh
			case b := <-dbCh:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)

		//without this line & without using mutexes, a race condition will happen due to concurrent
		//read and writes into the map
		//to detect where exactly the race happens run with the command $ go run --race .
		time.Sleep(100 * time.Millisecond)

	}
	wg.Wait()

	fmt.Println("***********************************")

	wg2 := &sync.WaitGroup{}
	ch := make(chan int, 5)

	wg2.Add(2)
	go func(ch chan int, wg2 *sync.WaitGroup) {
		//another way to read from from a channel
		// fmt.Println(<-ch)
		for msg := range ch {
			fmt.Println(msg)
		}
		wg2.Done()
	}(ch, wg2)
	go func(ch chan int, wg2 *sync.WaitGroup) {
		for i:=0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg2.Done()
	}(ch, wg2)
	wg2.Wait()
}

func queryCache(id int, m *sync.RWMutex)  (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Microsecond)
	for _ , b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}