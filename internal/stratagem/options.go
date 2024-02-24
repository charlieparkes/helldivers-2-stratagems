package stratagem

import "time"

type Opt func(*Stratagems)

func WithPressFor(pf time.Duration) Opt {
	return func(s *Stratagems) {
		s.pressFor = pf
	}
}

func WithDelay(d time.Duration) Opt {
	return func(s *Stratagems) {
		s.delay = d
	}
}

func WithRandomize(r bool) Opt {
	return func(s *Stratagems) {
		s.randomize = r
	}
}

func WithVerbose(v bool) Opt {
	return func(s *Stratagems) {
		s.verbose = v
	}
}
