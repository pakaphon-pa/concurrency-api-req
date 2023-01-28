package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

const MAX_GOROUTINES = 5

func main() {
	fmt.Println("Start Jobs...")
	defer timeTrack(time.Now(), "Job")
	pageChannel := make(chan int)
	wg := sync.WaitGroup{}
	initialWorker(pageChannel, &wg)
	getData(pageChannel)
	wg.Wait()
	fmt.Println("Process successfully...")
	// concurrent()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func getData(pageChannel chan int) {
	page := 1

	for {

		res := requestApi(page)

		pageChannel <- res
		if res == 10 {
			close(pageChannel)
			break
		}
		page++

	}

}

func requestApi(page int) int {
	time.Sleep(10 * time.Millisecond)

	return page
}

func initialWorker(pageChannel chan int, wg *sync.WaitGroup) {

	for i := 0; i < MAX_GOROUTINES; i++ {
		wg.Add(1)
		go func(pageChannel chan int) {
			defer wg.Done()
			for s := range pageChannel {

				fmt.Println("Process channel :" + strconv.Itoa(s))
				format(s)
			}
		}(pageChannel)
	}

}

func format(page int) {
	time.Sleep(1000 * time.Millisecond)
	save(page)
	fmt.Println("Format data in page : " + strconv.Itoa(page))

}

func save(page int) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Save data in page : " + strconv.Itoa(page))

}

func concurrent() {
	page := 1

	for {

		res := requestApi(page)
		format(res)

		if res == 5 {
			break
		}
		page++

	}
}
