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

func generateScatterItems() []opts.ScatterData {
	SSE, err := mysql.GetAllSSEs()
	if err != nil {
		log.Error("%v", err)
	}

	items := make([]opts.ScatterData, 0)
	for _, s := range SSE {
		items = append(items, opts.ScatterData{
			Name: string(s.Id),
			Value: s.SSE,
			Symbol:       "roundRect",
			SymbolSize:   20,
			SymbolRotate: 10,
		})
	}
	return items
}
func scatterBase() *charts.Scatter {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic scatter SSE"}),
	)

	scatter.SetXAxis(SSE).
		AddSeries("Category A", generateScatterItems()).
		AddSeries("Category B", generateScatterItems())

	return scatter
}

func scatterShowLabel() *charts.Scatter {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(charts.WithTitleOpts(
		opts.Title{
			Title: "label options SSE",
		}),
	)

	scatter.SetXAxis(SSE).
		AddSeries("Category A", generateScatterItems()).
		AddSeries("Category B", generateScatterItems()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:     true,
				Position: "right",
			}),
		)
	return scatter
}

//func generateScatterItemsDbScanX(idStr string) []opts.ScatterData {
//	dbscan, err := mongodb.Find(idStr)
//	if err != nil {
//		log.Error("%v", err)
//	}
//
//	//radarData := make([][],0)
//
//	//type x []int64
//
//	items := make([]opts.ScatterData, 0)
//	for i := 0; i < len(dbscan.HorizontalCoordinates); i++  {
//		items = append(items, opts.ScatterData{
//			Name: dbscan.Point,
//			Value: dbscan.HorizontalCoordinates[i],
//			Symbol:       "roundRect",
//			SymbolSize:   20,
//			SymbolRotate: 10,
//		})
//	}
//	return items
//}
//
//
//func generateScatterItemsDbScan(idStr string) []opts.ScatterData {
//	dbscan, err := mongodb.Find(idStr)
//	if err != nil {
//		log.Error("%v", err)
//	}
//
//	//radarData := make([][],0)
//	//
//	//type coo struct{
//	//	x int64
//	//	y int64
//	//}
//
//	items := make([]opts.ScatterData, 0)
//	for i := 0; i < len(dbscan.HorizontalCoordinates); i++  {
//		items = append(items, opts.ScatterData{
//			Name: dbscan.Point,
//			Value: dbscan.HorizontalCoordinates[i],
//			Symbol:       "roundRect",
//			SymbolSize:   20,
//			SymbolRotate: 10,
//		})
//	}
//	return items
//}
//
//func scatterShowLabelDbScan() *charts.Scatter {
//	scatter := charts.NewScatter()
//	scatter.SetGlobalOptions(charts.WithTitleOpts(
//		opts.Title{
//			Title: "label options",
//		}),
//	)
//
//	scatter.SetXAxis(generateScatterItemsDbScanX(idStr1)).
//		AddSeries("Category A", generateScatterItemsDbScan(idStr1)).
//		AddSeries("Category B", generateScatterItemsDbScan(idStr2)).
//		SetSeriesOptions(charts.WithLabelOpts(
//			opts.Label{
//				Show:     true,
//				Position: "right",
//			}),
//		)
//	return scatter
//}


func scatterSplitLine() *charts.Scatter {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "splitline options SSE",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "X",
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Y",
			SplitLine: &opts.SplitLine{
				Show: true,
			}}),
	)

	scatter.SetXAxis(SSE).
		AddSeries("Player A", generateScatterItems()).
		AddSeries("Player B", generateScatterItems())
	return scatter
}

type ScatterEcharts struct{}

func (ScatterEcharts) Echarts() {
	page := components.NewPage()
	page.AddCharts(
		scatterBase(),
		scatterShowLabel(),
		scatterSplitLine(),
		//scatterShowLabelDbScan(),
	)
	f, err := os.Create("charts/html/scatter.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
