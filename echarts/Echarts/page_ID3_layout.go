package Echarts

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/components"
)

func genID3Pages() *components.Page {
	page := components.NewPage()
	page.AddCharts(
		treeBase(),
	)
	return page
}

type PageID3LayoutEcharts struct{}

func (PageID3LayoutEcharts) Echarts() {
	page := genPages()
	f, err := os.Create("charts/html/page_ID3_layout.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
