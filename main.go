package main

import (
	"fmt"

	"github.com/praiakov/builder/builder"
	"github.com/praiakov/builder/director"
)

func main() {
	bookFinanceBuilder := builder.GetBuilder("finance")
	bookLawBuilder := builder.GetBuilder("law")

	direct := director.NewDirector(bookFinanceBuilder)
	bookfinance := direct.BuildBook()

	fmt.Printf("Book Name: %s\n", bookfinance.Name)
	fmt.Printf("Author: %s\n", bookfinance.Author)
	fmt.Printf("Price $: %.2f\n", bookfinance.Price)
	fmt.Printf("Type: %s\n", bookfinance.Booktype)
	fmt.Printf("Language: %s\n", bookfinance.Language)

	direct.SetBuilder(bookLawBuilder)
	bookLaw := direct.BuildBook()
	fmt.Printf("\n")
	fmt.Printf("Book Name: %s\n", bookLaw.Name)
	fmt.Printf("Author: %s\n", bookLaw.Author)
	fmt.Printf("Price: $ %.2f\n", bookLaw.Price)
	fmt.Printf("Type: %s\n", bookLaw.Booktype)
	fmt.Printf("Language: %s\n", bookLaw.Language)

}
