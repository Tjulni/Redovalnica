# Redovalnica – Go modul

Modul `redovalnica` omogoča upravljanje ocen študentov.  
Primeren je za uporabo v majhnih CLI aplikacijah ali kot del večjega sistema.

Modul ponuja:
- dodajanje ocen,
- izpis vseh ocen,
- izračun in izpis končnega uspeha.

Konfiguracijske spremenljivke:
- `StOcen` – najmanjše število ocen za izračun povprečja,
- `MinOcena` – najmanjša dovoljena ocena,
- `MaxOcena` – največja dovoljena ocena.

### Example usage

```go
redovalnica.StOcen = 4
redovalnica.MinOcena = 1
redovalnica.MaxOcena = 10

stud := map[string]redovalnica.Student{
    "123": {Ime: "Ana", Priimek: "Kovač"},
    "456": {Ime: "Marko", Priimek: "Novak"},
}

redovalnica.DodajOceno(stud, "123", 8)
redovalnica.DodajOceno(stud, "123", 10)
redovalnica.DodajOceno(stud, "456", 4)

fmt.Println("Ocene:")
redovalnica.IzpisVsehOcen(stud)

fmt.Println("\nKončni uspeh:")
redovalnica.IzpisiKoncniUspeh(stud)
```
Redovalnica generira naslednje vrednosti
Ocene:
REDOVALNICA:
123 - Ana Kovač: [8 10]
456 - Marko Novak: [4]

Končni uspeh:
Ana Kovač: povprečna ocena 0.0 -> Neuspešen študent
Marko Novak: povprečna ocena 0.0 -> Neuspešen študent