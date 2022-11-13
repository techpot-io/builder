package concretebuilder

import (
	"github.com/praiakov/builder/product"
)

type FinanceBuilder struct {
	name     string
	author   string
	booktype string
	price    float32
	language string
}

func NewFinanceBuilder() *FinanceBuilder {
	return &FinanceBuilder{}
}

func (b *FinanceBuilder) SetBookName() {
	b.name = "The Psychology of Money"
}

func (b *FinanceBuilder) SetBookAuthor() {
	b.author = "Morgan Housel"
}

func (b *FinanceBuilder) SetBookType() {
	b.booktype = "Finance"
}

func (b *FinanceBuilder) SetBookPrice() {
	b.price = 18.9
}

func (b *FinanceBuilder) SetBookLanguage() {
	b.language = "English"
}

func (b *FinanceBuilder) GetBook() product.Book {
	return product.Book{
		Name:     b.name,
		Author:   b.author,
		Booktype: b.booktype,
		Price:    b.price,
		Language: b.language,
	}
}
