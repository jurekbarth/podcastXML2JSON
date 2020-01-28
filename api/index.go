package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// AtomLink Type
type AtomLink struct {
	XMLName xml.Name `xml:"atom-link"`
	HREF    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}

// Author Type
type Author struct {
	XMLName xml.Name `xml:"itunes-owner"`
	Name    string   `xml:"itunes-name"`
	Email   string   `xml:"itunes-email"`
}

// Enclosure Type
type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`

	// URL is the downloadable url for the content. (Required)
	URL string `xml:"url,attr"`

	// Length is the size in Bytes of the download. (Required)
	Length int64 `xml:"-"`
	// LengthFormatted is the size in Bytes of the download. (Required)
	//
	// This field gets overwritten with the API when setting Length.
	LengthFormatted string `xml:"length,attr"`

	// Type is MIME type encoding of the download. (Required)
	Type EnclosureType `xml:"-"`
	// TypeFormatted is MIME type encoding of the download. (Required)
	//
	// This field gets overwritten with the API when setting Type.
	TypeFormatted string `xml:"type,attr"`
}

// EnclosureType Type
type EnclosureType int

const (
	M4A EnclosureType = iota
	M4V
	MP4
	MP3
	MOV
	PDF
	EPUB
)

// ICategory Type
type ICategory struct {
	XMLName     xml.Name `xml:"itunes-category"`
	Text        string   `xml:"text,attr"`
	ICategories []*ICategory
}

// IImage Type
type IImage struct {
	XMLName xml.Name `xml:"itunes-image"`
	HREF    string   `xml:"href,attr"`
}

// ISummary Type
type ISummary struct {
	XMLName xml.Name `xml:"itunes-summary"`
	Text    string   `xml:",cdata"`
}

// Image Type
type Image struct {
	XMLName     xml.Name `xml:"image"`
	URL         string   `xml:"url"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description,omitempty"`
	Width       int      `xml:"width,omitempty"`
	Height      int      `xml:"height,omitempty"`
}

// Item Type
type Item struct {
	XMLName          xml.Name   `xml:"item"`
	GUID             string     `xml:"guid"`
	Title            string     `xml:"title"`
	Link             string     `xml:"link"`
	Description      string     `xml:"description"`
	Author           *Author    `xml:"-"`
	AuthorFormatted  string     `xml:"author,omitempty"`
	Category         string     `xml:"category,omitempty"`
	Comments         string     `xml:"comments,omitempty"`
	Source           string     `xml:"source,omitempty"`
	PubDate          *time.Time `xml:"-"`
	PubDateFormatted string     `xml:"pubDate,omitempty"`
	Enclosure        *Enclosure

	// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
	IAuthor            string `xml:"itunes-author,omitempty"`
	ISubtitle          string `xml:"itunes-subtitle,omitempty"`
	ISummary           *ISummary
	IImage             *IImage
	IDuration          string `xml:"itunes-duration,omitempty"`
	IExplicit          string `xml:"itunes-explicit,omitempty"`
	IIsClosedCaptioned string `xml:"itunes-isClosedCaptioned,omitempty"`
	IOrder             string `xml:"itunes-order,omitempty"`
}

// Podcast Type
type Podcast struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`
	Link           string   `xml:"link"`
	Description    string   `xml:"description"`
	Category       string   `xml:"category,omitempty"`
	Cloud          string   `xml:"cloud,omitempty"`
	Copyright      string   `xml:"copyright,omitempty"`
	Docs           string   `xml:"docs,omitempty"`
	Generator      string   `xml:"generator,omitempty"`
	Language       string   `xml:"language,omitempty"`
	LastBuildDate  string   `xml:"lastBuildDate,omitempty"`
	ManagingEditor string   `xml:"managingEditor,omitempty"`
	PubDate        string   `xml:"pubDate,omitempty"`
	Rating         string   `xml:"rating,omitempty"`
	SkipHours      string   `xml:"skipHours,omitempty"`
	SkipDays       string   `xml:"skipDays,omitempty"`
	TTL            int      `xml:"ttl,omitempty"`
	WebMaster      string   `xml:"webMaster,omitempty"`
	Image          *Image
	TextInput      *TextInput
	AtomLink       *AtomLink

	// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
	IAuthor     string `xml:"itunes-author,omitempty"`
	ISubtitle   string `xml:"itunes-subtitle,omitempty"`
	ISummary    *ISummary
	IBlock      string `xml:"itunes-block,omitempty"`
	IImage      *IImage
	IDuration   string  `xml:"itunes-duration,omitempty"`
	IExplicit   string  `xml:"itunes-explicit,omitempty"`
	IComplete   string  `xml:"itunes-complete,omitempty"`
	INewFeedURL string  `xml:"itunes-new-feed-url,omitempty"`
	IOwner      *Author // Author is formatted for itunes as-is
	ICategories []*ICategory

	Items []*Item `xml:"item"`
	// contains filtered or unexported fields
}

// TextInput Type
type TextInput struct {
	XMLName     xml.Name `xml:"textInput"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Name        string   `xml:"name"`
	Link        string   `xml:"link"`
}

// RSS Type
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Podcast Podcast  `xml:"channel"`
}

func replaceColon(s string) string {
	return strings.Replace(s, ":", "-", 1)
}

// Handler serverless function entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	// url := "https://anchor.fm/s/119f3bc8/podcast/rss"
	re := regexp.MustCompile(`<(/)?[a-z]*?:`)
	url := "http://localhost:8000/test.xml"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	cleanedBody := []byte(re.ReplaceAllStringFunc(string(body), replaceColon))
	var podcastBody RSS

	err = xml.Unmarshal(cleanedBody, &podcastBody)
	if err != nil {
		fmt.Println(err)
	}

	js, err := json.Marshal(podcastBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
