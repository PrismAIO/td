package updates

import (
	"sort"

	"github.com/PrismAIO/td/tg"
)

func sortUpdatesByPts(u []tg.UpdateClass) {
	sort.Stable(ptsSorter(u))
}

type ptsSorter []tg.UpdateClass

func (p ptsSorter) Len() int {
	return len(p)
}

func (p ptsSorter) Less(i, j int) bool {
	type (
		ptsType int
		compare struct {
			typ       ptsType
			channelID int64
			ptsDiff   int
		}
	)
	// Sorting order
	//
	// 0) Common updates without PTS.
	// 1) Common PTS updates
	// 2) Channel PTS updates (by channelID and by pts)
	// 3) QTS updates
	const (
		other ptsType = iota
		commonPts
		channelPts
		commonQts
	)
	getType := func(u tg.UpdateClass) compare {
		if pts, ptsCount, ok := isCommonPtsUpdate(u); ok {
			return compare{typ: commonPts, ptsDiff: pts - ptsCount}
		}
		if channelID, pts, ptsCount, ok, _ := isChannelPtsUpdate(u); ok {
			return compare{typ: channelPts, channelID: channelID, ptsDiff: pts - ptsCount}
		}
		if qts, ok := isCommonQtsUpdate(u); ok {
			return compare{typ: commonQts, ptsDiff: qts}
		}
		return compare{typ: other}
	}

	a, b := getType(p[i]), getType(p[j])
	switch {
	case a.typ < b.typ:
		return true
	case a.typ > b.typ:
		return false
	}

	// a.typ == b.typ
	switch a.typ {
	case other:
		// Keep original order
	case commonPts, commonQts:
		return a.ptsDiff < b.ptsDiff
	case channelPts:
		if a.channelID < b.channelID {
			return true
		}
		if a.channelID > b.channelID {
			return false
		}
		return a.ptsDiff < b.ptsDiff
	}

	return false
}

func (p ptsSorter) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
