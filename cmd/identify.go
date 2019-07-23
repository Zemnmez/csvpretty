package main

import (
	"context"

	"github.com/zemnmez/cardauth"
)

func main() {
	if err := do(); err != nil {
		return
	}
}

func do() (err error) {
	_, err = cardauth.DefaultClient.Identify(context.Background())
	if err != nil {
		return
	}

	return
}
