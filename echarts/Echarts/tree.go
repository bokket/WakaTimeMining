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

func genTree() []opts.TreeData{
	ID3,err:=mysql.GetAllID3()
	if err!=nil {
		log.Error("%v",err)
	}

	var TreeNodes = []*opts.TreeData{
		{
			Name: ID3[0].Rchild,
			Value: ID3[0].Rvalue,
			Children: []*opts.TreeData{
				{
					Name: ID3[2].Rchild,
					Value: ID3[2].Rvalue,
				},
				{
					Name: ID3[2].Lchild,
					Value: ID3[2].Lvalue,
					Children: []*opts.TreeData{
						{
							Name: ID3[4].Rchild,
							Value: ID3[4].Rvalue,
						},
						{
							Name: ID3[4].Lchild,
							Value: ID3[4].Lvalue,
						},
					},
				},
			},
		},
		{
			Name: ID3[0].Lchild,
			Value: ID3[0].Lvalue,
			Children: []*opts.TreeData{
				{
					Name: ID3[1].Rchild,
					Value: ID3[1].Rvalue,
					Children: []*opts.TreeData{
						{
							Name: ID3[3].Rchild,
							Value: ID3[3].Rvalue,
						},
					},
				},
				{
					Name: ID3[1].Lchild,
					Value: ID3[1].Lvalue,
				},
			},
		},
	}

	var Tree = []opts.TreeData{
		{
			Name:     ID3[0].Root,
			Children: TreeNodes,
		},
	}

	return Tree
}

func treeBase() *charts.Tree {
	graph := charts.NewTree()
	graph.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Width: "100%", Height: "95vh"}),
		charts.WithTitleOpts(opts.Title{Title: "basic tree ID3"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: false}),
	)
	graph.AddSeries("tree", genTree()).
		SetSeriesOptions(
			charts.WithTreeOpts(
				opts.TreeChart{
					Layout:           "orthogonal",
					Orient:           "LR",
					InitialTreeDepth: -1,
					Leaves: &opts.TreeLeaves{
						Label: &opts.Label{Show: true, Position: "right", Color: "Black"},
					},
				},
			),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "top", Color: "Black"}),
		)
	return graph
}

type TreeEcharts struct{}

func (TreeEcharts) Echarts() {
	page := components.NewPage()
	page.AddCharts(
		treeBase(),
	)

	f, err := os.Create("charts/html/ID3.html")
	if err != nil {
		panic(err)

	}
	page.Render(io.MultiWriter(f))
}
