package user_tubes

import (
	"fmt"
	menu_tubes "github.com/shadrairfn/app-tubes-alpro/app/appMenu"
	struct_tubes "github.com/shadrairfn/app-tubes-alpro/app/appStruct"
)

func IsInput(arr *struct_tubes.TabArr, arr2 *struct_tubes.TabArr2, banyakMahasiswa *int) {
	/* input Nama, NIK, Tanggal lahir, Umur, Asal kota, Alamat dari Mahasiswa yang didaftarkan */
	var ProgramStud int
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

	fmt.Println()

	menu_tubes.MenuJurusan()
	fmt.Print("Pilih Jurusan: ")
	fmt.Scan(&ProgramStud)
	switch ProgramStud {
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
	i += 1
	fmt.Println()
	*banyakMahasiswa = i
}

func IsEdit(arr *struct_tubes.TabArr, arr2 *struct_tubes.TabArr2, banyakMahasiswa *int) {
	/* Mengedit mahasiswa yang sudah terdaftar dengan memasukan Nama, NIK, Tanggal lahir, Umur, Asal kota, Alamat yang baru */
	var menuEdit string
	var edit, nomorIndukEdit, kotaTinggalEdit string
	var namaJalan1Edit, namajalan2Edit string
	var tanggalLahirEdit, bulanLahirEdit, tahunLahirEdit, umurMahasiswaEdit int

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
	}
}

func LihatStatusKelulusan(arr struct_tubes.TabArr, banyakMahasiswa int, arr3 struct_tubes.TabKonf, CariNama string) {
	for i := 0; i < banyakMahasiswa; i++ {
		if CariNama == arr[i].NamaMahasiswa1 {
		fmt.Printf(" Nama		: %s		\n", arr[i].NamaMahasiswa1)
		fmt.Printf(" Nomor Induk	: %s 			\n", arr[i].NomorInduk)
		fmt.Printf(" Tanggal Lahir	: %d-%d-%d 			\n", arr[i].TanggalLahir, arr[i].BulanLahir, arr[i].TahunLahir)
		fmt.Printf(" Umur		: %d 				\n", arr[i].UmurMahasiswa)
		fmt.Printf(" Asal Kota	: %s 			\n", arr[i].KotaTinggal)
		fmt.Printf(" Alamat		: %s %s 		\n", arr[i].NamaJalan1, arr[i].NamaJalan2)
		fmt.Printf(" Status Kelulusan: %t\n", arr3[i].Konfirmasi)
		}
	}
}