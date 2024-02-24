package key

import "github.com/micmonay/keybd_event"

type Key int

const (
	UP    Key = keybd_event.VK_UP
	DOWN  Key = keybd_event.VK_DOWN
	LEFT  Key = keybd_event.VK_LEFT
	RIGHT Key = keybd_event.VK_RIGHT
)

func (k Key) Int() int {
	return int(k)
}

func (k Key) String() string {
	switch k {
	case UP:
		return "↑"
	case DOWN:
		return "↓"
	case LEFT:
		return "←"
	case RIGHT:
		return "→"
	default:
		return ""
	}
}
