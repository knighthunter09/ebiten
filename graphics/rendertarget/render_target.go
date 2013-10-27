package rendertarget

import (
	"github.com/hajimehoshi/go-ebiten/graphics/texture"
)

type RenderTarget struct {
	texture     *texture.Texture
	framebuffer interface{}
}

func NewWithFramebuffer(texture *texture.Texture, framebuffer interface{}) *RenderTarget {
	return &RenderTarget{
		texture:     texture,
		framebuffer: framebuffer,
	}
}

// TODO: Remove this
func (renderTarget *RenderTarget) Texture() *texture.Texture {
	return renderTarget.texture
}

// TODO: Remove this
func (renderTarget *RenderTarget) Framebuffer() interface{} {
	return renderTarget.framebuffer
}

func (renderTarget *RenderTarget) SetAsViewport(setter func(x, y, width, height int)) {
	renderTarget.texture.SetAsViewport(setter)
}