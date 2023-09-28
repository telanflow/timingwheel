package timingwheel

type option struct {
	timeWheel bool
	minHeap   bool
	skiplist  bool
	rbtree    bool
}

type Option func(c *option)

func WithTimeWheel() Option {
	return func(o *option) {
		o.timeWheel = true
	}
}

func WithMinHeap() Option {
	return func(o *option) {
		o.minHeap = true
	}
}

func WithSkipList() Option {
	return func(o *option) {
		o.skiplist = true
	}
}

func WithRbtree() Option {
	return func(o *option) {
		o.rbtree = true
	}
}
