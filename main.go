package main

import (
  "log"
  "regexp"
  "strings"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/golang/gddo/httputil/header"
)

type Request struct {
  Text string
}

type Response struct {
  Number string
  Value int
}

func romanTranslate(w http.ResponseWriter, r *http.Request) {
  // If the Content-Type header is present, check that it has the value application/json.
  if r.Header.Get("Content-Type") != "" {
    value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
    if value != "application/json" {
      msg := "Content-Type header is not application/json"
      http.Error(w, msg, http.StatusUnsupportedMediaType)
      return
    }
  }

  // Declare a new Request struct.
  var request Request

  // Try to decode the request body into the struct. If there is an error,
  // respond to the client with the error message and a 400 status code.
  err := json.NewDecoder(r.Body).Decode(&request)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  // Run romanNumerals func to fetch roman numbers and values
  var number, value = romanNumerals(request.Text)
  response := Response{number, value}

  // Enconde response to return as json
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(&response)
}

func romanNumerals(text string) (string, int) {
  var romans = findRomans(text)
  var max_roman = ""
  var max = 0

  for _, roman := range romans {
    value := romanDecode(roman)
    if value > max {
      max_roman = roman
      max = value
    }
  }
  return max_roman, max
}

func findRomans(text string) []string {
  if text == "" { return []string{} }
  romansList := []string{}

  roman := strings.ToUpper(text)
  r := regexp.MustCompile("M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})")
  regexGroups := r.FindAllStringSubmatch(roman, -1)

  for i := 1; i < len(regexGroups); i++ {
    romansList = append(romansList, regexGroups[i][0])
  }
  return romansList
}

func romanDecode(roman string) int {
  var decoder = map[rune]int {'I': 1,'V': 5,'X': 10,'L': 50,'C': 100,'D': 500,'M': 1000,}
  if len(roman) == 0 { return 0 }
  first := decoder[rune(roman[0])]
  if len(roman) == 1 { return first }
  next := decoder[rune(roman[1])]
  if next > first { return (next - first) + romanDecode(roman[2:]) }
  return first + romanDecode(roman[1:])
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/search", romanTranslate)
  log.Fatal(http.ListenAndServe(":8080", router))
}
