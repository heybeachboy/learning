package main

import (
	"fmt"
	"learning/Algorithm"
	"learning/MessageQueue"
	"log"
	"time"
)
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func main() {
	 data := Algorithm.CreateRandomList(10,10000)
	 fmt.Println("sort before:",data)
	  s := Algorithm.NewSort(data)
	  resp :=s.MergeSort()
	  fmt.Println("sort after:",resp)
	  fmt.Println("execute time:",s.Interval)
	  Rabbit :=MessageQueue.NewRabbitMqClient(MessageQueue.MqUrl,MessageQueue.Exchange,MessageQueue.QueueName)
	  err := Rabbit.Push([]byte(time.Now().String()))
	  err = Rabbit.Push([]byte(time.Now().String()))
	  err = Rabbit.Push([]byte(time.Now().String()))
	  Rabbit.Receiver("send")
	  fmt.Println("reply:",err)
	  var sum int32
	  //sum = 21608
	  //sum = 21191
	    sum = 21531
	  fmt.Println(fmt.Sprintf("%x",sum))

	  }
