package main

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {
	connect_env()
	env_map := read_env()

	config := oauth1.NewConfig(env_map["CONSUMER_KEY"], env_map["CONSUMER_SECRET_KEY"])
	token := oauth1.NewToken(env_map["TOKEN_KEY"], env_map["TOKEN_SECRET_KEY"])

	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	tweet := publish_tweet("teste", client)
	log.Println(tweet.Text)

	tweets := find_tweets_by_query("#desenvolvimento", 1, client)
	for _, val := range tweets.Statuses {
		log.Print("User: ", val.User.Name)
		log.Print("Tweet:", val.Text)

		retweet_tweet(val.ID, client)
	}
}

//Carrega o arquivo .env
func connect_env() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

//Lê o arquivo .env
func read_env() map[string]string {
	env_map, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err)
	}

	return env_map
}

//Encontrar um numero n de tweets a partir de uma query
func find_tweets_by_query(query string, count int8, client *twitter.Client) *twitter.Search {
	tweets, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "#desenvolvimento",
		Count: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	return tweets
}

//Retweetar um tweet a partir de seu ID
func retweet_tweet(tweet_id int64, client *twitter.Client) {
	_, _, err := client.Statuses.Retweet(tweet_id, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//Publicar um tweet
func publish_tweet(text string, client *twitter.Client) *twitter.Tweet {
	tweet, _, err := client.Statuses.Update(text, nil)
	if err != nil {
		log.Fatal(err)
	}

	return tweet
}
