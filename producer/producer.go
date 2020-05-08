package producer

import (
	"fmt"
	"time"

	"github.com/yimkh/procon/model"
)

//Producer is producer
func Producer(env *model.Environment, content int) {
	if model.Wait(&env.Empty) == false {
		fmt.Println("buffer is full")
		return
	}

	if model.Wait(&env.Mutex) == false {
		fmt.Println("buffer is used")
		model.Signal(&env.Empty)
		return
	}

	fmt.Println("produring")
	env.Buffer[env.In] = content
	in(&env.In)

	//produce time
	time.Sleep(2 * time.Second)
	fmt.Println("produred")

	model.Signal(&env.Mutex)
	model.Signal(&env.Full)

	fmt.Println(*env)
}

func in(in *int) {
	*in = (*in + 1) % 5
}
