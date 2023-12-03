package main

import(
	"fmt"
	"github.com/suhasbhagate/GoCode/GoCodeRevision/QueueDemo/queue"
)

func main(){
	var qu queue.Queue
	qu.Enqueue(100)
	qu.Enqueue(200)

	if !qu.IsEmpty(){
		fmt.Println(qu.Dequeue())
		fmt.Println(qu)
	}
	if !qu.IsEmpty(){
		fmt.Println(qu.Dequeue())
		fmt.Println(qu)
	}
	if !qu.IsEmpty(){
		fmt.Println(qu.Dequeue())
		fmt.Println(qu)
	}
}