package stratagem

import (
	"cmp"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/charlieparkes/helldivers-2-stratagems/internal/key"
	"github.com/charlieparkes/helldivers-2-stratagems/internal/random"
	"github.com/micmonay/keybd_event"
)

type Name string

type Keystrokes []key.Key

func (k Keystrokes) String() string {
	keys := []string{}
	for _, k := range k {
		keys = append(keys, k.String())
	}
	return strings.Join(keys, ", ")
}

type Stratagems struct {
	keyboard  keybd_event.KeyBonding
	patterns  map[Name]Keystrokes
	pressFor  time.Duration
	delay     time.Duration
	randomize bool
	randMax   int // max milliseconds to randomly add to delay and pressFor
	verbose   bool
}

func NewStratagems(opts ...Opt) *Stratagems {
	k, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	s := &Stratagems{
		keyboard:  k,
		patterns:  make(map[Name]Keystrokes),
		pressFor:  time.Millisecond * 100,
		delay:     time.Millisecond * 50,
		randomize: true,
		randMax:   100,
		verbose:   false,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Stratagems) Add(n Name, k ...key.Key) {
	var ks Keystrokes
	s.patterns[n] = append(ks, k...)
}

// Call in that stratagem!
// Must hold ctrl yourself, I don't want to ruin the fun, just make my StreamDeck a tool for democracy.
func (s *Stratagems) Call(n Name) error {
	if s.verbose {
		s.Print(n)
	}
	p, ok := s.patterns[n]
	if !ok {
		return fmt.Errorf("no strategem named %s", n)
	}
	for i, k := range p {
		s.Press(k, s.pressFor)
		if i < len(p)-1 {
			time.Sleep(s.delay + random.Duration(s.randMax))
		}
	}
	return nil
}

func (s *Stratagems) Press(k key.Key, pressFor time.Duration) {
	s.keyboard.SetKeys(k.Int())
	s.keyboard.Press()
	time.Sleep(pressFor + random.Duration(s.randMax))
	s.keyboard.Release()
	s.keyboard.Clear()
}

func (s *Stratagems) Print(n Name) {
	k, ok := s.patterns[n]
	if ok {
		fmt.Println(n, ": ", k)
	}
}

func (s *Stratagems) PrintAll() {
	for _, n := range sortedKeys(s.patterns) {
		fmt.Println("\t", n, ": ", s.patterns[n])
	}
}

func sortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}
