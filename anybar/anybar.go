package anybar

import (
	"fmt"
	"net"
	"context"
	"time"
)

type Anybar struct {
	port string
	icon string
}

func (a *Anybar) render() {
	address := fmt.Sprintf("localhost:%v", a.port)
	conn, err := net.Dial("udp", address)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Write([]byte(a.icon))
}

func (a *Anybar) reset() {
	a.icon = "black"
	a.render()
}

func InitInstance(port string, icon string) *Anybar {
	a := Anybar{port, icon}
	a.render()
	return &a
}

func (a *Anybar) ChangeColor(icon string) {
	a.icon = icon
	a.render()
}

func resetOnCancel(ctx context.Context, a *Anybar, ch chan struct{}) {
	go func() {
		select {
		case <-ctx.Done():
			a.reset()
			close(ch)
		}
	}()
}

func (a *Anybar) Blink(ctx context.Context, icon string, delay time.Duration) chan struct{} {
	ch := make(chan struct{})
	resetOnCancel(ctx, a, ch)
	go func() {
		a.ChangeColor(icon)
		time.Sleep(delay * time.Millisecond)
		a.reset()

		close(ch)
	}()
	return ch
}