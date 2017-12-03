package data

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/novatrixtech/silva-backend/model"
)

//LoadDataFromFiles append data from CSV files
//Our Data Collector and Cleaner
func LoadDataFromFiles(fileName string) (data []model.Word) {
	arquivo, err := os.Open(fileName)
	if err != nil {
		log.Println("[LoadDataFromFiles] Error when opening file: ", err.Error())
		return
	}
	defer arquivo.Close()

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		log.Println("[LoadDataFromFiles] Error when reading file. Error: ", err.Error())
		return
	}
	var matter, text string
	for _, linha := range conteudo {
		for indiceItem, item := range linha {
			if indiceItem == 0 {
				matter = item
			} else if indiceItem == 1 {
				text = item
			}
		}
		word := EvaluateWord(text, matter)
		//log.Printf("Linha[%d] é: matter=%s e word=%+v\r\n", indiceLinha, word.Matter.Name, word)
		data = append(data, word)
	}
	return
}
