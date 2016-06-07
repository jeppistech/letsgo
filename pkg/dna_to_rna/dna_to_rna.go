package dna_to_rna

// Detta är ett paket med en publik funktion, ToRNA(string)
// Översätter DNA-strängar till RNA

// Denna metod är private och kan inte anropas utanför denna fil
func charToRNA(dna_base string) string {
	switch dna_base {
	case "A":
		return "U"
	case "T":
		return "A"
	case "C":
		return "G"
	case "G":
		return "C"
	}
	return ""
}

func ToRNA(DNA string) string {
	var RNA string = ""
	for i := 0; i < len(DNA); i++ {
		RNA += charToRNA(string(DNA[i]))
	}
	return RNA
}
