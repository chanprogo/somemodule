package main

// IoHandlerFactory ...
type IoHandlerFactory interface {
	CreateIoHandler() IoHandler
}

// DefaultIoHandlerFactory ...
type DefaultIoHandlerFactory struct {
}

// CreateIoHandler ...
func (defaultFactory *DefaultIoHandlerFactory) CreateIoHandler() IoHandler {
	return new(DefaultIoHandler)
}
