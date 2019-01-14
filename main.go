package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info) // Usamos Hash256 para simplificar el cálculo hash de momento
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1] // Obtenemos el último bloque de la cadena
	newBlock := CreateBlock(data, prevBlock.Hash)  // Creamos el nuevo bloque
	chain.blocks = append(chain.blocks, newBlock)  // Añadimos el nuevo bloque al final de la cadena
}

// Función encargada de generar el primer bloque de la cadena de bloques (llamado Génesis)
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Función encargada de generar una nueva cadena de bloques a partir del bloque Génesis
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {

	chain := InitBlockChain() // Creamos nuestra cadena de bloques

	// Añadimos 3 nuevos bloques a nuestra cadena de bloques
	chain.AddBlock("Primer bloque despues del bloque Genesis")
	chain.AddBlock("Segundo bloque despues del bloque Genesis")
	chain.AddBlock("Tercer bloque despues del bloque Genesis")

	// Si modificamos por ejemplo el data del 2º bloque poneindo la s en minúscula, su hash cambia y por lo tanto al comparar sabriamos que esta actualizaciñon es inválida en caso de venir de nuestros nodos.
	//chain.AddBlock("Primer bloque despues del bloque Genesis")
	//chain.AddBlock("segundo bloque despues del bloque Genesis")
	//chain.AddBlock("Tercer bloque despues del bloque Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Hash anterior: %x\n", block.PrevHash)
		fmt.Printf("Datos en el bloque: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("")

	}

}
