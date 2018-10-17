package curves

import (
	"github.com/wieku/danser/bmath"
	"log"
)

type BSpline struct {
	points       []bmath.Vector2d
	subPoints []bmath.Vector2d
	path SliderAlgo
	ApproxLength float64
}

func NewBSpline(points []bmath.Vector2d) BSpline {
	bz := &BSpline{points: points}

	n := len(points)-2

	d := make([]bmath.Vector2d, n)
	d[0] = points[n].Sub(points[0])
	d[n-1] = points[n+1].Sub(points[n-1]).Scl(-1)

	A := make([]bmath.Vector2d, len(points))
	Bi := make([]float64, len(points))


	Bi[1] = -.25
	A[1] = points[2].Sub(points[0]).Sub(d[0]).Scl(1.0/4)//(Px[2] - Px[0] - dx[0])/4;   Ay[1] = (Py[2] - Py[0] - dy[0])/4;
	for i := 2; i < n-1; i++{
		Bi[i] = -1/(4 + Bi[i-1])
		A[i] = points[i+1].Sub(points[i-1]).Sub(A[i-1]).Scl(-1*Bi[i])//-(Px[i+1] - Px[i-1] - Ax[i-1])*Bi[i];
	}

	for i := n-2; i > 0; i--{
		d[i] = A[i].Add(d[i+1].Scl(Bi[i]))

	}
	
	bz.subPoints = append(bz.subPoints, points[0], points[1])
	for i := 2; i<len(points)-2; i++  {
		bz.subPoints = append(bz.subPoints, points[i].Sub(d[i-1]), points[i], points[i], points[i].Add(d[i-1]))
	}
	bz.subPoints = append(bz.subPoints, points[len(points)-2], points[len(points)-1])

	log.Println(bz.subPoints, "\n")

	bz.path = NewSliderAlgo("B", bz.subPoints, -1)

	return *bz
}

//It's not a neat solution, but it works
//This calculates point on bezier with constant velocity
func (bz BSpline) PointAt(t float64) bmath.Vector2d {
	return bz.path.PointAt(t)
}

func (bz BSpline) GetLength() float64 {
	return bz.ApproxLength
}

func (bz BSpline) GetStartAngle() float64 {
	return bz.points[0].AngleRV(bz.PointAt(1.0/bz.ApproxLength))
}

func (bz BSpline) GetEndAngle() float64 {
	return bz.points[len(bz.points)-1].AngleRV(bz.PointAt((bz.ApproxLength-1)/bz.ApproxLength))
}

func (ln BSpline) GetPoints(num int) []bmath.Vector2d {
	t0 := 1 / float64(num-1)

	points := make([]bmath.Vector2d, num)
	t := 0.0
	for i := 0; i < num; i += 1 {
		points[i] = ln.PointAt(t)
		t += t0
	}

	return points
}
