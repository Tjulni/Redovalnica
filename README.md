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