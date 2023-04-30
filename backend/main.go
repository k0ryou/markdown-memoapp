package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/k0ryou/markdown-parser-go/generator"
	"github.com/k0ryou/markdown-parser-go/lexer"
	"github.com/k0ryou/markdown-parser-go/parser"
	"github.com/k0ryou/markdown-parser-go/token"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type Memo struct {
	bun.BaseModel `bun:"table:memos,alias:t"`

	ID        int64  `bun:"id,pk,autoincrement"`
	Title     string `bun:"title,notnull"`
	Content   string `bun:"content,notnull"`
	CreatedAt time.Time
	UpdatedAt time.Time `bun:",nullzero"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
}

var db *bun.DB
var e *echo.Echo

func main() {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DBNAME")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)

	sqldb, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	defer sqldb.Close()

	db = bun.NewDB(sqldb, mysqldialect.New())
	defer db.Close()

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	ctx := context.Background()
	_, err = db.NewCreateTable().Model((*Memo)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	e = echo.New()

	e.GET("/api/memos", getMemoList)
	e.GET("/api/memos/:id", getMemo)
	e.POST("/api/memos", postMemo)
	e.PUT("/api/memos/:id", updateMemo)
	e.DELETE("/api/memos/:id", deleteMemo)
	e.Logger.Fatal(e.Start(":8989"))
}

func getMemoList(c echo.Context) error {
	isMd2Html := c.QueryParam("md2html") == "true"
	var memos []Memo
	ctx := context.Background()
	err := db.NewSelect().Model(&memos).Order("created_at").Scan(ctx)
	if err != nil {
		e.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, []error{errors.New("Cannot get memos")})
	}

	if isMd2Html { // md->html
		for idx, memo := range memos {
			memos[idx].Content = convertToHTMLString(memo.Content)
		}
	}

	return c.JSON(http.StatusOK, memos)
}

func getMemo(c echo.Context) error {
	isMd2Html := c.QueryParam("md2html") == "true"
	memoId := c.Param("id")
	var memo Memo
	ctx := context.Background()
	err := db.NewSelect().Model(&memo).Where("id = ?", memoId).Order("created_at").Scan(ctx)
	if err != nil {
		e.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, []error{errors.New("Cannot get memo")})
	}

	if isMd2Html {
		memo.Content = convertToHTMLString(memo.Content)
	}
	return c.JSON(http.StatusOK, memo)
}

func postMemo(c echo.Context) error {
	memo := new(Memo)
	if err := c.Bind(memo); err != nil {
		e.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, []error{errors.New("Cannot bind data")})
	}

	if memo.Title == "" {
		return c.JSON(http.StatusBadRequest, []error{errors.New("Title not found")})
	}

	if memo.Content == "" {
		return c.JSON(http.StatusBadRequest, []error{errors.New("Content not found")})
	}

	ctx := context.Background()
	_, err := db.NewInsert().Model(memo).Exec(ctx)
	if err != nil {
		e.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, []error{errors.New("Cannot insert memo")})
	}

	return c.JSON(http.StatusCreated, "")
}

func updateMemo(c echo.Context) error {
	memoId := c.Param("id")
	memo := new(Memo)
	if err := c.Bind(memo); err != nil {
		e.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, []error{errors.New("Cannot bind data")})
	}

	if memo.Title == "" {
		return c.JSON(http.StatusBadRequest, []error{errors.New("Title not found")})
	}

	if memo.Content == "" {
		return c.JSON(http.StatusBadRequest, []error{errors.New("Content not found")})
	}

	memo.Content = convertToHTMLString(memo.Content)
	ctx := context.Background()
	_, err := db.NewUpdate().Model(memo).Where("id = ?", memoId).Exec(ctx)
	if err != nil {
		e.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, []error{errors.New("Cannot update memo")})
	}

	return c.JSON(http.StatusNoContent, "")
}

func deleteMemo(c echo.Context) error {
	memoId := c.Param("id")
	var memo Memo
	ctx := context.Background()
	_, err := db.NewDelete().Model(&memo).Where("id = ?", memoId).Exec(ctx)
	if err != nil {
		e.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, []error{errors.New("Cannot delete memo")})
	}

	return c.JSON(http.StatusNoContent, "")
}

func convertToHTMLString(markdown string) string {
	mdArray := lexer.Analize(markdown)
	var asts = [][]token.Token{}
	for _, md := range mdArray {
		asts = append(asts, parser.Parse(md))
	}
	htmlString := generator.Generate(asts)
	return htmlString
}
