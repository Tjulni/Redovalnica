// Redovalnica – Go modul

// Modul `redovalnica` omogoča upravljanje ocen študentov.  
// Primeren je za uporabo v majhnih CLI aplikacijah ali kot del večjega sistema.

// Modul ponuja:
// - dodajanje ocen,
// - izpis vseh ocen,
// - izračun in izpis končnega uspeha.

// Konfiguracijske spremenljivke:
// - `StOcen` – najmanjše število ocen za izračun povprečja,
// - `MinOcena` – najmanjša dovoljena ocena,
// - `MaxOcena` – največja dovoljena ocena.

// Example usage
// // konfiguracija
// redovalnica.StOcen = 3
// redovalnica.MinOcena = 1
// redovalnica.MaxOcena = 10

// // seznam študentov
// studenti := map[string]redovalnica.Student{
// 	"123": {Ime: "Ana", Priimek: "Kovač"},
// 	"456": {Ime: "Marko", Priimek: "Novak"},
// }

// // dodajanje ocen
// redovalnica.DodajOceno(studenti, "123", 9)
// redovalnica.DodajOceno(studenti, "123", 8)
// redovalnica.DodajOceno(studenti, "456", 5)

// // izpis
// redovalnica.IzpisVsehOcen(studenti)
// redovalnica.IzpisiKoncniUspeh(studenti)

package redovalnica

import (
	"fmt"
)

var StOcen = -1
var MinOcena = -1
var MaxOcena = -1

// Student predstavlja enega študenta z imenom, priimkom in seznamom ocen.
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// DodajOceno doda oceno izbranemu študentu.
//
// Funkcija preveri:
//   - ali študent obstaja,
//   - ali je ocena znotraj dovoljenega intervala.
//
// Če pogoji niso izpolnjeni, funkcija izpiše napako in ne spremeni podatkov.
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	st, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		fmt.Println("Napaka: študent z vpisno številko", vpisnaStevilka, "ne obstaja.")
		return
	}

	if ocena < MinOcena || ocena > MaxOcena {
		fmt.Printf("Napaka: ocena mora biti med %d in %d.", MinOcena, MaxOcena)
		return
	}

	st.Ocene = append(st.Ocene, ocena)
	studenti[vpisnaStevilka] = st
}

// povprecje izračuna povprečje ocen izbranega študenta.
//
// Funkcija vrne:
//   - −1, če študent ne obstaja,
//   - 0, če študent nima dovolj ocen (glede na StOcen),
//   - povprečje, če pogoji so izpolnjeni.
func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	st, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		return -1.0
	}

	if len(st.Ocene) < StOcen {
		return 0
	}

	var vsota int
	for _, ocena := range st.Ocene {
		vsota += ocena
	}
	return float64(vsota) / float64(len(st.Ocene))
}

// IzpisVsehOcen izpiše vse študente z njihovimi ocenami.
//
// Podatki so izpisani v obliki:
//   vpisna_stevilka - Ime Priimek: [ocene]
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, st := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, st.Ime, st.Priimek, st.Ocene)
	}
}

// IzpisiKoncniUspeh izpiše povprečje in opis uspeha za vsakega študenta.
//
// Ocena uspeha temelji na povprečju:
//   - povprečje >= 9  → "Odličen študent!"
//   - povprečje >= 6  → "Povprečen študent"
//   - drugače         → "Neuspešen študent"
func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisna, student := range studenti {
		p := povprecje(studenti, vpisna)
		fmt.Printf("%s %s: povprečna ocena %.1f -> ", student.Ime, student.Priimek, p)

		if p >= 9 {
			fmt.Println("Odličen študent!")
		} else if p >= 6 {
			fmt.Println("Povprečen študent")
		} else {
			fmt.Println("Neuspešen študent")
		}
	}
}
