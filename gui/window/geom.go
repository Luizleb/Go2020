package window

import "github.com/go-gl/gl/v4.1-core/gl"

type points []float32

func DrawRect(x, y, w, h float32) uint32 {
	pts := rectPoints(x, y, w, h)
	return vaoRect(pts)
}

func rectPoints(x, y, w, h float32) points {
	ww := float32(Width)
	wh := float32(Height)
	return []float32{
		// A
		(-1 + 2*x/ww),
		(1 - 2*y/wh),
		0,
		//B
		(-1 + 2*(x+w)/ww),
		(1 - 2*y/wh),
		0,
		//C
		(-1 + 2*(x+w)/ww),
		(1 - 2*(y+h)/wh),
		0,
		// A
		(-1 + 2*x/ww),
		(1 - 2*y/wh),
		0,
		//C
		(-1 + 2*(x+w)/ww),
		(1 - 2*(y+h)/wh),
		0,
		//D
		(-1 + 2*x/ww),
		(1 - 2*(y+h)/wh),
		0,
	}
}

func vaoRect(p points) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(p), gl.Ptr(p), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexArrayAttrib(vao, 0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao

}
