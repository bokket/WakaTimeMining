package Echarts

import (
	"github.com/kelseyhightower/confd/log"
	"io"
	"os"
	"wakever/charts/db/mysql"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (

	scatter3DColor = []string{
		"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
		"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
	}
)

func genScatter3dData() []opts.Chart3DData {
	kmean, err := mysql.GetAllKmeans()
	if err != nil {
		log.Error("%v", err)
	}
	data := make([]opts.Chart3DData, 0)
	for _, k := range kmean {
		data = append(data, opts.Chart3DData{Value: []interface{}{
			k.HorizontalCoordinates,
			k.VerticalCoordinates,
			k.ThreeDimensionalCoordinates},
			Name: "CXX C GoLang",
		})
	}
	return data
}

func scatter3DBase() *charts.Scatter3D {
	scatter3d := charts.NewScatter3D()
	scatter3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "K-Means"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        100,
			InRange:    &opts.VisualMapInRange{Color: scatter3DColor},
		}),
	)

	scatter3d.AddSeries("3d clustering model", genScatter3dData())
	return scatter3d
}

type Scatter3dEcharts struct{}

func (Scatter3dEcharts) Echarts() {
	page := components.NewPage()
	page.AddCharts(
		scatter3DBase(),
	)

	f, err := os.Create("charts/html/scatter3d.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
