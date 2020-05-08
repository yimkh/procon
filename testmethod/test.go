package testmethod

import (
	"time"

	"github.com/yimkh/procon/consumer"
	"github.com/yimkh/procon/model"
	"github.com/yimkh/procon/producer"
)

//Testfull is to test full
func Testfull(env *model.Environment) {
	for i := 0; i < 6; i++ {
		producer.Producer(env, 10)
	}
}

//TestUsed is to test used
func TestUsed(env *model.Environment) {
	//go is goroutine
	go producer.Producer(env, 10)
	go producer.Producer(env, 9)
	go producer.Producer(env, 7)

	time.Sleep(5 * time.Second)
}

//TestEmpty is to test empty
func TestEmpty(env *model.Environment) {
	consumer.Consumer(env)
}

//TestProCon is to test producer  and consumer
func TestProCon(env *model.Environment) {
	go producer.Producer(env, 10)
	time.Sleep(3 * time.Second)

	go producer.Producer(env, 9)
	time.Sleep(3 * time.Second)

	go producer.Producer(env, 7)
	time.Sleep(3 * time.Second)

	go consumer.Consumer(env)
	time.Sleep(3 * time.Second)

	go consumer.Consumer(env)
	time.Sleep(3 * time.Second)

	time.Sleep(20 * time.Second)
}
