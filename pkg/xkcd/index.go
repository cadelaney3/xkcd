package xkcd

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"strings"
	//"strconv"
)

type Comic struct {
	Month string
	Num int
	Link string
	Year string
	News string
	SafeTitle string `json:"safe_title"`
	Transcript string
	Alt string
	Img string
	Title string
	Day string
}

type QueryResponse struct {
	URL string
	Transcript string
}

const (
	baseURL = "https://xkcd.com"
	jsonPath = "/info.0.json"
)

var index []*Comic

func LoadIndex(path string) error { 
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Error reading file from path %s: %s", path, err)
	}

	if err := json.Unmarshal(bytes, &index); err != nil {
		return fmt.Errorf("Error unmarshaling file: %s", err)
	}
	return nil
}

func CreateIndex(path string) {
	index := make([]*Comic, 0)

	var comic *Comic
	if latest, ok := GetLatestComicID(); ok {
		for i:=1; i<=latest; i++ {
			if comic, ok = GetComic(i); ok {
				index = append(index, comic)
			}
		}
	}

	bytes, err := json.Marshal(index)
	if err != nil {
		fmt.Println("Error marshaling index")
		return
	}
	
	if err := ioutil.WriteFile(path, bytes, 0644); err != nil {
		fmt.Println(err)
		fmt.Println("Error writing to path")
		return
	}

}

func GetLatestComicID() (id int, ok bool) {
	if comic, ok := GetComic(0); ok {
		return comic.Num, true
	}
	return -1, false	
}

func GetComic(id int) (comic *Comic, ok bool) {
	var url string
	if id == 0 {
		url = baseURL + jsonPath
	} else {
		url = baseURL + "/" + strconv.Itoa(id) + jsonPath
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Fetch url returned a code of: %d for id: %d\n", resp.StatusCode, id)
		return nil, false
	}

	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		fmt.Println("Error unmarshaling comic of id: ", id)
		return nil, false
	}

	ok = true
	return
}

func QueryIndex(terms []string) (response []*QueryResponse, ok bool) {
	for _, t := range terms {
		if val, err := strconv.Atoi(t); err == nil {
			if temp, ok := SearchByID(val); ok {
				response = append(response, temp)
			}
		}
	}
	response = append(response, SearchTranscript(terms)...)
	if len(response) > 0 {
		return response, true
	}

	return nil, false
}

func SearchTranscript(terms []string) (response []*QueryResponse) {
	search := strings.Join(terms, " ")
	for i:=0; i<len(index); i++ {
		if strings.Contains(index[i].Transcript, search) {
			response = append(response, 
				&QueryResponse{ 
					URL: baseURL + "/" + strconv.Itoa(index[i].Num),
					Transcript: index[i].Transcript,
				},
			)
		}
	}
	return
}

func SearchByID(id int) (response *QueryResponse, ok bool) {
	fmt.Println("len index: ", len(index))
	if id > 0 && id < len(index) {
		comic := index[id-1]
		if comic.Num == id {
			response = &QueryResponse{
				URL: baseURL + "/" + strconv.Itoa(id),
				Transcript: comic.Transcript,
			}
			ok = true
			return 
		}
		diff := comic.Num - id
		if index[id-1-diff].Num == id {
			response = &QueryResponse{
				URL: baseURL + "/" + strconv.Itoa(id),
				Transcript: index[id-1-diff].Transcript,
			}
			ok = true
			return
		}
	}
	return nil, false
}

func SearchByTitle(title string) (response *QueryResponse) {
	for _, v := range index {
		if strings.Contains(strings.ToLower(v.Title), strings.ToLower(title)) {
			response = &QueryResponse{
				URL: baseURL + "/" + strconv.Itoa(v.Num),
				Transcript: v.Transcript,
			}
			return
		}
	}
	return nil
}

func SearchByYear(year string) (response []*QueryResponse) {
	for _, v := range index {
		if strings.Contains(v.Year, year) {
			response = append(response,
				&QueryResponse{
					URL: baseURL + "/" + strconv.Itoa(v.Num),
					Transcript: v.Transcript,
				},
			)
		}
	}
	return
}