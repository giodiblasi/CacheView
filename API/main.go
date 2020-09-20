package main

import "fmt"

type console struct{}

func (target *console) OnDeleted(key string) {
	fmt.Printf("Deleted %s\n", key)
}

func (target *console) OnExpired(key string) {
	fmt.Printf("Exipred %s\n", key)
}

func (target *console) OnSet(key string) {
	fmt.Printf("Set %s\n", key)
}

func (target *console) OnError(err error) {
	fmt.Println(err)
}

func main() {
	fmt.Println("Hello, Redis.")
	cons := &console{}
	end := RedisListener(":5000", cons)
	<-end
}
