// Package blchart defines charts.
package blchart

import (
	"slices"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
	cairo "github.com/bit101/blcairo"
)

// ScatterChart is a line chart.
type ScatterChart struct {
	*Chart
	pointRadius            float64
	minX, minY, maxX, maxY float64
}

// NewScatterChart creates a new line chart.
func NewScatterChart(context *cairo.Context) *ScatterChart {
	s := &ScatterChart{
		Chart:       NewChart(context),
		pointRadius: 1,
	}
	s.decimals = 1
	s.autoScaleCompress = 0
	return s
}

// SetScale sets the min and max x and y ranges.
func (s *ScatterChart) SetScale(minX, minY, maxX, maxY float64) {
	s.autoScale = false
	s.minX = minX
	s.minY = minY
	s.maxX = maxX
	s.maxY = maxY
}

// SetPointRadius sets the radius of the points on the lines.
// Set to zero (default) for no points.
func (s *ScatterChart) SetPointRadius(pointRadius float64) {
	s.pointRadius = pointRadius
}

// Render draws the line chart.
func (s *ScatterChart) Render(vals geom.PointList) {
	s.startDraw()
	if s.autoScale {
		s.calculateAutoScale(vals)
	}
	s.context.SetSourceColor(s.fgColor)
	for _, val := range vals {
		x := blmath.Map(val.X, s.minX, s.maxX, s.x, s.x+s.width)
		y := blmath.Map(val.Y, s.minY, s.maxY, s.y+s.height, s.y)
		s.context.FillCircle(x, y, s.pointRadius)
	}
	s.endDraw()
	s.drawLabels(s.maxY, s.minY)
	s.drawBottomLabels(s.maxX, s.minX)
}

func (s *ScatterChart) calculateAutoScale(vals geom.PointList) {
	xVals := []float64{}
	yVals := []float64{}
	for _, point := range vals {
		xVals = append(xVals, point.X)
		yVals = append(yVals, point.Y)
	}

	s.minX = slices.Min(xVals)
	s.minY = slices.Min(yVals)
	s.maxX = slices.Max(xVals)
	s.maxY = slices.Max(yVals)

	xRange := s.maxX - s.minX
	yRange := s.maxY - s.minY

	s.minX = blmath.RoundTo(s.minX-xRange*s.autoScaleCompress, s.decimals)
	s.minY = blmath.RoundTo(s.minY-yRange*s.autoScaleCompress, s.decimals)
	s.maxX = blmath.RoundTo(s.maxX+xRange*s.autoScaleCompress, s.decimals)
	s.maxY = blmath.RoundTo(s.maxY+yRange*s.autoScaleCompress, s.decimals)
}
