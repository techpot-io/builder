package builder

import (
	"github.com/praiakov/builder/concretebuilder"
	"github.com/praiakov/builder/product"
)

type IBuilder interface {
	SetBookName()
	SetBookAuthor()
	SetBookType()
	SetBookPrice()
	SetBookLanguage()
	GetBook() product.Book
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "finance" {
		return concretebuilder.NewFinanceBuilder()
	}

	if builderType == "law" {
		return concretebuilder.NewLawBuilder()
	}

	return nil
}
