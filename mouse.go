package clicker

import (
	"context"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	Left       = "left"
	Right      = "right"
	Center     = "center"
	WheelDown  = "wheelDown"
	WheelUp    = "wheelUp"
	WheelLeft  = "wheelLeft"
	WheelRight = "wheelRight"
)

type Args struct {
	MouseButton string
	Double      bool
}

type Mouse struct {
	PosX int
	PosY int
}

func GetGlobalMousePos() (int, int) {
	return robotgo.Location()
}

func (m *Mouse) SetPosition(x, y int) {
	m.PosX = x
	m.PosY = y
}

func (m *Mouse) Move(x, y int) {
	m.SetPosition(x, y)
	robotgo.Move(x, y)
}

func (m *Mouse) Click(a Args) {
	m.Move(m.PosX, m.PosY)
	robotgo.Click(a.MouseButton, a.Double)
}

func (m Mouse) GetPosition() (int, int) {
	return m.PosX, m.PosY
}

func (m *Mouse) StartClicking(ctx context.Context, ClickInterval time.Duration, a Args) {
	ticker := time.NewTicker(ClickInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			m.Click(a)
		case <-ctx.Done():
			return
		}
	}
}
