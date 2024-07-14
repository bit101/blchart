// Package blcharts defines charts.
package blcharts

import (
	"slices"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
)

// Histogram is a line chart.
type Histogram struct {
	*Chart
}

// NewHistogram creates a new line chart.
func NewHistogram(context *cairo.Context) *Histogram {
	h := &Histogram{
		Chart: NewChart(context),
	}
	h.autoScaleCompress = 0.05
	return h
}

// Render draws the line chart.
func (h *Histogram) Render(vals []float64) {
	h.startDraw()
	h.context.Save()
	h.context.SetAntialias(cairo.AntialiasNone)
	border := 1.0
	graphWidth := h.width - border*2
	graphHeight := h.height - border*2
	minVal := slices.Min(vals)
	maxVal := slices.Max(vals)
	hVals := make([]float64, int(graphWidth))

	for _, val := range vals {
		index := int(blmath.Map(val, minVal, maxVal, 0, graphWidth-1))
		hVals[index]++
	}
	bottom := 0.0
	top := slices.Max(hVals)
	top += (top - bottom) * h.autoScaleCompress
	h.context.SetSourceColor(h.fgColor)
	for i, hVal := range hVals {
		rectH := blmath.Map(hVal, bottom, top, 0, graphHeight)
		h.context.FillRectangle(h.x+float64(i)+border, h.y+h.height-border-rectH, 1, rectH)
	}
	h.context.Restore()
	h.endDraw()
	h.drawLabels(top, bottom)
	h.drawBottomLabels(maxVal, minVal)
}
