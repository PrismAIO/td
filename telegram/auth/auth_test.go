package auth

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PrismAIO/td/internal/testutil"
	"github.com/PrismAIO/td/tg"
	"github.com/PrismAIO/td/tgmock"
)

const (
	testAppID   = 1
	testAppHash = "hash"
)

func testClient(invoker tg.Invoker) *Client {
	return &Client{
		api:     tg.NewClient(invoker),
		rand:    rand.Reader,
		appID:   testAppID,
		appHash: testAppHash,
	}
}

func mockClient(t *testing.T) (*tgmock.Mock, *Client) {
	mock := tgmock.New(t)
	return mock, NewClient(tg.NewClient(mock), testutil.ZeroRand{}, testAppID, testAppHash)
}

func mockTest(cb func(
	a *require.Assertions,
	mock *tgmock.Mock,
	client *Client,
)) func(t *testing.T) {
	return func(t *testing.T) {
		a := require.New(t)
		m, client := mockClient(t)

		cb(a, m, client)
	}
}
