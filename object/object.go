package object

import (
	"opengl_engine/vao"
	"github.com/go-gl/gl/v4.1-core/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
)

type ObjectInterface interface {
	GetVAO() uint32
	GetModelMatrix() mgl.Mat4
	GetModelMatrixUniform() int32
	GetTriCount() int32
}

type Object struct {
	vao uint32
	modelMatrix mgl.Mat4
	modelMatrixUniform int32
	triCount int32
}

func (o *Object) Create(program uint32, mesh []float32) {
	// Create the Array Object based on the mesh
	// o.vao = vao.Make(mesh)
	o.vao = vao.MakeIndexed(program, mesh)
	o.triCount = int32(len(mesh) / 3)

	// Model matrix
	o.modelMatrix = mgl.Ident4()
	o.modelMatrixUniform = gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(o.modelMatrixUniform, 1, false, &o.modelMatrix[0])
}

func (o Object) GetVAO() uint32 {
	return o.vao
}

func (o Object) GetModelMatrix() mgl.Mat4 {
	return o.modelMatrix
}

func (o Object) GetModelMatrixUniform() int32 {
	return o.modelMatrixUniform
}

func (o Object) GetTriCount() int32 {
	return o.triCount
}