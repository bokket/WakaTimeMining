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
	SSE = []string{"1","2","3","4","5","6","7","8"}
)

func generateSSE() []string {
	//type sse map[int]string

	sses:=make([]string,0)

	SSE, err := mysql.GetAllSSEs()
	if err != nil {
		log.Error("%v", err)
	}

	for i,_ :=range SSE {
		//sses[i]=string(s.Id)
		sses=append(sses,string(i))
		//println(sses[i])
	}

	return sses
}

func generateLineItems() []opts.LineData {
	SSE, err := mysql.GetAllSSEs()
	if err != nil {
		log.Error("%v", err)
	}

	items := make([]opts.LineData, 0)
	for i, s := range SSE {
		items = append(items, opts.LineData{
			Name: string(i),
			Value: s.SSE,
		})
	}
	return items
}

func lineBase() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "SSE", Subtitle: "This is to find the optimal k value that fits the kmeans model."}),
	)
	line.SetXAxis(SSE).
		AddSeries("SSE", generateLineItems())
	return line
}

func lineMarkPoint() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "markpoint options SSE",
		}),
	)

	line.SetXAxis(SSE).AddSeries("SSE", generateLineItems()).
		SetSeriesOptions(
			charts.WithMarkPointNameTypeItemOpts(
				opts.MarkPointNameTypeItem{Name: "Maximum", Type: "max"},
				opts.MarkPointNameTypeItem{Name: "Average", Type: "average"},
				opts.MarkPointNameTypeItem{Name: "Minimum", Type: "min"},
			),
			charts.WithMarkPointStyleOpts(
				opts.MarkPointStyle{Label: &opts.Label{Show: true}}),
		)
	return line
}

func lineArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "area options SSE",
		}),
	)

	line.SetXAxis(SSE).AddSeries("SSE", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(
				opts.Label{
					Show: true,
				}),
			charts.WithAreaStyleOpts(
				opts.AreaStyle{
					Opacity: 0.2,
				}),
		)
	return line
}

func lineOverlap() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "overlap rect-charts SSE"}),
	)

	line.SetXAxis(SSE).
		AddSeries("SSE", generateLineItems())
	line.Overlap(esEffectStyle())
	line.Overlap(scatterBase())
	return line
}

type LineEcharts struct{}

func (LineEcharts) Echarts() {
	page := components.NewPage()
	page.AddCharts(
		lineBase(),
		lineMarkPoint(),
		lineArea(),
		lineOverlap(),
	)
	f, err := os.Create("charts/html/line.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
