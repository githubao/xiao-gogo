// retry
// author: baoqiang
// time: 2019-04-02 11:15
package xretry

import (
	"errors"
	"fmt"
	"github.com/kamilsk/retry"
	"github.com/kamilsk/retry/strategy"
	"log"
	"math/rand"
	"time"
)

func HelloRetry() {
	fmt.Print("hello from local")

	if err := retry.Retry(nil, hardWork, strategy.Limit(20)); err != nil {
		panic(err)
	}
	fmt.Println("success")
}

func hardWork(uint) error {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)
	log.Printf("Got num: %d\n", i)
	if i != 9 {
		return errors.New("not big enough")
	}
	return nil
}
