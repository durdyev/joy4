package format

import (
	"github.com/durdyev/joy4/av/avutil"
	"github.com/durdyev/joy4/format/aac"
	"github.com/durdyev/joy4/format/flv"
	"github.com/durdyev/joy4/format/mp4"
	"github.com/durdyev/joy4/format/rtmp"
	"github.com/durdyev/joy4/format/rtsp"
	"github.com/durdyev/joy4/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
