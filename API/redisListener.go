package main

import (
	"strings"

	"github.com/gomodule/redigo/redis"
)

func getEvent(msg redis.Message, observer *EventObserver) {
	cmd := strings.Split(msg.Channel, ":")[1]
	key := string(msg.Data)
	if cmd == "set" {
		(*observer).OnSet(key)
		return
	}
	if cmd == "del" {
		(*observer).OnDeleted(key)
		return
	}
	if cmd == "expired" {
		(*observer).OnExpired(key)
		return
	}
}

/*RedisListener ...*/
func RedisListener(addr string, observer EventObserver) chan error {
	c, err := redis.Dial("tcp", addr)
	end := make(chan error)
	if err != nil {
		observer.OnError(err)
		end <- err
		return end
	}
	//defer c.Close()
	psc := redis.PubSubConn{Conn: c}

	if err := psc.PSubscribe("__keyevent*__:*"); err != nil {
		observer.OnError(err)
		end <- err
		return end
	}

	go func() {
		for {
			switch msg := psc.Receive().(type) {
			case error:
				observer.OnError(err)
				end <- msg
				break
			case redis.Message:
				getEvent(msg, &observer)
				break
			}
		}
	}()
	return end
}
