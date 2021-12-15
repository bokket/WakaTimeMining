package Echarts

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/components"
)

func genDbScanPages() *components.Page {
	page := components.NewPage()
	page.AddCharts(
		radarBase(),
		radarStyle(),
		radarLegendMulti(),
		radarLegendSingle(),
	)
	return page
}

type PageDbScanLayoutEcharts struct{}

func (PageDbScanLayoutEcharts) Echarts() {
	page := genPages()
	f, err := os.Create("charts/html/page_DbScan_layout.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
