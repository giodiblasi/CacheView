package main

/*EventObserver ...*/
type EventObserver interface {
	OnSet(key string)
	OnDeleted(key string)
	OnExpired(key string)
	OnError(err error)
}
