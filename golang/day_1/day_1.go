package day_1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Run() {
	content, err := ioutil.ReadFile("day_1/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	for _, l1 := range lines {
		if l1 == "" {
			continue
		}

		first, err := strconv.Atoi(l1)
		if err != nil {
			panic(err)
		}

		for _, l2 := range lines {
			if l2 == "" {
				continue
			}

			second, err := strconv.Atoi(l2)
			if err != nil {
				panic(err)
			}

			for _, l3 := range lines {
				if l3 == "" {
					continue
				}

				third, err := strconv.Atoi(l3)
				if err != nil {
					panic(err)
				}

				if first+second+third == 2020 {
					fmt.Println(first * second * third)

					os.Exit(0)
				}
			}

		}
	}
}
