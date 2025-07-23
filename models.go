package main

import "time"

// TemperatureData : Sensör verilerini temsil eden model
type TemperatureData struct { // sıcaklık
	ID        int       `json:"id" gorm:"primaryKey"` // otomatik birincil anahtar
	Value     float64   `json:"value"`                // Ölçülen sensör değeri
	CreatedAt time.Time `json:"created_at" `          // Kaydın oluşturulduğu zaman
}

// gorm:"autoCreateTime

type HumidityData struct { // nem
	ID        int       `json:"id" gorm:"primaryKey"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at" `
}

type PressureData struct { // basınç
	ID        int       `json:"id" gorm:"primaryKey"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at" `
}
