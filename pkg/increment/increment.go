package increment

import (
	"sync"
)

var GlobalVar = 0

func Increment(n int) {
	if n <= 0 {
		return
	}

	ch := make(chan struct{}, 1)

	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(n int) {
			defer wg.Done()

			<-ch

			for i := 0; i < n; i++ {
				GlobalVar++
			}

			ch <- struct{}{}
		}(i + 1)
	}

	ch <- struct{}{}

	wg.Wait()
}
