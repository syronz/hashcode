package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"pizza/pkg"
	"strconv"
	"strings"
)

const (
	DEPTH = 2
)

var Ings map[string]pkg.Ingredient
var Pizzas map[int]pkg.Pizza

func main() {

	Ings = make(map[string]pkg.Ingredient)
	ingCounter := 0

	Pizzas = make(map[int]pkg.Pizza)

	// file, err := os.Open("a_example")
	file, err := os.Open("b_little_bit_of_everything.in")
	// file, err := os.Open("./c_many_ingredients.in")
	// file, err := os.Open("./d_many_pizzas.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var n, t2, t3, t4 int
	scanner.Scan()
	firstLine := scanner.Text()
	arrStr := strings.Split(firstLine, " ")
	n, err = strconv.Atoi(arrStr[0])
	Handle(err)
	t2, err = strconv.Atoi(arrStr[1])
	Handle(err)
	t3, err = strconv.Atoi(arrStr[2])
	Handle(err)
	t3, err = strconv.Atoi(arrStr[3])
	Handle(err)

	PizzaArr := make([]pkg.Pizza, n)

	fmt.Println(">>>", n, t2, t3, t4)

	var lineNumber int
	for scanner.Scan() {
		line := scanner.Text()

		ingredients := strings.Split(line, " ")

		pizzaIngQty, errQty := strconv.Atoi(ingredients[0])
		if errQty != nil {
			log.Fatal("Error in getting the qty for pizza's gredient", err)
		}

		pizzaBin := big.NewInt(0)
		for i := 1; i < len(ingredients); i++ {
			// fmt.Println(".....", ingredients[i])
			ing, ok := Ings[ingredients[i]]
			var ingBin *big.Int
			if ok {
				ing.Counter++
				ingBin = ing.Num
				Ings[ingredients[i]] = ing
			} else {
				ingBin = pkg.Pow(2, ingCounter)
				tmp := pkg.Ingredient{
					Num:     ingBin,
					Counter: 1,
				}
				Ings[ingredients[i]] = tmp
				ingCounter++
			}

			pizzaBin.Add(pizzaBin, ingBin)
		}

		pizza := pkg.Pizza{
			Bin:    pizzaBin,
			IngQty: pizzaIngQty,
			Index:  lineNumber,
		}

		PizzaArr[lineNumber] = pizza

		Pizzas[lineNumber] = pizza
		lineNumber++

		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(Ings)
	// fmt.Println(Pizzas)
	// printPizzas(Pizzas)
	printIngredients(Ings)
	pkg.QuickSort(PizzaArr, 0, len(PizzaArr)-1)
	printPizzaArr(PizzaArr)

}

func printPizzas(pizzas map[int]pkg.Pizza) {
	for k, v := range pizzas {
		fmt.Printf(">>>> %v : [%b, qty: %v]\n", k, v.Bin, v.IngQty)
	}
}

func printIngredients(ings map[string]pkg.Ingredient) {
	for k, v := range ings {
		fmt.Printf("- %v : [%b, Counter: %v]\n", k, v.Num, v.Counter)
	}
}

func printPizzaArr(pizzas []pkg.Pizza) {
	for k, v := range pizzas {
		fmt.Printf(">>>> %v : [%b, qty: %v, line: %v]\n", k, v.Bin, v.IngQty, v.Index+2)
	}
}

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
