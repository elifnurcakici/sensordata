package main

import (
	"encoding/json"
	"net/http"
)

// Sıcaklık
func getLatestTemperature(w http.ResponseWriter, r *http.Request) { // w http.REsponseWriter : HTTP yanıtını yazmak için kullanılır,
	// r *http.Request : İstek bilgilerini yazmak için kullanılır.
	var d TemperatureData                               //TemperatureData türünde bir değişken tanımlanıyor.
	result := gormDB.Order("created_at desc").First(&d) // Veritabanında created_at sütununa göre en yeni kaydı getiriyor.
	if result.Error != nil {
		http.Error(w, "Veri bulunamadı", 500)
		return
	}
	json.NewEncoder(w).Encode(d) // Gelen d verisi JSON formatına çevrilip HTTP cevabı olarak istemciye gönderiliyor.
}

func getAllTemperature(w http.ResponseWriter, r *http.Request) {
	var data []TemperatureData
	result := gormDB.Order("created_at desc").Limit(10).Find(&data)
	if result.Error != nil {
		http.Error(w, "Veriler alınamadı", 500)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// Nem
func getLatestHumidity(w http.ResponseWriter, r *http.Request) {
	var d HumidityData
	result := gormDB.Order("created_at desc").First(&d)
	if result.Error != nil {
		http.Error(w, "Veri bulunamadı", 500)
		return
	}
	json.NewEncoder(w).Encode(d)
}

func getAllHumidity(w http.ResponseWriter, r *http.Request) {
	var data []HumidityData
	result := gormDB.Order("created_at desc").Limit(10).Find(&data)
	if result.Error != nil {
		http.Error(w, "Veriler alınamadı", 500)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// Basınç
func getLatestPressure(w http.ResponseWriter, r *http.Request) {
	var d PressureData
	result := gormDB.Order("created_at desc").First(&d)
	if result.Error != nil {
		http.Error(w, "Veri bulunamadı", 500)
		return
	}
	json.NewEncoder(w).Encode(d)
}

func getAllPressure(w http.ResponseWriter, r *http.Request) {
	var data []PressureData
	result := gormDB.Order("created_at desc").Limit(10).Find(&data)
	if result.Error != nil {
		http.Error(w, "Veriler alınamadı", 500)
		return
	}
	json.NewEncoder(w).Encode(data)
}
