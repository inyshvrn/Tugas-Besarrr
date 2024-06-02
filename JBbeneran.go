package main

import "fmt"

const NMAX int = 1000

var t_pendapatan, t_modal int
var arr id_barang //array barang
var irr transaksi
var n, nirr int //ndata

type barang struct {
	nama                                      string
	stok, modal, harga, terjual               int
	kategori                                  string
	total_modal, total_pendapatan, keuntungan int
}
type Transaction struct {
	ID           string
	Customer     string
	Amount       int
	Date         string
	total_barang int
	cart         [NMAX]barang
}

type id_barang [NMAX]barang
type transaksi [NMAX]Transaction

func main() {
	sambutan()
	login()
}

func sambutan() {
	fmt.Println("=============================================")
	fmt.Println("          SELAMAT DATANG DI ALISITI       ")
	fmt.Println("    Aplikasi jual beli untuk pegawai toko    ")
	fmt.Println("=============================================")
}

func login() {
	var loginberhasil bool = false
	var username string
	var pwUser string
	var pwSystem string = "alisiti123"

	fmt.Println("Log in")
	fmt.Println("=============================================")

	for loginberhasil == false {
		fmt.Println("Masukkan Username: ")
		fmt.Scan(&username)

		fmt.Println("Masukkan Password: ")
		fmt.Scan(&pwUser)

		if pwUser != pwSystem {
			fmt.Println("Password anda tidak valid")
			fmt.Println("=============================================")
			login()
		} else {
			fmt.Println("Berhasil log in!")
			fmt.Println("=============================================")
			program()
		}
	}
}

func program() {
	fmt.Println("---------------------------------------------")
	fmt.Println("1. Tampilkan data barang")
	fmt.Println("2. Edit data barang")
	fmt.Println("3. Tampilkan data transaksi")
	fmt.Println("4. Edit data transaksi")
	fmt.Println("5. Cari barang")
	fmt.Println("6. Tampilkan top 5 penjualan terbanyak")
	fmt.Println("7. Tampilkan total pendapatan, total modal dan total keuntungan")
	fmt.Println("8. KELUAR APLIKASI")
	fmt.Println("---------------------------------------------")
	// fmt.Println()
	fmt.Print("Silakan memilih nomor: ")

	var nomor, no1, idx int
	var x string

	fmt.Scan(&nomor)
	if nomor == 1 {
		fmt.Println("1. Pilih kategori")
		fmt.Println("2. Semua barang")
		fmt.Print("Silakan memilih nomor: ")
		fmt.Scan(&no1)

		if no1 == 1 {
			s_kategori(arr, n)
			program()
		} else if no1 == 2 {
			cetak(arr, n)
			program()
		}

	} else if nomor == 2 {
		edit(&arr, &n)
	} else if nomor == 3 {
		cetakt(irr, nirr)
		program()
	} else if nomor == 4 {
		edittrans(&irr, &nirr, &arr, &n)
	} else if nomor == 5 {
		fmt.Println("")
		fmt.Print("Masukkan nama barang yang anda cari: ")
		fmt.Scan(&x)
		idx = binarysearch(arr, n, x)
		if idx == -1 {
			fmt.Println("Barang tidak ada")
		} else {
			fmt.Printf("%-15s %-15s %-10s \n", "NAMA BARANG", "KATEGORI", "HARGA JUAL")
			fmt.Printf("%-15s %-15s %-10d \n", arr[idx].nama, arr[idx].kategori, arr[idx].harga)
			fmt.Println("---------------------------------------------")
			fmt.Println("1. Tampilkan rincian")
			fmt.Println("2. Kembali")
			fmt.Print("Silakan memilih nomor: ")
			fmt.Scan(&no1)
			if no1 == 1 {
				cetak_rinci(arr, n, idx)
				program()
			} else {
				program()
			}
		}

	} else if nomor == 6 {
		uruttop(arr, n)
		program()
	} else if nomor == 7 {
		fmt.Println("TOTAL PENDAPATAN	:", t_pendapatan)
		fmt.Println("TOTAL MODAL		:", t_modal)
		fmt.Println("TOTAL KEUNTUNGAN	:", t_pendapatan-t_modal)
	} else if nomor == 8 {
		return
	}
}

func cetakt(irr transaksi, n int) {
	if n == 0 {
		fmt.Println("Belum ada transaksi")
	} else {
		fmt.Println("-------------------------------------------------------")
		fmt.Printf("%-5s %-10s %-10s %-10s %-10s \n", "ID", "CUSTOMER", "AMOUNT", "DATE", "TOTAL BARANG")
		fmt.Println("-------------------------------------------------------")
		for i := 0; i < n; i++ {
			fmt.Printf("%-5s %-10s %-10d %-10s %-10d \n", irr[i].ID, irr[i].Customer, irr[i].Amount, irr[i].Date, irr[i].total_barang)
		}
	}
}

func cetak(arr id_barang, n int) {
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("%-15s %-20s %-15s \n", "NAMA BARANG", "KATEGORI", "HARGA JUAL")
	fmt.Println("-------------------------------------------------------")

	for i := 0; i < n; i++ {
		fmt.Printf("%-15s %-20s %-15d\n", arr[i].nama, arr[i].kategori, arr[i].harga)
	}
	fmt.Println("=======================================================")
}

func input(arr *id_barang, n *int) {
	var baru, x int
	var i int = *n
	var a string

	fmt.Println("(Produk dapat ditambahkan sekaligus jika mempunyai kategori yang sama)")
	fmt.Print("Ada berapa jenis banyak barang yang akan dimasukkan? ")
	fmt.Scan(&baru)

	*n += baru

	fmt.Println("KATEGORI")
	fmt.Println("1. Makanan dan Minuman")
	fmt.Println("2. Pakaian")
	fmt.Println("3. Personal Care")
	fmt.Print("Pilih kategori: ")
	fmt.Scan(&x)

	for i < *n {
		if x == 1 {
			arr[i].kategori = "Makanan dan Minuman"
		} else if x == 2 {
			arr[i].kategori = "Pakaian"
		} else if x == 3 {
			arr[i].kategori = "Personal Care"
		}

		fmt.Println("Masukkan nama barang: ")
		fmt.Scan(&a)
		arr[i].nama = a

		fmt.Println("Masukkan jumlah barang yang akan ditambah: ")
		fmt.Scan(&arr[i].stok)

		fmt.Println("Masukkan modal barang: ")
		fmt.Scan(&arr[i].modal)

		fmt.Println("Masukkan harga barang: ")
		fmt.Scan(&arr[i].harga)

		i++
		fmt.Println("Barang berhasil ditambahkan")
		fmt.Println("--------------------------------")
	}

	//urutkan
	if *n > 1 {
		urutnama(arr, *n)
	}

	fmt.Println("1. Lanjutkan tambah barang")
	fmt.Println("2. Selesai")
	fmt.Scan(&x)
	if x == 1 {
		input(arr, n)
	} else {
		program()
	}
	fmt.Println("--------------------------------")
}

func cetak_kategori(arr id_barang, n int, kat string) {
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("%-15s %-15s %-10s \n", "NAMA BARANG", "KATEGORI", "HARGA JUAL")
	fmt.Println("-------------------------------------------------------")
	for i := 0; i < n; i++ {
		if arr[i].kategori == kat {
			fmt.Printf("%-15s %-15s %-10d \n", arr[i].nama, arr[i].kategori, arr[i].harga)
		}
	}
}

func cetak_rinci(arr id_barang, n int, idx int) {
	fmt.Println("Nama Barang	:", arr[idx].nama)
	fmt.Println("Kategori		:", arr[idx].kategori)
	fmt.Println("Modal			:", arr[idx].modal)
	fmt.Println("Harga Jual		:", arr[idx].harga)
	fmt.Println("Barang terjual	:", arr[idx].terjual, "pcs")
	fmt.Println("Stok barang	:", arr[idx].stok, "pcs")
}

func s_kategori(arr id_barang, n int) {
	var no int
	fmt.Println("-------------------------")
	fmt.Println("KATEGORI")
	fmt.Println("-------------------------")
	fmt.Println("1. Makanan dan Minuman")
	fmt.Println("2. Pakaian")
	fmt.Println("3. Personal Care")
	fmt.Println("-------------------------")
	// fmt.Println()
	fmt.Print("Silahkan pilih nomor kategori: ")

	fmt.Scan(&no)
	if no == 1 {
		target := "Makanan dan Minuman"
		cetak_kategori(arr, n, target)
	} else if no == 2 {
		target := "Pakaian"
		cetak_kategori(arr, n, target)
	} else if no == 3 {
		target := "Personal Care"
		cetak_kategori(arr, n, target)
	}
}

func edit(arr *id_barang, n *int) {
	var x, y, z int
	var target, a string
	fmt.Println("MENU EDIT BARANG")
	fmt.Println("-------------------")
	fmt.Println("1. Tambah barang")
	fmt.Println("2. Hapus barang")
	fmt.Println("3. Edit data barang")
	fmt.Println("4. Kembali")
	fmt.Println("-------------------")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&x)
	if x == 3 {
		fmt.Println("Masukkan nama barang yang akan diedit:")
		fmt.Scan(&target)
		idx := search(*arr, *n, target)
		if idx == -1 {
			fmt.Println("Barang tidak ditemukan")
			edit(arr, n)
		} else {
			fmt.Println("PILIHAN EDIT")
			fmt.Println("1. Edit kategori barang")
			fmt.Println("2. Edit nama barang")
			fmt.Println("3. Tambah/kurangi stok barang")
			fmt.Println("4. Edit modal barang")
			fmt.Println("5. Edit harga barang")
			fmt.Print("Pilih untuk mengedit: ")
			fmt.Scan(&x)
			if x == 1 {
				fmt.Println()
				fmt.Println("-------------------------")
				fmt.Println("        KATEGORI")
				fmt.Println("-------------------------")
				fmt.Println("1. Makanan dan Minuman")
				fmt.Println("2. Pakaian")
				fmt.Println("3. Personal Care")
				fmt.Println("-------------------------")
				fmt.Println()
				fmt.Print("Pilih nomor kategori: ")
				fmt.Scan(&z)
				if z == 1 {
					arr[idx].kategori = "Makanan dan Minuman"
				} else if z == 2 {
					arr[idx].kategori = "Pakaian"
				} else if z == 3 {
					arr[idx].kategori = "Personal Care"
				}
			} else if x == 2 {
				fmt.Println("Masukkan nama barang: ")
				fmt.Scan(&a)
				arr[idx].nama = a
			} else if x == 3 {
				fmt.Println("Masukkan stok barang: ")
				fmt.Scan(&y)
				arr[idx].stok += y
			} else if x == 4 {
				fmt.Println("Masukkan modal barang: ")
				fmt.Scan(&y)
				arr[idx].modal = y
			} else if x == 5 {
				fmt.Println("Masukkan harga barang: ")
				fmt.Scan(&y)
				arr[idx].harga = y
			}
			fmt.Println("Data barang berhasil diedit")
		}
	} else if x == 2 {
		fmt.Println("Masukkan nama barang yang akan dihapus:")
		fmt.Scan(&target)
		idx := search(*arr, *n, target)
		if idx == -1 {
			fmt.Println("Barang tidak ditemukan")
			edit(arr, n)
		} else {
			hapus(arr, n, idx)
			fmt.Println("Data barang berhasil dihapus")
			edit(arr, n)
		}

	} else if x == 1 {
		input(arr, n)
		edit(arr, n)
	} else if x == 4 {
		program()
	}
}

func search(arr id_barang, n int, target string) int {
	var idx int = -1
	for i := 0; i < n; i++ {
		if arr[i].nama == target {
			idx = i
		}
	}
	return idx
}

func binarysearch(arr id_barang, n int, target string) int {
	var r, l int = n - 1, 0
	var m int

	for r > l {
		m = l + (r-l)/2
		if arr[m].nama == target {
			return m
		}
		if arr[m].nama < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

func hapus(arr *id_barang, n *int, idx int) {
	for i := idx; i < *n; i++ {
		arr[i] = arr[i+1]
	}
	*n--
}

func urutnama(arr *id_barang, n int) {
	//selection sort
	//btw, ini belum aku cek
	var max, i, j int
	var temp barang

	for i < n {
		max = i
		j = i
		for j < n { // cari maksimal
			if arr[j].nama < arr[max].nama {
				max = j
			}
			j++
		}
		// tukar posisi max
		temp = arr[i]
		arr[i] = arr[max]
		arr[max] = temp
		i++
	}
}

func shiftright(arr *id_barang, n *int, x int) {
	var i int
	*n++
	i = *n - 1
	for i > x {
		arr[i] = arr[i-1]
		i--
	}
}

func sisipkan(arr *id_barang, n int, temp barang) {
	idx := binarysearch(*arr, n, temp.nama)
	shiftright(arr, &n, idx)
	arr[idx] = temp
}

func edittrans(irr *transaksi, nirr *int, arr *id_barang, n *int) {
	var x int
	var target string
	fmt.Println("MENU EDIT TRANSAKSI")
	fmt.Println("-------------------")
	fmt.Println("1. Tambah transaksi")
	fmt.Println("2. Hapus data transaksi")
	fmt.Println("3. Kembali")
	fmt.Println("-------------------")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&x)
	if x == 1 {
		inputT(irr, nirr, arr, n)
	} else if x == 2 {
		fmt.Println("Masukkan id transaksi yang akan dihapus:")
		fmt.Scan(&target)
		idx := searchT(*irr, *nirr, target)
		if idx == -1 {
			fmt.Println("Id transaksi tidak ditemukan")
			edittrans(irr, nirr, arr, n)
		} else {
			// hapusT(irr, nirr, arr, n, idx)
			fmt.Println("Data transaksi berhasil dihapus")
			edittrans(irr, nirr, arr, n)
		}
	} else if x == 3 {
		program()
	}
}

func searchT(irr transaksi, nir int, target string) int {
	var idx int = -1
	for i := 0; i < nir; i++ {
		if irr[i].ID == target {
			idx = i
		}
	}
	return idx
}

func binarysearchid(irr transaksi, n int, target string) int {
	var r, l int = n - 1, 0
	var m int

	for r > l {
		m = l + (r-l)/2
		if irr[m].ID == target {
			return m
		}
		if irr[m].ID < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

func hapusT(irr *transaksi, nirr *int, arr *id_barang, n *int, idx int) {
	for j := 0; j < irr[idx].total_barang; j++ {
		idxbarang := binarysearchid(*irr, *nirr, irr[idx].cart[j].nama)

		arr[idxbarang].terjual -= irr[idx].cart[j].terjual
		arr[idxbarang].stok += arr[idxbarang].terjual
		arr[idxbarang].total_modal -= irr[idx].cart[j].modal
		arr[idxbarang].total_pendapatan -= irr[idx].cart[j].terjual * arr[idxbarang].harga
		arr[idxbarang].keuntungan -= arr[idxbarang].total_pendapatan - arr[idxbarang].total_modal

		t_modal -= irr[idx].cart[j].modal
		t_pendapatan -= irr[idx].cart[j].terjual * arr[idxbarang].harga
	}

	for i := idx; i < *n; i++ {
		irr[i] = irr[i+1]
	}
	*nirr--
}

func inputT(irr *transaksi, nirr *int, arr *id_barang, n *int) {

	var i int = *nirr
	var a, b string
	var x int

	*nirr += 1

	fmt.Println("Masukkan ID transaksi: ")
	fmt.Scan(&b)
	irr[i].ID = b

	fmt.Println("Masukkan nama Customer: ")
	fmt.Scan(&a)
	irr[i].Customer = a

	fmt.Println("Masukkan tanggal transaksi: ")
	fmt.Scan(&irr[i].Date)

	fmt.Println("Masukkan total barang yang dibeli: ")
	fmt.Scan(&x)
	irr[i].total_barang = x

	fmt.Println("MASUKKAN BARANG APA SAJA YANG DIBELI ")
	for j := 0; j < irr[i].total_barang; j++ {
		fmt.Println("Masukkan nama barang yang dibeli: ")
		fmt.Scan(&irr[i].cart[j].nama)

		idxbarang := binarysearch(*arr, *n, irr[i].cart[j].nama)

		fmt.Println("Masukkan jumlah barang yang dibeli: ")
		fmt.Scan(&irr[i].cart[j].terjual)
		if irr[i].cart[j].terjual > arr[idxbarang].stok {
			irr[i].cart[j].terjual = arr[idxbarang].stok
		}

		irr[i].cart[j].modal = arr[idxbarang].modal * irr[i].cart[j].terjual
		arr[idxbarang].terjual += irr[i].cart[j].terjual
		arr[idxbarang].stok -= arr[idxbarang].terjual
		arr[idxbarang].total_modal += irr[i].cart[j].modal
		arr[idxbarang].total_pendapatan += irr[i].cart[j].terjual * arr[idxbarang].harga
		arr[idxbarang].keuntungan += arr[idxbarang].total_pendapatan - arr[idxbarang].total_modal
		irr[i].Amount += arr[idxbarang].harga * irr[i].cart[j].terjual

		t_modal += irr[i].cart[j].modal
		t_pendapatan += irr[i].cart[j].terjual * arr[idxbarang].harga

		fmt.Println("-------------------------------")
		fmt.Println("Barang ditambahkan ke keranjang")
		fmt.Println("-------------------------------")
	}

	//urutkan
	urutid(irr, *nirr)

	fmt.Println("Data transaksi berhasil ditambahkan")
	fmt.Println("--------------------------------")
	fmt.Println("1. Lanjutkan tambah data transaksi")
	fmt.Println("2. Selesai")
	fmt.Scan(&x)
	if x == 1 {
		inputT(irr, nirr, arr, n)
	} else {
		program()
	}
}

func urutid(irr *transaksi, nirr int) {
	//selection sort
	var max, i, j int
	var temp Transaction

	for i < nirr {
		max = i
		j = i
		for j < nirr { // cari maksimal
			if irr[j].ID < irr[max].ID {
				max = j
			}
			j++
		}
		// tukar posisi max
		temp = irr[i]
		irr[i] = irr[max]
		irr[max] = temp
		i++
	}
}

func uruttop(arr id_barang, n int) {
	//ini insertion sort
	var array id_barang = arr
	var temp barang
	var i, j int
	for i < n {
		j = i
		temp = array[j]
		for j > 0 && temp.terjual > array[j].terjual {
			array[j] = array[j-1]
			j--
		}
		array[j] = temp
		i++
	}

	fmt.Println("TOP 5 BEST SELLER")
	for k := 0; k < 5; k++ {
		fmt.Printf("%d. %-10s dengan produk terjual sebanyak: %d\n", k+1, array[k].nama, array[k].terjual)
	}
}
