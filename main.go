package main

import (
	"encoding/csv"
	"log"
	"os"
)

type User struct {
    fullName  string
    email   string
    location string
}

func main() {

    records, err := readData("input_example.csv")

	nameExistMap := make(map[string]bool)
    
	if err != nil {
        log.Fatal(err)
    }

	f, err := os.Create("output_example.csv")
    defer f.Close()

	 if err != nil {

        log.Fatalln("failed to open file", err)
    }

    for _, record := range records {

        user := User{
            fullName:  record[0],
            email:   record[1],
            location: record[2],
        }

        if _, exist := nameExistMap[user.fullName]; exist {
        	continue
    	} else {
        	nameExistMap[user.fullName] = true
			w := csv.NewWriter(f)
    		defer w.Flush()
			extraData := []string{record[0], record[1], record[2]}
    		w.Write(extraData)
    		}
    	}	
        
	
}

func readData(fileName string) ([][]string, error) {

    f, err := os.Open(fileName)

    if err != nil {
        return [][]string{}, err
    }

    defer f.Close()

    r := csv.NewReader(f)

    if _, err := r.Read(); err != nil {
        return [][]string{}, err
    }

    records, err := r.ReadAll()

    if err != nil {
        return [][]string{}, err
    }

    return records, nil
}

