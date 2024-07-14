// Package blcharts defines charts.
package blcharts

import (
	"math"
	"slices"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
)

// BarChart is a line chart.
type BarChart struct {
	*Chart
	spacing float64
}

// NewBarChart creates a new line chart.
func NewBarChart(context *cairo.Context) *BarChart {
	return &BarChart{
		Chart:   NewChart(context),
		spacing: 2.0,
	}
}

// SetSpacing sets the space between each bar.
func (b *BarChart) SetSpacing(spacing float64) {
	b.spacing = spacing
}

// Render draws the line chart.
func (b *BarChart) Render(vals []float64) {
	b.startDraw()
	b.context.Save()
	if b.spacing == 0 {
		b.context.SetAntialias(cairo.AntialiasNone)
	}
	top := b.maxVal
	bottom := b.minVal
	if b.autoScale {
		b.minVal = slices.Min(vals)
		b.maxVal = slices.Max(vals)
		valRange := b.maxVal - b.minVal
		top = b.maxVal + valRange*b.autoScaleCompress
		bottom = b.minVal
	}
	border := 1.0
	minWidth := 0.5
	graphWidth := math.Max(b.width-b.spacing, minWidth) - border*2
	fVals := float64(len(vals))
	barWidth := (graphWidth - fVals*b.spacing) / fVals
	b.context.SetSourceColor(b.fgColor)
	for i, val := range vals {
		x := blmath.Map(float64(i), 0, float64(len(vals)), b.x, b.x+graphWidth) + b.spacing + border
		h := blmath.Map(val, bottom, top, 0, b.height)
		b.context.FillRectangle(x, b.y+b.height-h, barWidth, h)
	}
	b.context.Restore()
	b.endDraw()
}
