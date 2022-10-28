package main

import (
	"context"
	"fmt"
)

type database map[string]bool

var db database = database{
	"jane": true,
}

func main() {
	ctx := context.Background()

	processRequest(ctx, "jane")
}

func processRequest(ctx context.Context, userid string) {
	// TODO: send userID information to checkMemberShip through context for
	// map lookup.
	ctxValue := context.WithValue(ctx, "userid", userid)
	ch := checkMemberShip(ctxValue)
	status := <-ch
	fmt.Printf("membership status of userid : %s : %v\n", userid, status)
}

// checkMemberShip - takes context as input.
// extracts the user id information from context.
// spins a goroutine to do map lookup
// sends the result on the returned channel.
func checkMemberShip(ctx context.Context) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		// do some database lookup
		status := db[ctx.Value("userid").(string)]
		ch <- status
	}()
	return ch
}
