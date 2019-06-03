// gls
// author: baoqiang
// time: 2019-04-19 14:45
package gls

import (
	"fmt"
	"sync"
)

var local ThreadLocal

func init() {
	local.m = make(map[uint]map[string]interface{}, 0)
}

func HelloGls() {
	fmt.Print("hello from local")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()
			//defer local.Clean()

			defer func() {
				//goid, _ := gls.GetGoroutineId()
				goid := GetGoid()
				fmt.Printf("[goroutine %d] idx: %d, got number: %v\n", goid, idx, local.Get("number"))
			}()

			local.Put("number", idx+100)
		}(i)
	}

	wg.Wait()
}
