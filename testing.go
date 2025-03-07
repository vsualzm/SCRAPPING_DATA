package main

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// // Struktur response JSON
// type Response struct {
// 	ResponCode  int    `json:"responCode"`
// 	Msg         string `json:"Msg"`
// 	TotalData   int    `json:"total_data"`
// 	CurrentPage int    `json:"current_page"`
// 	TotalPage   int    `json:"total_page"`
// 	Data        []Pos  `json:"data"`
// }

// type Pos struct {
// 	NamaDirian string        `json:"nama_dirian"`
// 	Kabupaten  string        `json:"kabupaten"`
// 	Alamat     string        `json:"alamat"`
// 	JamLayanan string        `json:"jamlayanan"`
// 	Layanan    []LayananType `json:"layanan"`
// 	IdDirian   string        `json:"id_dirian"`
// 	JnsKantor  string        `json:"jnskantor"`
// 	Propinsi   string        `json:"propinsi"`
// 	Kecamatan  string        `json:"kecamatan"`
// 	Kelurahan  string        `json:"kelurahan"`
// }

// type LayananType struct {
// 	Layanan string `json:"layanan"`
// }

// func main() {

// 	// cheking connection database

// 	dsn := "host=localhost port=5432 user=postgres password=1234 dbname=map_db sslmode=disable"
// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Cek koneksi dengan Ping
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("Gagal terhubung ke database:", err)
// 	}

// 	fmt.Println("Berhasil terhubung ke database!")

// }

// // func main() {
// // 	// Koneksi ke PostgreSQL
// // 	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=yourpassword dbname=yourdb sslmode=disable")
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	defer db.Close()

// // 	// Fetch API
// // 	url := "https://postoffice.posindonesia.co.id/backend/externalweb/carikantor"
// 	// requestBody, _ := json.Marshal(map[string]interface{}{
// 	// 	"perPage":     9,
// 	// 	"currentPage": 1,
// 	// 	"cari":        "",
// 	// 	"jnsktr":      "",
// 	// })
// // 	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
// // 	req.Header.Set("Content-Type", "application/json")

// // 	client := &http.Client{}
// // 	resp, err := client.Do(req)
// // 	if err != nil {
// // 		log.Fatal("Error fetching API:", err)
// // 	}
// // 	defer resp.Body.Close()

// // 	body, _ := ioutil.ReadAll(resp.Body)

// // 	// Parse JSON response
// // 	var apiResponse Response
// // 	err = json.Unmarshal(body, &apiResponse)
// // 	if err != nil {
// // 		log.Fatal("Error decoding JSON:", err)
// // 	}

// // 	// Simpan ke database
// // 	for _, pos := range apiResponse.Data {
// // 		// Konversi layanan ke JSONB
// // 		layananJSON, _ := json.Marshal(pos.Layanan)

// // 		_, err := db.Exec(`
// // 			INSERT INTO MAP_POS (nama_dirian, kabupaten, alamat, jamlayanan, layanan, id_dirian, jnskantor, propinsi, kecamatan, kelurahan)
// // 			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
// // 			ON CONFLICT (id_dirian) DO UPDATE
// // 			SET nama_dirian = EXCLUDED.nama_dirian,
// // 			    kabupaten = EXCLUDED.kabupaten,
// // 			    alamat = EXCLUDED.alamat,
// // 			    jamlayanan = EXCLUDED.jamlayanan,
// // 			    layanan = EXCLUDED.layanan,
// // 			    jnskantor = EXCLUDED.jnskantor,
// // 			    propinsi = EXCLUDED.propinsi,
// // 			    kecamatan = EXCLUDED.kecamatan,
// // 			    kelurahan = EXCLUDED.kelurahan;
// // 		`, pos.NamaDirian, pos.Kabupaten, pos.Alamat, pos.JamLayanan, layananJSON, pos.IdDirian, pos.JnsKantor, pos.Propinsi, pos.Kecamatan, pos.Kelurahan)

// // 		if err != nil {
// // 			log.Println("Error inserting data:", err)
// // 		}
// // 	}

// // 	fmt.Println("Data berhasil disimpan ke database!")
// // }

// berhasil

// package main

// import (
// 	"bytes"
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	_ "github.com/lib/pq"
// )

// // Struktur untuk respons API
// type ResponseAPI struct {
// 	ResponCode  int      `json:"responCode"`
// 	Msg         string   `json:"Msg"`
// 	TotalData   int      `json:"total_data"`
// 	CurrentPage int      `json:"current_page"`
// 	TotalPage   int      `json:"total_page"`
// 	Data        []Office `json:"data"`
// }

// // Struktur untuk data kantor pos
// type Office struct {
// 	NamaDirian string    `json:"nama_dirian"`
// 	Kabupaten  string    `json:"kabupaten"`
// 	Alamat     string    `json:"alamat"`
// 	JamLayanan string    `json:"jamlayanan"`
// 	Layanan    []Layanan `json:"layanan"`
// 	IDDirian   string    `json:"id_dirian"`
// 	JnsKantor  string    `json:"jnskantor"`
// 	Propinsi   string    `json:"propinsi"`
// 	Kecamatan  string    `json:"kecamatan"`
// 	Kelurahan  string    `json:"kelurahan"`
// }

// // Struktur untuk layanan
// type Layanan struct {
// 	Layanan string `json:"layanan"`
// }

// func main() {
// 	// 1. Panggil API
// 	url := "https://postoffice.posindonesia.co.id/backend/externalweb/carikantor"
// 	requestBody := map[string]interface{}{
// 		"perPage":     9,
// 		"currentPage": 3,
// 		"cari":        "",
// 		"jnsktr":      "",
// 	}
// 	reqBodyJSON, _ := json.Marshal(requestBody)

// 	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBodyJSON))
// 	if err != nil {
// 		log.Fatal("Gagal mengakses API:", err)
// 	}
// 	defer resp.Body.Close()

// 	body, _ := ioutil.ReadAll(resp.Body)

// 	// 2. Decode JSON
// 	var responseData ResponseAPI
// 	err = json.Unmarshal(body, &responseData)
// 	if err != nil {
// 		log.Fatal("Gagal decode JSON:", err)
// 	}

// 	// 3. Simpan ke database
// 	db, err := sql.Open("postgres", "host=localhost user=postgres password=1234 dbname=map_db port=5432 sslmode=disable")
// 	if err != nil {
// 		log.Fatal("Gagal terhubung ke database:", err)
// 	}
// 	defer db.Close()

// 	log.Println("cek response data", responseData)

// 	// Insert data kantor pos ke database
// 	for _, office := range responseData.Data {
// 		// Konversi layanan ke JSON string
// 		layananJSON, _ := json.Marshal(office.Layanan)

// 		_, err = db.Exec(`
// 			INSERT INTO map_pos (nama_dirian, kabupaten, alamat, jamlayanan, layanan, id_dirian, jnskantor, propinsi, kecamatan, kelurahan)
// 			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
// 			ON CONFLICT (id_dirian) DO NOTHING`, // Hindari duplikasi data
// 			office.NamaDirian, office.Kabupaten, office.Alamat, office.JamLayanan, string(layananJSON),
// 			office.IDDirian, office.JnsKantor, office.Propinsi, office.Kecamatan, office.Kelurahan)

// 		if err != nil {
// 			log.Println("Gagal menyimpan data:", err)
// 		}
// 	}

// 	fmt.Println("Data berhasil disimpan ke database!")
// }
