package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, pool)
	var cnt int64
	result := make([]user, n, n)

	for cnt = 0; cnt < n; cnt++ {
		id := cnt
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			result[id] = getOne(id)
			<-sem
			wg.Done()
		}()
	}
	wg.Wait()
	return result
}
