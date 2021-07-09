package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"
	"time"

	cache "github.com/patrickmn/go-cache"
	memcache "github.com/rainycape/memcache"
	bitly "github.com/zpnk/go-bitly"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint64(len(alphabet)) + 1
)

var mc *memcache.Client

var c *cache.Cache

func main() {
	var err error

	c = cache.New(5*time.Minute, 10*time.Minute)

	mc, err = memcache.New("127.0.0.1:11211")
	if err != nil {
		panic(err)
	}
	fmt.Println(Encode(length))

	httphandler()

}

func Encode(number uint64) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(11)

	for ; number > 0; number = number / length {
		encodedBuilder.WriteByte(alphabet[(number % length)])
	}

	return encodedBuilder.String()
}

func Decode(encoded string) (uint64, error) {
	var number uint64

	for i, symbol := range encoded {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)

		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("invalid character: " + string(symbol))
		}
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}

func fetchSHortURL(url string) string {
	val, found := c.Get(url)
	value, err := mc.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("memcache value", value)
	fmt.Println(val)
	if found {
		res, ok := val.(string)
		if ok && res != "" {
			return res
		}
	}

	b := bitly.New("a95165eb986a6eecb6d812d5012c4d49b17f4d37")

	shortURL, err := b.Links.Shorten(url)
	if err != nil {
		log.Print("error    ", err)
	}

	//	fmt.Print(shortURL.LongURL, shortURL.URL, err)
	c.Set(shortURL.LongURL, shortURL.URL, cache.NoExpiration)
	mc.Set(&memcache.Item{Key: shortURL.LongURL, Value: []byte(shortURL.URL), Expiration: 0})
	log.Printf("Output %s", shortURL.URL)
	return shortURL.URL
}

func httphandler() {
	http.HandleFunc("/GET", ShortenURL)
	http.ListenAndServe(":8080", nil)
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	urlval := r.URL.Query()
	urlRequest := urlval.Get("hi")
	shorturl := fetchSHortURL(urlRequest)
	fmt.Fprint(w, shorturl)
	json.NewEncoder(w).Encode(shorturl + "SHASHANK")
}
