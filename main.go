package main

import (
	"fmt"

	admin_tubes "github.com/shadrairfn/app-tubes-alpro/app/appAdmin"
	menu_tubes "github.com/shadrairfn/app-tubes-alpro/app/appMenu"
	struct_tubes "github.com/shadrairfn/app-tubes-alpro/app/appStruct"
	user_tubes "github.com/shadrairfn/app-tubes-alpro/app/appUser"
)
func main() {
	var arr struct_tubes.TabArr
	var arr2 struct_tubes.TabArr2
	var arr3 struct_tubes.TabKonf
	var banyakMahasiswa int
	var chooseMenu, chooseMenuAdmin, chooseMenuUser int

	banyakMahasiswa = 0
	for {
		menu_tubes.MenuAdminUser()
		fmt.Scan(&chooseMenu)
		if chooseMenu > 2 {
			break
		}
		switch chooseMenu {
		case 1:
			menu_tubes.MenuAdmin()
			fmt.Scan(&chooseMenuAdmin)
			switch chooseMenuAdmin {
			case 1:
				admin_tubes.CekPendaftar(arr, arr2, banyakMahasiswa)
			case 2:
				admin_tubes.BeriKonfirmasi(arr, arr2, banyakMahasiswa, &arr3)
			}
		case 2:
			menu_tubes.MenuUser()
			fmt.Scan(&chooseMenuUser)
			switch chooseMenuUser {
			case 1:
				user_tubes.IsInput(&arr, &arr2, &banyakMahasiswa)
			case 2:
				user_tubes.IsEdit(&arr, &arr2, &banyakMahasiswa)
			case 3:
				var CariNama string
					fmt.Print("Masukan Nama Anda: ")
					fmt.Scan(&CariNama)
					user_tubes.LihatStatusKelulusan(arr, banyakMahasiswa, arr3, CariNama)
				}
		}
	}
}