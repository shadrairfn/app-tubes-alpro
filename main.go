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
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&chooseMenu)
		fmt.Println()
		if chooseMenu > 2 {
			break
		}
		switch chooseMenu {
		case 1:
			menu_tubes.MenuAdmin()
			fmt.Print("Pilih Menu: ")
			fmt.Scan(&chooseMenuAdmin)
			fmt.Println()
			switch chooseMenuAdmin {
			case 1:
				admin_tubes.CekPendaftar(arr, arr2, banyakMahasiswa, arr3)
				fmt.Println()
			case 2:
				admin_tubes.BeriKonfirmasi(arr, arr2, banyakMahasiswa, &arr3)
				fmt.Println()
			case 3:
				admin_tubes.EditMahasiswa(&arr, &arr2, &banyakMahasiswa, &arr3)
				fmt.Println()
			}
		case 2:
			menu_tubes.MenuUser()
			fmt.Print("Pilih Menu: ")
			fmt.Scan(&chooseMenuUser)
			fmt.Println()
			switch chooseMenuUser {
			case 1:
				user_tubes.IsInput(&arr, &arr2, &banyakMahasiswa)
				fmt.Println()
			case 2:
				var CariNama string
				fmt.Print("Masukan Nama Anda: ")
				fmt.Scan(&CariNama)
				fmt.Println()
				user_tubes.CekStatusMahasiswa(arr, banyakMahasiswa, CariNama)
				fmt.Println()
			case 3:
				var CariNama string
				fmt.Print("Masukan Nama Anda: ")
				fmt.Scan(&CariNama)
				fmt.Println()
				user_tubes.LihatStatusKelulusan(arr, arr2, banyakMahasiswa, arr3, CariNama)
				fmt.Println()
			}
		}
	}
}

/* TEST CASE

2
1
shadra
123123
8 9 2004
19
bandung
malik
1
70
2
1
mehdi
456456
8 9 2004
19
bandung
jl_malik
1
75
2
1
irfan
123123
8 9 2004
19
bandung
jl_kidam_malik
1
80
2
1
rafif
123123
8 9 2004
19
bandung
jlmalik
2
77
2
1
milka
123123
8 9 2004
19
bandung
jl_moha
2
81

*/