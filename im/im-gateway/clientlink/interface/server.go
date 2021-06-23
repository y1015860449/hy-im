package _interface

type Server interface {
	Start() error
	Stop() error
}
