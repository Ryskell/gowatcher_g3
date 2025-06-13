package main

import (
	"errors"
	"fmt"
	"gowatcher_g3/internal/checker"
	"sync"
)

func main() {
	targets := []string{
		"https://www.google.com",
		"https://www.notarealwebsite.abc",
		"https://github.com",
		"https://www.movie.database/film/details",
		"https://www.gaming.news/release/new-game",
		"https://www.health.clinic/appointment/online",
		"https://www.car.manufacturer/model/electric",
		"https://www.home.decor/ideas/living-room",
		"https://www.environmental.org/project/clean-water",
		"https://www.space.agency/mission/mars",
		"https://www.fashion.magazine/trend/summer",
		"https://www.tech.conference/schedule/day1",
		"https://www.food.blog/recipe/dessert",
		"https://www.online.course/programming/python",
		"https://www.travel.guide/city/paris",
		"https://www.music.label/artist/new-album",
		"https://www.sports.club/events/match",
		"https://www.photography.tips/technique/lighting",
		"https://www.diy.tools/review/drill",
		"https://www.pet.vet/service/vaccination",
		"https://www.gardening.store/seeds/flower",
		"https://www.finance.advice/retirement/planning",
		"https://www.history.podcast/episode/ww2",
		"https://www.language.exchange/partner/find",
		"https://www.book.review/author/classic",
		"https://www.movie.review/genre/comedy",
		"https://www.gaming.forum/topic/strategy",
	}

	// creation waitgroup qui est un compteur
	var wg sync.WaitGroup

	wg.Add(len(targets))

	for _, url := range targets {
		// pour chaque URL on lance une routine
		// la fonction anonyme recoit une copie u de l'URL (important pour eviter un piege classique de
		//capture de variablez dans la boucle
		go func(u string) {
			//garantit qu'à la fin de la fonction, le compteur wg sera décrémenté de 1,
			//signalant que cette goroutine
			defer wg.Done()
			result := checker.CheckURL(u)
			if result.Err != nil {
				var unreachable *checker.UnreachableURLError
				if errors.As(result.Err, &unreachable) {
					fmt.Printf("%s set inacessible : %v/n", unreachable.URL, unreachable.Err)
				} else {
					fmt.Printf("KO %s : erreur - %v \n", result.Target, result.Err)
				}
			} else {
				fmt.Printf("OK %s : %v\n", result.Target, result.Status)
			}
		}(url)
	}
	wg.Wait()
}
