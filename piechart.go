// Package blcharts defines charts.
package blcharts

import (
	"fmt"
	"math"

	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
)

// PieChart is a line chart.
type PieChart struct {
	*Chart
	spacing   float64
	colors    []blcolor.Color
	catLabels []string
}

// NewPieChart creates a new line chart.
func NewPieChart(context *cairo.Context) *PieChart {
	return &PieChart{
		Chart: NewChart(context),
	}
}

// SetColors sets the color to be used in the pie chart.
func (p *PieChart) SetColors(colors ...blcolor.Color) {
	p.colors = colors
}

// SetCatLabels sets a list of category labels that will be shown on each slice.
// If no labels are set, slices will be labeled with the numeric value for each slice.
// Currently can't have both for space and layout reasons.
func (p *PieChart) SetCatLabels(catLabels ...string) {
	p.catLabels = catLabels
}

// Render draws the line chart.
func (p *PieChart) Render(vals []float64) {
	total := 0.0
	for _, val := range vals {
		total += val
	}
	p.startDraw()
	angle := 0.0
	radius := math.Min(p.width, p.height) * 0.4
	p.context.Save()
	p.context.Translate(p.x+p.width/2, p.y+p.height/2)
	p.context.SetLineWidth(1)
	for i, val := range vals {
		arc := val / total * blmath.Tau
		p.setSectorColor(i, vals)
		p.context.FillCircleSector(0, 0, radius, angle, blmath.Tau, false)
		if len(p.catLabels) > 0 {
			p.renderCatLabel(angle, arc, radius, p.catLabels[i])
		} else {
			p.renderLabel(angle, arc, radius, val)
		}
		angle += arc
	}

	p.context.SetSourceColor(p.fgColor)
	p.context.StrokeCircle(0, 0, radius)
	p.context.Restore()

	p.endDraw()
}

func (p *PieChart) setSectorColor(i int, vals []float64) {
	if len(p.colors) > 0 {
		p.context.SetSourceColor(p.colors[i%len(p.colors)])
	} else {
		p.context.SetSourceGray(0.2 + float64(i)/float64(len(vals))*0.7)
	}
}

func (p *PieChart) renderLabel(angle, arc, radius, val float64) {
	if p.showLabels {
		p.context.Save()
		p.context.SetFontSize(p.labelFontSize)
		centerAngle := angle + arc/2
		p.context.SetSourceColor(p.fgColor)
		label := fmt.Sprint(blmath.RoundTo(val, p.decimals))
		x := math.Cos(centerAngle) * (radius + 10)
		y := math.Sin(centerAngle) * (radius + 10)
		extents := p.context.TextExtents(label)
		y += extents.Height / 2
		if centerAngle > math.Pi/2 && centerAngle < math.Pi*3/2 {
			x -= extents.Width
		}
		p.context.FillTextAny(label, x, y)
		p.context.Restore()
	}
}

func (p *PieChart) renderCatLabel(angle, arc, radius float64, label string) {
	if p.showLabels {
		p.context.Save()
		p.context.SetFontSize(p.labelFontSize)
		centerAngle := angle + arc/2
		p.context.SetSourceColor(p.fgColor)
		x := math.Cos(centerAngle) * (radius + 10)
		y := math.Sin(centerAngle) * (radius + 10)
		extents := p.context.TextExtents(label)
		y += extents.Height / 2
		if centerAngle > math.Pi/2 && centerAngle < math.Pi*3/2 {
			x -= extents.Width
		}
		p.context.FillTextAny(label, x, y)
		p.context.Restore()
	}
}
