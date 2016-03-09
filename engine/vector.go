package engine

type Vector struct{
	X, Y float32
}

func Vec(x, y float32) Vector {
	return Vector{x, y}
}

func (v Vector) Add(n Vector) Vector {
	return Vec(v.X+n.X,v.Y+n.Y)
}

func (v Vector) Sub(n Vector) Vector {
	return Vec(v.X-n.X,v.Y-n.Y)
}

func (v Vector) Mult(n float32) Vector {
	return Vec(v.X*n,v.Y*n)
}

func (v Vector) Div(n float32) Vector {
	return Vec(v.X/n,v.Y/n)
}

func (v *Vector) Set(n Vector) {
	v = &n
}
