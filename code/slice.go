package main

import "fmt"

func main() {
	languages := []string{"python", "ruby", "javascript"}
	languages = append(languages, "go")

	fmt.Println("Linguagens: ", languages)
	fmt.Println("Linguagens: ", languages[1:3])

	for i, linguagem := range languages {
		fmt.Printf(
			"Linguagem na posição %d do slice, valor: %s\n",
			i,
			linguagem,
		)
	}
}
