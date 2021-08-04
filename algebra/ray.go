package algebra

import "fmt"

type Ray struct {
	Origin Vector3
	Direction Vector3
}

func (r Ray) GetPosition(t MnFloat) Vector3 {
	return r.Origin.Add(r.Direction.MultiplyScalar(t))
}

func (r Ray) Dump() Ray {
	fmt.Println("Ray:")
	fmt.Println(fmt.Sprintf("  Origin: %f, %f, %f", r.Origin.X, r.Origin.Y, r.Origin.Z))
	fmt.Println(fmt.Sprintf("  Direction: %f, %f, %f", r.Direction.X, r.Direction.Y, r.Direction.Z))

	return r
}
