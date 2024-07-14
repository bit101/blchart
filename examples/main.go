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
	surface := cairo.NewSurface(700, 800)
	context := cairo.NewContext(surface)
	context.ClearWhite()
	random.RandSeed()

	vals := random.FloatArray(6, 20, 100)

	bars := blcharts.NewBarChart(context)
	// bars.SetFgColor(blcolor.RandomRGB())
	// bars.SetBgColor(blcolor.RandomRGB())
	bars.SetSpacing(20)
	bars.Move(20, 20)
	bars.Render(vals)

	vals = random.FloatArray(20, 20, 100)
	// bars.SetFgColor(blcolor.RandomRGB())
	// bars.SetBgColor(blcolor.RandomRGB())
	bars.SetSpacing(2)
	bars.Move(360, 20)
	bars.Render(vals)

	vals = random.FloatArray(20, 20, 100)
	lines := blcharts.NewLineChart(context)
	// lines.SetFgColor(blcolor.RandomRGB())
	// lines.SetBgColor(blcolor.RandomRGB())
	lines.SetLineWidth(5)
	lines.Move(20, 280)
	lines.Render(vals)

	vals = random.FloatArray(20, 20, 100)
	// lines.SetFgColor(blcolor.RandomRGB())
	// lines.SetBgColor(blcolor.RandomRGB())
	lines.SetAutoScaleCompress(0.5)
	lines.SetFill(true)
	lines.Move(360, 280)
	lines.Render(vals)

	vals = random.FloatArray(20, 20, 100)
	// lines.SetFgColor(blcolor.RandomRGB())
	// lines.SetBgColor(blcolor.RandomRGB())
	lines.SetAutoScaleCompress(0.1)
	lines.SetFill(false)
	lines.SetPointRadius(3)
	lines.SetLineWidth(1)
	lines.Move(20, 540)
	lines.Render(vals)

	vals = random.FloatArray(20, 20, 100)
	// lines.SetFgColor(blcolor.RandomRGB())
	// lines.SetBgColor(blcolor.RandomRGB())
	lines.SetAutoScaleCompress(1)
	lines.SetPointRadius(6)
	lines.SetLineWidth(0)
	lines.Move(360, 540)
	lines.Render(vals)

	surface.WriteToPNG(fileName)
	render.ViewImage(fileName)
}
