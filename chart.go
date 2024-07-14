// Package blcharts defines charts.
package blcharts

import (
	"github.com/bit101/bitlib/blcolor"
	cairo "github.com/bit101/blcairo"
)

// Chart is the basic struct for any chart.
type Chart struct {
	context           *cairo.Context
	x                 float64
	y                 float64
	width             float64
	height            float64
	bgColor           blcolor.Color
	fgColor           blcolor.Color
	borderColor       blcolor.Color
	minVal            float64
	maxVal            float64
	autoScale         bool
	autoScaleCompress float64
}

// NewChart creates a new chart.
func NewChart(context *cairo.Context) *Chart {
	return &Chart{
		context:           context,
		x:                 0,
		y:                 0,
		width:             320,
		height:            240,
		bgColor:           blcolor.Grey(0.9),
		fgColor:           blcolor.Grey(0.1),
		borderColor:       blcolor.Black,
		autoScale:         true,
		autoScaleCompress: 0.1,
	}
}

// SetScale sets the min and max values for a chart. Sets autoScale to false.
func (c *Chart) SetScale(minVal, maxVal float64) {
	c.autoScale = false
	c.minVal = minVal
	c.maxVal = maxVal
}

// SetAutoScale sets autoscale to true.
// This is the default, but it can be reset if you set specific values.
func (c *Chart) SetAutoScale() {
	c.autoScale = true
}

// SetAutoScaleCompress compresses the range when autoscaling,
// so values don't crowd the edges of the chart.
// Default value is 0.1. Higher values limit the range more.
// A value of 0 will cause lines/bars to extend to the top and bottom edges of the chart.
func (c *Chart) SetAutoScaleCompress(margin float64) {
	c.autoScaleCompress = margin
}

// Move moves the chart to a new position.
func (c *Chart) Move(x, y float64) {
	c.x = x
	c.y = y
}

// Resize sets a new size for the chart.
func (c *Chart) Resize(width, height float64) {
	c.width = width
	c.height = height
}

// SetFgColor sets the foreground color.
func (c *Chart) SetFgColor(fgColor blcolor.Color) {
	c.fgColor = fgColor
}

// SetBgColor sets the background color.
func (c *Chart) SetBgColor(bgColor blcolor.Color) {
	c.bgColor = bgColor
}

// SetBorderColor sets the background color.
func (c *Chart) SetBorderColor(borderColor blcolor.Color) {
	c.borderColor = borderColor
}

// startDraw renders the background of the chart to the context and sets clipping.
// For now there will be no concept of redrawing where previous content is cleared.
// Best to just clear the entire context and redraw everything if you need to redraw a chart.
func (c *Chart) startDraw() {
	c.context.Save()
	c.context.SetAntialias(cairo.AntialiasNone)
	c.context.SetSourceColor(c.borderColor)
	c.context.FillRectangle(c.x, c.y, c.width, c.height)
	c.context.SetSourceColor(c.bgColor)
	c.context.FillRectangle(c.x+1, c.y+1, c.width-2, c.height-2)
	c.context.Restore()
	c.context.Rectangle(c.x+1, c.y+1, c.width-2, c.height-2)
	c.context.Clip()
}

// endDraw ends the drawing phase, resetting clipping.
func (c *Chart) endDraw() {
	c.context.ResetClip()
}
