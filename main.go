package main

import (
	"log"
	"net/http"
)

func main() {
	gormDB = connectDB() // connectDB( )fonksiyonu ile veritabanına bağlanılıyor
	// GORM kütüphanesi ile veritabanı işlemleri gormDB nesnesi üzerinden yapılır.

	// Arka planda sahte veri üretme (3 sensör)
	go generateFakeData()

	// API endpointleri (REST)
	http.HandleFunc("/data/temperature/all", getAllTemperature)
	http.HandleFunc("/data/humidity/all", getAllHumidity)
	http.HandleFunc("/data/pressure/all", getAllPressure)

	http.HandleFunc("/data/temperature/latest", getLatestTemperature)
	http.HandleFunc("/data/humidity/latest", getLatestHumidity)
	http.HandleFunc("/data/pressure/latest", getLatestPressure)

	// WebSocket endpointi
	http.HandleFunc("/ws", handleWebSocket) // /ws adresine gelen istekleri handleWebSocket fonksiyonu ile karşılar.
	// Bu fonksiyon gerçek zamnalı veri iletimi için websocket bağlantılarını yönetir.

	// Statik dosyaları (HTML/JS) servis et
	http.Handle("/", http.FileServer(http.Dir("./static"))) // ./static klasöründeki dosyaları web üzerinden sunar.

	log.Println("Server 8080 portunda çalışıyor...")
	http.ListenAndServe(":8080", nil)
}
