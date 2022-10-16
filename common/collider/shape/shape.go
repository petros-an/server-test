package shape

import (
	"math"

	"github.com/petros-an/server-test/common/vector"
)

type Shape2D interface {
	GetPosition() vector.Vector2D
	SetPSR2D(psr vector.PSR2D, setPosition bool, setScale bool, setRotation bool)
	IsPointInside(position vector.Vector2D) bool
	OverlapsShape(shape Shape2D) bool
	overlapsRectangle(rect Rectangle) bool
	overlapsEllipse(ell Ellipse) bool
	overlapsPolygon(pol Polygon) bool
	overlapsInfiniteLine(line InfiniteLine) bool
	overlapsHalfPlane(plane HalfPlane) bool
}

type Rectangle struct {
	vector.PSR2D
}

func (this Rectangle) GetPosition() vector.Vector2D {
	return this.Position
}

func (this *Rectangle) SetPSR2D(psr vector.PSR2D, setPosition bool, setScale bool, setRotation bool) {
	this.PSR2D.SetPSR2D(psr, setPosition, setScale, setRotation)
}

func (this Rectangle) IsPointInside(position vector.Vector2D) bool {
	return position.RemoveTransformationSelf(this.PSR2D).MagnitudeManhattan() <= (0.5 * 0.5)
}

func (this Rectangle) OverlapsShape(shape Shape2D) bool {
	return shape.overlapsRectangle(this)
}

func (this Rectangle) overlapsRectangle(other Rectangle) bool {
	rotationDiff := this.RotationDifference(other.PSR2D)
	this.Scale.AbsValuesSelf()
	other.Scale.AbsValuesSelf()
	// check if rotation differece is 0/90/180/270
	// if x is 0 then y is +-1 and vice versa because rotation should be normalized
	if rotationDiff.Y == 0 || rotationDiff.X == 0 {
		if rotationDiff.X == 0 {
			other.Scale.InvertXYSelf()
		}
		this.Scale.AddSelf(other.Scale)
		return this.IsPointInside(other.Position)
	}
	return PolygonFromRectangle(this).overlapsPolygon(PolygonFromRectangle(other))
}

func (this Rectangle) overlapsEllipse(ell Ellipse) bool {
	return EllipseOverlapsRectangle(ell, this)
}

func (this Rectangle) overlapsPolygon(pol Polygon) bool {
	return pol.overlapsPolygon(PolygonFromRectangle(this))
}

func (this Rectangle) overlapsInfiniteLine(line InfiniteLine) bool {
	return InfiniteLineOverlapsRectangle(line, this)
}

func (this Rectangle) overlapsHalfPlane(plane HalfPlane) bool {
	return plane.OverlapsShape(&this)
}

type Ellipse struct {
	vector.PSR2D
}

func (this Ellipse) GetPosition() vector.Vector2D {
	return this.Position
}

func (this *Ellipse) SetPSR2D(psr vector.PSR2D, setPosition bool, setScale bool, setRotation bool) {
	this.PSR2D.SetPSR2D(psr, setPosition, setScale, setRotation)
}

func (this Ellipse) IsPointInside(position vector.Vector2D) bool {
	return (position.RemoveTransformationSelf(this.PSR2D)).MagnitudeSq() <= (0.5 * 0.5)
}

func (this Ellipse) OverlapsShape(shape Shape2D) bool {
	return shape.overlapsEllipse(this)
}

func (this Ellipse) overlapsRectangle(rect Rectangle) bool {
	return EllipseOverlapsRectangle(this, rect)
}

func (this Ellipse) overlapsEllipse(other Ellipse) bool {
	if this.HasUniformScale() && other.HasUniformScale() {
		return this.Position.Sub(other.Position).MagnitudeSq()*4 <= (this.Scale.X+other.Scale.X)*(this.Scale.X+other.Scale.X)
	}
	return PolygonOverlapsEllipse(PolygonFromEllipse(this), other)
}

func (this Ellipse) overlapsPolygon(pol Polygon) bool {
	return PolygonOverlapsEllipse(pol, this)
}

func (this Ellipse) overlapsInfiniteLine(line InfiniteLine) bool {
	return InfiniteLineOverlapsEllipse(line, this)
}

func (this Ellipse) overlapsHalfPlane(plane HalfPlane) bool {
	return plane.OverlapsShape(&this)
}

func HalfUnitCircleOverlapsPolygon(pol Polygon) bool {
	n := pol.NumberOfSides()
	points := pol.TransformedPoints()
	for i := 0; i < n; i++ {
		currentPoint := points[i]
		nextPoint := points[(i+1)%n]
		if currentPoint.MagnitudeSq() <= (0.5 * 0.5) {
			return true
		}
		if HalfUnitCircleOverlapsFiniteLine(FiniteLine{currentPoint, nextPoint}) {
			return true
		}
	}
	return pol.IsPointInside(vector.Vector2D{X: 0, Y: 0})
}

type Polygon struct {
	vector.PSR2D

	OriginalPoints []vector.Vector2D

	transformedPoints    []vector.Vector2D
	PointsAreTransformed bool
}

func (this Polygon) GetPosition() vector.Vector2D {
	return this.Position
}

func (this *Polygon) SetPSR2D(psr vector.PSR2D, setPosition bool, setScale bool, setRotation bool) {
	this.PSR2D.SetPSR2D(psr, setPosition, setScale, setRotation)
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

func (this Polygon) IsPointInside(position vector.Vector2D) bool {
	n := this.NumberOfSides()

	points := this.TransformedPoints()
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

func (this Polygon) OverlapsShape(shape Shape2D) bool {
	return shape.overlapsPolygon(this)
}

func (this Polygon) overlapsRectangle(rect Rectangle) bool {
	return this.overlapsPolygon(PolygonFromRectangle(rect))
}

func (this Polygon) overlapsEllipse(ell Ellipse) bool {
	return PolygonOverlapsEllipse(this, ell)
}

func (this Polygon) overlapsPolygon(other Polygon) bool {
	n1 := this.NumberOfSides()
	n2 := other.NumberOfSides()
	points1 := this.TransformedPoints()
	points2 := other.TransformedPoints()
	for i := 0; i < n1; i++ {
		line1 := FiniteLine{Start: points1[i], End: points1[(i+1)%n1]}
		for j := 0; j < n2; j++ {
			if v := line1.finiteLinesIntersectionPoint(FiniteLine{Start: points2[j], End: points2[(j+1)%n2]}); v != nil {
				return true
			}
		}
	}
	return false
}

func (this Polygon) overlapsInfiniteLine(line InfiniteLine) bool {
	return InfiniteLineOverlapsPolygon(line, this)
}

func (this Polygon) overlapsHalfPlane(plane HalfPlane) bool {
	return plane.OverlapsShape(&this)
}

func (this Polygon) NumberOfSides() int {
	return len(this.TransformedPoints())
}

// func regularPolygonEquation(numSides float64, position vector.Vector2D) float64 {
// 	r := position.Magnitude()
// 	th := position.AngleR()
// 	return r - math.Cos(math.Pi/numSides)/math.Cos(th-2*math.Pi/numSides*math.Floor((numSides*th+math.Pi)/(2*math.Pi)))

// }

func HalfUnitCircleOverlapsRectangle(rect Rectangle) bool {

	// put circle to origin and rotate as such that rectangle has no rotation anymore
	if rect.Rotation.Y != 0 {
		rect.Position.MulConjSelf(rect.Rotation)
	}

	rect.Scale.AbsValuesSelf().DivSelf(2)
	rect.Position.AbsValuesSelf().SubSelf(rect.Scale)

	// put rectangle to 1st quarter (x,y>0) by abs the position as it is the same everywhere, then get the bottom left corner by sub size/2 from position
	// if the bottom left corner is inside the unit circle then there is overlap
	return rect.Position.MagnitudeSq() <= (0.5 * 0.5)
}

func EllipseOverlapsRectangle(this Ellipse, rect Rectangle) bool {
	if this.HasUniformScale() {
		if this.Scale.X == 1 {
			rect.Position.SubSelf(this.Position)
		} else {
			rect.Scale.DivSelf(this.Scale.X)
			rect.Position.SubSelf(this.Position).DivSelf(this.Scale.X)
		}
	} else {
		rotationDiff := rect.RotationDifference(this.PSR2D)
		if rotationDiff.Y == 0 {
			rect.Scale.DivVSelf(this.Scale)
			rect.Position.SubSelf(this.Position).DivVSelf(this.Scale)
		} else if rotationDiff.X == 0 {
			rect.Scale.DivVSelf(*this.Scale.InvertXYSelf())
			rect.Position.SubSelf(this.Position).DivVSelf(this.Scale)
		} else {
			return PolygonFromEllipse(this).overlapsPolygon(PolygonFromRectangle(rect))
		}
	}
	return HalfUnitCircleOverlapsRectangle(rect)
}

func PolygonOverlapsEllipse(pol Polygon, ell Ellipse) bool {
	n := pol.NumberOfSides()
	points := pol.TransformedPoints()
	for i := 0; i < n; i++ {
		points[i].RemoveTransformationSelf(ell.PSR2D)
	}
	return HalfUnitCircleOverlapsPolygon(pol)
}

func determinantV(v1 vector.Vector2D, v2 vector.Vector2D) float64 {
	return v1.X*v2.Y - v1.Y*v2.X
}
func determinant(a float64, b float64, c float64, d float64) float64 {
	return a*c - b*d
}

type FiniteLine struct {
	Start vector.Vector2D
	End   vector.Vector2D
}

func (this FiniteLine) Length() float64 {
	return this.Start.Sub(this.End).Magnitude()
}

func (this FiniteLine) LengthSq() float64 {
	return this.Start.Sub(this.End).MagnitudeSq()
}

func (this FiniteLine) finiteLinesIntersectionPoint(other FiniteLine) *vector.Vector2D {
	det := determinantV(this.End.Sub(this.Start), other.Start.Sub(other.End))
	if det == 0 {
		return nil
	}
	t := determinantV(other.Start.Sub(this.Start), other.Start.Sub(other.End)) / det
	u := determinantV(this.End.Sub(this.Start), other.Start.Sub(this.Start)) / det
	if t < 0 || u < 0 || t > 1 || u > 1 {
		return nil
	}
	return this.Start.MulSelf(1 - t).AddSelf(this.End.Mul(t))
}

func HalfUnitCircleOverlapsFiniteLine(line FiniteLine) bool {
	line.End.SubSelf(line.Start)

	// (lineStart->origin) dot (lineStart->lineEnd) / |lineStart->lineEnd|^2 = (SC) dot (SE) / |SE|^2 =
	// cos() * |SO|*|SE|/|SE|^2 =
	// cos() * |SO|/|SE| =
	// (dotNormal*|SE|/|SO|) * |SO|/|SE| =
	// dotNormal                                 (lineStart->origin) dot (lineStart->lineEnd) -> -lineStart dot (lineEnd-lineStart)
	dotNormal := -line.Start.Dot(line.End) / line.End.MagnitudeSq()
	if dotNormal < 0 || dotNormal > 1 {
		return false
	}

	// distSq(S + SE*ration, O) <= 0.5^2
	// |SO + SE*ration|^2 <= 0.5^2,  SO -> O - S
	// |SE*ration - S|^2 <= 0.5^2
	return line.End.MulSelf(dotNormal).SubSelf(line.Start).MagnitudeSq() <= (0.5 * 0.5)
}

type InfiniteLine struct {
	Direction vector.Vector2D
	Offset    float64

	// y = (Direction.Y/Direction.X) * x + Offset/Direction.X  ->   Direction.X * y - Direction.Y * x = Offset
}

func (this InfiniteLine) RegularPoint() vector.Vector2D {
	if this.Direction.X != 0 {
		return vector.Vector2D{X: 0, Y: this.Offset / this.Direction.X}
	}
	return vector.Vector2D{X: -this.Offset / this.Direction.Y, Y: 0}
}

func (this InfiniteLine) Regular2Points() (vector.Vector2D, vector.Vector2D) {
	p := this.RegularPoint()
	return p, p.Add(this.Direction)
}

func (this InfiniteLine) GetPosition() vector.Vector2D {
	return this.RegularPoint()
}

func (this InfiniteLine) Equation(point vector.Vector2D) float64 {
	return this.Direction.X*point.Y - this.Direction.Y*point.X - this.Offset
}

func NewInfiniteLineFrom2Points(point1 vector.Vector2D, point2 vector.Vector2D) InfiniteLine {
	new := InfiniteLine{
		Direction: point2.Sub(point1).Normalized(),
	}
	new.SetOffsetFromPosition(point1)
	return new
}

func (this *InfiniteLine) SetOffsetFromPosition(position vector.Vector2D) *InfiniteLine {
	this.Offset = this.Direction.X*position.Y - this.Direction.Y*position.X
	return this
}

func (this *InfiniteLine) SetPSR2D(psr vector.PSR2D, setPosition bool, setScale bool, setRotation bool) {
	if setRotation {
		this.Direction = psr.Rotation
	}
	if setPosition {
		this.SetOffsetFromPosition(psr.Position)
	}
}

func (this InfiniteLine) InfiniteLineOverlapsInfiniteLine(other InfiniteLine) bool {
	return !this.Direction.Equals(other.Direction) && !this.Direction.Equals(other.Direction.Opp()) && math.Abs(this.Offset) == math.Abs(other.Offset)
}

func (this InfiniteLine) InfiniteLinesIntersectionPoint(other InfiniteLine) *vector.Vector2D {
	det := -determinantV(this.Direction, other.Direction)
	detY := -determinant(this.Offset, this.Direction.Y, other.Offset, other.Direction.Y)
	detX := determinant(this.Direction.X, this.Offset, other.Direction.X, other.Offset)
	if det == 0 {
		if detX == 0 && detY == 0 {
			v := this.RegularPoint()
			return &v
		}
		return nil
	}
	return &vector.Vector2D{X: detX / det, Y: detY / det}
}

func (this InfiniteLine) ClosestPointToOrigin() vector.Vector2D {
	if this.Offset == 0 {
		return vector.Vector2D{X: 0, Y: 0}
	}
	pointOnLine := this.RegularPoint()
	return this.Direction.Mul(-pointOnLine.Dot(this.Direction)).Add(pointOnLine)
}

func (this InfiniteLine) ClosestPointTo(position vector.Vector2D) vector.Vector2D {
	pointOnLine := this.RegularPoint()
	return this.Direction.Mul(position.Sub(pointOnLine).Dot(this.Direction)).Add(pointOnLine)
}

func (this InfiniteLine) IsPointInside(point vector.Vector2D) bool {
	return this.Equation(point) == 0
}

func (this InfiniteLine) OverlapsHalfUnitCircle() bool {
	return this.ClosestPointToOrigin().MagnitudeSq() <= (0.5 * 0.5)
}

func (this InfiniteLine) OverlapsShape(shape Shape2D) bool {
	return shape.overlapsInfiniteLine(this)
}

func (this InfiniteLine) overlapsRectangle(rect Rectangle) bool {
	return InfiniteLineOverlapsRectangle(this, rect)
}

func (this InfiniteLine) overlapsEllipse(ell Ellipse) bool {
	return InfiniteLineOverlapsEllipse(this, ell)
}

func (this InfiniteLine) overlapsPolygon(pol Polygon) bool {
	return InfiniteLineOverlapsPolygon(this, pol)
}

func (this InfiniteLine) overlapsInfiniteLine(other InfiniteLine) bool {
	return this.InfiniteLinesIntersectionPoint(other) != nil
}

func (this InfiniteLine) OverlapsCenteredCircle(radius float64) bool {
	return this.ClosestPointToOrigin().MagnitudeSq() <= (radius * radius)
}

func InfiniteLineOverlapsEllipse(line InfiniteLine, ell Ellipse) bool {
	if !ell.HasUniformScale() {
		point1, point2 := line.Regular2Points()
		point1.RemoveTransformationSelf(ell.PSR2D)
		point2.RemoveTransformationSelf(ell.PSR2D)
		return NewInfiniteLineFrom2Points(point1, point2).OverlapsHalfUnitCircle()
	}
	line.SetOffsetFromPosition(line.RegularPoint().Sub(ell.Position))
	return line.OverlapsCenteredCircle(ell.Scale.X / 2)
}

func InfiniteLineOverlapsFiniteLine(infIine InfiniteLine, finLine FiniteLine) bool {
	if infIine.IsPointInside(finLine.Start) {
		return true
	}
	point := infIine.InfiniteLinesIntersectionPoint(NewInfiniteLineFrom2Points(finLine.Start, finLine.End))
	if point == nil {
		return false
	}
	lineLenSq := finLine.LengthSq()
	return point.Sub(finLine.Start).MagnitudeSq() <= lineLenSq && point.Sub(finLine.End).MagnitudeSq() <= lineLenSq
}

func InfiniteLineOverlapsRectangle(line InfiniteLine, rect Rectangle) bool {
	return InfiniteLineOverlapsPolygon(line, PolygonFromRectangle(rect))
}

func InfiniteLineOverlapsPolygon(line InfiniteLine, pol Polygon) bool {
	points := pol.TransformedPoints()
	n := len(points)
	for i := 0; i < n; i++ {
		currentPoint := points[i]
		nextPoint := points[(i+1)%n]
		if InfiniteLineOverlapsFiniteLine(line, FiniteLine{currentPoint, nextPoint}) {
			return true
		}
	}
	return false
}

type HalfPlane struct {
	InfiniteLine
}

func (this HalfPlane) IsPointInside(point vector.Vector2D) bool {
	return this.Equation(point) <= 0
}

func (this HalfPlane) OverlapsShape(shape Shape2D) bool {
	return this.InfiniteLine.OverlapsShape(shape) || this.IsPointInside(shape.GetPosition())
}

func (this HalfPlane) overlapsHalfPlane(other HalfPlane) bool {
	return this.InfiniteLine.OverlapsShape(&other) || this.IsPointInside(other.RegularPoint()) || other.IsPointInside(this.RegularPoint())
}
