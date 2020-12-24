package params

import (
	"io/ioutil"
	"log"
)

const (
	host         = "localhost:5000"
	Api          = "http://" + host + "/bolt"
	Vkwriterfile = "vkwriter.bot"
	users        = "/usertg/"
	channels     = "/channels/"
	PubNames     = "/pubNames/"
	username     = "/usernametg/"
	pubSubTg     = "/pubSubTg/"
	feedSubTg    = "/feedSubTg/"
	userSubTg    = "/userSubTg/"
	lastPost     = "/vkpublastpost/"
	Feed         = "/feeds/"
	best         = "/best/"
	viralhash    = "/viralhash/"
	links        = "/links/"
	video        = "/video/"
	ShortUrl     = "https://www.googleapis.com/urlshortener/v1/url?key="
	TgApi        = "https://api.telegram.org"
	//MyakotkaId   = -1001067277325
	BaseUri   = Api + "/"
	Publics   = Api + PubNames
	Feeds     = Api + Feed
	Links     = Api + links
	Video     = Api + video
	Users     = Api + users
	Channels  = Api + channels
	Bests     = Api + best
	ViralHash = Api + viralhash
	UserName  = Api + username
	Subs      = Api + pubSubTg
	FeedSubs  = Api + feedSubTg
	UserSubs  = Api + userSubTg
	LastPost  = Api + lastPost
	//YaToken    = "1309689c-a584-41f2-a0e8-299747bb6326"
	Example    = "\nExample: \nhttps://www.reddit.com/r/gifs/top/\nhttps://vk.com/evil_incorparate\n\nMore examples: http://telegra.ph/telefeedbot-05-12\n\n Also you can try this bot: @memefeed"
	SomeErr    = "Something going wrong. Try letter or contact with mainteiner."
	Hello      = "Send me a link on domain/rss.\n\n" + Example
	Psst       = "As soon as there are new articles here - i will send them, but with some delay."
	NotFound   = "Rss feed not found\nPls send me direct link on rss\n\n"
	NewChannel = "Add this as admin in channel\nSend me link on channel, example: https://t.me/channel\n\n"
	SubsHelp   = "Commands:\nAdd url:\n@channelname url(s)\nDelete url(s):\n@channelname delete url"
	Rate       = "Please rate me here ❤️❤️❤️:\nhttps://storebot.me/bot/telefeedbot"
	Donate	   = "Like tihs bot? Buy a coffe to the maintainer! Contact @Azhinu to supoort this project."
	Help       = "Manual: https://telegra.ph/telefeedbot-05-12 \nSupport chat: https://t.me/joinchat/AAAAAEMFJOGkHNVp8qKQ1g"
	TopLinks   = `
	Top rss feeds (by subscribers):

`
)

var (
	Telefeedfile = "./telefeed.bot"
	ChannelsFatherfile = "./channelsfather.bot"
	TokensFile = "./vk_tokens.bot"
	Stores             = [...]string{"@telefeedcontent1", "@telefeedcontent2", "@telefeedcontent3"}
	StoreIds           = [...]int64{-1001140338639, -1001144965998, -1001122084977, -1001121449455, -1001147806509, -1001069985583, -1001128095164}
	GooglKeys          = [...]string{"AIzaSyCTmUsTGqjl7iWJLiJisrejgTNamp7bfIA", "AIzaSyAZaQiAkSmYFZLMjUOxtKOj3R29TPs81X0"}
)

func init() {
	/*
		f, err := os.OpenFile("./testlogfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.SetOutput(f)
		}
		defer f.Close()*/

	log.SetOutput(ioutil.Discard)
}
