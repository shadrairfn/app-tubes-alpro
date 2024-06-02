package sorting_tubes

import struct_tubes "github.com/shadrairfn/app-tubes-alpro/app/appStruct"

func NilaiAscending(arr *struct_tubes.TabArr, arr2 *struct_tubes.TabArr2, banyakMahasiswa int) {
	for pass := 1; pass < banyakMahasiswa; pass++ {
		idx := pass - 1
		for j := pass; j < banyakMahasiswa; j++ {
			if (*arr2)[idx].NilaiTes > (*arr2)[j].NilaiTes {
				idx = j
			}
		}
		temp := (*arr2)[idx].NilaiTes
		temp1 := (*arr)[idx].NamaMahasiswa
		temp2 := (*arr)[idx].NomorInduk

		(*arr2)[idx].NilaiTes = (*arr2)[pass - 1].NilaiTes
		(*arr)[idx].NamaMahasiswa = (*arr)[pass - 1].NamaMahasiswa
		(*arr)[idx].NomorInduk = (*arr)[pass - 1].NomorInduk
		
		(*arr2)[pass - 1].NilaiTes = temp
		(*arr)[pass - 1].NamaMahasiswa = temp1
		(*arr)[pass - 1].NomorInduk = temp2
	}
}

func NilaiDescending(arr *struct_tubes.TabArr, arr2 *struct_tubes.TabArr2, banyakMahasiswa int) {
	for pass := 1; pass < banyakMahasiswa; pass++ {
		j := pass - 1
		temp := (*arr2)[pass].NilaiTes
		temp1 := (*arr)[pass].NamaMahasiswa
		temp2 := (*arr)[pass].NomorInduk
		for j >= 0 && temp > (*arr2)[j].NilaiTes {
			(*arr)[j+1].NamaMahasiswa = (*arr)[j].NamaMahasiswa
			(*arr)[j+1].NomorInduk = (*arr)[j].NomorInduk
			(*arr2)[j+1].NilaiTes = (*arr2)[j].NilaiTes
			j--
		}
		(*arr2)[j+1].NilaiTes = temp
		(*arr)[j+1].NamaMahasiswa = temp1
		(*arr)[j+1].NomorInduk = temp2
	}
}