package main

import (
	"log"
	"net/http"
	"wakever/charts/Echarts"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	charts := []Echarts.Echarter{
		Echarts.EffectscatterEcharts{},
		Echarts.Scatter3dEcharts{},
		Echarts.LineEcharts{},
		Echarts.Line3dEcharts{},
		Echarts.RadarEcharts{},
		Echarts.ScatterEcharts{},
		//Echarts.GraphEcharts{},
		Echarts.ThemeriverEcharts{},
		Echarts.SankeyEcharts{},
		Echarts.TreeEcharts{},
		Echarts.PageSSELayoutEcharts{},
		Echarts.ParallelEcharts{},
		Echarts.FunnelEcharts{},
	}

	for _, c := range charts {
		c.Echarts()
	}

	fs := http.FileServer(http.Dir("charts/html"))
	log.Println("running server at http://localhost:8089")
	log.Fatal(http.ListenAndServe("localhost:8089", logRequest(fs)))
}
