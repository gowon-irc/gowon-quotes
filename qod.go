package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	quotesURL = "http://quotes.rest/qod.json"
)

// JSON represents the return from quotes.rest
type JSON struct {
	Contents Contents `json:"contents"`
}

func (j JSON) String() string {
	return j.Contents.String()
}

// Contents represents the list of quotes
type Contents struct {
	Quotes []Quote `json:"quotes"`
}

func (c Contents) String() string {
	if len(c.Quotes) < 1 {
		return "No quotes found"
	}
	return c.Quotes[0].String()
}

// Quote represents a single quote
type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func (q Quote) String() string {
	return fmt.Sprintf("%s\n- %s", q.Quote, q.Author)
}

func qod() (msg string, err error) {
	res, err := http.Get(quotesURL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	j := JSON{}
	err = json.Unmarshal(body, &j)

	if err != nil {
		return "", err
	}

	return j.String(), nil
}
