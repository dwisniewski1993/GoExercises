package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
)

// Tranzakcja składa się z zawartości paczki i hasza
type Transaction struct {
	Hash    []byte
	Payload []byte
}

// Tworzenie nowej tranzakcji
func NewTransaction(payload []byte) *Transaction {
	log.Printf("transaction.go [INFO]: Creating new transaction")
	tx := Transaction{Payload: payload}
	tx.Hash, _ = GetTransactionHash(tx)

	return &tx
}

// Hash tranzakcji wyliczony z zawartości
func GetTransactionHash(tx Transaction) ([]byte, error) {
	txBytes, err := tx.MarshalBinary()
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	h.Write(txBytes)

	return h.Sum(nil), nil
}

// Przeliczanie na bity
func (tx *Transaction) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, tx.Payload)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
