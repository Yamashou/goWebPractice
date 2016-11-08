package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
      <html>
        <head>
          <titleチャット</title>
        </head>
        <body>
          チャットしましょう！
        </body>
      </html>

      `))
	})
	//web サーバーを開始します
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
