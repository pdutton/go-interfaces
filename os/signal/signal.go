package signal

import (
	"context"
	"os"
	"os/signal"
)

type Signal interface {
	Ignore(sig ...os.Signal)
	Ignored(sig os.Signal) bool
	Notify(c chan<- os.Signal, sig ...os.Signal)
	NotifyContext(parent context.Context, signals ...os.Signal) (ctxt context.Context, stop context.CancelFunc)
	Reset(sig ...os.Signal)
	Stop(c chan<- os.Signal)
}

type signalFacade struct{}

func NewSignal() Signal {
	return signalFacade{}
}

func (_ signalFacade) Ignore(sig ...os.Signal) {
	signal.Ignore(sig...)
}

func (_ signalFacade) Ignored(sig os.Signal) bool {
	return signal.Ignored(sig)
}

func (_ signalFacade) Notify(c chan<- os.Signal, sig ...os.Signal) {
	signal.Notify(c, sig...)
}

func (_ signalFacade) NotifyContext(parent context.Context, signals ...os.Signal) (ctxt context.Context, stop context.CancelFunc) {
	return signal.NotifyContext(parent, signals...)
}

func (_ signalFacade) Reset(sig ...os.Signal) {
	signal.Reset(sig...)
}

func (_ signalFacade) Stop(c chan<- os.Signal) {
	signal.Stop(c)
}
