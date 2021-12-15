package Echarts

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/components"
)

func genPages() *components.Page {
	page := components.NewPage()
	page.AddCharts(
		lineBase(),
		lineMarkPoint(),
		lineArea(),
		lineOverlap(),
		scatterBase(),
		scatterShowLabel(),
		scatterSplitLine(),
		esEffectStyle(),
	)
	return page
}

type PageSSELayoutEcharts struct{}

func (PageSSELayoutEcharts) Echarts() {
	page := genPages()
	f, err := os.Create("charts/html/page_SSE_layout.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
