package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	rna "../pkg/dna_to_rna"
)

var (
	defaultDNA = "ACGTGCTATTAGTCAT"

	dna = ""
)

func init() {
	RootCmd.AddCommand(cmdConvert)

	cmdConvert.Flags().StringVarP(&dna, "dna", "d", defaultDNA, "the DNA string to convert")
}

var cmdConvert = &cobra.Command{
	Use:   "convert",
	Short: "Convert a DNA string to m-RNA",
	Long:  `Convert a DNA string to message-RNA`,
	Run: func(cmd *cobra.Command, args []string) {

		// Deklarera och ge värde på samma gång med operatorn ":="
		rnastring := rna.ToRNA(dna)

		// Visa värdet åt användaren
		fmt.Println("Your RNA:", rnastring)
	},
}
