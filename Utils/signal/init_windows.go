package signal

import (
	"syscall"
)

func (set *SignalSet) RegisterExit(handler SignalHandler) {
	if _, found := set.m[syscall.SIGKILL]; !found {
		set.m[syscall.SIGKILL] = handler
	}
	if _, found := set.m[syscall.SIGTERM]; !found {
		set.m[syscall.SIGTERM] = handler
	}
	if _, found := set.m[syscall.SIGQUIT]; !found {
		set.m[syscall.SIGQUIT] = handler
	}
	if _, found := set.m[syscall.SIGINT]; !found {
		set.m[syscall.SIGINT] = handler
	}
}
func (set *SignalSet) RegisterReload(handler SignalHandler) {

}
