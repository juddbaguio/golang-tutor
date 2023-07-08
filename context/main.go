package main

import (
	"context"
	"log"
)

type userRepoTxKey struct{}
type productRepoTxKey struct{}

func main() {
	ctx := context.WithValue(context.Background(), userRepoTxKey{}, "judd")
	ctx = context.WithValue(ctx, productRepoTxKey{}, "baguio")

	log.Println(ctx.Value(userRepoTxKey{}))
	log.Println(ctx.Value(productRepoTxKey{}))

}
