package main

import (
	"strconv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	db "crawler2/database"
	_ "github.com/go-sql-driver/mysql"
)


type Crawlerdata struct {
	Data []Data `json:"data"`
}

type Data struct {
	Attr Attributes `json:"attributes"`
}

type Attributes struct {
	id			int		`json:"id"`
	Rank		int     `json:"rank"`
	Address		string  `json:"holder"`
	Amount		string  `json:"amount"`
	Percentage	float64 `json:"percent"`
}

func parseUrls(url string, ch chan bool) Crawlerdata {
	fmt.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Http get err:", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Http status code:", resp.StatusCode)
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	cd := Crawlerdata{}
	json.Unmarshal(raw, &cd)
	fmt.Println(cd.Data[0].Attr.Address)
	return cd
}

func main() {
	var cd2 []Attributes
	var id int64

	defer db.SqlDB.Close()

	ch := make(chan bool)
	for j := 0; j < 10; j++ {
		j0 := j
		for i := 0; i < 5; i++ {
			i0 := i
			go func() {
				crawler := parseUrls("https://scopi.wemixnetwork.com/api/v1/chain/1003/token/0x3e61e305390bb45983879f318ee6d0d5d2805e65/holder?pagesize=20&page="+strconv.Itoa(1+i0+5*j0), ch)
				for c := 0; c < 20; c++ {
					c0 := c
					result , err := db.SqlDB.Exec("INSERT INTO users(numb, address, amount, percentage) VALUES(?, ?, ?, ?)", crawler.Data[c0].Attr.Rank, crawler.Data[c0].Attr.Address, crawler.Data[c0].Attr.Amount, crawler.Data[c0].Attr.Percentage)
					if err != nil {
						log.Fatal("insert:", err)
						return
					}
					id, err = result.LastInsertId()
					if err != nil {
						log.Fatalln(err)
					}
				}
				ch <- true
			}()
		}
		for coun := 0; coun < 5; coun++ {
			<-ch
		}
	}
	rows, err := db.SqlDB.Query("SELECT id, numb, address, amount, percentage FROM users")
	if err != nil {
		log.Println("query:", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		attr := Attributes{}
		err = rows.Scan(&attr.id, &attr.Rank, &attr.Address, &attr.Amount, &attr.Percentage)
		if err != nil {
			log.Println("scan:", err)
			return
		}
		fmt.Println("get data:", attr)
		cd2 = append(cd2, attr)
	}

	router := initRouter()
	router.Run(":8080")
}