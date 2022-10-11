package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	urlMain := "http://localhost:8090/docs/"
	resp, err := http.Get(urlMain)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var sfile []string  // List of all files shown on the site
	var dlfile []string // List of files having the letter "A" at left most
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		rf := scanner.Text()
		/*
			The scanner.Text() returns as following

			<pre>
			<a href="Estimation.pdf">Estimation.pdf</a>
			<a href="ForcAsting.pdf">ForcAsting.pdf</a>
			. . .
			<a href="Households%20pro-environment.pdf">Households pro-environment.pdf</a>
			<a href="treeApplications.pdf">treeApplications.pdf</a>
			</pre>

			To extract just the filenames the following codes are processed
		*/

		if len(rf) > 4 && rf[len(rf)-4:] == "pre>" {
			fmt.Printf("head/tail\n")
		} else {
			s := scanner.Text()
			readfile := s[strings.Index(s, ">")+1 : strings.Index(s, "</a>")]
			sfile = append(sfile, readfile)
			fmt.Println("..read filename: ", readfile)
		}
	}
	// Just the filenames are shown
	for i := 0; i < len(sfile); i++ {
		fmt.Println(sfile[i])
	}
	var winner []int    // Array, holds the left most index of letter "A" of all filenames. If not -1 is set
	var n = len(sfile)  // Holds the number of filenames
	var leftmost = 1000 // The global minimum index of the letter "A" found in filenames.
	var curLeft = -1    // Gets the index of left most letter "A".
	if n > 0 {          // Checks whether the server has at least one file

		for j := 0; j < n; j++ {
			curLeft = strings.Index(strings.ToUpper(sfile[j]), "A")
			if curLeft > -1 && curLeft < leftmost {
				leftmost = curLeft
			}
			winner = append(winner, curLeft)
		}
		if leftmost < 1000 { // If there is at least one letter "A"
			for j := 0; j < n; j++ {
				if winner[j] == leftmost {
					dlfile = append(dlfile, sfile[j])
					fmt.Println("Downloadable file: ", sfile[j])
				}
			}

			//**** download files begin
			var w sync.WaitGroup
			for i := 0; i < len(dlfile); i++ {
				w.Add(1)
				filename := dlfile[i]
				url := urlMain + filename
				if err := DownloadFile(filename, url, w); err != nil {
					panic(err)
				}
				fmt.Println(url)
			}
			w.Wait()
			//**** download files end
		} else {
			fmt.Println("\nNone of the filenames involve the letter <A>")
		}
	} else {
		fmt.Println("\n No file to process!")
	}
}

func DownloadFile(filepath string, url string, wg sync.WaitGroup) error {

	resp, err := http.Get(url)
	defer wg.Done()
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
