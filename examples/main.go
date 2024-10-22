// Package main renders an image or video
package main

import (
	"math"
	"slices"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blchart"
)

//revive:disable:unused-parameter
const (
	tau = blmath.Tau
	pi  = math.Pi
)

func main() {
	fileName := "out.png"
	surface := cairo.NewSurface(1100, 860)
	context := cairo.NewContext(surface)
	context.ClearWhite()
	random.RandSeed()

	vals := random.FloatArray(6, 50, 100)

	bars := blchart.NewBarChart(context)
	// bars.SetFgColor(blcolor.RandomRGB())
	// bars.SetBgColor(blcolor.RandomRGB())
	bars.SetSpacing(20)
	bars.SetChartLabel("Profit Margin Jan-Mar 2024")
	bars.Move(40, 20)
	bars.Render(vals)

	vals = random.FloatArray(20, 20, 100)
	// bars.SetFgColor(blcolor.RandomRGB())
	// bars.SetBgColor(blcolor.RandomRGB())
	bars.SetChartLabel("Rainfall in inches, 2021")
	bars.SetSpacing(2)
	bars.Move(400, 20)
	bars.Render(vals)

	vals = random.FloatArray(20, 20, 100)
	lines := blchart.NewLineChart(context)
	// lines.SetFgColor(blcolor.RandomRGB())
	// lines.SetBgColor(blcolor.RandomRGB())
	lines.SetLineWidth(5)
	lines.SetChartLabel("Approval rating")
	lines.Move(40, 300)
	lines.Render(vals)

	vals = random.FloatArray(20, 20, 100)
	// lines.SetFgColor(blcolor.RandomRGB())
	// lines.SetBgColor(blcolor.RandomRGB())
	lines.SetAutoScaleCompress(0.5)
	lines.SetFill(true)
	lines.SetChartLabel("Average temperature")
	lines.Move(400, 300)
	lines.Render(vals)

	vals = random.FloatArray(20, 20, 100)
	// lines.SetFgColor(blcolor.RandomRGB())
	// lines.SetBgColor(blcolor.RandomRGB())
	lines.SetAutoScaleCompress(0.1)
	lines.SetFill(false)
	lines.SetPointRadius(3)
	lines.SetLineWidth(1)
	lines.SetChartLabel("Readership increase")
	lines.Move(40, 580)
	lines.Render(vals)

	vals = random.FloatArray(6, 10, 100)
	for i, val := range vals {
		vals[i] = val * val
	}
	slices.Sort(vals)
	pie := blchart.NewPieChart(context)
	pie.SetChartLabel("Relative Awesomeness")
	// colors := []blcolor.Color{}
	// for i := 0.0; i < 6; i++ {
	// 	h := i / 6 * 330
	// 	colors = append(colors, blcolor.HSV(h, 1, 1))
	// }
	// slices.SortFunc(colors, func(a, b blcolor.Color) int {
	// 	return random.Element([]int{-1, 1})
	// })
	// pie.SetColors(colors...)
	pie.SetCatLabels("Stuff", "Things", "Foo", "Bar", "Cats", "Dogs")
	pie.Move(400, 580)
	pie.Render(vals)

	points := geom.NewPointList()
	for i := range 1000 {
		x := float64(i)
		// x = random.GaussRange(0, 1000)
		y := x + random.FloatRange(-1, 1)*math.Cos(x/2000*tau)*500
		y = random.GaussRange(0, 1000) + x/2
		points.AddXY(x, y)
	}
	scatter := blchart.NewScatterChart(context)
	scatter.SetChartLabel("Curious arrangement of points")
	scatter.Move(760, 20)
	scatter.Render(points)

	points = logisticGraph(2.93, 4.0, 0.002)
	scatter.SetPointRadius(0.25)
	scatter.SetDecimals(5)
	scatter.SetChartLabel("Bifurcation diagram")
	scatter.Move(760, 300)
	scatter.Render(points)

	vals = []float64{}
	for range 10000 {
		vals = append(vals, random.GaussRange(0, 600))
		vals = append(vals, random.GaussRange(300, 900))
	}

	histo := blchart.NewHistogram(context)
	histo.SetChartLabel("Histogram with twin peaks")
	histo.Move(760, 580)
	histo.Render(vals)

	surface.WriteToPNG(fileName)
	render.ViewImage(fileName)
}

func logisticGraph(minR, maxR, res float64) geom.PointList {
	points := geom.NewPointList()
	for r := minR; r < maxR; r += res {
		x := 0.3
		for range 40 {
			x = r * x * (1 - x)
		}
		for range 500 {
			x = r * x * (1 - x)
			points.AddXY(r, x)
		}
	}
	return points
}
