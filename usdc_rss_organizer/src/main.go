package main

import (
	"flag"
	"fmt"
	//	"net/url"
	//	"regexp"
	"bufio"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("function %s took %s", name, elapsed)
} // https://blog.stathat.com/2012/10/10/time_any_function_in_go.html

// TODO: use Strings to slice the numbers in the title, the first word before a space
// maybe use goroutines to call slicer to generate list asynchronously
// slicer needs to ALSO separate the string before the first space
func slicer(indictment string) string {
	reader := bufio.NewReader(os.Stdin)
	indictment, _ = reader.ReadString('\n')

	slice := strings.Split(indictment, " v. ")
	indictment = fmt.Sprintf("%s has been indicted by %s\n.", slice[1], slice[0])
	return indictment
}

type List struct {
	indictmentList []string `xml:"title"`
	dateList       []string `xml:"pubDate"`
}

type Indictment struct {
	caseNumber string
	indicter   string
	indictee   string
	date       string
}

func main() {
	defer timeTrack(time.Now(), "Execution took")

	var inputFile = flag.String("infile", "https://ecf.dcd.uscourts.gov/cgi-bin/rss_outside.pl", "Path to XML")
	if err != nil {
		fmt.Print("Failed to fetch rss feed\n")
		os.Exit(1)
	}

	b, _ := ioutil.ReadAll(inputFile)

	var q Query
	xml.Unmarshal(b, &q)

	var name string

	/* fix when there's time

		fmt.Println("Enter a name, or leave blank to display everything\n")
		fmt.Scanf("%d", &name)

		if len(name) == 0 {
			fmt.Printf("You want all entries on %d\n", &name)

		}

		if len(name) > 0 || time.After(time.Second * 5) {

		}*/
}
