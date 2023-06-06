package customers

import (
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) CreateCustomer(c *Customer) error {
	return r.db.Create(c).Error
}

func (r Repository) DeleteCustomer(c *Customer, d *DeleteCustomerRequest) error {
	if err := r.db.First(&c, d.CustomerID).Error; err != nil {
		return err
	}

	return r.db.Delete(&c).Error
}

func (r Repository) FindAllCustomers(firstname string, email string, limit string, page string) ([]Customer, error) {
	parseDataAndSave("https://reqres.in/api/users?page=2", r.db)
	var customer []Customer
	limits, _ := strconv.Atoi(limit)
	pages, _ := strconv.Atoi(page)
	offset := (pages - 1) * limits

	err := r.db.Limit(limits).Offset(offset).Find(&customer).Error
	search := r.db
	if firstname != "" {
		search = search.Where("first_name = ?", firstname)
	}
	if email != "" {
		search = search.Where("email = ?", email)
	}
	err = search.Limit(limits).Offset(offset).Find(&customer).Error
	if err != nil {
		return nil, err
	}

	return customer, err
}
func parseDataAndSave(url string, db *gorm.DB) error {
	// Mengirim permintaan GET ke URL JSON
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Membaca isi respon sebagai byte
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Membuat variabel untuk menyimpan respons JSON
	var jsonResponse Response

	// Unmarshal byte menjadi struct Response
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return err
	}

	// Periksa apakah ada data untuk disimpan
	if len(jsonResponse.Data) == 0 {
		return nil
	}

	// Iterasi setiap data dan lakukan proses insert ke database
	for _, userData := range jsonResponse.Data {
		user := Customer{
			Email:     userData.Email,
			FirstName: userData.FirstName,
			LastName:  userData.LastName,
			Avatar:    userData.Avatar,
		}
		db.Create(&user)
	}

	return nil
}
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
