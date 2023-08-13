package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Cocktail struct {
	ID int `json:id`
	Name string `json:name`
	Ingredients []string `json:ingredients`
	Description string `json:description`
}

type Ingredient struct {
	ID int `json:id`
	Name string `json:name`
	Description string `json:description`
	IsAlcohol bool `json:isAlcohol`
}

var pool *pgxpool.Pool

func getAllIngredients() []Ingredient {
	fmt.Println("getAllIngredients")
	rows, err := pool.Query(context.Background(), "select * from ingredients;")
	if err != nil {
		fmt.Printf("failed to getAllIngredients")
		return nil
	}

	ingredients := []Ingredient{}
	for rows.Next(){
		var ing Ingredient
		err := rows.Scan(&ing.ID, &ing.Name, &ing.Description, &ing.IsAlcohol,)
		if err != nil {
			fmt.Println("failed to scan data")
			fmt.Println(err)
			break
		}
		ingredients = append(ingredients, ing)
	}
	return ingredients
}

func getAllCocktails() []Cocktail {
	fmt.Println("getAllCocktails")
	rows, err := pool.Query(context.Background(), "select * from cocktails;")
	if err != nil {
		fmt.Printf("failed to getAllCocktails")
		return nil
	}

	cocktails := []Cocktail{}
	for rows.Next(){
		var c Cocktail
		err := rows.Scan(&c.ID, &c.Name, &c.Ingredients, &c.Description,)
		if err != nil {
			fmt.Println("failed to scan data")
			fmt.Println(err)
			break
		}
		cocktails = append(cocktails, c)
	}
	return cocktails
}

func responseAllIngredients(ctx *gin.Context) {
	var ingredients []Ingredient
	ingredients = getAllIngredients()
	ctx.IndentedJSON(http.StatusOK, ingredients)
}

func responseAllCocktails(ctx *gin.Context) {
	var cocktails []Cocktail
	cocktails = getAllCocktails()
	ctx.IndentedJSON(http.StatusOK, cocktails)
}

func main() {
	fmt.Println("Start Program")
	connConfig, err := pgx.ParseConfig(DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse db config")
		os.Exit(1)
	}

	poolConfig, err := pgxpool.ParseConfig("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse pool config")
		os.Exit(1)
	}
	poolConfig.ConnConfig = connConfig

	pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to db")
		os.Exit(1)
	}

	defer pool.Close()

	router := gin.Default()
	router.GET("/ingredients", responseAllIngredients)
	router.GET("/cocktails", responseAllCocktails)
	router.Run("localhost:9090")
}