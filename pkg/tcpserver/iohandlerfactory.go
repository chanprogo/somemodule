package iohandler

// IoHandlerFactory ...
type IoHandlerFactory interface {
	CreateIoHandler() IoHandler
}

// DefaultIoHandlerFactory ...
type DefaultIoHandlerFactory struct {
}

func (defaultFactory *DefaultIoHandlerFactory) CreateIoHandler() IoHandler {
	return new(DefaultIoHandler)
}
