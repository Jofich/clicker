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
	WheelLeft  = "wheelLeft"  // not supported by all OS
	WheelRight = "wheelRight" // not supported by all OS
)

// click parameters needed by robotgo
type Args struct {
	MouseButton string
	Double      bool
}

// Mouse provides methods for controlling the cursor and clicks.
// Wrapper over the main mouse. No one forbids having several such 
// structures, but be careful, because with this approach, 
// there is a risk of losing clicks
type Mouse struct {
	PosX int // may be negative for working with multiple screens
	PosY int // may be negative for working with multiple screens
}

//returns main mouse position
func GetGlobalMousePos() (int, int) {
	return robotgo.Location()
}

//set postion for Mouse struct
func (m *Mouse) SetPosition(x, y int) {
	m.PosX = x
	m.PosY = y
}

//move the main mouse and specifies new coordinates for the structure
func (m *Mouse) Move(x, y int) {
	m.SetPosition(x, y)
	robotgo.Move(x, y)
}

// move the main mouse to the set coordinates and click
func (m *Mouse) Click(a Args) {
	m.Move(m.PosX, m.PosY)
	robotgo.Click(a.MouseButton, a.Double)
}

//return position of Mouse struct
func (m Mouse) GetPosition() (int, int) {
	return m.PosX, m.PosY
}

// start clicking on the specified position
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
