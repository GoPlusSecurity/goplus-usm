package client

import "time"

type ClientOption interface {
	apply(*clientOption)
}

type clientOption struct {
	timeOut time.Duration
	key     string
}

type FuncClientOption struct {
	f func(*clientOption)
}

func (fdo *FuncClientOption) apply(do *clientOption) {
	fdo.f(do)
}

func newFuncDialOption(f func(*clientOption)) *FuncClientOption {
	return &FuncClientOption{
		f: f,
	}
}

// WithTimeOut returns a FuncClientOption that configures a timeout for dialing a
// ClientConn initially.
// default 10 * Second
func WithTimeOut(timeOut time.Duration) *FuncClientOption {
	return newFuncDialOption(func(c *clientOption) {
		c.timeOut = timeOut
	})
}

// WithKey returns a FuncClientOption that configures a key for dialing a
// ClientConn by token auth.
func WithKey(key string) *FuncClientOption {
	return newFuncDialOption(func(c *clientOption) {
		c.key = key
	})
}
