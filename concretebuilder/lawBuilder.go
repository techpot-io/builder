package concretebuilder

import (
	"github.com/praiakov/builder/product"
)

type LawBuilder struct {
	name     string
	author   string
	booktype string
	price    float32
	language string
}

func NewLawBuilder() *LawBuilder {
	return &LawBuilder{}
}

func (b *LawBuilder) SetBookName() {
	b.name = "Visual Law"
}

func (b *LawBuilder) SetBookAuthor() {
	b.author = "Bernardo de Azevedo e Souza"
}

func (b *LawBuilder) SetBookType() {
	b.booktype = "Law"
}

func (b *LawBuilder) SetBookPrice() {
	b.price = 15
}

func (b *LawBuilder) SetBookLanguage() {
	b.language = "Portuguse"
}

func (b *LawBuilder) GetBook() product.Book {
	return product.Book{
		Name:     b.name,
		Author:   b.author,
		Booktype: b.booktype,
		Price:    b.price,
		Language: b.language,
	}
}
