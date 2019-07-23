package cardauth // import "github.com/zemnmez/cardauth"

import (
	"context"
	"fmt"

	"github.com/sf1/go-card/smartcard"
)

type Client struct{}

func (i *Identifier) Init() (err error) {
	return nil
}

// Identify returns a unique, secure identifier for
// a presented card
func (c Client) Identify(ctx context.Context) (identifier []byte, err error) {
	crd, err := smartcard.EstablishContext()
	defer ctx.Release() // if an error occurs
	if err != nil {
		return
	}

	reader, err := crd.WaitForCardPresent()
	if err != nil {
		return
	}

	fmt.Printf("Card ATR: %s\n", card.ATR())
}
