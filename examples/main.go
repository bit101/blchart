// Package main renders an image or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcharts"
)

//revive:disable:unused-parameter
const (
	tau = blmath.Tau
	pi  = math.Pi
)

func main() {
	fileName := "out.png"
	surface := cairo.NewSurface(740, 860)
	context := cairo.NewContext(surface)
	context.ClearWhite()
	random.RandSeed()

	vals := random.FloatArray(6, 50, 100)

	bars := blcharts.NewBarChart(context)
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
	lines := blcharts.NewLineChart(context)
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
	pie := blcharts.NewPieChart(context)
	pie.SetChartLabel("Needs category labels, eh?")
	// pie.SetColors(
	// 	blcolor.Aquamarine,
	// 	blcolor.Blueviolet,
	// 	blcolor.Coral,
	// 	blcolor.Darkblue,
	// 	blcolor.Gold,
	// 	blcolor.Hotpink,
	// 	blcolor.Lawngreen,
	// 	blcolor.Magenta,
	// )
	pie.Move(400, 580)
	pie.Render(vals)

	surface.WriteToPNG(fileName)
	render.ViewImage(fileName)
}
