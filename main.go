package main

import (
	"cenozavr/model"
	"fmt"
	"github.com/antchfx/htmlquery"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

var db *gorm.DB
var err error

func main() {

	//Parsing
	//Init source doc
	filePath := "ozon.rucategorymoloko-9283.html"
	doc, err := htmlquery.LoadDoc(filePath)
	if err != nil {
		panic(err)
	}

	//Finding price
	prlist := htmlquery.Find(doc, "//div[@class='ui-q2']/span[1]")
	var pricelist []int

	for _, p := range prlist {

		stprice := strings.TrimSuffix(htmlquery.InnerText(p), " ₽")
		stprice = strings.ReplaceAll(stprice, " ", "")

		price, err := strconv.Atoi(stprice)
		if err != nil {
			panic(err)
		}

		pricelist = append(pricelist, price)
		//fmt.Printf("%d %d\n", i, price)
		if err != nil {
			fmt.Print("ERROR")
		}
	}

	//Finding Url_img
	ilist := htmlquery.Find(doc, "//div[@class='s9j']/img")
	var imglist []string

	for _, x := range ilist {
		imglink := htmlquery.SelectAttr(x, "srcset")
		imglink = strings.TrimSuffix(imglink, " 2x")

		imglist = append(imglist, imglink)
		//fmt.Printf("%d %s\n", i, imglink)
		if err != nil {
			fmt.Println("ERROR")
		}
	}

	//Finding Id+Name+Url
	list := htmlquery.Find(doc, "//a[@class='tile-hover-target s4j']")
	var idList []int
	var nameList []string
	var urlList []string

	for _, n := range list {
		//fmt.Printf("%s %s (%s)\n", productcode(htmlquery.SelectAttr(n, "href")), htmlquery.InnerText(n), htmlquery.SelectAttr(n, "href"))
		idList = append(idList, productcode(htmlquery.SelectAttr(n, "href")))
		nameList = append(nameList, htmlquery.InnerText(n))
		urlList = append(urlList, htmlquery.SelectAttr(n, "href"))
		if err != nil {
			fmt.Print("ERROR")
		}
	}

	//Database connection
	dsn := "host=localhost user=admin password=admin dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	db.AutoMigrate(&model.Goods{})

	for i := 0; i < 3; i++ {
		var (
			moloko = &model.Goods{
				Model:   gorm.Model{},
				Id:      idList[i],
				Name:    nameList[i],
				Url:     urlList[i],
				Url_img: imglist[i],
				Price:   pricelist[i],
			}
		)

		db.Create(&moloko)
	}

}

func productcode(link string) int {

	code := strings.SplitAfterN(link, "/", 6)
	newcode := string(code[4])[len(code[4])-10:]
	pcode := strings.TrimSuffix(newcode, "/")

	icode, err := strconv.Atoi(pcode)
	if err != nil {
		panic(err)
	}
	return icode
}
