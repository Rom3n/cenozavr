package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"strconv"
	"strings"
)

func main() {
	//database.Init()
	parse()
}

func productcode(link string) string {
	code := strings.SplitAfterN(link, "/", 6)
	newcode := string(code[4])[len(code[4])-10:]
	pcode := strings.TrimSuffix(newcode, "/")
	return pcode
}

func parse() {

	//URL := "https://www.ozon.ru/category/moloko-9283/"
	//doc, err := htmlquery.LoadURL(URL)
	//if err != nil {
	//	panic(err)
	//}

	filePath := "ozon.rucategorymoloko-9283.html"
	doc, err := htmlquery.LoadDoc(filePath)
	if err != nil {
		panic(err)
	}

	pricelist := htmlquery.Find(doc, "//div[@class='ui-q2']/span[1]")

	for i, p := range pricelist {

		stprice := strings.TrimSuffix(htmlquery.InnerText(p), " ₽")
		stprice = strings.ReplaceAll(stprice, " ", "")

		price, err := strconv.Atoi(stprice)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d %d\n", i, price)
		if err != nil {
			fmt.Print("ERROR")
		}
	}

	imglist := htmlquery.Find(doc, "//div[@class='s9j']/img")
	for i, x := range imglist {
		imglink := htmlquery.SelectAttr(x, "srcset")
		imglink = strings.TrimSuffix(imglink, " 2x")
		fmt.Printf("%d %s\n", i, imglink)
		if err != nil {
			fmt.Println("ERROR")
		}
	}

	list := htmlquery.Find(doc, "//a[@class='tile-hover-target s4j']")
	for _, n := range list {
		fmt.Printf("%s %s (%s)\n", productcode(htmlquery.SelectAttr(n, "href")), htmlquery.InnerText(n), htmlquery.SelectAttr(n, "href"))
		if err != nil {
			fmt.Print("ERROR")
		}
	}
}
