package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func server(w http.ResponseWriter, r *http.Request) [][]string {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))

	}

	return records
}

func echo(w http.ResponseWriter, r *http.Request) {
	var response string
	records := server(w, r)
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	fmt.Fprint(w, response)
}

func invert(w http.ResponseWriter, r *http.Request) {
	records := server(w, r)
	sliceX := len(records)
	sliceY := len(records[0])
	transpose := make([][]string, sliceX)
	var response string

	for i := range transpose {
		transpose[i] = make([]string, sliceX)
	}

	for i := 0; i < sliceX; i++ {
		for j := 0; j < sliceY; j++ {
			transpose[j][i] = records[i][j]
		}
	}

	for _, row := range transpose {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	fmt.Fprintf(w, "Invert:\n%s\n", response)
}

func flatten(w http.ResponseWriter, r *http.Request) {
	response := []string{}
	records := server(w, r)
	for _, row := range records {
		for _, v := range row {
			response = append(response, v)
		}
	}

	fmt.Fprintf(w, "Flatten: %s\n", strings.Join(response, ","))
}

func sum(w http.ResponseWriter, r *http.Request) {

	records := server(w, r)
	intResponse := 0

	for _, row := range records {
		for _, v := range row {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

			intResponse += n
		}
	}

	fmt.Fprintf(w, "Sum: %d\n", intResponse)

}

func multiply(w http.ResponseWriter, r *http.Request) {

	records := server(w, r)
	intResponse := 1

	for _, row := range records {
		for _, v := range row {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

			intResponse *= n
		}
	}

	fmt.Fprintf(w, "Multiply: %d\n", intResponse)

}

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/invert", invert)
	http.HandleFunc("/flatten", flatten)
	http.HandleFunc("/sum", sum)
	http.HandleFunc("/multiply", multiply)
	http.ListenAndServe(":8080", nil)

}
