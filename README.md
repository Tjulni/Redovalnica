# Redovalnica
---

### 📌 Example usage (kot Go datoteka za Godoc)

Ustvari datoteko `example_test.go` v paketu `redovalnica`:

```go
package redovalnica_test

import (
    "github.com/Tjulni/Redovalnica/redovalnica"
)

func Example() {
    redovalnica.StOcen = 2
    redovalnica.MinOcena = 1
    redovalnica.MaxOcena = 10

    studenti := map[string]redovalnica.Student{
        "001": {Ime: "Luka", Priimek: "Zajc"},
    }

    redovalnica.DodajOceno(studenti, "001", 7)
    redovalnica.DodajOceno(studenti, "001", 9)

    redovalnica.IzpisVsehOcen(studenti)
    redovalnica.IzpisiKoncniUspeh(studenti)

    // Output:
    // REDOVALNICA:
    // 001 - Luka Zajc: [7 9]
    // Luka Zajc: povprečna ocena 8.0 -> Povprečen študent
}