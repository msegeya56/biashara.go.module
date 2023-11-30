

package entity

import (
	"encoding/json"
	"io"

	"errors"
	"fmt"
	"net/http"

	// "github.com/msegeya56/ecommerce.go.module/pkg/tools/com
)

type Product struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"id"`
    Name        string   `gorm:"column:name;type:varchar(255)" json:"name,omitempty" dgraph:"product.name" xml:"name" form:"name"`
    Description string   `gorm:"column:description;type:varchar(255)" json:"description,omitempty" dgraph:"product.description" xml:"description" form:"description"`
    Price       float64  `gorm:"column:price;type:double precision" json:"price,omitempty" dgraph:"product.price" xml:"price" form:"price"`
    Stock       uint     `gorm:"column:stock;type:integer" json:"stock,omitempty" dgraph:"product.stock" xml:"stock" form:"stock"`
    CategoryID  uint     `gorm:"column:category_id;type:integer" json:"category_id,omitempty" dgraph:"product.category_id" xml:"category_id" form:"category_id"`
    // Category    Category `gorm:"foreignKey:CategoryID" json:"category,omitempty" dgraph:"product.category" xml:"category" form:"category"`
    Next        string   `json:"next,omitempty" dgraph:"product.next" xml:"next" form:"next"`
    Link        string   `json:"link,omitempty" dgraph:"product.link" xml:"link" form:"link"`
    Previous    string   `json:"previous,omitempty" dgraph:"product.previous" xml:"previous" form:"previous"`
    // Add other fields here with the same struct tags as above
}


// ProductGateway represents the gateway for the Product entity.
type ProductGateway struct{
	gorm.Model
}


// NewProductGateway creates a new instance of the ProductGateway.
func NewProductGateway() *ProductGateway {
	return &ProductGateway{}
}



type Tag struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"id"`
	Name      string `gorm:"column:name;type:varchar(255)" json:"name"`
	ProductID uint   `gorm:"column:product_id" json:"product_id"`
}

type Review struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"id"`
	Text      string `gorm:"column:text;type:text" json:"text"`
	ProductID uint   `gorm:"column:product_id" json:"product_id"`
}

type Rating struct {
	ID        uint    `gorm:"primaryKey;column:id" json:"id"`
	Score     float64 `gorm:"column:score" json:"score"`
	ProductID uint    `gorm:"column:product_id" json:"product_id"`
}

type Image struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"id"`
	URL       string `gorm:"column:url;type:varchar(255)" json:"url"`
	ProductID uint   `gorm:"column:product_id" json:"product_id"`
}

func (p *Product) ToJson() string {
	jsonBytes, _ := json.Marshal(p)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func (p *Product) FromJson(data string) *Product {
	err := json.Unmarshal([]byte(data), p)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return p
}

func (p *Product) FromIOReadCloser(r io.ReadCloser) (*Product, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(p)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return p, nil
}

func (p *Product) FromRequestBody(r *http.Request) (*Product, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	// defer r.Body.close()

	for {
		err := decoder.Decode(p)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return p, nil
}

func (p *Product) FromResponseBody(r *http.Response) (*Product, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(p)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return p, nil
}

