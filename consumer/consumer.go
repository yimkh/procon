package consumer

import (
	"fmt"
	"time"

	"github.com/yimkh/procon/model"
)

//Consumer is consumer
func Consumer(env *model.Environment) {
	if model.Wait(&env.Full) == false {
		fmt.Println("buffer is empty")
		return
	}

	if model.Wait(&env.Mutex) == false {
		fmt.Println("buffer is used")
		model.Signal(&env.Full)
		return
	}

	fmt.Println("consuming")

	get(env.Buffer, env.Out)
	out(&env.Out)

	//consume time
	time.Sleep(2 * time.Second)
	fmt.Println("consuming")

	model.Signal(&env.Mutex)
	model.Signal(&env.Empty)

	fmt.Println(env)
}

func out(out *int) {
	*out = (*out + 1) % 5
}

func get(Buffer [5]int, n int) {
	fmt.Println(Buffer[n])
	Buffer[n] = 0
}
