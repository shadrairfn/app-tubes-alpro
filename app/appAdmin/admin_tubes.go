package adminTubes

import (
	"fmt"
	menu_tubes "github.com/shadrairfn/app-tubes-alpro/app/appMenu"
	struct_tubes "github.com/shadrairfn/app-tubes-alpro/app/appStruct"
)

func CekPendaftar(arr struct_tubes.TabArr, arr2 struct_tubes.TabArr2, banyakMahasiswa int) {
	var sorting int
	var sortingNilai, sortingJurusan int
	for i := 0; i < banyakMahasiswa; i++ {
		fmt.Println("==================================================")
		fmt.Printf(" Mahasiswa ke %d: 				\n", i+1)
		fmt.Printf(" Nama		: %s 		\n", arr[i].NamaMahasiswa1)
		fmt.Printf(" Nomor Induk	: %s 			\n", arr[i].NomorInduk)
		fmt.Printf(" Tanggal Lahir	: %d-%d-%d 			\n", arr[i].TanggalLahir, arr[i].BulanLahir, arr[i].TahunLahir)
		fmt.Printf(" Umur		: %d 				\n", arr[i].UmurMahasiswa)
		fmt.Printf(" Asal Kota	: %s 			\n", arr[i].KotaTinggal)
		fmt.Printf(" Alamat		: %s %s 		\n", arr[i].NamaJalan1, arr[i].NamaJalan2)
		fmt.Printf(" Program Studi	: %s %s\n", arr2[i].Program, arr2[i].Jurusan)
		fmt.Printf(" Nilai Tes 	: %0.1f\n", arr2[i].NilaiTes)
		fmt.Println("==================================================")
		fmt.Println()
	}
	menu_tubes.MenuSortingAdmin()
	var j int
	var temp float64
	var temp1, temp2 string
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&sorting)
	switch sorting {
	case 1:
		menu_tubes.SortingNilaiTes()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&sortingNilai)
		switch sortingNilai {
		case 1:			
			for pass := 1; pass < banyakMahasiswa; pass++ {
				temp := arr2[pass].NilaiTes
				temp1 := arr[pass].NamaMahasiswa1
				temp2 := arr[pass].NomorInduk
			
				j := pass - 1
				for j >= 0 && arr2[j].NilaiTes > temp {
					arr2[j+1].NilaiTes = arr2[j].NilaiTes
					arr[j+1].NamaMahasiswa1 = arr[j].NamaMahasiswa1
					arr[j+1].NomorInduk = arr[j].NomorInduk
					j--
				}
				arr2[j+1].NilaiTes = temp
				arr[j+1].NamaMahasiswa1 = temp1
				arr[j+1].NomorInduk = temp2
			}			

			for k := 0; k < banyakMahasiswa; k++ {
				fmt.Printf("%s || %s = %0.2f\n", arr[k].NamaMahasiswa1, arr[k].NomorInduk, arr2[k].NilaiTes)
			}

		case 2:
			for pass := 0; pass < banyakMahasiswa; pass++ {
				j = pass
				temp = arr2[j+1].NilaiTes
				temp1 = arr[j+1].NamaMahasiswa1
				temp2 = arr[j+1].NomorInduk
				for j >= 0 && temp > arr2[j].NilaiTes {
					arr[j+1] = arr[j]
					arr[j+1] = arr[j]
					arr2[j+1] = arr2[j]
					j -= 1
				}
				arr2[j+1].NilaiTes = temp
				arr[j+1].NamaMahasiswa1 = temp1
				arr[j+1].NomorInduk = temp2
			}

			for k := 0; k < banyakMahasiswa; k++ {
				fmt.Printf("%s || %s = %0.2f\n", arr[k].NamaMahasiswa1, arr[k].NomorInduk, arr2[k].NilaiTes)
			}
		}
	case 2:
		menu_tubes.SortingJurusan()
		fmt.Print("Pilih Menu Program dan Jurusan: ")
		fmt.Scan(&sortingJurusan)
		switch sortingJurusan {
		case 1:
			for k := 0; k < banyakMahasiswa; k++ {
				if arr2[k].Program == "S1" && arr2[k].Jurusan == "Informatika" {
					fmt.Printf("%s %s = S1 Informatika\n", arr[k].NamaMahasiswa1, arr[k].NomorInduk)
				}
			}
		case 2:
			for k := 0; k < banyakMahasiswa; k++ {
				if arr2[k].Program == "S1" && arr2[k].Jurusan == "Sistem Informasi" {
					fmt.Printf("%s %s = S1 Sistem Informasi\n", arr[k].NamaMahasiswa1, arr[k].NomorInduk)
				}
			}
		case 3:
			for k := 0; k < banyakMahasiswa; k++ {
				if arr2[k].Program == "D4" && arr2[k].Jurusan == "Informatika" {
					fmt.Printf("%s %s = D4 Informatika\n", arr[k].NamaMahasiswa1, arr[k].NomorInduk)
				}
			}
		case 4:
			for k := 0; k < banyakMahasiswa; k++ {
				if arr2[k].Program == "D4" && arr2[k].Jurusan == "Sistem Informasi" {
					fmt.Printf("%s %s = D4 Sistem Informasi\n", arr[k].NamaMahasiswa1, arr[k].NomorInduk)
				}
			}
		case 5:
			for k := 0; k < banyakMahasiswa; k++ {
				if arr2[k].Program == "D3" && arr2[k].Jurusan == "Informatika" {
					fmt.Printf("%s %s = D3 Informatika\n", arr[k].NamaMahasiswa1, arr[k].NomorInduk)
				}
			}
		case 6:
			for k := 0; k < banyakMahasiswa; k++ {
				if arr2[k].Program == "D3" && arr2[k].Jurusan == "Sistem Informasi" {
					fmt.Printf("%s %s = D3 Sistem Informasi\n", arr[k].NamaMahasiswa1, arr[k].NomorInduk)
				}
			}
		}
	}
}

func BeriKonfirmasi(arr struct_tubes.TabArr, arr2 struct_tubes.TabArr2, banyakMahasiswa int, arr3 *struct_tubes.TabKonf) {
	var kelulusan string
	for i := 0; i < banyakMahasiswa; i++ {
		fmt.Printf(" Mahasiswa ke %d: 				\n", i+1)
		fmt.Printf(" Nama		: %s		\n", arr[i].NamaMahasiswa1)
		fmt.Printf(" Nomor Induk	: %s 			\n", arr[i].NomorInduk)
		fmt.Printf(" Tanggal Lahir	: %d-%d-%d 			\n", arr[i].TanggalLahir, arr[i].BulanLahir, arr[i].TahunLahir)
		fmt.Printf(" Umur		: %d 				\n", arr[i].UmurMahasiswa)
		fmt.Printf(" Asal Kota	: %s 			\n", arr[i].KotaTinggal)
		fmt.Printf(" Alamat		: %s %s 		\n", arr[i].NamaJalan1, arr[i].NamaJalan2)
		fmt.Printf(" Program Studi	: %s %s\n", arr2[i].Program, arr2[i].Jurusan)
		fmt.Printf(" Nilai Tes 	: %0.1f\n", arr2[i].NilaiTes)
		fmt.Print(" Status Kelulusana (ya/tidak): ")
		fmt.Scan(&kelulusan)
		if kelulusan == "ya" {
			arr3[i].Konfirmasi = true
		} else {
			arr3[i].Konfirmasi = false
		}
	}
}

func EditMahasiswa(arr *struct_tubes.TabArr, arr2 *struct_tubes.TabArr2, banyakMahasiswa *int) {
	/* Mengedit mahasiswa yang sudah terdaftar dengan memasukan Nama, NIK, Tanggal lahir, Umur, Asal kota, Alamat yang baru */
	var menuEdit string
	var edit, nomorIndukEdit, kotaTinggalEdit string
	var namaJalan1Edit, namajalan2Edit string
	var tanggalLahirEdit, bulanLahirEdit, tahunLahirEdit, umurMahasiswaEdit int
	var noIndukHapus string

	var programStud int
	var nilaiTesEdit float64

	fmt.Print("Menu: (tambah / edit / hapus): ")
	fmt.Scan(&menuEdit)

	switch menuEdit {
	case "tambah":
		i := *banyakMahasiswa
		fmt.Printf("Nama			     : ")
		fmt.Scan(&arr[i].NamaMahasiswa1)

		fmt.Printf("Nomor Induk Kependudukan     : ")
		fmt.Scan(&arr[i].NomorInduk)

		fmt.Printf("Tanggal Lahir (Ex : 8 9 2004): ")
		fmt.Scan(&arr[i].TanggalLahir, &arr[i].BulanLahir, &arr[i].TahunLahir)

		fmt.Print("Umur			     : ")
		fmt.Scan(&arr[i].UmurMahasiswa)

		fmt.Printf("Asal Kota (Ex : Bandung)     : ")
		fmt.Scan(&arr[i].KotaTinggal)

		fmt.Printf("Alamat Rumah (2 suku kata)   : ")
		fmt.Scan(&arr[i].NamaJalan1, &arr[i].NamaJalan2)

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

		i += 1
		fmt.Println()
		*banyakMahasiswa = i
	case "edit":
		fmt.Print("Nama mahasiswa yang akan diganti:  ")
		fmt.Scan(&edit)
		fmt.Printf("Masukan data baru\banyakMahasiswa")
		for i := 0; i < *banyakMahasiswa; i++ {
			if arr[i].NamaMahasiswa1 == edit {
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
					fmt.Println()
					arr[i].NamaMahasiswa1 = edit
					arr[i].NomorInduk = nomorIndukEdit
					arr[i].KotaTinggal = kotaTinggalEdit
					arr[i].TanggalLahir = tahunLahirEdit
					arr[i].BulanLahir = bulanLahirEdit
					arr[i].TahunLahir = tahunLahirEdit
					arr[i].NamaJalan1 = namaJalan1Edit
					arr[i].NamaJalan2 = namajalan2Edit
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
			fmt.Printf(" Mahasiswa ke %d: 				\banyakMahasiswa", i+1)
			fmt.Printf(" Nama		: %s		\banyakMahasiswa", arr[i].NamaMahasiswa1)
			fmt.Printf(" Nomor Induk	: %s 			\banyakMahasiswa", arr[i].NomorInduk)
			fmt.Printf(" Tanggal Lahir	: %d-%d-%d 			\banyakMahasiswa", arr[i].TanggalLahir, arr[i].BulanLahir, arr[i].TahunLahir)
			fmt.Printf(" Umur		: %d 				\banyakMahasiswa", arr[i].UmurMahasiswa)
			fmt.Printf(" Asal Kota	: %s 			\banyakMahasiswa", arr[i].KotaTinggal)
			fmt.Printf(" Alamat		: %s %s 		\banyakMahasiswa", arr[i].NamaJalan1, arr[i].NamaJalan2)
			fmt.Printf(" Program Studi	: %s %s\banyakMahasiswa", arr2[i].Program, arr2[i].Jurusan)
			fmt.Printf(" Nilai Tes 	: %0.1f\banyakMahasiswa", arr2[i].NilaiTes)
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