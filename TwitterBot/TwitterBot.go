package TwitterBot

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/bluele/mecab-golang"
	"github.com/cotton392/ctn_ai/markov"
	"net/url"
	"os"
	"strconv"
)

func GetTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))

	return api
}

func GetTweetText(username string, tweetCount int) []string {
	res := []string{}
	api := GetTwitterApi()
	values := url.Values{}
	values.Add("screen_name", username)
	values.Add("count", strconv.Itoa(tweetCount))
	values.Add("trim_user", "true")
	values.Add("exclude_replies", "true")
	values.Add("include_rts", "false") // tweet取得に際しての設定

	tweets, err := api.GetUserTimeline(values) // ユーザータイムラインを取得
	if err != nil {
		fmt.Printf("Tweet get error.")
		os.Exit(-1)
	}
	for _, s := range tweets {
		res = append(res, s.FullText)
	} // resにツイート本文を追加

	return res
}

func TweetText() {
	api := GetTwitterApi()
	tweets := GetTweetText("cotton392", 30)
	markovBlocks := [][]string{}
	m, err := mecab.New("-Owakati")
	if err != nil {
		fmt.Printf("Mecab instance error. err: %v", err)
	}
	defer m.Destroy()

	for _, tweet := range tweets {
		_data := markov.ParseToNode(m, tweet)
		elems := markov.GetMarkovBlocks(_data)
		markovBlocks = append(markovBlocks, elems...)
	}

	tweetElemSet := markov.MarkovChainExec(markovBlocks)
	text := markov.TextGenerate(tweetElemSet)
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("-----------------------------------")
	fmt.Println(tweet.Text)
}
