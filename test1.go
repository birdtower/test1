package test1

import (
	"fmt"
	"math/rand"
)

type Patient struct {
	Name   string
	Health int
}

type Drug struct {
	Name string
	Tox  int
}

type Event struct {
	Name string
	Tox  int
}

//GiveMed will modify the .health field of a single element of the slice of type patient according to
//the .tox value of single element of the slice of type drug
func GiveMed(p []Patient, d []Drug, k int, n int) {
	fmt.Printf("Hai somministrato %v a %v.\n", d[n-1].Name, p[k-1].Name)
	p[k-1].Health = p[k-1].Health + d[n-1].Tox
	if n == 2 || n == 6 {
		fmt.Printf("%v non si sente troppo bene.\n", p[k-1].Name)
		if n == 6 {
			fmt.Printf("%v ha un'aura verde fosforescente.\n", p[k-1].Name)
		}
	} else if n == 5 {
		fmt.Printf("%v si sente esattamente come prima.\n", p[k-1].Name)
	} else {
		fmt.Printf("%v si sente meglio.\n", p[k-1].Name)
	}
}

//SpawnEvent will choose a single random element from a slice of type event and
//modify the .health field of a single element of the slice of type patient according to
//the .tox value of single element of the slice of type event
func SpawnEvent(e []Event, p []Patient, k int) {
	n := rand.Intn(10)
	fmt.Print(e[n].Name)
	p[k-1].Health += e[n].Tox
	return
}

//RemovePat will remove a single element from a slice of type patient and return a new slice
//with all the other elements in the same order
func RemovePat(p []Patient, id int) []Patient {
	p = append(p[:id], p[id+1:]...)
	return p
}
