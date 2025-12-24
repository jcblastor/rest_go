package main

import (
	"fmt"
	"log"

	"github.com/jcblastor/rest_go/pkg/product"
	"github.com/jcblastor/rest_go/pkg/storage"
)

func main() {
	/*
		if err := serviceProduct.Migrate(); err != nil {
			log.Fatalf("product.Migrate: %v", err)
		}

		storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
		serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

		if err := serviceInvoiceHeader.Migrate(); err != nil {
			log.Fatalf("invoiceHeader.Migrate: %v", err)
		}

		storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
		serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

		if err := serviceInvoiceItem.Migrate(); err != nil {
			log.Fatalf("invoiceItem.Migrate: %v", err)
		}
	*/
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	products, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(products)
}
