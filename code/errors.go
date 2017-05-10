package main

import (
	"errors"
	"fmt"
)

var ErrEmptyData = errors.New("The data to create is empty")

func createOnDatabase(dataToCreate string) error {
	if dataToCreate == "" {
		return ErrEmptyData
	}

	return nil
}

func main() {
	if err := createOnDatabase(""); err != nil {
		fmt.Println("Falha ao criar no banco: ", err)
		return
	}

	fmt.Println("Dados salvos com sucesso!")
}
