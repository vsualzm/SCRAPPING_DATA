package main

// import (
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
// 	// Koneksi ke database
// 	db, err := sql.Open("postgres", "host=localhost user=postgres password=1234 dbname=map_db port=5432 sslmode=disable")
// 	if err != nil {
// 		log.Fatal("Gagal terhubung ke database:", err)
// 	}
// 	defer db.Close()

// 	for currentPage := 1; currentPage <= 972; currentPage++ {
// 		fmt.Printf("Fetching data for page %d...\n", currentPage)

// 		// Panggil API
// 		url := "https://postoffice.posindonesia.co.id/backend/externalweb/carikantor"
// 		requestBody := map[string]interface{}{
// 			"perPage":     9,
// 			"currentPage": currentPage,
// 			"cari":        "",
// 			"jnsktr":      "",
// 		}
// 		reqBodyJSON, _ := json.Marshal(requestBody)

// 		resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBodyJSON))
// 		if err != nil {
// 			log.Println("Gagal mengakses API:", err)
// 			continue // Lewati iterasi ini jika terjadi error
// 		}
// 		defer resp.Body.Close()

// 		body, _ := ioutil.ReadAll(resp.Body)

// 		// Decode JSON
// 		var responseData ResponseAPI
// 		err = json.Unmarshal(body, &responseData)
// 		if err != nil {
// 			log.Println("Gagal decode JSON:", err)
// 			continue
// 		}

// 		// Insert data kantor pos ke database
// 		for _, office := range responseData.Data {
// 			// Konversi layanan ke JSON string
// 			layananJSON, _ := json.Marshal(office.Layanan)

// 			_, err = db.Exec(`
// 				INSERT INTO map_pos (nama_dirian, kabupaten, alamat, jamlayanan, layanan, id_dirian, jnskantor, propinsi, kecamatan, kelurahan)
// 				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
// 				ON CONFLICT (id_dirian) DO NOTHING`, // Hindari duplikasi data
// 				office.NamaDirian, office.Kabupaten, office.Alamat, office.JamLayanan, string(layananJSON),
// 				office.IDDirian, office.JnsKantor, office.Propinsi, office.Kecamatan, office.Kelurahan)

// 			if err != nil {
// 				log.Println("Gagal menyimpan data:", err)
// 			}
// 		}

// 		fmt.Printf("Page %d processed successfully!\n", currentPage)

// 		// Tunggu 3 detik sebelum request berikutnya
// 		time.Sleep(1 * time.Second)
// 	}

// 	fmt.Println("Proses selesai!")
// }

// mencari detailKTR
// func main() {

// 	// url := "https://postoffice.posindonesia.co.id/backend/externalweb/detailktr"
// 	// step 1 hit ke url dengan request ambil dari databse id_dirian
// 	// step 2 response data json dari api iru

// 	// {
// 	// 	"responCode": 200,
// 	// 	"Msg": "Data Tersedia",
// 	// 	"data": [
// 	// 		{
// 	// 			"nama_dirian": "CPM MAHKAMAHKONSTITUSI",
// 	// 			"alamat": "Gd. Mahkamah Agung Jl. Medan Merdeka Barat",
// 	// 			"telpon_old": "1500161",
// 	// 			"latitude": "-6.1735296",
// 	// 			"longitude": "106.8220006",
// 	// 			"jenis": "MR",
// 	// 			"jenis_dirian": "Mailing Room",
// 	// 			"propinsi": "DKI JAKARTA",
// 	// 			"kabupaten": "KOTA ADM. JAKARTA PUSAT",
// 	// 			"kecamatan": "GAMBIR",
// 	// 			"kelurahan": "GAMBIR"
// 	// 		}
// 	// 	]
// 	// }

// 	// jns kantor dapet dari map_pos

// }

// ini udah keren
// package main

// import (
// 	"bytes"
// 	"database/sql"
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// // Struktur untuk menampung data dari database
// type Office struct {
// 	IDDirian  string `json:"id_dirian"`
// 	JnsKantor string `json:"jns_kantor"`
// }

// // Struktur request untuk API
// type RequestAPI struct {
// 	IDDirian string `json:"id_dirian"`
// }

// // Struktur respons dari API
// type ResponseAPI struct {
// 	ResponCode int        `json:"responCode"`
// 	Msg        string     `json:"Msg"`
// 	Data       []Kordinat `json:"data"`
// }

// // Struktur data koordinat untuk database
// type Kordinat struct {
// 	NamaDirian  string `json:"nama_dirian"`
// 	Alamat      string `json:"alamat"`
// 	TelponOld   string `json:"telpon_old"`
// 	Latitude    string `json:"latitude"`
// 	Longitude   string `json:"longitude"`
// 	Jenis       string `json:"jenis"`
// 	JenisDirian string `json:"jenis_dirian"`
// 	Propinsi    string `json:"propinsi"`
// 	Kabupaten   string `json:"kabupaten"`
// 	Kecamatan   string `json:"kecamatan"`
// 	Kelurahan   string `json:"kelurahan"`
// }

// func main() {
// 	// Koneksi ke database PostgreSQL
// 	db, err := sql.Open("postgres", "host=localhost user=postgres password=1234 dbname=map_db port=5432 sslmode=disable")
// 	if err != nil {
// 		log.Fatal("Gagal terhubung ke database:", err)
// 	}
// 	defer db.Close()

// 	// Ambil data dari tabel map_pos
// 	rows, err := db.Query("SELECT id_dirian, jnskantor FROM map_pos WHERE jnskantor = 'KCP' AND id > 5815")
// 	if err != nil {
// 		log.Fatal("Gagal mengambil data dari database:", err)
// 	}
// 	defer rows.Close()

// 	// Looping setiap record yang diambil dari database
// 	count := 0
// 	for rows.Next() {
// 		var office Office
// 		if err := rows.Scan(&office.IDDirian, &office.JnsKantor); err != nil {
// 			log.Println("Gagal membaca data dari database:", err)
// 			continue
// 		}

// 		// Buat request ke API eksternal
// 		requestData := RequestAPI{IDDirian: office.IDDirian}
// 		reqBodyJSON, _ := json.Marshal(requestData)

// 		resp, err := http.Post("https://postoffice.posindonesia.co.id/backend/externalweb/detailktr", "application/json", bytes.NewBuffer(reqBodyJSON))
// 		if err != nil {
// 			log.Println("Gagal mengakses API untuk id_dirian", office.IDDirian, ":", err)
// 			continue
// 		}
// 		defer resp.Body.Close()

// 		body, _ := ioutil.ReadAll(resp.Body)

// 		// Decode JSON response
// 		var responseData ResponseAPI
// 		err = json.Unmarshal(body, &responseData)
// 		if err != nil {
// 			log.Println("Gagal decode JSON untuk id_dirian", office.IDDirian, ":", err)
// 			continue
// 		}

// 		// Cek jika respons berhasil dan ada data
// 		if responseData.ResponCode == 200 && len(responseData.Data) > 0 {
// 			data := responseData.Data[0] // Ambil data pertama

// 			// Simpan data ke database
// 			_, err = db.Exec(`
// 				INSERT INTO data_kordinat (id_dirian, jns_kantor, alamat, telpon_old, latitude, longitude, jenis, jenis_dirian, propinsi, kabupaten, kecamatan, kelurahan)
// 				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
// 				ON CONFLICT (id_dirian) DO NOTHING`,
// 				office.IDDirian, office.JnsKantor, data.Alamat, data.TelponOld, data.Latitude, data.Longitude, data.Jenis,
// 				data.JenisDirian, data.Propinsi, data.Kabupaten, data.Kecamatan, data.Kelurahan,
// 			)

// 			if err != nil {
// 				log.Println("Gagal menyimpan data untuk id_dirian", office.IDDirian, ":", err)
// 			} else {
// 				count++
// 				log.Println("Sukses menyimpan data untuk id_dirian:", office.IDDirian)
// 			}
// 		} else {
// 			log.Println("Data tidak tersedia untuk id_dirian:", office.IDDirian)
// 		}

// 		// Tunggu 3 detik sebelum request berikutnya
// 		time.Sleep(1 * time.Second)
// 	}

// 	log.Println("Proses selesai, total data sukses disimpan:", count)
// }

// package main

// import (
// 	_ "github.com/lib/pq"
// )

// // Struktur untuk menampung data dari database
// type Office struct {
// 	IDDirian  string `json:"id_dirian"`
// 	JnsKantor string `json:"jns_kantor"`
// }

// // Struktur request untuk API
// type RequestAPI struct {
// 	IDDirian string `json:"id_dirian"`
// }

// // Struktur respons dari API
// type ResponseAPI struct {
// 	ResponCode int        `json:"responCode"`
// 	Msg        string     `json:"Msg"`
// 	Data       []Kordinat `json:"data"`
// }

// // Struktur data koordinat untuk database
// type Kordinat struct {
// 	NamaDirian  string `json:"nama_dirian"`
// 	Alamat      string `json:"alamat"`
// 	TelponOld   string `json:"telpon_old"`
// 	Latitude    string `json:"latitude"`
// 	Longitude   string `json:"longitude"`
// 	Jenis       string `json:"jenis"`
// 	JenisDirian string `json:"jenis_dirian"`
// 	Propinsi    string `json:"propinsi"`
// 	Kabupaten   string `json:"kabupaten"`
// 	Kecamatan   string `json:"kecamatan"`
// 	Kelurahan   string `json:"kelurahan"`
// }

// func main() {
// 	// Koneksi ke database PostgreSQL
// 	db, err := sql.Open("postgres", "host=localhost user=postgres password=1234 dbname=map_db port=5432 sslmode=disable")
// 	if err != nil {
// 		log.Fatal("Gagal terhubung ke database:", err)
// 	}
// 	defer db.Close()

// 	// Dataset yang ingin diambil
// 	dataSet := []string{
// 		"12110A2", "35691B1", "50518A", "774453B2",
// 		"68122A", "79588B1", "85342B1", "84152B1", "85333B1", "85364B1",
// 		"85393B1", "86262B1", "86686B1", "92255B1", "97656B1", "98439B1",
// 	}

// 	// Membuat placeholder untuk parameter IN query
// 	placeholders := make([]string, len(dataSet))
// 	args := make([]interface{}, len(dataSet))
// 	for i, id := range dataSet {
// 		placeholders[i] = fmt.Sprintf("$%d", i+1)
// 		args[i] = id
// 	}

// 	// Query hanya mengambil ID yang ada di dataset
// 	query := fmt.Sprintf("SELECT id_dirian, jnskantor FROM map_pos WHERE id_dirian IN (%s)", strings.Join(placeholders, ","))
// 	rows, err := db.Query(query, args...)
// 	if err != nil {
// 		log.Fatal("Gagal mengambil data dari database:", err)
// 	}
// 	defer rows.Close()

// 	// Looping setiap record yang diambil dari database
// 	count := 0
// 	for rows.Next() {
// 		var office Office
// 		if err := rows.Scan(&office.IDDirian, &office.JnsKantor); err != nil {
// 			log.Println("Gagal membaca data dari database:", err)
// 			continue
// 		}

// 		// Buat request ke API eksternal
// 		requestData := RequestAPI{IDDirian: office.IDDirian}
// 		reqBodyJSON, _ := json.Marshal(requestData)

// 		resp, err := http.Post("https://postoffice.posindonesia.co.id/backend/externalweb/detailktr", "application/json", bytes.NewBuffer(reqBodyJSON))
// 		if err != nil {
// 			log.Println("Gagal mengakses API untuk id_dirian", office.IDDirian, ":", err)
// 			continue
// 		}
// 		defer resp.Body.Close()

// 		var responseData ResponseAPI
// 		err = json.NewDecoder(resp.Body).Decode(&responseData)
// 		if err != nil {
// 			log.Println("Gagal decode JSON untuk id_dirian", office.IDDirian, ":", err)
// 			continue
// 		}

// 		// Cek jika respons berhasil dan ada data
// 		if responseData.ResponCode == 200 && len(responseData.Data) > 0 {
// 			data := responseData.Data[0] // Ambil data pertama

// 			// Simpan data ke database
// 			_, err = db.Exec(`
// 				INSERT INTO data_kordinat (id_dirian, jns_kantor, alamat, telpon_old, latitude, longitude, jenis, jenis_dirian, propinsi, kabupaten, kecamatan, kelurahan)
// 				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
// 				ON CONFLICT (id_dirian) DO NOTHING`,
// 				office.IDDirian, office.JnsKantor, data.Alamat, data.TelponOld, data.Latitude, data.Longitude, data.Jenis,
// 				data.JenisDirian, data.Propinsi, data.Kabupaten, data.Kecamatan, data.Kelurahan,
// 			)

// 			if err != nil {
// 				log.Println("Gagal menyimpan data untuk id_dirian", office.IDDirian, ":", err)
// 			} else {
// 				count++
// 				log.Println("Sukses menyimpan data untuk id_dirian:", office.IDDirian)
// 			}
// 		} else {
// 			log.Println("Data tidak tersedia untuk id_dirian:", office.IDDirian)
// 		}

// 		// Tunggu 1 detik sebelum request berikutnya
// 		time.Sleep(1 * time.Second)
// 	}

// 	log.Println("Proses selesai, total data sukses disimpan:", count)
// }
