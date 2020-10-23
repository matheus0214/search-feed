package search

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

var datas = "datas/datas.json"
var wg sync.WaitGroup
var matchers = make(map[string]Matcher)

/*
   Create log.txt to set errors

   Get all feeds and search for this specified to search in file datas
*/
func Run() {
	fileerror, errors := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errors != nil {
		log.Fatalln(errors)
	}

	defer fileerror.Close()

	log.SetOutput(fileerror)

	// Types feeds search
	searchs, errFeeds := GetFeeds()
	if errFeeds != nil {
		log.Fatalln(errFeeds)
	}

	// File with datas
	file, err := os.Open(datas)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	// Values parsed
	var valuesSearch []*Values
	err = json.NewDecoder(file).Decode(&valuesSearch)
	if err != nil {
		log.Fatalln(err)
	}

	wg.Add(len(searchs))

	results := make(chan *Result)

	for _, value := range searchs {
		match, exist := matchers[value.Type]
		if !exist {
			log.Fatalf("Not register feed <%v>", value.Type)
		}

		go func(match Matcher, v []*Values) {
			res, errSearch := match.Search(v)
			if errSearch != nil {
				log.Fatalln(err)
			}

			for _, val := range res {
				results <- val
			}

			wg.Done()
		}(match, valuesSearch)

		go func() {
			wg.Wait()

			close(results)
		}()

	}

	for v := range results {
		fmt.Printf("Type: %s\nContent: %s\n", v.Type, v.Content)
	}

}

func Register(typeMatcher string, match Matcher) {
	fmt.Println("Register", typeMatcher, "matcher")
	if _, ok := matchers[typeMatcher]; ok {
		log.Fatalln("Matcher already register")
	}

	matchers[typeMatcher] = match
}
