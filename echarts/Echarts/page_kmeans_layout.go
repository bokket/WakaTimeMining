package Echarts

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/components"
)

func genKmeansPages() *components.Page {
	page := components.NewPage()
	page.AddCharts(
		scatter3DBase(),
		scatterBase(),
		line3DBase(),
		line3DAutoRotate(),
	)
	return page
}

type PageKmeansLayoutEcharts struct{}

func (PageKmeansLayoutEcharts) Echarts() {
	page := genPages()
	f, err := os.Create("charts/html/page_kmeans_layout.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

