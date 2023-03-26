package webhook

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func PostWebhook(webhook_url string) {
	// 送信するデータを作成
	values := url.Values{}
	values.Add("payload", "{\"text\": \"Hello, World!\"}")

	// POST送信
	resp, err := http.PostForm(webhook_url, values)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// レスポンスを表示
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
