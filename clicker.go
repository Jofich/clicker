package clicker

import (
	"context"
	"time"
)

func NewClicker() Clicker {
	return &Mouse{}
}

// implements Mouse methods
type Clicker interface {
	SetPosition(x, y int)
	Move(x, y int)
	Click(a Args)
	GetPosition() (int, int)
	StartClicking(ctx context.Context, ClickInterval time.Duration, a Args)
}
