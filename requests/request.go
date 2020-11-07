package requests

type Request interface {
	GetMethod() string
	GetApiParams() Parameter
}
