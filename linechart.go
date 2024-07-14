// Package blcharts defines charts.
package blcharts

import (
	"math"
	"slices"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
	cairo "github.com/bit101/blcairo"
)

// LineChart is a line chart.
type LineChart struct {
	*Chart
	fill        bool
	lineWidth   float64
	pointRadius float64
}

// NewLineChart creates a new line chart.
func NewLineChart(context *cairo.Context) *LineChart {
	return &LineChart{
		Chart:       NewChart(context),
		fill:        false,
		lineWidth:   1,
		pointRadius: 0,
	}
}

// SetFill sets whether or not the area below the lines will be filled.
func (l *LineChart) SetFill(fill bool) {
	l.fill = fill
}

// SetLineWidth sets the width of the line drawn.
func (l *LineChart) SetLineWidth(lineWidth float64) {
	l.lineWidth = lineWidth
}

// SetPointRadius sets the radius of the points on the lines.
// Set to zero (default) for no points.
func (l *LineChart) SetPointRadius(pointRadius float64) {
	l.pointRadius = pointRadius
}

// Render draws the line chart.
func (l *LineChart) Render(vals []float64) {
	l.startDraw()
	top := l.maxVal
	bottom := l.minVal
	if l.autoScale {
		l.minVal = slices.Min(vals)
		l.maxVal = slices.Max(vals)
		valRange := l.maxVal - l.minVal
		top = l.maxVal + valRange*l.autoScaleCompress
		bottom = l.minVal - valRange*l.autoScaleCompress
		top = math.Round(top)
		bottom = math.Round(bottom)
	}
	l.context.SetSourceColor(l.fgColor)
	l.context.SetLineWidth(l.lineWidth)
	points := geom.NewPointList()
	for i, val := range vals {
		x := blmath.Map(float64(i), 0, float64(len(vals)-1), l.x, l.x+l.width)
		y := blmath.Map(val, bottom, top, l.y+l.height, l.y)
		points.AddXY(x, y)
	}
	if l.lineWidth > 0 || l.fill {
		for _, p := range points {
			l.context.LineTo(p.X, p.Y)
		}
	}
	if l.fill {
		l.context.LineTo(l.x+l.width, l.y+l.height)
		l.context.LineTo(l.x, l.y+l.height)
		l.context.Fill()
	} else {
		l.context.Save()
		l.context.SetLineJoin(cairo.LineJoinRound)
		l.context.Stroke()
		l.context.Restore()
	}
	if l.pointRadius > 0 {
		l.context.Points(points, l.pointRadius)
	}
	l.endDraw()
	l.drawLabels(top, bottom)
}
