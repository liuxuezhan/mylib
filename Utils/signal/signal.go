package signal

import (
	"Utils/wlog"
	"errors"
	"os"
	"os/signal"
	"runtime"
)

type SignalHandler func()

type SignalSet struct {
	m map[os.Signal]SignalHandler
}

func SignalSetNew() *SignalSet {
	ss := new(SignalSet)
	ss.m = make(map[os.Signal]SignalHandler)
	return ss
}

func (set *SignalSet) Handle(sig os.Signal) (err error) {
	if _, found := set.m[sig]; found {
		set.m[sig]()
		return nil
	} else {
		return errors.New("No handler available for signal ")
	}
}
func Init(exit, reload func()) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	sigHandler := SignalSetNew()
	sigHandler.RegisterExit(exit)
	sigHandler.RegisterReload(reload)
	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan)
	go func(sigChan chan os.Signal) {
		for true {
			select {
			case sig := <-sigChan:
				wlog.Warning("name : ", sig.String())
				err := sigHandler.Handle(sig)
				if err != nil {
					wlog.Error(err.Error())
				}

			}
		}
	}(sigChan)
}
