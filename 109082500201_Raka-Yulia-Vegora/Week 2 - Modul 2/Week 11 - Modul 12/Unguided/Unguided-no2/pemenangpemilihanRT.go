package main

import "fmt"

func main() {
	votes := make(map[int]int)
	totalInput := 0
	validVotes := 0

	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		totalInput++
		if n >= 1 && n <= 20 {
			votes[n]++
			validVotes++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalInput)
	fmt.Printf("Suara sah: %d\n", validVotes)

	// Cari suara terbanyak pertama dan kedua
	firstMax := 0
	secondMax := 0

	for _, count := range votes {
		if count > firstMax {
			secondMax = firstMax
			firstMax = count
		} else if count > secondMax {
			secondMax = count
		}
	}

	// Cari kandidat ketua (suara terbanyak, nomor terkecil jika sama)
	ketua := -1
	for candidate := 1; candidate <= 20; candidate++ {
		if votes[candidate] == firstMax {
			if ketua == -1 {
				ketua = candidate
			}
		}
	}

	// Cari kandidat wakil (suara terbanyak kedua, nomor terkecil jika sama)
	// Tapi bukan kandidat ketua
	wakil := -1
	// Hitung berapa kandidat dengan suara firstMax
	countFirst := 0
	for _, count := range votes {
		if count == firstMax {
			countFirst++
		}
	}

	if countFirst >= 2 {
		// Ada tie di posisi pertama, wakil = kandidat terkecil berikutnya dengan firstMax
		for candidate := 1; candidate <= 20; candidate++ {
			if votes[candidate] == firstMax && candidate != ketua {
				wakil = candidate
				break
			}
		}
	} else {
		// Wakil dari suara terbanyak kedua
		for candidate := 1; candidate <= 20; candidate++ {
			if votes[candidate] == secondMax {
				wakil = candidate
				break
			}
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}