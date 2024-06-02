package admin_tubes

import (
	"fmt"

	menu_tubes "github.com/shadrairfn/app-tubes-alpro/app/appMenu"
	sorting_tubes "github.com/shadrairfn/app-tubes-alpro/app/appSorting"
	struct_tubes "github.com/shadrairfn/app-tubes-alpro/app/appStruct"
)

func CekPendaftar(arr struct_tubes.TabArr, arr2 struct_tubes.TabArr2, banyakMahasiswa int, arr3 struct_tubes.TabKonf) {
	var sorting int
	var sortingJurusan int
	for i := 0; i < banyakMahasiswa; i++ {
		fmt.Println("==================================================")
		fmt.Printf(" Mahasiswa ke %d: 				\n", i+1)
		fmt.Printf(" Nama		: %s 		\n", arr[i].NamaMahasiswa)
		fmt.Printf(" Nomor Induk	: %s 			\n", arr[i].NomorInduk)
		fmt.Printf(" Tanggal Lahir	: %d-%d-%d 			\n", arr[i].TanggalLahir, arr[i].BulanLahir, arr[i].TahunLahir)
		fmt.Printf(" Umur		: %d 				\n", arr[i].UmurMahasiswa)
		fmt.Printf(" Asal Kota	: %s 			\n", arr[i].KotaTinggal)
		fmt.Printf(" Alamat		: %s 		\n", arr[i].NamaJalan)
		fmt.Printf(" Program Studi	: %s %s\n", arr2[i].Program, arr2[i].Jurusan)
		fmt.Printf(" Nilai Tes 	: %0.1f\n", arr2[i].NilaiTes)
		fmt.Println("==================================================")
		fmt.Println()
	}
	menu_tubes.MenuSortingAdmin()
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&sorting)
	fmt.Println()
	switch sorting {
	case 1:	
		sorting_tubes.NilaiAscending(&arr, &arr2, banyakMahasiswa)
		
		fmt.Println("==========================")
		for k := 0; k < banyakMahasiswa; k++ {
			fmt.Printf("Nama	: %s\nNo.Induk: %s\nNilai	: %0.2f\n", arr[k].NamaMahasiswa, arr[k].NomorInduk, arr2[k].NilaiTes)
			fmt.Println()
		}
		fmt.Println("==========================")

	case 2:
		sorting_tubes.NilaiDescending(&arr, &arr2, banyakMahasiswa)

		fmt.Println("==========================")
		for k := 0; k < banyakMahasiswa; k++ {
			fmt.Printf("Nama	: %s\nNo.Induk: %s\nNilai	: %0.2f\n", arr[k].NamaMahasiswa, arr[k].NomorInduk, arr2[k].NilaiTes)
			fmt.Println()
		}
		fmt.Println("==========================")

	case 3:
		menu_tubes.SortingJurusan()
		fmt.Print("Pilih Menu Program dan Jurusan: ")
		fmt.Scan(&sortingJurusan)
		switch sortingJurusan {
		case 1:
			fmt.Println("=========================")
			for k := 0; k < banyakMahasiswa; k++ {
				if arr2[k].Program == "S1" && arr2[k].Jurusan == "Informatika" {
					fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: S1 Informatika\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
					fmt.Println()
				}
			}
			fmt.Println("=========================")
		case 2:
			fmt.Println("=========================")
			for k := 0; k < banyakMahasiswa; k++ {
				if arr2[k].Program == "S1" && arr2[k].Jurusan == "Sistem Informasi" {
					fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: S1 Sistem Informasi\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
					fmt.Println()
				}
			}
			fmt.Println("=========================")
		case 3:
			fmt.Println("=========================")
			for k := 0; k < banyakMahasiswa; k++ {
				if arr2[k].Program == "D4" && arr2[k].Jurusan == "Informatika" {
					fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: D4 Informatika\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
					fmt.Println()
				}
			}
			fmt.Println("=========================")
		case 4:
			fmt.Println("=========================")
			for k := 0; k < banyakMahasiswa; k++ {
				if arr2[k].Program == "D4" && arr2[k].Jurusan == "Sistem Informasi" {
					fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: D4 Sistem Informasi\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
					fmt.Println()
				}
			}
			fmt.Println("=========================")
		case 5:
			fmt.Println("=========================")
			for k := 0; k < banyakMahasiswa; k++ {
				if arr2[k].Program == "D3" && arr2[k].Jurusan == "Informatika" {
					fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: D3 Informatika\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
					fmt.Println()
				}
			}
			fmt.Println("=========================")
		case 6:
			fmt.Println("=========================")
			for k := 0; k < banyakMahasiswa; k++ {
				if arr2[k].Program == "D3" && arr2[k].Jurusan == "Sistem Informasi" {
					fmt.Printf("Nama		: %s\nNomor Induk	: %s\nProgram Studi	: D3 Sistem Informasi\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
					fmt.Println()
				}
			}
			fmt.Println("=========================")
		}
	case 4:
		fmt.Println("=========================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr3[k].Konfirmasi == true {
				fmt.Printf("Nama	: %s\nNo.Induk: %s\nStatus	: Lulus\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println()
			}
		}
		fmt.Println("=========================")
	case 5:
		fmt.Println("=========================")
		for k := 0; k < banyakMahasiswa; k++ {
			if arr3[k].Konfirmasi == false {
				fmt.Printf("Nama	: %s\nNo.Induk: %s\nStatus	: Tidak Lulus\n", arr[k].NamaMahasiswa, arr[k].NomorInduk)
				fmt.Println()
			}
		}
		fmt.Println("=========================")
	}
}

func BeriKonfirmasi(arr struct_tubes.TabArr, arr2 struct_tubes.TabArr2, banyakMahasiswa int, arr3 *struct_tubes.TabKonf) {
	var kelulusan string
	for i := 0; i < banyakMahasiswa; i++ {
		fmt.Printf(" Mahasiswa ke %d: 				\n", i+1)
		fmt.Printf(" Nama		: %s		\n", arr[i].NamaMahasiswa)
		fmt.Printf(" Nomor Induk	: %s 			\n", arr[i].NomorInduk)
		fmt.Printf(" Tanggal Lahir	: %d-%d-%d 			\n", arr[i].TanggalLahir, arr[i].BulanLahir, arr[i].TahunLahir)
		fmt.Printf(" Umur		: %d 				\n", arr[i].UmurMahasiswa)
		fmt.Printf(" Asal Kota	: %s 			\n", arr[i].KotaTinggal)
		fmt.Printf(" Alamat		: %s 		\n", arr[i].NamaJalan)
		fmt.Printf(" Program Studi	: %s %s\n", arr2[i].Program, arr2[i].Jurusan)
		fmt.Printf(" Nilai Tes 	: %0.1f\n", arr2[i].NilaiTes)
		fmt.Println()
		fmt.Print(" Status Kelulusana (ya/tidak): ")
		fmt.Scan(&kelulusan)
		if kelulusan == "ya" {
			arr3[i].Konfirmasi = true
		} else {
			arr3[i].Konfirmasi = false
		}
	}
}

func EditMahasiswa(arr *struct_tubes.TabArr, arr2 *struct_tubes.TabArr2, banyakMahasiswa *int, arr3 *struct_tubes.TabKonf) {
	/* Mengedit mahasiswa yang sudah terdaftar dengan memasukan Nama, NIK, Tanggal lahir, Umur, Asal kota, Alamat yang baru */
	var menuEdit string
	var edit, nomorIndukEdit, kotaTinggalEdit string
	var namaJalan1Edit, namajalan2Edit string
	var tanggalLahirEdit, bulanLahirEdit, tahunLahirEdit, umurMahasiswaEdit int
	var noIndukHapus string
	var kelulusan string
	var programStud int
	var nilaiTesEdit float64

	fmt.Print("Menu: (tambah / edit / hapus): ")
	fmt.Scan(&menuEdit)

	switch menuEdit {
	case "tambah":
		i := *banyakMahasiswa
		fmt.Printf("Nama			     : ")
		fmt.Scan(&arr[i].NamaMahasiswa)

		fmt.Printf("Nomor Induk Kependudukan     : ")
		fmt.Scan(&arr[i].NomorInduk)

		fmt.Printf("Tanggal Lahir (Ex : 8 9 2004): ")
		fmt.Scan(&arr[i].TanggalLahir, &arr[i].BulanLahir, &arr[i].TahunLahir)

		fmt.Print("Umur			     : ")
		fmt.Scan(&arr[i].UmurMahasiswa)

		fmt.Printf("Asal Kota (Ex : Bandung)     : ")
		fmt.Scan(&arr[i].KotaTinggal)

		fmt.Printf("Alamat Rumah 		   : ")
		fmt.Scan(&arr[i].NamaJalan)

		menu_tubes.MenuJurusan()
		fmt.Print("Pilih Jurusan: ")
		fmt.Scan(&programStud)
		switch programStud {
		case 1:
			arr2[i].Program = "S1"
			arr2[i].Jurusan = "Informatika"
		case 2:
			arr2[i].Program = "S1"
			arr2[i].Jurusan = "Sistem Informasi"
		case 3:
			arr2[i].Program = "D4"
			arr2[i].Jurusan = "Informatika"
		case 4:
			arr2[i].Program = "D4"
			arr2[i].Jurusan = "Sistem Informasi"
		case 5:
			arr2[i].Program = "D3"
			arr2[i].Jurusan = "Informatika"
		case 6:
			arr2[i].Program = "D3"
			arr2[i].Jurusan = "Sistem Informasi"
		}

		fmt.Print("Masukan nilai anda: ")
		fmt.Scan(&arr2[i].NilaiTes)
		fmt.Println()
		fmt.Print(" Status Kelulusana (ya/tidak): ")
		fmt.Scan(&kelulusan)
		if kelulusan == "ya" {
			arr3[i].Konfirmasi = true
		} else {
			arr3[i].Konfirmasi = false
		}

		i += 1
		fmt.Println()
		*banyakMahasiswa = i
	case "edit":
		fmt.Print("Nama mahasiswa yang akan diganti:  ")
		fmt.Scan(&edit)
		fmt.Printf("Masukan data baru\banyakMahasiswa")
		for i := 0; i < *banyakMahasiswa; i++ {
			if arr[i].NamaMahasiswa == edit {
				fmt.Printf("Nama			     : ")
				fmt.Scan(&edit)

				fmt.Printf("Nomor Induk Kependudukan     : ")
				fmt.Scan(&nomorIndukEdit)

				fmt.Print("Tanggal Lahir: ")
				fmt.Scan(&tanggalLahirEdit, &bulanLahirEdit, &tahunLahirEdit)

				fmt.Print("Umur			     : ")
				fmt.Scan(&umurMahasiswaEdit)

				fmt.Printf("Asal Kota (Ex : Bandung)     : ")
				fmt.Scan(&kotaTinggalEdit)

				fmt.Printf("Alamat Rumah (2 suku kata)   : ")
				fmt.Scan(&namaJalan1Edit, &namajalan2Edit)

				menu_tubes.MenuJurusan()
				fmt.Print("Pilih Jurusan: ")
				fmt.Scan(&programStud)
				switch programStud {
					case 1:
						arr2[i].Program = "S1"
						arr2[i].Jurusan = "Informatika"
					case 2:
						arr2[i].Program = "S1"
						arr2[i].Jurusan = "Sistem Informasi"
					case 3:
						arr2[i].Program = "D4"
						arr2[i].Jurusan = "Informatika"
					case 4:
						arr2[i].Program = "D4"
						arr2[i].Jurusan = "Sistem Informasi"
					case 5:
						arr2[i].Program = "D3"
						arr2[i].Jurusan = "Informatika"
					case 6:
						arr2[i].Program = "D3"
						arr2[i].Jurusan = "Sistem Informasi"
					}

					fmt.Print("Masukan nilai anda: ")
					fmt.Scan(&nilaiTesEdit)
					fmt.Print(" Status Kelulusana (ya/tidak): ")
					fmt.Scan(&kelulusan)
						if kelulusan == "ya" {
							arr3[i].Konfirmasi = true
						} else {
							arr3[i].Konfirmasi = false
						}
					fmt.Println()
					arr[i].NamaMahasiswa = edit
					arr[i].NomorInduk = nomorIndukEdit
					arr[i].KotaTinggal = kotaTinggalEdit
					arr[i].TanggalLahir = tahunLahirEdit
					arr[i].BulanLahir = bulanLahirEdit
					arr[i].TahunLahir = tahunLahirEdit
					arr[i].NamaJalan = namaJalan1Edit
					arr[i].UmurMahasiswa = umurMahasiswaEdit
					arr2[i].NilaiTes = nilaiTesEdit
			}	
		}
		case "hapus":
		fmt.Print("Ketikan nomor induk yang akan dihapus: ")
		fmt.Scan(&noIndukHapus)
		hasilFind := FindDelete(*arr, *banyakMahasiswa, noIndukHapus) // Memanggil fungsi find delete untuk mencari index data yang akan dihapus //
		
		if hasilFind != -1 {
			for i := hasilFind; i < *banyakMahasiswa; i++ {
				arr[i] = arr[i+1]
				arr2[i] = arr2[i+1]
			}
			fmt.Println("Penghapusan berhasil !")
			*banyakMahasiswa = *banyakMahasiswa - 1
			fmt.Println("Data hasil penghapusan: ")
			for i := 0; i < *banyakMahasiswa; i++ {
			fmt.Println("==================================================")
			fmt.Printf(" Mahasiswa ke %d: 				\n", i+1)
			fmt.Printf(" Nama		: %s		\n", arr[i].NamaMahasiswa)
			fmt.Printf(" Nomor Induk	: %s 			\n", arr[i].NomorInduk)
			fmt.Printf(" Tanggal Lahir	: %d-%d-%d 			\n", arr[i].TanggalLahir, arr[i].BulanLahir, arr[i].TahunLahir)
			fmt.Printf(" Umur		: %d 				\n", arr[i].UmurMahasiswa)
			fmt.Printf(" Asal Kota	: %s 			\n", arr[i].KotaTinggal)
			fmt.Printf(" Alamat		: %s  		\n", arr[i].NamaJalan)
			fmt.Printf(" Program Studi	: %s %s\n", arr2[i].Program, arr2[i].Jurusan)
			fmt.Printf(" Nilai Tes 	: %0.1f\n", arr2[i].NilaiTes)
			fmt.Println("==================================================")
			fmt.Println()
			}
		} else {
			fmt.Println("Penghapusan gagal karena tidak ditemukan data yang cocok")
		}
	}	
}

func FindDelete(arr struct_tubes.TabArr, banyakMahasiswa int, x string) int {
	// { Fungsi mengembalikan index data yang akan dihapus } //
	idx := -1
	i := 0
	for i < banyakMahasiswa && idx == -1 {
		if arr[i].NomorInduk == x {
			idx = i
		}
		i += 1
	}
	return idx
}