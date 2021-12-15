package Echarts

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
	"os"
)

var graphNodes = []opts.GraphNode{
	{Name: "Point1"},
	{Name: "Point2"},
}


func genGraphNodes(idStr string) []opts.GraphNode {
	//ap,err:=mysql.GetAllApriori()
	//if err!=nil {
	//	log.Error("%v",err)
	//}

	graphNodes := make([]opts.GraphNode,0)

	for i := 1; i < 15; i++  {
		//println(u.Id)
		//f:=strings.TrimPrefix(u.FrequentSet,"frozenset({")
		//f=strings.TrimSuffix(f,"})")
		//c = util.GetAbsString(u.FrequentSet,"frozenset({","})")
		//f = util.GetAbsString(u.ConsequentSet,"frozenset({","})")
		//
		//parse,err:=strconv.ParseFloat(u.Credibility,32)
		//if err!=nil {
		//
		//}

		//graphNodes=append(graphNodes,opts.GraphNode{
		//	Name: ,
		//	X: float32(dbscan.HorizontalCoordinates[i]),
		//	Y: float32(dbscan.VerticalCoordinates[i]),
		//})
	}
	return graphNodes
}

func genNodes() []opts.GraphNode {
	graphNodes1:=genGraphNodes(idStr1)
	graphNodes2:=genGraphNodes(idStr2)

	for i:=0;i<len(graphNodes2);i++ {
		graphNodes1=append(graphNodes1,graphNodes2[i])
	}
	return graphNodes1
}

func genLinks() []opts.GraphLink {

	links := make([]opts.GraphLink, 0)
	for i := 0; i < len(genNodes()); i++ {
		for j := 0; j < len(genNodes()); j++ {
			links = append(links, opts.GraphLink{Source: genNodes()[i].Name, Target: genNodes()[j].Name})
		}
	}
	return links
}

func graphBase() *charts.Graph {
	//graphNodes:=genGraphNodes(idStr1)

	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic graph example DbScan"}),
	)
	graph.AddSeries("graph", genNodes(), genLinks(),
		charts.WithGraphChartOpts(
			opts.GraphChart{Force: &opts.GraphForce{Repulsion: 8000}},
		),
	)
	return graph
}

func graphCircle() *charts.Graph {
	//graphNodes:=genGraphNodes(idStr1)

	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Circular layout DbScan"}),
	)

	graph.AddSeries("graph", genNodes(), genLinks()).
		SetSeriesOptions(
			charts.WithGraphChartOpts(
				opts.GraphChart{
					Force:  &opts.GraphForce{Repulsion: 8000},
					Layout: "circular",
				}),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "right"}),
		)
	return graph
}


type GraphEcharts struct{}

func (GraphEcharts) Echarts() {
	page := components.NewPage()
	page.AddCharts(
		graphBase(),
		graphCircle(),
	)

	f, err := os.Create("charts/html/graph.html")
	if err != nil {
		panic(err)

	}
	page.Render(io.MultiWriter(f))
}
