package main

import (
	"database/sql"
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

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
)

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type UserInputFromHTTPPost struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ① GoプログラムからMySQLへ接続
var db *sql.DB

func init() {
	loadEnv()
	// ①-1 connecting to the test_database
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	connStr := fmt.Sprintf("%s:%s@(%s)/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	log.Println(connStr)
	log.Printf(string(connStr))
	// ①-2 opening sql
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// ①-3 checking connection
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// ②-1 get "name" from http.Request
		name := r.URL.Query().Get("name") // To be filled
		if name == "" {
			log.Println("fail: name is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// ②-2 GETリクエストのクエリパラメータから条件を満たすデータを取得
		rows, err := db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// ②-3
		users := make([]UserResForHTTPGet, 0)
		for rows.Next() {
			var u UserResForHTTPGet
			if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)

				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		// ②-4 レスポンス用ユーザースライスをJSONへ変換→HTTPレスポンスボディに書き込み
		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)

	case http.MethodPost:
		// jsonのdecode（復号）
		var decodedJSON UserInputFromHTTPPost //復号内容を入れる構造体
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
		} else //ageに20〜80の範囲にない値が渡された場合
		if decodedJSON.Age < 20 || decodedJSON.Age > 80 {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("age is under 20 or over 80")
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
				"INSERT INTO user (id, name, age) "+
					"VALUES (?,?,?)",
				id.String(),
				decodedJSON.Name,
				decodedJSON.Age); err != nil {
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
				u.Age = decodedJSON.Age

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

func main() {

	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", handler)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
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
