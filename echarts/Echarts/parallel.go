package Echarts

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/kelseyhightower/confd/log"
	"io"
	"os"
	"wakever/charts/db/mysql"
)

func generateParallelAxistList()  []opts.ParallelAxis{
	ap,err:=mysql.GetAllApriori()
	if err!=nil {
		log.Error("%v",err)
	}

	parallelAxisList := make([]opts.ParallelAxis,0)

	for i := 1; i < 15; i++ {

		parallelAxisList =append(parallelAxisList,opts.ParallelAxis{
			Dim: i,
			Name: ap[i].FrequentSet,
		})
	}

	//{Dim: 0, Name: "Date", Inverse: true, Max: 31, NameLocation: "start"},
	//{Dim: 1, Name: "AQI"},
	//{Dim: 2, Name: "PM2.5"},
	//{Dim: 3, Name: "PM10"},
	//{Dim: 4, Name: "CO"},
	//{Dim: 5, Name: "NO2"},
	//{Dim: 6, Name: "SO2"},
	//{Dim: 7, Name: "Level", Type: "category", Data: data},

	return parallelAxisList
}

func generateParallelVaule() [][]interface{}{
	ap,err:=mysql.GetAllApriori()
	if err!=nil {
		log.Error("%v",err)
	}

	parallelData := make([][]interface{},0)
	for i := 1; i < 15; i++ {
		parallelData=append(parallelData,[]interface{}{
			ap[i].Credibility,
			ap[i].FrequentSet,
		})
	}
	return parallelData
}

func generateParallelData(data [][]interface{}) []opts.ParallelData {
	items := make([]opts.ParallelData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.ParallelData{Value: data[i]})
	}
	return items
}

func parallelBase() *charts.Parallel {
	parallel := charts.NewParallel()
	parallel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "basic Parallel example",
		}),
		charts.WithParallelAxisList(generateParallelAxistList()),
		charts.WithLegendOpts(opts.Legend{Show: true}),
	)

	parallel.AddSeries("Beijing", generateParallelData(generateParallelVaule()))
	return parallel
}

func parallelComponent() *charts.Parallel {
	parallel := charts.NewParallel()
	parallel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "with Component",
		}),
		charts.WithParallelComponentOpts(opts.ParallelComponent{
			Left:   "15%",
			Right:  "13%",
			Bottom: "10%",
			Top:    "20%",
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithParallelAxisList(generateParallelAxistList()),
	)

	parallel.AddSeries("Beijing", generateParallelData(generateParallelVaule()))
	return parallel
}

func parallelMulti() *charts.Parallel {
	parallel := charts.NewParallel()
	parallel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Multi Series",
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithParallelAxisList(generateParallelAxistList()),
	)

	parallel.AddSeries("Beijing", generateParallelData(generateParallelVaule())).
		AddSeries("Guangzhou", generateParallelData(generateParallelVaule())).
		AddSeries("Shanghai", generateParallelData(generateParallelVaule()))
	return parallel
}

type ParallelEcharts struct{}

func (ParallelEcharts) Echarts() {
	page := components.NewPage()
	page.AddCharts(
		parallelBase(),
		parallelComponent(),
		parallelMulti(),
	)
	f, err := os.Create("charts/html/parallel.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
