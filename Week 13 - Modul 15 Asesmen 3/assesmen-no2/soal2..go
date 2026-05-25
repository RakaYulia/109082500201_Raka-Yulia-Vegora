package main

import "fmt"

const NMAX = 1001

type player struct {
	fristName string
	lastName  string
	goals     int
	assists   int
}

type arrPlayer [NMAX]player

func shouldBefore(a, b player) bool {
	if a.goals != b.goals {
		return a.goals > b.goals
	}

	if a.assists != b.assists {
		return a.assists > b.assists
	}

	nameA := a.fristName + " " + a.lastName
	nameB := b.fristName + " " + b.lastName

	return nameA < nameB
}

func selectionSort(T arrPlayer, n int) arrPlayer {
	for i := 1; i < n; i++ {
		key := T[i]
		j := i - 1
		for j >= 0 && shouldBefore(key, T[j]) {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = key
	}
	return T
}

func main() {
	var T arrPlayer
	var n int

	fmt.Println("Masukan Data Input :")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].fristName, &T[i].lastName, &T[i].goals, &T[i].assists)
	}

	T = selectionSort(T, n)

	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n",
			T[i].fristName,
			T[i].lastName,
			T[i].goals,
			T[i].assists,
		)
	}
}