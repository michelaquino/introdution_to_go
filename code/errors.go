package main

import (
	"errors"
	"fmt"
)

var ErrEmptyData = errors.New("The data to save is empty")

func saveOnDatabase(dataToCreate string) error {
	if dataToCreate == "" {
		return ErrEmptyData
	}

	return nil
}

func main() {
	if err := saveOnDatabase(""); err != nil {
		fmt.Println("Falha ao salvar no banco de dados: ", err)
		return
	}

	fmt.Println("Dados salvos com sucesso!")
}
