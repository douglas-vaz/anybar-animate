package main

import (
	"github.com/douglas-vaz/anybar-animate/anybar"
	"flag"
	"context"
	"time"
)

func main() {
	icon := flag.String("icon", "black", "Provide one of the anybar icon values")
	port := flag.String("port", "1738", "UDP port the Anybar instance is listenening to")
	flag.Parse()
	a := anybar.InitInstance(*port, "black")

	ctx, _ := context.WithTimeout(context.Background(), 1000 * time.Millisecond)
	<- a.Blink(ctx, *icon, 2000)

}

