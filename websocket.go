package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // Bağlı tüm WebSocket istemcilerini tutan bir map (clients)
var upgrader = websocket.Upgrader{}          // HTTP'yi WebSocket'e çevirir (upgrader)
//.Upgrader : HTTP bağlantılarını WebSocket bağlantısına yükseltmek için kullanılan yapı

// WebSocket bağlantısı
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // istekleri engelleyen kısıtlamayı kaldırmak için burada true döndürülüyor,
	// yani tüm kaynaklardan gelen bağlantılara izin veriliyor
	ws, err := upgrader.Upgrade(w, r, nil) // normal http bağlantısı burada websockete çevriliyor
	if err != nil {
		log.Println("WebSocket hatası:", err)
		return
	}
	defer ws.Close() // fonskiyon bitince ws bağlantısı kapatılıyor

	clients[ws] = true // bu yeni ws bağlantısı clients'e ekleniyor. Yani artık bu client e veri gönderilebilir.
	for {
		_, _, err := ws.ReadMessage() // ws den mesaj okunmaya çalışıyor
		if err != nil {               // eğer alınamazsa yani bağlantı kapanırsa
			delete(clients, ws) // bu client clients nesnesinden siliniyor
			break
		}
	}
}

// Yeni verileri tüm istemcilere gönderir
func broadcast(sensorType string, data interface{}) { // sensör türü ve datayı alıyor
	msg := map[string]interface{}{ // bir json objesi oluşturuyor
		"type": sensorType,
		"data": data,
	}
	jsonData, _ := json.Marshal(msg) // msg JSON a dönüştürülüyor.
	for client := range clients {    // clients map indeki tüm bağlı istemcilere
		client.WriteMessage(websocket.TextMessage, jsonData) // bu msg mesajı gönderiliyor
	}
}

// Arka planda her sensör için sahte veri üretir ve yayar
func generateFakeData() {
	for {
		// Sıcaklık (20-30°C)
		temp := insertTemperature(20 + rand.Float64()*10) // insertTemperature ile veritabanına kaydediliyor.
		broadcast("temperature", temp)                    // tüm bağlı istemcilere gönderiliyor.

		// Nem (40-60%)
		hum := insertHumidity(40 + rand.Float64()*20)
		broadcast("humidity", hum)

		// Basınç (950-1050 hPa)
		press := insertPressure(950 + rand.Float64()*100)
		broadcast("pressure", press)

		time.Sleep(5 * time.Second) // her döngü sonunda 5 saniye bekliyor
	}
}
