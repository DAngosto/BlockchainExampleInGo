package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

type BlockChain struct {
	Blocks []*Block
}

/*
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info) // Usamos Hash256 para simplificar el cálculo hash de momento
	b.Hash = hash[:]
}
*/

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}

	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	//block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1] // Obtenemos el último bloque de la cadena
	newBlock := CreateBlock(data, prevBlock.Hash)  // Creamos el nuevo bloque
	chain.Blocks = append(chain.Blocks, newBlock)  // Añadimos el nuevo bloque al final de la cadena
}

// Función encargada de generar el primer bloque de la cadena de bloques (llamado Génesis)
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Función encargada de generar una nueva cadena de bloques a partir del bloque Génesis
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
