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

func themeRiverTime() *charts.ThemeRiver {
	tr := charts.NewThemeRiver()
	tr.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "ThemeRiver Apriori",
		}),
		charts.WithSingleAxisOpts(opts.SingleAxis{
			Type:   "time",
			Bottom: "10%",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "axis",
		}),
	)

	ap,err:=mysql.GetAllApriori()
	if err!=nil {
		log.Error("%v",err)
	}

	data := make([]opts.ThemeRiverData,0)


	for i:=1;i<15;i++ {
		//items=append(items,opts.EffectScatterData{
		//	Value: u.SSE,
		//	//Name: string(u.Id),
		//})
		parseV, err := strconv.ParseFloat(ap[i].Credibility, 64)
		if err != nil {
			log.Error("%v", err)
		}

		data = append(data, opts.ThemeRiverData{
			Value: parseV,
			Date:  ap[i].ConsequentSet,
			Name:  ap[i].FrequentSet,
		})
	}

	tr.AddSeries("themeRiver", data)
	return tr
}

type ThemeriverEcharts struct{}

func (ThemeriverEcharts) Echarts() {
	page := components.NewPage()
	page.AddCharts(
		themeRiverTime(),
	)

	f, err := os.Create("charts/html/themeriver.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
