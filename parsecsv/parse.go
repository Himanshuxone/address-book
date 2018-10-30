// package used to process csv and string the result in the structs
package parsecsv

import (
	"encoding/csv"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// readCSV is used to implement Read method for transforming the csv file before it is read by csv reader
type readCSV struct{ r io.Reader }

// Result contains the contact information of a person
type Result struct {
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Address   Address `json:"address"`
	Code      float64 `json:"code"`
}

// Address contains city, state and code
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

var (
	result   Result
	response []*Result
)

// This function is an implementation of read method to parse the file and remove unwanted chaarcters inside the
// csv file if found
func (read *readCSV) Read(b []byte) (n int, err error) {
	x := make([]byte, len(b))
	if n, err = read.r.Read(x); err != nil {
		return n, err
	}

	// create a regex to remove unwanted characters
	reg, err := regexp.Compile("[^a-zA-Z0-9.,\\s]+")

	if err != nil {
		log.Println(err)
	}

	processedString := reg.ReplaceAllString(string(x), "")
	copy(b, []byte(processedString)) // copy the processed string to the main bytes which will be read by csv reader
	return n, nil
}

// Search is used to find contact information of the variable url parameter(firstname).
func Search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var res []*Result
	for _, v := range response {
		if strings.EqualFold(v.FirstName, vars["firstname"]) {
			res = append(res, v)
		}
	}

	bytes, err := json.Marshal(res)

	CheckError(err)
	w.Write(bytes)
}

// ProcessCSV is reading the csv line by line and storing the result into struct for future requirement
// It is passing file to readCSV struct which implements io.reader and is parsing the file to remove unwanted charcaters
// in raw csv before passing to the csv reader. For more information look for Read() method in the same package
func ProcessCSV() {
	pwd, err := os.Getwd()
	CheckError(err)
	file, err := os.Open(filepath.Join(pwd, "dummy.csv"))
	defer file.Close()
	CheckError(err)
	reader := &readCSV{file}
	wg.Add(1)
	go func() {
		defer wg.Done()
		csvReader := csv.NewReader(reader)
		csvReader.LazyQuotes = true
		csvReader.Comma = ','
		csvReader.Comment = '#'
		csvReader.TrimLeadingSpace = true

		// Since the file does not contains any header we might ignore the process to read first line as header
		// if _, err := r.Read(); err != nil { //read header
		//     log.Fatal(err)
		// }

		// read csv one line at a time and store the data inside result struct at the same time
		for {
			rec, err := csvReader.Read() // It is better to read the lines one by one else in case of large csv it will
			// create problem when we readAll to the memory and then range over the lines.
			if err != nil {
				if err == io.EOF {
					break
				}

				if err, ok := err.(*csv.ParseError); ok && err.Err == csv.ErrFieldCount {
					log.Println(err)
				}
			}

			// This will check if teh length of the records slice is more than 5
			// As their are 6 columns
			if len(rec) > 5 {
				code, _ := strconv.ParseFloat(strings.TrimSpace(rec[5]), 64)
				response = append(response, &Result{
					FirstName: strings.TrimSpace(strings.Title(rec[0])),
					LastName:  strings.TrimSpace(rec[1]),
					Address: Address{
						Street: strings.TrimSpace(rec[2]),
						City:   strings.TrimSpace(rec[3]),
						State:  strings.TrimSpace(rec[4]),
					},
					Code: code,
				})
			}
		}
	}()

	// wait for the go routine to finish so that the file open for reading can be closed before parent exists
	// Else when running benchmark on the current function will take too much time and open too many files
	// Also one think can be check which is ulimit -a to see the resources limit in ubuntu
	wg.Wait()
	log.Println("Process Completed")
}

// CheckError is used to log error information inside the logs.txt
func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}
