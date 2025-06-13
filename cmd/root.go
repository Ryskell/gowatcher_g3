package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gowatcher",
	Short: "Gowatcher est un outil pour vérifier l'accessibilité des URLs.",
	Long:  "Un outil CLI en Go pour vérfifier letat d'URLs, gérer la concurrence et exporter les résultats.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: %v/n", err)
		os.Exit(1)
	}
}
