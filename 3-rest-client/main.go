package main

import (
				"fmt"
				"net/http"
				"encoding/json"
				"os"
				"strconv"
)

func main() {
				resp, err := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
				if err != nil {
								fmt.Println(err)
								os.Exit(1)
				}
				defer resp.Body.Close()
				fmt.Println("status code is " + resp.Status)


				var s moviedata

				if err:= json.NewDecoder(resp.Body).Decode(&s); err != nil {
					fmt.Errorf("errorparsingbody%v", err)
					return
				}
				var number float64

         number, _ = strconv.ParseFloat(s.ImdbRating, 64)
				 var numbera int = int(number*10)
				 fmt.Printf("The movie : Batman Begins was released in %s - the IMBD rating is %d%% with %s votes\n", s.Year, numbera, s.ImdbVotes)
}
