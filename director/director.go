package director

import (
	"github.com/praiakov/builder/builder"
	"github.com/praiakov/builder/product"
)

type Director struct {
	builder builder.IBuilder
}

func NewDirector(b builder.IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b builder.IBuilder) {
	d.builder = b
}

func (d *Director) BuildBook() product.Book {
	d.builder.SetBookName()
	d.builder.SetBookAuthor()
	d.builder.SetBookType()
	d.builder.SetBookPrice()
	d.builder.SetBookLanguage()

	return d.builder.GetBook()
}
