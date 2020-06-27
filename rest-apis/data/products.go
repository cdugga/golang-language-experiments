package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"time"
)

type Product struct {
	ID 		int `json:"id"`
	Name 	string `json:"name" validate:"required"`
	Description string `json:"description"`
	Price 		float32 `json:"price validate:"lt=0"`
	SKU 		string `json:"sku"`
	CreatedOn 	string `json:"-"`
	UpdatedOn 	string `json:"-"`
	DeletedOn 	string `json:"-"`
}

type Products []*Product

func (p *Product) Validate() error{
	validate := validator.New()
	return validate.Struct(p)
}


func (p*Products) ToJSON(w io.Writer) error{
	e := json.NewEncoder(w)
	return e.Encode(p)
}


func (p*Product) FromJSON(R io.Reader) error{
	e := json.NewDecoder(R)
	return e.Decode(p)
}


func GetProducts() Products {
	return productList
}

func AddProduct(p *Product){
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error{
	_, pos, err := FindProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p

	return nil

}

var ErrProductNotFound = fmt.Errorf("Product not found")

func FindProduct(id int) (*Product, int, error){
	for i,p := range productList {
		if p.ID == id{
			return p,i, nil
		}
	}
	return nil,0,ErrProductNotFound
}


func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
