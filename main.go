package main

import (
	"./blockchain"
	"fmt"
)

func main() {

	chain := blockchain.InitBlockChain() // Creamos nuestra cadena de bloques

	// Añadimos 3 nuevos bloques a nuestra cadena de bloques
	chain.AddBlock("Primer bloque despues del bloque Genesis")
	chain.AddBlock("Segundo bloque despues del bloque Genesis")
	chain.AddBlock("Tercer bloque despues del bloque Genesis")

	// Si modificamos por ejemplo el data del 2º bloque poniendo la s en minúscula, su hash cambia y por lo tanto al comparar sabriamos que esta actualizaciñon es inválida en caso de venir de nuestros nodos.
	//chain.AddBlock("Primer bloque despues del bloque Genesis")
	//chain.AddBlock("segundo bloque despues del bloque Genesis")
	//chain.AddBlock("Tercer bloque despues del bloque Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Hash anterior: %x\n", block.PrevHash)
		fmt.Printf("Datos en el bloque: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("")
	}

}
