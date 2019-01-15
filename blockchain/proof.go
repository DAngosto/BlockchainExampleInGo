package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// 1º - Conseguir el valor de data del bloque

// 2º - Crear un contador que empiece en 0 y aumente teoricamente hasta infinito

// 3º - Crear un hash del data + contador

// Pasos 1, 2 y 3 ---> --Función Run ---> InitData--

// 4º - Comprobar el hash para ver si cumple los requirimientos

// Requirimientos:
// Los primeros bytes deben contener 0's

// En este caso es estático pero en una implementación real este valor iría incrementando por el algoritmo que se esté usando en función del número de mineros y de la potencia computacional total
const Difficulty = 16

type ProofOfWork struct {
	Block  *Block
	Target *big.Int // Es un  número que repsenta los requirimientos descritos arriba que vienen influenciados por la dificultad puesta
}

// Recibimos un bloque de la cadena de bloques para generar una nueva prueba de trabajo para dicho bloque
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	//fmt.Printf("%", )
	//fmt.Println("target: ", target)

	pow := &ProofOfWork{b, target}

	return pow
}

// Recibimos el valor del contador para realizar la unión de bytes para la formación del data que irá a ser hasheado en la función Run
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToByte(int64(nonce)),
			ToByte(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte
	nonce := 0
	for nonce < math.MaxInt64 { // Loop virtual infinito para sacar el hash que cumpla los requirimientos
		data := pow.InitData(nonce) // Preparamos el data para que cumpla los requirimientos
		hash = sha256.Sum256(data)  // Lo hasheamos

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 { // Comprobamos si el hash generado cumple los requirimientos, si no es así
			break // se aumenta el contador y se vuelve a realizar el proceso hasta obtenerlo.
		} else { // En este caso se comprueba si el hash generado tiene tantos 0's al
			nonce++ // principio como se pedían (Requirimientos de arriba)
		}

	}
	fmt.Println()

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1

}

// Recibe un entero de 64 bits y lo pasa a Bytes con el tipo de ordenación BigEndian
func ToByte(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
