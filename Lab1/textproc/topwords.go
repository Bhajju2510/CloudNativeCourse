package textproc

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strings"
)

func topWords(path string, K int) []WordCount {
    // Your code here.....
    //open the file indicated by path

    dat, err := os.Open(path)
    if err != nil {
        checkError(err)}

    defer dat.Close()

    //read line by line
    scanner := bufio.NewScanner(dat)

    if err := scanner.Err(); err != nil {
        checkError(err)
    }

    //Reading each line and appending to a string
    var strText string
    for scanner.Scan() {
        strText = strText + " " + scanner.Text()
    }
    //record the count of words using map
    //convert each entry in the map into a word count object
    wrds := []WordCount{}
    for index, element := range getWordCount(strText) {
        wrd := WordCount{Word: index, Count: element}
        //append wordcount object to a slice
        wrds = append(wrds, wrd)
    }

    //Edgecase where K is greater than length
    if K > len(wrds) {
        K = len(wrds)
    }

    //sort the slice by object count
    sortWordCounts(wrds)

    var a []WordCount

    //return top k of slices
    for i := 0; i < K; i++ {
        a = append(a, wrds[i])
    }

    fmt.Println("------------Finally--------------")
    fmt.Println(a)
    return a

    //return top k of slices
}

func getWordCount(str string) map[string]int {
    wordList := strings.Fields(str)
    counts := make(map[string]int)
    for _, word := range wordList {
        //_, ok := counts[word]
        //if ok {
            counts[word] += 1
        //} else {
          //  counts[word] = 1
        //}
    }
    return counts
}

//--------------- DO NOT MODIFY----------------!
// A struct that represents how many times a word is observed in a document
type WordCount struct {
    Word  string
    Count int
}

// Method to convert struct to string format
func (wc WordCount) String() string {
    return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
func sortWordCounts(wordCounts []WordCount) {
    sort.Slice(wordCounts, func(i, j int) bool {
        wc1 := wordCounts[i]
        wc2 := wordCounts[j]
        if wc1.Count == wc2.Count {
            return wc1.Word < wc2.Word
        }
        return wc1.Count > wc2.Count
    })
}
func checkError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

