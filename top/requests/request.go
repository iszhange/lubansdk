package requests

type Request interface {
	SetParam(key, value string)
	GetMethod() string
	GetApiParams() Parameter
}
