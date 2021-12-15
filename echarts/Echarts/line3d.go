package Echarts

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/kelseyhightower/confd/log"
	"io"
	"os"
	"strconv"
	"wakever/charts/db/mysql"
)

var line3DColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func genLine3dData() []opts.Chart3DData {
	kmean, err := mysql.GetAllKmeans()
	if err != nil {
		log.Error("%v", err)
	}

	data := make([][3]float64, 0)
	for _, u := range kmean {

		parseX, err := strconv.ParseFloat(u.HorizontalCoordinates, 64)
		if err != nil {
			log.Error("%v", err)
		}

		parseY, err := strconv.ParseFloat(u.VerticalCoordinates, 64)
		if err != nil {
			log.Error("%v", err)
		}

		parseZ, err := strconv.ParseFloat(u.ThreeDimensionalCoordinates, 64)
		if err != nil {
			log.Error("%v", err)
		}

		data = append(data,
			[3]float64{
				parseX,
				parseY,
				parseZ,
			},
		)
	}

	ret := make([]opts.Chart3DData, 0, len(data))
	for _, d := range data {
		ret = append(ret, opts.Chart3DData{Value: []interface{}{d[0], d[1], d[2]}})
	}
	return ret
}

func line3DBase() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic line3d K-Means"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        30,
			InRange:    &opts.VisualMapInRange{Color: line3DColor},
		}),
	)

	line3d.AddSeries("line3D SSE", genLine3dData())
	return line3d
}

func line3DAutoRotate() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "auto rotating K-Means"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        30,
			InRange:    &opts.VisualMapInRange{Color: line3DColor},
		}),

		charts.WithGrid3DOpts(opts.Grid3D{
			ViewControl: &opts.ViewControl{
				AutoRotate: true,
			},
		}),
	)

	line3d.AddSeries("line3D SSE", genLine3dData())
	return line3d
}

type Line3dEcharts struct{}

func (Line3dEcharts) Echarts() {
	page := components.NewPage()
	page.AddCharts(
		line3DBase(),
		line3DAutoRotate(),
		scatter3DBase(),

		esEffectStyle(),
		lineBase(),
		lineMarkPoint(),
		lineArea(),
		lineOverlap(),
		scatterBase(),
		scatterShowLabel(),
		scatterSplitLine(),

	)

	f, err := os.Create("charts/html/kmeans.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
