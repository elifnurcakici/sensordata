package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB //global veritabanı bağlantısı

// PostgreSQL'e bağlanır ve tabloyu otomatik oluşturur
func connectDB() *gorm.DB {
	// PostgreSQL bağlantı bilgileri
	dsn := "host=localhost user=postgres password=843038 dbname=sensordb port=5432 sslmode=disable TimeZone=Asia/Istanbul"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // gorm.Open(...): PostgreSQL bağlantısını açar.
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı:", err)
	}

	// Eğer yoksa tabloyu otomatik oluşturur ya da gerekiyorsa günceller
	if err := db.AutoMigrate(&TemperatureData{}, &HumidityData{}, &PressureData{}); err != nil {
		log.Fatal("Tablo Oluştutulamadı:", err)
	}

	return db
}

// Veri ekleme fonksiyonları
func insertTemperature(val float64) TemperatureData {
	data := TemperatureData{Value: val} // 1. Yeni bir TemperatureData nesnesi oluşturuluyor.
	gormDB.Create(&data)                // 2. Bu nesne veritabanına kaydediliyor.
	return data                         // 3. Kaydedilen nesne geri döndürülüyor.
}

func insertHumidity(val float64) HumidityData {
	data := HumidityData{Value: val} // 1. Nem verisi için bir struct oluşturuluyor.
	gormDB.Create(&data)             // 2. Veritabanına kaydediliyor.
	return data
}

func insertPressure(val float64) PressureData {
	data := PressureData{Value: val} // 1. Basınç verisi için bir struct oluşturuluyor.
	gormDB.Create(&data)             // 2. Veritabanına kaydediliyor.
	return data                      // 3. Kaydedilen veri geri döndürülüyor.
}
