package dao

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/oklog/ulid/v2"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type UserNameFromHTTPPost struct {
	Name string `json:"name"`
}
type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

//type User struct {
//	Id    string `json:"id"`
//	Name  string `json:"name"`
//	Point int    `json:"age"`
//}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// jsonのdecode（復号）
		var decodedJSON UserNameFromHTTPPost //復号内容を入れる構造体
		//.Decodeの返り値はerr。それを入れる
		if err := json.NewDecoder(r.Body).Decode(&decodedJSON); err != nil { //errorが入っている場合エラーメッセージ
			log.Printf("fail: decoding json in method POST, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//log.Printf("nam =%s, age = %d\n", decodedJSON.Name, decodedJSON.Age)

		if decodedJSON.Name == "" { //nameが空文字だった場合
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("name is empty")
			return
		} else //長さが50を超えるような長い文字列が渡された場合
		if len(decodedJSON.Name) > 50 {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("name to long")
			return
		} else {
			// ULID (Universally Unique Lexicographically Sortable Identifier) を作成
			timeNow := time.Now()
			entropy := ulid.Monotonic(rand.New(rand.NewSource(timeNow.UnixNano())), 0)
			id := ulid.MustNew(ulid.Timestamp(timeNow), entropy)
			tx, err := db.Begin()
			if err != nil {
				log.Printf("fail: db.Begin(), %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			//INSERTを実行
			if _, err = tx.Exec(
				"INSERT INTO users (id, name) "+
					"VALUES (?,?)",
				id.String(),
				decodedJSON.Name,
			); err != nil {
				log.Printf("fail: tx.Exec():INSERT, %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
				if err = tx.Rollback(); err != nil {
					log.Printf("fail: tx.Rollback(), %v\n", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				} else {
					if err = tx.Commit(); err != nil {
						log.Printf("fail: tx.Commit(), %v\n", err)
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
				}

				return
			} else {
				w.WriteHeader(http.StatusOK)
				log.Printf(`{ "id" : "%s" }`, id.String())
				w.Header().Set("Content-Type", "application/json")
				var u UserResForHTTPGet
				u.Id = id.String()
				u.Name = decodedJSON.Name

				// レスポンス用ユーザースライスをJSONへ変換→HTTPレスポンスボディに書き込み
				bytes, err := json.Marshal(u)
				if err != nil {
					log.Printf("fail: json.Marshal, %v\n", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Write(bytes)
			}
		}
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return

	}

}

func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}

func loadEnv() {
	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")

	// もし err がnilではないなら、"failed: loading .env:"が出力されます。
	if err != nil {
		fmt.Printf("failed: loading .env: %v", err)
	}

}
