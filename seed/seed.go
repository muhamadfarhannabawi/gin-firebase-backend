package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/muhamadfarhannabawi/gin-firebase-backend/config"
	"github.com/muhamadfarhannabawi/gin-firebase-backend/models"
)

func main() {
	godotenv.Load()
	config.InitDatabase()

	products := []models.Product{
		{Name: "Sandal Gunung Pria", Price: 75000, Category: "Sandal Pria", Stock: 40,
			Description: "Sandal gunung kuat dan nyaman untuk aktivitas outdoor", ImageURL: "https://picsum.photos/400"},

		{Name: "Sandal Jepit Wanita", Price: 25000, Category: "Sandal Wanita", Stock: 100,
			Description: "Sandal jepit ringan dengan desain kekinian", ImageURL: "https://picsum.photos/401"},

		{Name: "Sandal Anak Karakter", Price: 30000, Category: "Sandal Anak", Stock: 60,
			Description: "Sandal anak dengan gambar karakter lucu", ImageURL: "https://picsum.photos/402"},

		{Name: "Sandal Kulit Pria Premium", Price: 120000, Category: "Sandal Pria", Stock: 25,
			Description: "Sandal kulit asli dengan kualitas premium", ImageURL: "https://picsum.photos/403"},

		{Name: "Sandal Rumah Wanita", Price: 20000, Category: "Sandal Wanita", Stock: 80,
			Description: "Sandal rumah empuk dan nyaman dipakai sehari-hari", ImageURL: "https://picsum.photos/404"},
	}

	for _, p := range products {
		config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
