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

//var dimensions = []string{"Visit", "Add", "Order", "Payment", "Deal"}

func genFunnelKvItems() []opts.FunnelData {
	ap,err:=mysql.GetAllApriori()
	if err!=nil {
		log.Error("%v",err)
	}

	dimensions := make([]string,0)

	for i := 1; i < len(ap); i++ {
		dimensions=append(dimensions,string(ap[i].ConsequentSet))
	}

	items := make([]opts.FunnelData, 0)
	for i := 0; i < len(dimensions); i++ {
		items = append(items, opts.FunnelData{
			Name: dimensions[i],
			Value: ap[i].Credibility,
		})
	}
	return items
}
func funnelBase() *charts.Funnel {
	funnel := charts.NewFunnel()
	funnel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic funnel Apriori"}),
	)

	funnel.AddSeries("Analytics", genFunnelKvItems())
	return funnel
}

// TODO: check the different from echarts side
func funnelShowLabel() *charts.Funnel {
	funnel := charts.NewFunnel()
	funnel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "show label Apriori"}),
	)

	funnel.AddSeries("Analytics", genFunnelKvItems()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:     true,
				Position: "left",
			},
		))
	return funnel
}

type FunnelEcharts struct{}

func (FunnelEcharts) Echarts() {
	page := components.NewPage()
	page.AddCharts(
		funnelBase(),
		funnelShowLabel(),
	)

	f, err := os.Create("charts/html/apriori.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
