package ed

type ESPSenderInterface interface {
	Send(e Email) (interface{}, error)
}
