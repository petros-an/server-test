package shape

import (
	"math"

	"github.com/petros-an/server-test/common/vector"
)

type Shape2D interface {
	GetPSR2D() vector.PSR2D
	SetPSR2D(psr vector.PSR2D, setPosition bool, setScale bool, setRotation bool)
	IsPointInside(position vector.Vector2D) bool
	OverlapsWithShape(shape Shape2D) bool
	OverlapsWithRectangle(rect Rectangle) bool
	OverlapsWithEllipse(ell Ellipse) bool
	OverlapsWithPolygon(pol Polygon) bool
}

type Rectangle struct {
	vector.PSR2D
}

func (rect Rectangle) GetPSR2D() vector.PSR2D {
	return rect.PSR2D
}

func (rect *Rectangle) SetPSR2D(psr vector.PSR2D, setPosition bool, setScale bool, setRotation bool) {
	rect.PSR2D.SetPSR2D(psr, setPosition, setScale, setRotation)
}

func (rect Rectangle) IsPointInside(position vector.Vector2D) bool {
	return position.RemoveTransformationSelf(rect.PSR2D).MagnitudeManhattan() <= 1
}

func (rect Rectangle) OverlapsWithShape(shape Shape2D) bool {
	return shape.OverlapsWithRectangle(rect)
}

func (rect Rectangle) OverlapsWithRectangle(otherRect Rectangle) bool {
	return RectangleOverlapsWithRectangle(rect, otherRect)
}

func (rect Rectangle) OverlapsWithEllipse(ell Ellipse) bool {
	return EllipseOverlapsWithRectangle(ell, rect)
}

func (rect Rectangle) OverlapsWithPolygon(pol Polygon) bool {
	return PolygonOverlapsWithPolygon(pol, PolygonFromRectangle(rect))
}

type Ellipse struct {
	vector.PSR2D
}

func (ell Ellipse) GetPSR2D() vector.PSR2D {
	return ell.PSR2D
}

func (ell *Ellipse) SetPSR2D(psr vector.PSR2D, setPosition bool, setScale bool, setRotation bool) {
	ell.PSR2D.SetPSR2D(psr, setPosition, setScale, setRotation)
}

func (ell Ellipse) IsPointInside(position vector.Vector2D) bool {
	return (position.RemoveTransformationSelf(ell.PSR2D)).MagnitudeSq() <= 1
}

func (ell Ellipse) OverlapsWithShape(shape Shape2D) bool {
	return shape.OverlapsWithEllipse(ell)
}

func (ell Ellipse) OverlapsWithRectangle(rect Rectangle) bool {
	return EllipseOverlapsWithRectangle(ell, rect)
}

func (ell Ellipse) OverlapsWithEllipse(otherEll Ellipse) bool {
	return PolygonOverlapsWithEllipse(PolygonFromEllipse(otherEll), ell)
}

func (ell Ellipse) OverlapsWithPolygon(pol Polygon) bool {
	return PolygonOverlapsWithEllipse(pol, ell)
}

func unitCircleOverlapsWithFiniteLine(lineStart vector.Vector2D, lineEnd vector.Vector2D) bool {
	lineEnd.SubSelf(lineStart)

	// (lineStart->origin) dot (lineStart->lineEnd) / |lineStart->lineEnd|^2 = (SC) dot (SE) / |SE|^2 =
	// cos() * |SO|*|SE|/|SE|^2 =
	// cos() * |SO|/|SE| =
	// (ratio*|SE|/|SO|) * |SO|/|SE| = ratio
	ratio := -lineStart.Dot(lineEnd) / lineEnd.MagnitudeSq()
	if ratio < 0 || ratio > 1 {
		return false
	}

	// distSq(S + SE*ration, O) <= 1
	// |S + SE*ration|^2 <= 1
	// |SE*ration + S|^2 <= 1
	return lineEnd.MulSelf(ratio).AddSelf(lineStart).MagnitudeSq() <= 1
}

func unitCircleOverlapsWithPolygon(pol Polygon) bool {
	n := pol.NumberOfSides()
	points := pol.TransformedPoints()
	for i := 0; i < n; i++ {
		currentPoint := &points[i]
		nextPoint := &points[(i+1)%n]
		if currentPoint.MagnitudeSq() <= 1 {
			return true
		}
		if unitCircleOverlapsWithFiniteLine(*currentPoint, *nextPoint) {
			return true
		}
	}
	return pol.IsPointInside(vector.Vector2D{X: 0, Y: 0})
}

type Polygon struct {
	vector.PSR2D

	OriginalPoints    []vector.Vector2D
	transformedPoints []vector.Vector2D

	PointsAreTransformed bool
}

func (pol Polygon) GetPSR2D() vector.PSR2D {
	return pol.PSR2D
}

func (pol *Polygon) SetPSR2D(psr vector.PSR2D, setPosition bool, setScale bool, setRotation bool) {
	pol.PSR2D.SetPSR2D(psr, setPosition, setScale, setRotation)
}

func (this *Polygon) TransformedPoints() []vector.Vector2D {
	if this.PointsAreTransformed {
		return this.transformedPoints
	}
	n := len(this.OriginalPoints)
	this.transformedPoints = make([]vector.Vector2D, n)
	for i := 0; i < n; i++ {
		this.transformedPoints[i] = this.OriginalPoints[i].ApplyTransformation(this.PSR2D)
	}
	this.PointsAreTransformed = true
	return this.transformedPoints
}

func PolygonFromRectangle(rect Rectangle) Polygon {
	return Polygon{
		PSR2D: rect.PSR2D,
		OriginalPoints: []vector.Vector2D{
			{X: 0.5, Y: 0.5},
			{X: -0.5, Y: 0.5},
			{X: -0.5, Y: -0.5},
			{X: 0.5, Y: -0.5},
		},
	}
}

func PolygonFromEllipse(ell Ellipse) Polygon {
	return Polygon{
		PSR2D: ell.PSR2D,
		OriginalPoints: []vector.Vector2D{
			{X: 1, Y: 0},
			{X: 0.86602540378, Y: 0.5},
			{X: 0.5, Y: 0.86602540378},
			{X: 0, Y: 1},
			{X: -0.5, Y: 0.86602540378},
			{X: -0.86602540378, Y: 0.5},
			{X: -1, Y: 0},
			{X: -0.86602540378, Y: -0.5},
			{X: -0.5, Y: -0.86602540378},
			{X: 0, Y: -1},
			{X: 0.5, Y: -0.86602540378},
			{X: 0.86602540378, Y: -0.5},
		},
	}
}

func (pol Polygon) IsPointInside(position vector.Vector2D) bool {
	n := pol.NumberOfSides()

	points := pol.TransformedPoints()
	for i := 0; i < n; i++ {
		currentPoint := points[i]
		nextPoint := points[(i+1)%n]

		if nextPoint.SubSelf(currentPoint).Rotate90Self().Dot(position.Sub(currentPoint)) <= 0 {
			return false
		}
	}
	// log.Println("circle center inside")
	return true
}

func (pol Polygon) OverlapsWithShape(shape Shape2D) bool {
	return shape.OverlapsWithPolygon(pol)
}

func (pol Polygon) OverlapsWithRectangle(rect Rectangle) bool {
	return PolygonOverlapsWithPolygon(pol, PolygonFromRectangle(rect))
}

func (pol Polygon) OverlapsWithEllipse(ell Ellipse) bool {
	return PolygonOverlapsWithEllipse(pol, ell)
}

func (pol Polygon) OverlapsWithPolygon(otherPol Polygon) bool {
	return PolygonOverlapsWithPolygon(pol, otherPol)
}

func (pol Polygon) NumberOfSides() int {
	return len(pol.TransformedPoints())
}

// func regularPolygonEquation(numSides float64, position vector.Vector2D) float64 {
// 	r := position.Magnitude()
// 	th := position.AngleR()
// 	return r - math.Cos(math.Pi/numSides)/math.Cos(th-2*math.Pi/numSides*math.Floor((numSides*th+math.Pi)/(2*math.Pi)))

// }

func unitCircleOverlapsWithRectangle(rect Rectangle) bool {

	// put circle to origin and rotate as such that rectangle has no rotation anymore
	rect.Position.MulConjSelf(rect.Rotation)

	rect.Scale.AbsValuesSelf().DivSelf(2)
	rect.Position.AbsValuesSelf().SubSelf(rect.Scale)

	// put rectangle to 1st quarter (x,y>0) by abs the position as it is the same everywhere, then get the bottom left corner by sub size/2 from position
	// if the bottom left corner is inside the unit circle then there is overlap
	return rect.Position.MagnitudeSq() <= 1
}

func RectangleOverlapsWithRectangle(rect1 Rectangle, rect2 Rectangle) bool {
	rotationDiff := rect1.RotationDifference(rect2.PSR2D)
	rect1.Scale.AbsValuesSelf()
	rect2.Scale.AbsValuesSelf()
	// check if rotation differece is 0/90/180/270
	// if x is 0 then y is +-1 and vice versa because rotation should be normalized
	if rotationDiff.Y == 0 || rotationDiff.X == 0 {
		if rotationDiff.X == 0 {
			rect2.Scale.InvertXYSelf()
		}
		rect1.Scale.AddSelf(rect2.Scale)
		return rect1.IsPointInside(rect2.Position)
	}
	return PolygonOverlapsWithPolygon(PolygonFromRectangle(rect1), PolygonFromRectangle(rect2))
}

func EllipseOverlapsWithRectangle(ell Ellipse, rect Rectangle) bool {
	if ell.Scale.X == ell.Scale.Y {
		if ell.Scale.X != 1 {
			rect.Scale.DivSelf(ell.Scale.X)
			rect.Position.UnScaleFromCenterSelf(ell.Scale, ell.Position)
		}
	} else {
		rotationDiff := rect.RotationDifference(ell.PSR2D)
		if math.Abs(rotationDiff.X) == 1 && rotationDiff.Y == 0 {
			rect.Scale.DivVSelf(ell.Scale)
			rect.Position.UnScaleFromCenterSelf(ell.Scale, ell.Position)
		} else if rotationDiff.X == 0 && math.Abs(rotationDiff.Y) == 1 {
			rect.Scale.DivVSelf(ell.Scale.InvertXY())
			rect.Position.UnScaleFromCenterSelf(ell.Scale, ell.Position)
		}
		return PolygonOverlapsWithPolygon(PolygonFromEllipse(ell), PolygonFromRectangle(rect))
	}
	rect.Position.SubSelf(ell.Position)
	return unitCircleOverlapsWithRectangle(rect)
}

func PolygonOverlapsWithEllipse(pol Polygon, ell Ellipse) bool {
	n := pol.NumberOfSides()
	points := pol.TransformedPoints()
	for i := 0; i < n; i++ {
		points[i].RemoveTransformationSelf(ell.PSR2D)
	}
	return unitCircleOverlapsWithPolygon(pol)
}

func PolygonOverlapsWithPolygon(pol1 Polygon, pol2 Polygon) bool {
	n1 := pol1.NumberOfSides()
	n2 := pol2.NumberOfSides()
	points1 := pol1.TransformedPoints()
	points2 := pol2.TransformedPoints()
	for i := 0; i < n1; i++ {
		line1 := FinitLine{start: points1[i], end: points1[(i+1)%n1]}
		for j := 0; j < n2; j++ {
			if v := line1.finiteLinesIntersectionPoint(FinitLine{start: points2[j], end: points2[(j+1)%n2]}); v != nil {
				return true
			}
		}
	}
	return false
}

func determinant(v1 vector.Vector2D, v2 vector.Vector2D) float64 {
	return v1.X*v2.Y - v1.Y*v2.X
}

type FinitLine struct {
	start vector.Vector2D
	end   vector.Vector2D
}

func (this FinitLine) finiteLinesIntersectionPoint(other FinitLine) *vector.Vector2D {
	det := determinant(this.end.Sub(this.start), other.start.Sub(other.end))
	t := determinant(other.start.Sub(this.start), other.start.Sub(other.end)) / det
	u := determinant(this.end.Sub(this.start), other.start.Sub(this.start)) / det
	if t < 0 || u < 0 || t > 1 || u > 1 {
		return nil
	}
	v := this.start.MulSelf(1 - t).AddSelf(this.end.Mul(t))
	return v
}
