package _interface

type EventBus interface {
	Publish(event string, payload any)
	Subscribe(event string, handler func(payload any))
}
