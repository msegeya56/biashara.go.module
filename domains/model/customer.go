package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"github.com/msegeya56/biashara.go.module/domains/entity"
	"gorm.io/gorm"
	"net/http"
)



type Customer struct {
    gorm.Model  
	FirstName       string  `json:"first_name" gorm:"first_name"`
    LastName        string  `json:"last_name" gorm:"last_name"`
    MiddleName      string  `json:"middle_name" gorm:"middle_name"`
    Address         string  `json:"address" gorm:"address"`
    Latitude        float64 `json:"latitude" gorm:"latitude"`
    Longitude       float64 `json:"longitude" gorm:"longitude"`
    Alias           string  `json:"alias" gorm:"alias"`
    Email           string  `json:"email" gorm:"email"`
    Dob             string  `json:"dob" gorm:"dob"`
    Name            string  `json:"name"  gorm:"name"`
    ShippingAddress string  `json:"shipping_address" gorm:"shipping_address"`
    BillingAddress  string  `json:"billing_address" gorm:"billing_address"`
    AccountBalance  float64 `json:"account_balance" gorm:"amount_balance"`  
    Next            string  `json:"next"  gorm:"next"`
    Link            string  `json:"link"  gorm:"link"`
    Previous        string  `json:"previous" gorm:"previous"`
}

type CustomerReply struct {
	Data       *entity.Customer
	Collection []entity.Customer
	streams    <-chan entity.Customer
	Error      error
}

func (c *Customer) ToJson() string {
	jsonBytes, _ := json.Marshal(c)
	x := fmt.Sprintf("%v", string(jsonBytes))

	fmt.Println(x)
	return x

}
func (c *Customer) FromJson(data string) *Customer {
	err := json.Unmarshal([]byte(data), c)

	if err != nil {
		fmt.Println("Error converting json string to struct ERROR :", err.Error())
	}
	fmt.Println("struct Object ", data)

	return c
}

func (c *Customer) FromIOReadCloser(r io.ReadCloser) (*Customer, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r)

	defer r.Close()

	for {
		err := decoder.Decode(c)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return c, nil
}

func (c *Customer) FromRequestBody(r *http.Request) (*Customer, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(c)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return c, nil
}

func (c *Customer) FromResponseBody(r *http.Response) (*Customer, error) {

	if r == nil {

		return nil, errors.New("Http Request is nil")
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	for {
		err := decoder.Decode(c)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding gorm:", err)
			return nil, err
		}

		// Proces each person object as it is decoded

	}

	return c, nil
}
