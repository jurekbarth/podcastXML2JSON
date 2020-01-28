package handler

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// AtomLink Type
type AtomLink struct {
	XMLName xml.Name `xml:"atom-link" json:"-"`
	HREF    string   `xml:"href,attr" json:"href"`
	Rel     string   `xml:"rel,attr" json:"rel"`
	Type    string   `xml:"type,attr" json:"type"`
}

// Author Type
type Author struct {
	XMLName xml.Name `xml:"itunes-owner" json:"-"`
	Name    string   `xml:"itunes-name" json:"itunesName"`
	Email   string   `xml:"itunes-email" json:"itunesEmail"`
}

// Enclosure Type
type Enclosure struct {
	XMLName xml.Name `xml:"enclosure" json:"-"`
	URL     string   `xml:"url,attr" json:"url"`
	Length  string   `xml:"length,attr" json:"length"`
	Type    string   `xml:"type,attr" json:"type"`
}

// ICategory Type
type ICategory struct {
	XMLName     xml.Name     `xml:"itunes-category" json:"-"`
	Text        string       `xml:"text,attr" json:"text"`
	ICategories []*ICategory `xml:"itunes-category" json:"subCategories,omitempty"`
}

// IImage Type
type IImage struct {
	XMLName xml.Name `xml:"itunes-image" json:"-"`
	HREF    string   `xml:"href,attr" json:"href"`
}

// ISummary Type
type ISummary struct {
	XMLName xml.Name `xml:"itunes-summary" json:"-"`
	Text    string   `xml:",cdata" json:"text"`
}

// Image Type
type Image struct {
	XMLName     xml.Name `xml:"image" json:"-"`
	URL         string   `xml:"url" json:"url"`
	Title       string   `xml:"title" json:"title"`
	Link        string   `xml:"link" json:"link"`
	Description string   `xml:"description,omitempty" json:"description,omitempty"`
	Width       int      `xml:"width,omitempty" json:"width,omitempty"`
	Height      int      `xml:"height,omitempty" json:"height,omitempty"`
}

// Item Type
type Item struct {
	XMLName          xml.Name   `xml:"item" json:"-"`
	GUID             string     `xml:"guid" json:"guid"`
	Title            string     `xml:"title" json:"title"`
	Link             string     `xml:"link" json:"link"`
	Description      string     `xml:"description" json:"description"`
	Author           string     `xml:"author,omitempty" json:"author,omitempty"`
	Category         string     `xml:"category,omitempty" json:"category,omitempty"`
	Comments         string     `xml:"comments,omitempty" json:"comments,omitempty"`
	Source           string     `xml:"source,omitempty" json:"source,omitempty"`
	PubDateFormatted string     `xml:"pubDate,omitempty" json:"publishDate,omitempty"`
	Enclosure        *Enclosure `json:"enclosure"`

	IAuthor            string    `xml:"itunes-author,omitempty" json:"itunesAuthor,omitempty"`
	ISubtitle          string    `xml:"itunes-subtitle,omitempty" json:"itunesSubtitle,omitempty"`
	ISummary           *ISummary `json:"itunesSummary,omitempty"`
	IImage             *IImage   `json:"itunesImage,omitempty"`
	IDuration          string    `xml:"itunes-duration,omitempty" json:"itunesDuration,omitempty"`
	IExplicit          string    `xml:"itunes-explicit,omitempty" json:"itunesExplicit,omitempty"`
	IIsClosedCaptioned string    `xml:"itunes-isClosedCaptioned,omitempty" json:"itunesIsClosedCaptioned,omitempty"`
	IOrder             string    `xml:"itunes-order,omitempty" json:"itunesOrder,omitempty"`
	ISeason            string    `xml:"itunes-season,omitempty" json:"itunesSeason,omitempty"`
	IEpisode           string    `xml:"itunes-episode,omitempty" json:"itunesEpisode,omitempty"`
	IEpisodeType       string    `xml:"itunes-episodeType,omitempty" json:"itunesEpisodeType,omitempty"`
}

// Podcast Type
type Podcast struct {
	XMLName        xml.Name   `xml:"channel" json:"-"`
	Title          string     `xml:"title" json:"title"`
	Link           string     `xml:"link" json:"link"`
	Description    string     `xml:"description" json:"description"`
	Category       string     `xml:"category,omitempty" json:"category,omitempty"`
	Cloud          string     `xml:"cloud,omitempty" json:"cloud,omitempty"`
	Copyright      string     `xml:"copyright,omitempty" json:"copyright,omitempty"`
	Docs           string     `xml:"docs,omitempty" json:"docs,omitempty"`
	Generator      string     `xml:"generator,omitempty" json:"generator,omitempty"`
	Language       string     `xml:"language,omitempty" json:"language,omitempty"`
	LastBuildDate  string     `xml:"lastBuildDate,omitempty" json:"lastBuildDate,omitempty"`
	ManagingEditor string     `xml:"managingEditor,omitempty" json:"managingEditor,omitempty"`
	PubDate        string     `xml:"pubDate,omitempty" json:"publishDate,omitempty"`
	Rating         string     `xml:"rating,omitempty" json:"rating,omitempty"`
	SkipHours      string     `xml:"skipHours,omitempty" json:"skipHours,omitempty"`
	SkipDays       string     `xml:"skipDays,omitempty" json:"skipDays,omitempty"`
	TTL            int        `xml:"ttl,omitempty" json:"timeToLive,omitempty"`
	WebMaster      string     `xml:"webMaster,omitempty" json:"webMaster,omitempty"`
	Image          *Image     `json:"image,omitempty"`
	TextInput      *TextInput `json:"textInput,omitempty"`
	AtomLink       *AtomLink  `json:"atomLink,omitempty"`

	// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
	IAuthor     string       `xml:"itunes-author,omitempty" json:"itunesAuthor,omitempty"`
	ISubtitle   string       `xml:"itunes-subtitle,omitempty" json:"itunesSubtitle,omitempty"`
	ISummary    *ISummary    `json:"itunesSummary,omitempty"`
	IBlock      string       `xml:"itunes-block,omitempty" json:"itunesBlock,omitempty"`
	IImage      *IImage      `json:"itunesImage,omitempty"`
	IDuration   string       `xml:"itunes-duration,omitempty" json:"itunesDuration,omitempty"`
	IExplicit   string       `xml:"itunes-explicit,omitempty" json:"itunesExplicit,omitempty"`
	IComplete   string       `xml:"itunes-complete,omitempty" json:"itunesComplete,omitempty"`
	INewFeedURL string       `xml:"itunes-new-feed-url,omitempty" json:"itunesNewFeedUrl,omitempty"`
	IOwner      *Author      `json:"itunesOwner,omitempty"`
	ICategories []*ICategory `xml:"itunes-category,omitempty" json:"itunesCategories,omitempty"`

	Items []*Item `xml:"item" json:"items"`
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
	XMLName xml.Name `xml:"rss" json:"-"`
	Podcast Podcast  `xml:"channel" json:"podcast"`
}

func replaceColon(s string) string {
	return strings.Replace(s, ":", "-", 1)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handler serverless function entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	// url := "https://anchor.fm/s/119f3bc8/podcast/rss"
	// url := "http://localhost:8000/test.xml"
	feedURL := r.URL.Query().Get("feed")
	if feedURL == "" {
		handleError(w, errors.New("feed param empty"))
	}
	re := regexp.MustCompile(`<(/)?[a-z]*?:`)

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, feedURL, nil)

	handleError(w, err)

	res, err := client.Do(req)
	handleError(w, err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	handleError(w, err)

	// Fixes golang xml namespace issue: https://github.com/golang/go/issues/8535
	cleanedBody := []byte(re.ReplaceAllStringFunc(string(body), replaceColon))

	var podcastBody RSS
	err = xml.Unmarshal(cleanedBody, &podcastBody)
	handleError(w, err)

	js, err := json.Marshal(podcastBody.Podcast)
	handleError(w, err)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Write(js)
}
