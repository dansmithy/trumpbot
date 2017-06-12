package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func getResponseString() string {
	responses := []string{
		"It is my opinion that many of the leaks coming out of the White House are fabiricated lies made up by the #FakeNews media",
		"This is the single greatest witch hunt of a politician in American history!",
		"James Comey better hope that there are no 'tapes' of our conversations before he starts leaking to the press!",
		"The Russia-Trump collusion story is a total hoax, when will this taxpayer funder charade end?",
		"President Obama 'is the worst president in U.S. history!'",
	}
	return responses[rand.Intn(len(responses))]
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", getResponseString())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
