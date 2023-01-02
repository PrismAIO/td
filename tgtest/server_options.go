package tgtest

import (
	"io"
	"time"

	"go.uber.org/zap"

	"github.com/PrismAIO/td/clock"
	"github.com/PrismAIO/td/internal/crypto"
	"github.com/PrismAIO/td/internal/mt"
	"github.com/PrismAIO/td/internal/mtproto"
	"github.com/PrismAIO/td/internal/proto"
	"github.com/PrismAIO/td/internal/tmap"
	"github.com/PrismAIO/td/tg"
	"github.com/PrismAIO/td/transport"
)

// ServerOptions of Server.
type ServerOptions struct {
	// DC ID of this server. Default to 2.
	DC int
	// Random is random source. Defaults to rand.Reader.
	Random io.Reader
	// Logger is instance of zap.Logger. No logs by default.
	Logger *zap.Logger
	// Codec constructor.
	// Defaults to nil (underlying transport server detects protocol automatically).
	Codec func() transport.Codec
	// Clock to use. Defaults to clock.System.
	Clock clock.Clock
	// MessageID generator. Creates a new proto.MessageIDGen by default.
	// Clock will be used for creation.
	MessageID mtproto.MessageIDSource
	// Types map, used in verbose logging of incoming message.
	Types *tmap.Map
	// ReadTimeout is a connection read timeout.
	ReadTimeout time.Duration
	// ReadTimeout is a connection write timeout.
	WriteTimeout time.Duration
}

func (opt *ServerOptions) setDefaults() {
	if opt.DC == 0 {
		opt.DC = 2
	}
	if opt.Random == nil {
		opt.Random = crypto.DefaultRand()
	}
	if opt.Logger == nil {
		opt.Logger = zap.NewNop()
	}

	// Ignore opt.Codec, will be handled by transport.NewCustomServer.

	if opt.Clock == nil {
		opt.Clock = clock.System
	}
	if opt.MessageID == nil {
		opt.MessageID = proto.NewMessageIDGen(opt.Clock.Now)
	}
	if opt.Types == nil {
		opt.Types = tmap.New(
			tg.TypesMap(),
			mt.TypesMap(),
			proto.TypesMap(),
		)
	}
	if opt.ReadTimeout == 0 {
		opt.ReadTimeout = 30 * time.Second
	}
	if opt.WriteTimeout == 0 {
		opt.WriteTimeout = 30 * time.Second
	}
}
