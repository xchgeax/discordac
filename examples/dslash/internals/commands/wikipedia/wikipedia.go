package wikipedia

import (
	"encoding/xml"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/vlaetansky/discordslash"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var Command = &discordslash.SlashedCommand{
	Specification: &discordgo.ApplicationCommand{
		Name:        "wikipedia",
		Description: "Search wikipedia",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "query",
				Description: "What to search",
				Required:    true,
			},
		},
	},
	Handler: func(cc *discordslash.CommandContext) {
		if query, ok := cc.GetOption("query"); ok {
			cc.Respond("Searching...")
			go findWikipedia(query.StringValue(), cc)
		}
	},
}

func findWikipedia(query string, cc *discordslash.CommandContext) {
	resp, err := http.Get(
		fmt.Sprintf("https://en.wikipedia.org/w/api.php?format=xml&action=query&prop=extracts&exsentences=7&exintro&explaintext&exsectionformat=plain&redirects=1&titles=%v", url.QueryEscape(query)))
	if err != nil {
		logrus.WithError(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var r Response
	if err := xml.Unmarshal(body, &r); err != nil {
		logrus.WithError(err).Error("Couldn't parse XML response")
	}

	text := r.Query.Pages.Page.Extract.Text
	if text != "" {
		err = cc.EditResponse(text)
	} else {
		err = cc.EditResponse(fmt.Sprintf("%v: no results.", query))
	}
	if err != nil {
		logrus.WithError(err)
		return
	}
}

type Response struct {
	XMLName       xml.Name `xml:"api"`
	Text          string   `xml:",chardata"`
	Batchcomplete string   `xml:"batchcomplete,attr"`
	Query         struct {
		Text      string `xml:",chardata"`
		Redirects struct {
			Text string `xml:",chardata"`
			R    struct {
				Text string `xml:",chardata"`
				From string `xml:"from,attr"`
				To   string `xml:"to,attr"`
			} `xml:"r"`
		} `xml:"redirects"`
		Pages struct {
			Text string `xml:",chardata"`
			Page struct {
				Text    string `xml:",chardata"`
				Idx     string `xml:"_idx,attr"`
				Pageid  string `xml:"pageid,attr"`
				Ns      string `xml:"ns,attr"`
				Title   string `xml:"title,attr"`
				Extract struct {
					Text  string `xml:",chardata"`
					Space string `xml:"space,attr"`
				} `xml:"extract"`
			} `xml:"page"`
		} `xml:"pages"`
	} `xml:"query"`
}
