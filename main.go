package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

// Struktur untuk menampung data dari database
type Office struct {
	IDDirian  string `json:"id_dirian"`
	JnsKantor string `json:"jns_kantor"`
}

// Struktur request untuk API
type RequestAPI struct {
	IDDirian string `json:"id_dirian"`
}

// Struktur respons dari API
type ResponseAPI struct {
	ResponCode int        `json:"responCode"`
	Msg        string     `json:"Msg"`
	Data       []Kordinat `json:"data"`
}

// Struktur data koordinat untuk database
type Kordinat struct {
	NamaDirian  string `json:"nama_dirian"`
	Alamat      string `json:"alamat"`
	TelponOld   string `json:"telpon_old"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Jenis       string `json:"jenis"`
	JenisDirian string `json:"jenis_dirian"`
	Propinsi    string `json:"propinsi"`
	Kabupaten   string `json:"kabupaten"`
	Kecamatan   string `json:"kecamatan"`
	Kelurahan   string `json:"kelurahan"`
}

func main() {
	// Koneksi ke database PostgreSQL
	db, err := sql.Open("postgres", "host=localhost user=postgres password=1234 dbname=map_db port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	defer db.Close()

	// Ambil data dari tabel map_pos
	rows, err := db.Query("SELECT id_dirian, jnskantor FROM map_pos WHERE jnskantor = 'KCU'")
	if err != nil {
		log.Fatal("Gagal mengambil data dari database:", err)
	}
	defer rows.Close()

	// Looping setiap record yang diambil dari database
	count := 0
	for rows.Next() {
		var office Office
		if err := rows.Scan(&office.IDDirian, &office.JnsKantor); err != nil {
			log.Println("Gagal membaca data dari database:", err)
			continue
		}

		// Buat request ke API eksternal
		requestData := RequestAPI{IDDirian: office.IDDirian}
		reqBodyJSON, _ := json.Marshal(requestData)

		resp, err := http.Post("https://postoffice.posindonesia.co.id/backend/externalweb/detailktr", "application/json", bytes.NewBuffer(reqBodyJSON))
		if err != nil {
			log.Println("Gagal mengakses API untuk id_dirian", office.IDDirian, ":", err)
			continue
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		// Decode JSON response
		var responseData ResponseAPI
		err = json.Unmarshal(body, &responseData)
		if err != nil {
			log.Println("Gagal decode JSON untuk id_dirian", office.IDDirian, ":", err)
			continue
		}

		// Cek jika respons berhasil dan ada data
		if responseData.ResponCode == 200 && len(responseData.Data) > 0 {
			data := responseData.Data[0] // Ambil data pertama

			// Simpan data ke database
			_, err = db.Exec(
				`INSERT INTO data_kordinat_new (id_dirian, jns_kantor, alamat, telpon_old, latitude, longitude, jenis, jenis_dirian, propinsi, kabupaten, kecamatan, kelurahan)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
				ON CONFLICT (id_dirian) DO NOTHING`,
				office.IDDirian, office.JnsKantor, data.Alamat, data.TelponOld, data.Latitude, data.Longitude, data.Jenis,
				data.JenisDirian, data.Propinsi, data.Kabupaten, data.Kecamatan, data.Kelurahan,
			)

			if err != nil {
				log.Println("Gagal menyimpan data untuk id_dirian", office.IDDirian, ":", err)
			} else {
				count++
				log.Println("Sukses menyimpan data untuk id_dirian:", office.IDDirian)
			}
		} else {
			log.Println("Data tidak tersedia untuk id_dirian:", office.IDDirian)
		}

		// Tunggu 1 detik sebelum request berikutnya
		time.Sleep(1 * time.Second)
	}

	log.Println("Proses selesai, total data sukses disimpan:", count)
}
