package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gowatcher_g3/internal/checker"
	"gowatcher_g3/internal/config"
	"sync"
)

var (
	inputFilePath string
	//outputFilePath
)
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Lance la fonction pour vérifier les URLs",
	Long:  "La commande 'check parcourt une liste prédéfinie d'URLs et affiche leur statut d'accessibilité",
	Run: func(cmd *cobra.Command, args []string) {

		if inputFilePath == "" {
			fmt.Println("Erreur sur le chemin du fichier d'entrée (--input)")
			return

		}
		targets, err := config.LoadTargetsFromFile(inputFilePath)
		if err != nil {
			fmt.Printf("erreur lors du chargement des URLs: %v/n", err)
			return
		}
		if len(targets) == 0 {
			fmt.Println("Aucune url a verfifier trouvee dans le fichier dentree")
			return
		}
		// creation waitgroup qui est un compteur
		var wg sync.WaitGroup
		resultsChan := make(chan checker.CheckResult, len(targets))

		wg.Add(len(targets))

		for _, url := range targets {
			// pour chaque URL on lance une routine
			// la fonction anonyme recoit une copie u de l'URL (important pour eviter un piege classique de
			//capture de variablez dans la boucle
			go func(t config.InputTarget) {
				//garantit qu'à la fin de la fonction, le compteur wg sera décrémenté de 1,
				//signalant que cette goroutine
				defer wg.Done()
				result := checker.CheckURL(t)
				resultsChan <- result // envoyer le resultat au channel

			}(url)
		}
		wg.Wait()
		close(resultsChan) //fermer canal apres que tous les resultats on ete envoye
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
