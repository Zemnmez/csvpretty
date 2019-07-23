package cardauth // import "github.com/zemnmez/cardauth"

import (
	"context"
	"fmt"
	"sync"

	"github.com/sf1/go-card/smartcard"
	"github.com/zemnmez/cardauth/internal/apdu"
)

// DefaultClient is a Client singleton
var DefaultClient Client

// Client exposes an API for cardauth
type Client struct {
	initOnce *sync.Once
}

// Init forces the initialization of this Client.
// This is also done implicitly when any method is called
func (c *Client) Init() (err error) {
	c.initOnce.Do(func() {
		err = c.init()
	})

	return
}

func (c *Client) init() (err error) {
	return nil
}

// Identify returns a unique, secure identifier for
// a presented card
func (c Client) Identify(ctx context.Context) (identifier []byte, err error) {
	crd, err := smartcard.EstablishContext()
	defer crd.Release() // if an error occurs
	if err != nil {
		return
	}

	reader, err := crd.WaitForCardPresent()
	if err != nil {
		return
	}

	card, err := reader.Connect()
	defer card.Disconnect()

	fmt.Printf("Card ATR: %s\n", card.ATR())

	command := smartcard.SelectCommand(0xa0, 0x00, 0x00, 0x00, 0x62, 0x03, 0x01, 0xc, 0x01, 0x01)
	response, err := card.TransmitAPDU(command)
	if err != nil {
		return
	}
	fmt.Printf("Response: %s\n", response)
	return
}
