package test1

import (
	"fmt"
	"math/rand"
)

type patient struct {
	name   string
	health int
}

type drug struct {
	name string
	tox  int
}

type event struct {
	name string
	tox  int
}

//giveMed will modify the .health field of a single element of the slice of type patient according to
//the .tox value of single element of the slice of type drug
func giveMed(p []patient, d []drug, k int, n int) {
	fmt.Printf("Hai somministrato %v a %v.\n", d[n-1].name, p[k-1].name)
	p[k-1].health = p[k-1].health + d[n-1].tox
	if n == 2 || n == 6 {
		fmt.Printf("%v non si sente troppo bene.\n", p[k-1].name)
		if n == 6 {
			fmt.Printf("%v ha un'aura verde fosforescente.\n", p[k-1].name)
		}
	} else if n == 5 {
		fmt.Printf("%v si sente esattamente come prima.\n", p[k-1].name)
	} else {
		fmt.Printf("%v si sente meglio.\n", p[k-1].name)
	}
}

//spawnEvent will choose a single random element from a slice of type event and
//modify the .health field of a single element of the slice of type patient according to
//the .tox value of single element of the slice of type event
func spawnEvent(e []event, p []patient, k int) {
	n := rand.Intn(10)
	fmt.Print(e[n].name)
	p[k-1].health += e[n].tox
	return
}

//removePat will remove a single element from a slice of type patient and return a new slice
//with all the other elements in the same order
func removePat(p []patient, id int) []patient {
	p = append(p[:id], p[id+1:]...)
	return p
}
