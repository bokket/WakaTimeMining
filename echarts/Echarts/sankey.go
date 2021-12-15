package Echarts

import (
	"github.com/kelseyhightower/confd/log"
	"io"
	"os"
	"strconv"
	"wakever/charts/db/mysql"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)


func sankeyBase() *charts.Sankey {
	ap,err:=mysql.GetAllApriori()
	if err!=nil {
		log.Error("%v",err)
	}

	sankeyNode := make([]opts.SankeyNode, 0)

	sankeyLink := make([]opts.SankeyLink,0)

	for i := 1; i < 15; i++ {
		sankeyNode=append(sankeyNode,opts.SankeyNode{
			Name: ap[i].FrequentSet,
		})

		parseV, err := strconv.ParseFloat(ap[i].Credibility, 32)
		if err != nil {
			log.Error("%v", err)
		}
		sankeyLink=append(sankeyLink,opts.SankeyLink{
			Source: ap[i].FrequentSet,
			Target: ap[i].ConsequentSet,
			Value: float32(parseV),
		})
	}
	//	{Name: "category1"},
	//	{Name: "category2"},
	//	{Name: "category3"},
	//	{Name: "category4"},
	//	{Name: "category5"},
	//	{Name: "category6"},
	//}
	//	{Source: "category1", Target: "category2", Value: 10},
	//	{Source: "category2", Target: "category3", Value: 15},
	//	{Source: "category3", Target: "category4", Value: 20},
	//	{Source: "category5", Target: "category6", Value: 25},
	//}


	sankey := charts.NewSankey()
	sankey.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Sankey-basic Apriori",
		}),
	)

	sankey.AddSeries("sankey apriori", sankeyNode, sankeyLink, charts.WithLabelOpts(opts.Label{Show: true}))
	return sankey
}

type SankeyEcharts struct{}

func (SankeyEcharts) Echarts() {
	page := components.NewPage()
	page.AddCharts(
		sankeyBase(),
	)

	f, err := os.Create("charts/html/sankey.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
