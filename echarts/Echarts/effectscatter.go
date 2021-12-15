package Echarts

import (
	"github.com/kelseyhightower/confd/log"
	mysql "wakever/charts/db/mysql"

	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func generateEffectScatterItems() []opts.EffectScatterData {
	gotUsers,err:=mysql.GetAllSSEs()
	if err!=nil {
		log.Error("%v",err)
	}
	items := make([]opts.EffectScatterData, 0)
	for _,u := range gotUsers {
		items=append(items,opts.EffectScatterData{
			Value: u.SSE,
			//Name: string(u.Id),
		})
	}
	return items
}

func esEffectStyle() *charts.EffectScatter {
	es := charts.NewEffectScatter()
	es.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "SSE",
		}),
	)

	es.SetXAxis(SSE).
		AddSeries("Dunk", generateEffectScatterItems(),
			charts.WithRippleEffectOpts(opts.RippleEffect{
				Period:    4,
				Scale:     10,
				BrushType: "stroke",
			})).
		AddSeries("Shoot", generateEffectScatterItems(),
			charts.WithRippleEffectOpts(opts.RippleEffect{
				Period:    3,
				Scale:     6,
				BrushType: "fill",
			}),
		)
	return es
}

type Echarter interface {
	Echarts()
}

type EffectscatterEcharts struct{}

func (EffectscatterEcharts) Echarts() {
	page := components.NewPage()
	page.AddCharts(
		esEffectStyle(),
	)

	f, err := os.Create("charts/html/effectscatter.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
