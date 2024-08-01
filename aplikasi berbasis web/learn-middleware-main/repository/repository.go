package repository

import (
	"errors"
	"example-middleware/model"
)

// gunakan untuk mengolah data
type IRepository interface {
	Get() []model.Data //Metode ini mengembalikan sebuah slice (daftar) dari model.Data
	GetById(id int) model.Data //Metode ini mengambil satu data berdasarkan id yang diberikan.
	Update(model.Data) error //Metode ini digunakan untuk memperbarui data yang sudah ada.
						//Menerima satu model.Data sebagai argumen untuk menggantikan data yang ada.
						//Mengembalikan sebuah error jika terjadi kesalahan selama proses pembaruan.
	Delete(id int) error //Metode ini digunakan untuk menghapus data berdasarkan id yang diberikan
	Create(model.Data) error // Metode ini digunakan untuk menambahkan data baru ke entitas penyimpanan.
							//Menerima satu model.Data sebagai argumen yang akan ditambahkan.
}

// in memmory
// sebuah map yang menggunakan tipe int sebagai kunci dan model.Data sebagai nilai.
//Map ini digunakan untuk menyimpan data dengan menggunakan id (integer) sebagai kunci.
type repository struct {
	data map[int]model.Data
}

//membuat objek yang bertanggung jawab atas penyimpanan dan manajemen data. Struktur data yang 
//digunakan untuk penyimpanan adalah map dengan kunci berupa integer (biasanya digunakan untuk 
//mengidentifikasi entri) dan nilai berupa model.Data.
func NewRepository() IRepository {
	return &repository{
		data: make(map[int]model.Data),
	}
}

func (r *repository) Get() []model.Data {
	// Membuat sebuah slice kosong untuk menyimpan data
	var datas []model.Data

	// Iterasi melalui semua entri dalam map r.data
	for _, data := range r.data {
		// Menambahkan setiap data ke dalam slice datas
		datas = append(datas, data)
	}
	// Mengembalikan slice datas yang berisi semua data dari repositori
	return datas
}

func (r *repository) GetById(id int) model.Data {
	return r.data[id]
}

func (r *repository) Update(newData model.Data) error {
	dataById := r.data[newData.ID]

	if dataById.ID == 0 {
		return errors.New("data not found")
	}

	dataById.Name = newData.Name
	dataById.Address = newData.Address
	dataById.PhoneNumber = newData.PhoneNumber
	dataById.Email = newData.Email

	r.data[newData.ID] = dataById
	return nil
}

func (r *repository) Create(data model.Data) error {
	if r.data[data.ID].ID != 0 {
		return errors.New("data already exist")
	}

	r.data[data.ID] = data
	return nil
}

func (r *repository) Delete(id int) error {
	delete(r.data, id)
	return nil
}
