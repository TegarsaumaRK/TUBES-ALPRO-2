package main

import "fmt"

const NMAXMEJA int = 100
const NMAXRESERVASI int = 100

var nextKodeMeja int = 1

type Meja struct {
	kodeMeja  int
	kapasitas int
	status    string
}
type TabMeja [NMAXMEJA]Meja

type Reservasi struct {
	kodeReservasi int
	namaPelanggan string
	jumlahOrang   int
	kodeMeja      int
}
type TabReservasi [NMAXRESERVASI]Reservasi

func main() {
	var dataMeja TabMeja
	var dataReservasi TabReservasi
	var nMeja, nReservasi int
	var pilihanUtama, pilihanManajemen, pilihanReservasi int

	nMeja = 0
	nReservasi = 0
	pilihanUtama = 0

	for pilihanUtama != 3 {
		fmt.Println("======================================")
		fmt.Println(" APLIKASI RESERVASI MEJA RESTORAN")
		fmt.Println("======================================")
		fmt.Println("1. Manajemen Kursi / Meja")
		fmt.Println("2. Reservasi Meja")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihanUtama)

		if pilihanUtama == 1 {
			pilihanManajemen = 0

			for pilihanManajemen != 7 {
				fmt.Println()
				fmt.Println("======================================")
				fmt.Println(" MENU MANAJEMEN KURSI / MEJA")
				fmt.Println("======================================")
				fmt.Println("1. Tambah Data Meja")
				fmt.Println("2. Hapus Data Meja")
				fmt.Println("3. Cari Data Meja")
				fmt.Println("4. Update Data Meja")
				fmt.Println("5. Urutkan Meja Berdasarkan Kapasitas")
				fmt.Println("6. Tampilkan Daftar Data Meja")
				fmt.Println("7. Kembali ke Menu Utama")
				fmt.Print("Pilih menu: ")
				fmt.Scan(&pilihanManajemen)

				if pilihanManajemen == 1 {
					tambahDataMeja(&dataMeja, &nMeja)
				} else if pilihanManajemen == 2 {
					hapusDataMeja(&dataMeja, &nMeja)
				} else if pilihanManajemen == 3 {
					cariDataMeja(dataMeja, nMeja)
				} else if pilihanManajemen == 4 {
					updateDataMeja(&dataMeja, nMeja)
				} else if pilihanManajemen == 5 {
					selectionSortMeja(&dataMeja, nMeja)
				} else if pilihanManajemen == 6 {
					tampilDataMeja(dataMeja, nMeja)
				} else if pilihanManajemen == 7 {
					fmt.Println("Kembali ke menu utama.")
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}

		} else if pilihanUtama == 2 {
			pilihanReservasi = 0

			for pilihanReservasi != 7 {
				fmt.Println()
				fmt.Println("======================================")
				fmt.Println(" MENU RESERVASI MEJA")
				fmt.Println("======================================")
				fmt.Println("1. Cari Meja Tersedia Berdasarkan Kapasitas")
				fmt.Println("2. Buat Reservasi")
				fmt.Println("3. Batalkan Reservasi")
				fmt.Println("4. Urutkan Reservasi Berdasarkan Jumlah Orang")
				fmt.Println("5. Tampilkan Data Reservasi")
				fmt.Println("6. Cari Data Reservasi")
				fmt.Println("7. Kembali ke Menu Utama")
				fmt.Print("Pilih menu: ")
				fmt.Scan(&pilihanReservasi)

				if pilihanReservasi == 1 {
					cariMejaTersedia(dataMeja, nMeja)
				} else if pilihanReservasi == 2 {
					buatReservasi(&dataMeja, nMeja, &dataReservasi, &nReservasi)
				} else if pilihanReservasi == 3 {
					batalkanReservasi(&dataMeja, nMeja, &dataReservasi, &nReservasi)
				} else if pilihanReservasi == 4 {
					insertionSortReservasi(&dataReservasi, nReservasi)
				} else if pilihanReservasi == 5 {
					tampilDataReservasi(dataReservasi, nReservasi)
				} else if pilihanReservasi == 6 {
					cariDataReservasi(dataReservasi, nReservasi)
				} else if pilihanReservasi == 7 {
					fmt.Println("Kembali ke menu utama.")
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		} else if pilihanUtama == 3 {
			fmt.Println()
			fmt.Println("Program selesai. Terima kasih.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}

		fmt.Println()
	}
}

func tambahDataMeja(T *TabMeja, n *int) {
// IS: T berisi n data meja, T belum penuh
// FS: data meja baru ditambahkan dengan kode otomatis, n bertambah 1
	var kapasitas int

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println(" TAMBAH DATA MEJA")
	fmt.Println("======================================")

	if *n >= NMAXMEJA {
		fmt.Println("Data meja sudah penuh.")
	} else {
		fmt.Print("Masukkan kapasitas meja: ")
		fmt.Scan(&kapasitas)

		T[*n].kodeMeja = nextKodeMeja
		T[*n].kapasitas = kapasitas
		T[*n].status = "tersedia"
		nextKodeMeja++
		*n = *n + 1

		fmt.Println("Data meja berhasil ditambahkan.")
		fmt.Println("Kode Meja yang diberikan:", T[*n-1].kodeMeja)
	}
}

func binarySearchMeja(T TabMeja, n int, kode int) int {
// IS: T berisi n data meja yang terurut ascending by kodeMeja, kode adalah nilai yang dicari
// FS: mengembalikan indeks meja jika ditemukan, -1 jika tidak ada
	var left, right, mid, ketemu int

	left = 0
	right = n - 1
	ketemu = -1

	for left <= right && ketemu == -1 {
		mid = (left + right) / 2
		if T[mid].kodeMeja == kode {
			ketemu = kode - 1
		} else if kode < T[mid].kodeMeja {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ketemu
}

func hapusDataMeja(T *TabMeja, n *int) {
// IS: T berisi n data meja
// FS: meja dengan kode yang diminta dihapus, elemen digeser, n berkurang 1
	var kode, posisi, i int

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println(" HAPUS DATA MEJA")
	fmt.Println("======================================")

	if *n == 0 {
		fmt.Println("Data meja masih kosong.")
	} else {
		tampilDataMeja(*T, *n)
		fmt.Print("Masukkan kode meja yang ingin dihapus: ")
		fmt.Scan(&kode)

		posisi = sequentialSearchMeja(*T, *n, kode)

		if posisi == -1 {
			fmt.Println("Data meja tidak ditemukan.")
		} else {
			i = posisi
			for i <= *n-2 {
				T[i] = T[i+1]
				i = i + 1
			}
			T[*n-1].kodeMeja = 0
			T[*n-1].kapasitas = 0
			T[*n-1].status = ""
			*n = *n - 1
			fmt.Println("Data meja berhasil dihapus.")
		}
	}
}

func cariDataMeja(T TabMeja, n int) {
// IS: T berisi n data meja yang terurut by kodeMeja (binary search)
// FS: menampilkan detail meja yang dicari, atau pesan tidak ditemukan
	var kode, posisi int

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println(" CARI DATA MEJA")
	fmt.Println("======================================")

	if n == 0 {
		fmt.Println("Data meja masih kosong.")
	} else {
		fmt.Print("Masukkan kode meja yang dicari: ")
		fmt.Scan(&kode)

		posisi = binarySearchMeja(T, n, kode)

		if posisi == -1 {
			fmt.Println("Data meja tidak ditemukan.")
		} else {
			fmt.Println()
			fmt.Println("Data meja ditemukan.")
			fmt.Println("--------------------------------------")
			fmt.Println("Kode Meja :", T[posisi].kodeMeja)
			fmt.Println("Kapasitas :", T[posisi].kapasitas, "kursi")
			fmt.Println("Status    :", T[posisi].status)
			fmt.Println("--------------------------------------")
		}
	}
}

func updateDataMeja(T *TabMeja, n int) {
// IS: T berisi n data meja
// FS: kapasitas atau status meja yang dipilih berhasil diubah
	var kode, posisi, kapasitasBaru, pilihanUpdate int
	var statusBaru string

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println(" UPDATE DATA MEJA")
	fmt.Println("======================================")

	if n == 0 {
		fmt.Println("Data meja masih kosong.")
	} else {
		tampilDataMeja(*T, n)
		fmt.Print("Masukkan kode meja yang ingin diupdate: ")
		fmt.Scan(&kode)

		posisi = sequentialSearchMeja(*T, n, kode)

		if posisi == -1 {
			fmt.Println("Data meja tidak ditemukan.")
		} else {
			fmt.Println()
			fmt.Println("Data meja ditemukan.")
			fmt.Println("--------------------------------------")
			fmt.Println("Kode Meja :", T[posisi].kodeMeja)
			fmt.Println("Kapasitas :", T[posisi].kapasitas, "kursi")
			fmt.Println("Status    :", T[posisi].status)
			fmt.Println("--------------------------------------")

			fmt.Println()
			fmt.Println("Pilih data yang ingin diupdate:")
			fmt.Println("1. Kapasitas Meja")
			fmt.Println("2. Status Meja")
			fmt.Println("3. Kapasitas dan Status Meja")
			fmt.Println("4. Kembali")
			fmt.Print("Pilih menu: ")
			fmt.Scan(&pilihanUpdate)

			if pilihanUpdate == 1 {
				fmt.Print("Masukkan kapasitas baru: ")
				fmt.Scan(&kapasitasBaru)
				T[posisi].kapasitas = kapasitasBaru
				fmt.Println("Kapasitas meja berhasil diupdate.")
			} else if pilihanUpdate == 2 {
				fmt.Print("Masukkan status baru tersedia/dipesan: ")
				fmt.Scan(&statusBaru)
				T[posisi].status = statusBaru
				fmt.Println("Status meja berhasil diupdate.")
			} else if pilihanUpdate == 3 {
				fmt.Print("Masukkan kapasitas baru: ")
				fmt.Scan(&kapasitasBaru)
				fmt.Print("Masukkan status baru tersedia/dipesan: ")
				fmt.Scan(&statusBaru)
				T[posisi].kapasitas = kapasitasBaru
				T[posisi].status = statusBaru
				fmt.Println("Kapasitas dan status meja berhasil diupdate.")
			} else if pilihanUpdate == 4 {
				fmt.Println("Update data meja dibatalkan.")
			} else {
				fmt.Println("Pilihan update tidak valid.")
			}
		}
	}
}

func selectionSortMeja(T *TabMeja, n int) {
// IS: T berisi n data meja
// FS: T diurutkan ascending atau descending berdasarkan kapasitas sesuai pilihan
	var pilihan int

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println(" URUTKAN MEJA BERDASARKAN KAPASITAS")
	fmt.Println("======================================")

	if n == 0 {
		fmt.Println("Data meja masih kosong.")
	} else {
		fmt.Println("1. Urutkan Ascending  (kecil ke besar)")
		fmt.Println("2. Urutkan Descending (besar ke kecil)")
		fmt.Print("Pilih urutan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			selectionSortAsc(T, n)
		} else if pilihan == 2 {
			selectionSortDesc(T, n)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func selectionSortAsc(T *TabMeja, n int) {
// IS: T berisi n data meja dengan urutan sembarang
// FS: T terurut ascending berdasarkan kapasitas
	var pass, idx, i int
	var temp Meja

	for pass = 1; pass <= n-1; pass++ {
		idx = pass - 1
		for i = pass; i <= n-1; i++ {
			if T[idx].kapasitas > T[i].kapasitas {
				idx = i
			}
		}
		temp = T[pass-1]
		T[pass-1] = T[idx]
		T[idx] = temp
	}
	fmt.Println("Data meja berhasil diurutkan dari kecil ke besar.")
}

func selectionSortDesc(T *TabMeja, n int) {
// IS: T berisi n data meja dengan urutan sembarang
// FS: T terurut descending berdasarkan kapasitas
	var pass, idx, i int
	var temp Meja

	for pass = 1; pass <= n-1; pass++ {
		idx = pass - 1
		for i = pass; i <= n-1; i++ {
			if T[idx].kapasitas < T[i].kapasitas {
				idx = i
			}
		}
		temp = T[pass-1]
		T[pass-1] = T[idx]
		T[idx] = temp
	}
	fmt.Println("Data meja berhasil diurutkan dari besar ke kecil.")
}

func tampilDataMeja(T TabMeja, n int) {
// IS: T berisi n data meja
// FS: menampilkan semua data meja dalam bentuk tabel
	var i int
	var garis string

	garis = "------------------------------------------------------"

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println(" DAFTAR DATA MEJA")
	fmt.Println("======================================================")

	if n == 0 {
		fmt.Println("Data meja masih kosong.")
	} else {
		fmt.Println(garis)
		fmt.Printf("%-3s | %-9s | %-13s | %-10s\n", "No", "Kode Meja", "Kapasitas", "Status")
		fmt.Println(garis)
		i = 0
		for i < n {
			fmt.Printf("%-3d | %-9d | %-3d %-9s | %-10s\n", i+1, T[i].kodeMeja, T[i].kapasitas, "kursi", T[i].status)
			i = i + 1
		}
		fmt.Println(garis)
		fmt.Println("Jumlah data meja:", n)
	}
}

func cariMejaTersedia(T TabMeja, n int) {
// IS: T berisi n data meja, user memasukkan jumlah orang
// FS: menampilkan daftar meja yang tersedia dan kapasitasnya mencukupi
	var jumlahOrang, i int
	var ditemukan bool

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println(" CARI MEJA TERSEDIA")
	fmt.Println("======================================")

	if n == 0 {
		fmt.Println("Data meja masih kosong.")
	} else {
		fmt.Print("Masukkan jumlah orang: ")
		fmt.Scan(&jumlahOrang)

		ditemukan = false

		fmt.Println()
		fmt.Println("Daftar meja yang tersedia:")
		fmt.Println("------------------------------------------------")
		fmt.Printf("%-10s %-12s %-10s\n", "Kode", "Kapasitas", "Status")
		fmt.Println("------------------------------------------------")

		for i = 0; i < n; i++ {
			if T[i].kapasitas >= jumlahOrang && T[i].status == "tersedia" {
				fmt.Printf("%-10d %-12d %-10s\n", T[i].kodeMeja, T[i].kapasitas, T[i].status)
				ditemukan = true
			}
		}

		fmt.Println("------------------------------------------------")

		if !ditemukan {
			fmt.Println("Tidak ada meja yang sesuai.")
		}
	}
}

func buatReservasi(T *TabMeja, nMeja int, R *TabReservasi, nReservasi *int) {
// IS: T berisi n data meja, R berisi nReservasi data reservasi
// FS: reservasi baru dibuat, status meja berubah jadi dipesan, nReservasi bertambah 1
	var nama string
	var jumlahOrang, kodeReservasiBaru int
	var idxTerbaik, i int

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println(" BUAT RESERVASI")
	fmt.Println("======================================")

	if *nReservasi >= NMAXRESERVASI {
		fmt.Println("Data reservasi sudah penuh.")
	} else {
		fmt.Print("Masukkan nama pelanggan : ")
		fmt.Scan(&nama)

		fmt.Print("Masukkan jumlah orang   : ")
		fmt.Scan(&jumlahOrang)
		
		idxTerbaik = -1

		for i = 0; i < nMeja; i++ {
			if T[i].status == "tersedia" && T[i].kapasitas >= jumlahOrang {
				if idxTerbaik == -1 {
					idxTerbaik = i
				} else if T[i].kapasitas < T[idxTerbaik].kapasitas {
					idxTerbaik = i
				}
			}
		}

		if idxTerbaik == -1 {
			fmt.Println()
			fmt.Println("Tidak ada meja yang tersedia untuk jumlah orang tersebut.")
		} else {
			kodeReservasiBaru = *nReservasi + 1

			R[*nReservasi].kodeReservasi = kodeReservasiBaru
			R[*nReservasi].namaPelanggan = nama
			R[*nReservasi].jumlahOrang = jumlahOrang
			R[*nReservasi].kodeMeja = T[idxTerbaik].kodeMeja
			*nReservasi = *nReservasi + 1

			T[idxTerbaik].status = "dipesan"

			fmt.Println()
			fmt.Println("Reservasi berhasil dibuat.")
			fmt.Println("--------------------------------------")
			fmt.Println("Kode Reservasi :", kodeReservasiBaru)
			fmt.Println("Nama Pelanggan :", nama)
			fmt.Println("Jumlah Orang   :", jumlahOrang)
			fmt.Println("Kode Meja      :", T[idxTerbaik].kodeMeja)
			fmt.Println("Kapasitas Meja :", T[idxTerbaik].kapasitas)
			fmt.Println("--------------------------------------")
		}
	}
}

func sequentialSearchReservasi(R TabReservasi, n int, nama string) int {
// IS: R berisi n data reservasi, nama adalah nilai yang dicari
// FS: mengembalikan indeks reservasi jika ditemukan, -1 jika tidak ada
	var i, idx int

	idx = -1
	i = 0
	for i < n && idx == -1 {
		if R[i].namaPelanggan == nama {
			idx = i
		}
		i = i + 1
	}
	return idx
}

func sequentialSearchReservasiKode(R TabReservasi, n int, x int) int {
// IS: R berisi n data reservasi, x adalah nilai yang dicari
// FS: mengembalikan indeks reservasi jika ditemukan, -1 jika tidak ada
	var i, idx int

	idx = -1
	i = 0
	for i < n && idx == -1 {
		if R[i].kodeReservasi == x {
			idx = i
		}
		i = i + 1
	}
	return idx
}

func batalkanReservasi(T *TabMeja, nMeja int, R *TabReservasi, nReservasi *int) {
// IS: T berisi nMeja data meja, R berisi nReservasi data reservasi
// FS: reservasi dihapus, status meja kembali tersedia, nReservasi berkurang 1
	var kodeReservasi, i int
	var idx, kodeR int

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println(" BATALKAN RESERVASI")
	fmt.Println("======================================")

	if *nReservasi == 0 {
		fmt.Println("Belum ada data reservasi.")
	} else {
		tampilDataReservasi(*R, *nReservasi)
		fmt.Print("Masukkan kode reservasi : ")
		fmt.Scan(&kodeReservasi)
		kodeR = sequentialSearchReservasiKode(*R, *nReservasi, kodeReservasi)
		idx = sequentialSearchMeja(*T , nMeja, R[kodeR].kodeMeja)
		T[idx].status = "tersedia"
		for i = kodeReservasi; i < *nReservasi-1; i++ {
			R[i] = R[i+1]
		}
		*nReservasi = *nReservasi - 1
		fmt.Println("Reservasi berhasil dibatalkan.")
	}
}

func insertionSortReservasi(R *TabReservasi, n int) {
// IS: R berisi n data reservasi
// FS: R diurutkan ascending atau descending berdasarkan jumlah orang sesuai pilihan
	var pilihan int

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println(" URUTKAN DATA RESERVASI")
	fmt.Println("======================================")

	if n <= 1 {
		fmt.Println("Data reservasi tidak perlu diurutkan.")
	} else {
		fmt.Println("1. Urutkan Ascending  (kecil ke besar)")
		fmt.Println("2. Urutkan Descending (besar ke kecil)")
		fmt.Print("Pilih urutan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			insertionSortAsc(R, n)
		} else if pilihan == 2 {
			insertionSortDesc(R, n)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func insertionSortAsc(R *TabReservasi, n int) {
// IS: R berisi n data reservasi dengan urutan sembarang
// FS: R terurut ascending berdasarkan jumlah orang
	var pass, i int
	var temp Reservasi

	for pass = 1; pass < n; pass++ {
		temp = R[pass]
		i = pass
		for i > 0 && temp.jumlahOrang < R[i-1].jumlahOrang {
			R[i] = R[i-1]
			i--
		}
		R[i] = temp
	}
	fmt.Println("Data reservasi berhasil diurutkan dari kecil ke besar.")
}

func insertionSortDesc(R *TabReservasi, n int) {
// IS: R berisi n data reservasi dengan urutan sembarang
// FS: R terurut descending berdasarkan jumlah orang
	var pass, i int
	var temp Reservasi

	for pass = 1; pass < n; pass++ {
		temp = R[pass]
		i = pass
		for i > 0 && temp.jumlahOrang > R[i-1].jumlahOrang {
			R[i] = R[i-1]
			i--
		}
		R[i] = temp
	}
	fmt.Println("Data reservasi berhasil diurutkan dari besar ke kecil.")
}

func tampilDataReservasi(R TabReservasi, n int) {
// IS: R berisi n data reservasi
// FS: menampilkan semua data reservasi dalam bentuk tabel
	var i int
	var garis string

	garis = "---------------------------------------"

	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println(" DATA RESERVASI")
	fmt.Println("=======================================")

	if n == 0 {
		fmt.Println("Belum ada data reservasi.")
	} else {
		fmt.Println(garis)
		fmt.Printf("%-5s %-10s %-15s %-10s\n", "Kode", "Meja", "Nama", "Orang")
		fmt.Println(garis)

		for i = 0; i < n; i++ {
			fmt.Printf("%-5d %-10d %-15s %-10d\n", R[i].kodeReservasi, R[i].kodeMeja, R[i].namaPelanggan, R[i].jumlahOrang)
		}

		fmt.Println(garis)
		fmt.Println("Jumlah reservasi :", n)
	}
}

func cariDataReservasi(R TabReservasi, n int) {
// IS: R berisi n data meja
// FS: menampilkan detail Reservasi yang dicari, atau pesan tidak ditemukan
	var posisi int
	var nama string

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println(" CARI DATA Reservasi")
	fmt.Println("======================================")

	if n == 0 {
		fmt.Println("Data Reservasi masih kosong.")
	} else {
		fmt.Print("Masukkan Nama Pelanggan yang dicari: ")
		fmt.Scan(&nama)

		posisi = sequentialSearchReservasi(R, n, nama)

		if posisi == -1 {
			fmt.Println("Data Reservasi tidak ditemukan.")
		} else {
			fmt.Println()
			fmt.Println("Data meja ditemukan.")
			fmt.Println("--------------------------------------")
			fmt.Println("Nama Pelanggan :", R[posisi].namaPelanggan)
			fmt.Println("Kode Reservasi :", R[posisi].kodeReservasi)
			fmt.Println("Kode Meja      :", R[posisi].kodeMeja)
			fmt.Println("Jumlah Orang   :", R[posisi].jumlahOrang, "Orang")
			fmt.Println("--------------------------------------")
		}
	}
}
// IS: T berisi n data meja, kode adalah nilai yang dicari (sequential)
// FS: mengembalikan indeks meja jika ditemukan, -1 jika tidak ada
func sequentialSearchMeja(T TabMeja, n int, kode int) int {
	var i, ketemu int

	ketemu = -1
	i = 0
	for i < n && ketemu == -1 {
		if T[i].kodeMeja == kode {
			ketemu = i
		}
		i = i + 1
	}
	return ketemu
}