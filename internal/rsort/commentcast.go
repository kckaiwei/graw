// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package rsort

import "github.com/turnage/graw/internal/data"

type commentsThingImpl struct {
	e *data.Comment
}

func (g commentsThingImpl) Name() string { return g.e.Name }

func (g commentsThingImpl) Birth() uint64 { return g.e.CreatedUTC }

func commentsAsThings(gs []*data.Comment) []redditThing {
	things := make([]redditThing, len(gs))
	for i, g := range gs {
		things[i] = &commentsThingImpl{g}
	}
	return things
}
