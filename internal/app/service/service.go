package service

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Service struct {
}

type QuoteRu struct {
	Date, Author, Name, Quote string
}

type QuoteEn struct {
	Date, Author, Name, Quote string
}

func New() *Service {
	return &Service{}
}

func (s *Service) Output() (interface{}, interface{}, error) {
	db, errSQL := sql.Open("mysql", "user:user@tcp(db:3306)/db_quotes")
	if errSQL != nil {
		panic(errSQL)
	}
	defer db.Close()

	selRu, err := db.Query("SELECT * FROM `quotes`")
	if err != nil {
		panic(err)
	}
	defer selRu.Close()

	selEn, err := db.Query("SELECT * FROM `quotesEn`")
	if err != nil {
		panic(err)
	}
	defer selEn.Close()

	var quoteArrayRu = []QuoteRu{}
	var quoteArrayEn = []QuoteEn{}

	for selRu.Next() {
		var quote QuoteRu
		err := selRu.Scan(&quote.Date, &quote.Author, &quote.Name, &quote.Quote)
		if err != nil {
			panic(err)
		}
		quoteArrayRu = append(quoteArrayRu, quote)
	}

	for selEn.Next() {
		var quote QuoteEn
		err := selEn.Scan(&quote.Date, &quote.Author, &quote.Name, &quote.Quote)
		if err != nil {
			panic(err)
		}
		quoteArrayEn = append(quoteArrayEn, quote)
	}

	return quoteArrayRu, quoteArrayEn, nil
}

func (s *Service) Save(author, name, quote, lang string) error {
	db, err := sql.Open("mysql", "user:user@tcp(db:3306)/db_quotes")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	date := time.Now()

	if lang == "ru" {
		_, err = db.Exec(fmt.Sprintf("INSERT INTO `quotes` (`date`,`author`,`name`,`quote`) VALUES ('%s','%s','%s','%s')", date.Format("01-02-2006 15:04:05"), author, name, quote))
		if err != nil {
			panic(err)
		}
	}

	if lang == "en" {
		_, err = db.Exec(fmt.Sprintf("INSERT INTO `quotesEn` (`date`,`author`,`name`,`quote`) VALUES ('%s','%s','%s','%s')", date.Format("01-02-2006 15:04:05"), author, name, quote))
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (s *Service) ValidateLanguage(author, name, quote string) (string, error) {
	var lang string

	en := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	ru := "йцукенгшщзхъфывапролджэёячсмитьбюЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЁЯЧСМИТЬБЮ"

	if strings.Contains(en, string(author[0])) && strings.Contains(en, string(name[0])) && strings.Contains(en, string(quote[0])) {
		lang = "en"
	} else if strings.Contains(ru, string([]rune(author)[0:1])) && strings.Contains(ru, string([]rune(name)[0:1])) && strings.Contains(ru, string([]rune(quote)[0:1])) {
		lang = "ru"
	} else {
		lang = "-"
	}
	return lang, nil
}

func (s *Service) ValidateAdmin(login, password string) (bool, error) {
	if login == "admin" && password == "0000" {
		return true, nil
	}
	return false, nil
}

func (s *Service) Delete(author, quote string) error {
	db, errSQL := sql.Open("mysql", "user:user@tcp(db:3306)/db_quotes")
	if errSQL != nil {
		panic(errSQL)
	}
	defer db.Close()

	_, errDel := db.Exec(fmt.Sprintf("DELETE FROM `quotes` WHERE `author`='%s' AND `quote`='%s'", author, quote))
	if errDel != nil {
		panic(errDel)
	}

	_, errDel = db.Exec(fmt.Sprintf("DELETE FROM `quotesEn` WHERE `author`='%s' AND `quote`='%s'", author, quote))
	if errDel != nil {
		panic(errDel)
	}

	return nil
}
