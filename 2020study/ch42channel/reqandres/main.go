package main

import (
	"fmt"
	"sync"
)

type Request struct {
	Num  int
	Resp chan Response
}

type Response struct {
	Num      int
	WorkerID int
}

// 보통 실무에서 int형을 바로 넘겨주는 것이 아닌 구조체를 만들어서 넘겨줌
// 아래는 id를 넘겨준다.
func PlusOneService(reqs <-chan Request, workerID int) {
	for req := range reqs {
		go func(req Request) {
			defer close(req.Resp)
			req.Resp <- Response{req.Num + 1, workerID}
		}(req)
	}
}

func main() {
	reqs := make(chan Request)
	defer close(reqs)
	for i := 0; i < 3; i++ {
		go PlusOneService(reqs, i)
	}
	var wg sync.WaitGroup
	for i := 3; i < 53; i += 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resp := make(chan Response)
			reqs <- Request{i, resp}
			fmt.Println(i, "=>", <-resp)
		}(i)
	}
	wg.Wait()
}
