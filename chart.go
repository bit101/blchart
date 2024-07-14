// Package blcharts defines charts.
package blcharts

import (
	"fmt"
	"math"

	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/bitlib/blmath"
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
	textColor         blcolor.Color
	borderColor       blcolor.Color
	minVal            float64
	maxVal            float64
	autoScale         bool
	autoScaleCompress float64
	decimals          int
	showLabels        bool
	labelFontSize     float64
	rotateLabels      bool
	chartLabel        string
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
		textColor:         blcolor.Black,
		borderColor:       blcolor.Black,
		autoScale:         true,
		autoScaleCompress: 0.1,
		decimals:          0,
		showLabels:        true,
		labelFontSize:     12,
		rotateLabels:      true,
		chartLabel:        "",
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

// SetDecimals sets how much precision data labels will have.
func (c *Chart) SetDecimals(decimals int) {
	c.decimals = decimals
}

// ShowLabels sets whether or not value labels will be shown.
func (c *Chart) ShowLabels(show bool) {
	c.showLabels = show
}

// RotateLabels sets whether or not value labels will be rotated.
func (c *Chart) RotateLabels(rotate bool) {
	c.rotateLabels = rotate
}

// SetLabelFontSize sets the size of the font used for labels.
func (c *Chart) SetLabelFontSize(size float64) {
	c.labelFontSize = size
}

// SetChartLabel sets the label for the chart.
func (c *Chart) SetChartLabel(label string) {
	c.chartLabel = label
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

// SetBorderColor sets the border color.
func (c *Chart) SetBorderColor(borderColor blcolor.Color) {
	c.borderColor = borderColor
}

// SetTextColor sets the text color.
func (c *Chart) SetTextColor(textColor blcolor.Color) {
	c.textColor = textColor
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
	if c.chartLabel != "" {
		c.context.Save()
		c.context.SetFontSize(c.labelFontSize)
		extents := c.context.TextExtents(c.chartLabel)
		c.context.FillText(
			c.chartLabel,
			c.x+c.width/2-extents.Width/2,
			c.y+c.height+c.labelFontSize+2,
		)
		c.context.Restore()
	}
}

func (c *Chart) drawLabels(top, bottom float64) {
	if c.showLabels {
		c.context.Save()
		c.context.SetFontSize(c.labelFontSize)
		c.context.SetSourceColor(c.textColor)

		// top
		label := fmt.Sprint(blmath.RoundTo(top, c.decimals))
		extents := c.context.TextExtents(label)
		if c.rotateLabels {
			c.context.Save()
			c.context.Translate(c.x, c.y)
			c.context.Rotate(-math.Pi / 2)
			c.context.FillText(label, -extents.Width, -5)
			c.context.Restore()
		} else {
			c.context.FillText(label, c.x-extents.Width-5, c.y+extents.Height/2)
		}

		// bottom
		label = fmt.Sprint(blmath.RoundTo(bottom, 1))
		extents = c.context.TextExtents(label)
		if c.rotateLabels {
			c.context.Save()
			c.context.Translate(c.x, c.y+c.height)
			c.context.Rotate(-math.Pi / 2)
			c.context.FillText(label, 0, -5)
			c.context.Restore()
		} else {
			c.context.FillText(label, c.x-extents.Width-5, c.y+c.height+extents.Height/2)
		}
		c.context.Restore()
	}
}

func (c *Chart) drawBottomLabels(maxX, minX float64) {
	if c.showLabels {
		c.context.Save()
		c.context.SetFontSize(c.labelFontSize)
		c.context.SetSourceColor(c.textColor)

		// left
		label := fmt.Sprint(blmath.RoundTo(minX, c.decimals))
		extents := c.context.TextExtents(label)
		c.context.FillText(label, c.x, c.y+c.height+c.labelFontSize+2)

		// right
		label = fmt.Sprint(blmath.RoundTo(maxX, c.decimals))
		extents = c.context.TextExtents(label)
		c.context.FillText(label, c.x+c.width-extents.Width, c.y+c.height+c.labelFontSize+2)

		c.context.Restore()
	}
}
