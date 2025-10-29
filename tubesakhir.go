package main

import "fmt"

const NMAX int = 100

type suhuHarian struct {
	Hari int
	Suhu int
}

type tabSuhu [NMAX]suhuHarian

func inputData(H *tabSuhu, n *int) {
	var i int
	fmt.Println("Masukkan jumlah hari (0 untuk batal): ")
	fmt.Scan(n)
	if *n == 0 {
		fmt.Println("Input dibatalkan. Kembali ke menu.")
		return
	}
	for i = 0; i < *n; i++ {
		fmt.Printf("Masukkan suhu hari ke-%d (-999 untuk batal): ", i+1)
		var suhu int
		fmt.Scan(&suhu)
		if suhu == -999 {
			fmt.Println("Input suhu dibatalkan. Kembali ke menu.")
			*n = i // hanya menyimpan data sampai titik dibatalkan
			return
		}
		H[i].Hari = i + 1
		H[i].Suhu = suhu
	}
}

func tampilData(H tabSuhu, n int) {
	fmt.Println("\nData Suhu Harian:")
	for i := 0; i < n; i++ {
		if H[i].Suhu != -999 {
			fmt.Printf("Hari ke-%d: %d °C\n", H[i].Hari, H[i].Suhu)
		}
	}
}

func cariMaks(H tabSuhu, n int) int {
	max := -999
	for i := 0; i < n; i++ {
		if H[i].Suhu != -999 && (max == -999 || H[i].Suhu > max) {
			max = H[i].Suhu
		}
	}
	return max
}

func cariMin(H tabSuhu, n int) int {
	min := 9999
	for i := 0; i < n; i++ {
		if H[i].Suhu != -999 && H[i].Suhu < min {
			min = H[i].Suhu
		}
	}
	return min
}

func selectionSortDescending(data *tabSuhu, n int) {
	for pass := 0; pass < n-1; pass++ {
		maxIdx := pass
		for j := pass + 1; j < n; j++ {
			if data[j].Suhu != -999 && data[maxIdx].Suhu < data[j].Suhu {
				maxIdx = j
			}
		}
		data[pass], data[maxIdx] = data[maxIdx], data[pass]
	}
}

func insertionSortAscending(data *tabSuhu, n int) {
	for i := 1; i < n; i++ {
		temp := data[i]
		j := i
		for j > 0 && data[j-1].Suhu != -999 && data[j-1].Suhu > temp.Suhu {
			data[j] = data[j-1]
			j--
		}
		data[j] = temp
	}
}

func sequentialSearch(suhu tabSuhu, n int, hari int) int {
	for i := 0; i < n; i++ {
		if suhu[i].Hari == hari && suhu[i].Suhu != -999 {
			return i
		}
	}
	return -1
}

func binarySearch(suhu tabSuhu, n int, hari int) int {
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if suhu[mid].Suhu == -999 {
			left, right := mid-1, mid+1
			found := -1
			for left >= low || right <= high {
				if left >= low && suhu[left].Suhu != -999 {
					found = left
					left = -1 // keluar dari loop
				}
				if right <= high && suhu[right].Suhu != -999 {
					found = right
					right = high + 1 // keluar dari loop
				}
				left--
				right++
			}
			if found == -1 {
				return -1
			}
			mid = found
		}
		if suhu[mid].Hari < hari {
			low = mid + 1
		} else if suhu[mid].Hari > hari {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func sortHariAscending(data *tabSuhu, n int) {
	for i := 1; i < n; i++ {
		temp := data[i]
		j := i
		for j > 0 && data[j-1].Hari > temp.Hari {
			data[j] = data[j-1]
			j--
		}
		data[j] = temp
	}
}

func main() {
	var H tabSuhu
	var n, pilihan, hari, nilai int
	selesai := false

	for !selesai {
		fmt.Println("\n--- MENU APLIKASI SUHU ---")
		fmt.Println("1. Input Data Suhu")
		fmt.Println("2. Tampilkan Data Suhu")
		fmt.Println("3. Cari Suhu Tertinggi dan Terendah")
		fmt.Println("4. Cari Suhu Berdasarkan Hari")
		fmt.Println("5. Edit Data Suhu")
		fmt.Println("6. Hapus Data Suhu")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilihan)

		if pilihan == 0 {
			selesai = true
		} else if pilihan == 1 {
			inputData(&H, &n)
		} else if pilihan == 2 {
			var algoritma int
			fmt.Println("Pilih algoritma pengurutan:")
			fmt.Println("1. Ascending")
			fmt.Println("2. Descending")
			fmt.Println("0. Kembali ke menu")
			fmt.Print("Pilihan: ")
			fmt.Scan(&algoritma)

			if algoritma == 1 {
				insertionSortAscending(&H, n)
				tampilData(H, n)
			} else if algoritma == 2 {
				selectionSortDescending(&H, n)
				tampilData(H, n)
			} else if algoritma == 0 {
				fmt.Println("Kembali ke menu.")
			} else {
				fmt.Println("Algoritma tidak valid.")
			}
		} else if pilihan == 3 {
			fmt.Printf("\nSuhu Tertinggi: %d°C\n", cariMaks(H, n))
			fmt.Printf("Suhu Terendah: %d°C\n", cariMin(H, n))
		} else if pilihan == 4 {
			fmt.Print("Cari suhu hari ke- (0 untuk batal): ")
			fmt.Scan(&hari)
			if hari != 0 {
				idx := sequentialSearch(H, n, hari)
				if idx != -1 {
					fmt.Printf("Suhu pada hari ke-%d adalah %d°C\n", hari, H[idx].Suhu)
				} else {
					fmt.Println("Tidak ditemukan.")
				}
			} else {
				fmt.Println("Pencarian dibatalkan. Kembali ke menu.")
			}
		} else if pilihan == 5 {
			fmt.Print("Edit suhu hari ke- (0 untuk batal): ")
			fmt.Scan(&hari)
			if hari != 0 {
				idx := sequentialSearch(H, n, hari)
				if idx != -1 {
					fmt.Print("Masukkan nilai baru: ")
					fmt.Scan(&nilai)
					H[idx].Suhu = nilai
				} else {
					fmt.Println("Hari tidak ditemukan.")
				}
			} else {
				fmt.Println("Edit dibatalkan. Kembali ke menu.")
			}
		} else if pilihan == 6 {
			fmt.Print("Hapus suhu hari ke- (0 untuk batal): ")
			fmt.Scan(&hari)
			if hari != 0 {
				sortHariAscending(&H, n)
				idx := binarySearch(H, n, hari)
				if idx != -1 {
					H[idx].Suhu = -999
					fmt.Printf("Suhu di hari ke-%d berhasil dihapus.\n", hari)
				} else {
					fmt.Println("Hari tidak ditemukan.")
				}
			} else {
				fmt.Println("Penghapusan dibatalkan. Kembali ke menu.")
			}
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}
