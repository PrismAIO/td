package telegram

import (
	"sync"

	"github.com/PrismAIO/td/internal/mt"
	"github.com/PrismAIO/td/internal/proto"
	"github.com/PrismAIO/td/internal/tmap"
	"github.com/PrismAIO/td/tg"
)

// Port is default port used by telegram.
const Port = 443

var (
	typesMap  *tmap.Map
	typesOnce sync.Once
)

func getTypesMapping() *tmap.Map {
	typesOnce.Do(func() {
		typesMap = tmap.New(
			tg.TypesMap(),
			mt.TypesMap(),
			proto.TypesMap(),
		)
	})
	return typesMap
}
