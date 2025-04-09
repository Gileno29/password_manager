package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars   = "0123456789"
	specialChars = "!@#$%^&*()-_=+,.?/:;{}[]~"
)

func EncryptPasswordWithKey(password string, publicKey *rsa.PublicKey) ([]byte, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Segundo passo: encriptar o hash com a chave pública
	encryptedHash, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, hashedPassword)
	if err != nil {
		return nil, err
	}

	return encryptedHash, nil
}

// DecryptPasswordWithKey para verificação (usando chave privada)
func DecryptPasswordWithKey(encryptedHash []byte, privateKey *rsa.PrivateKey) (string, error) {
	// Primeiro: descriptografar com a chave privada
	decryptedHash, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedHash)
	if err != nil {
		return "", err
	}

	// Retorna o hash bcrypt para verificação posterior
	return string(decryptedHash), nil
}

func ParsePublicKey(publicKeyPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pub.(*rsa.PublicKey), nil
}

func GeneratePassword(length int, useUpper, useDigits, useSpecial bool) (string, error) {
	if length < 8 {
		return "", errors.New("o comprimento da senha deve ser pelo menos 8 caracteres")
	}

	var charPool string
	charPool += lowerChars // Sempre inclui letras minúsculas

	if useUpper {
		charPool += upperChars
	}
	if useDigits {
		charPool += digitChars
	}
	if useSpecial {
		charPool += specialChars
	}

	// Garante que a senha contenha pelo menos um caractere de cada tipo selecionado
	password := make([]byte, length)
	var err error

	// Primeiro caractere é sempre minúsculo
	password[0], err = randomChar(lowerChars)
	if err != nil {
		return "", err
	}

	position := 1

	// Garante pelo menos um caractere de cada tipo selecionado
	if useUpper {
		password[position], err = randomChar(upperChars)
		if err != nil {
			return "", err
		}
		position++
	}

	if useDigits {
		password[position], err = randomChar(digitChars)
		if err != nil {
			return "", err
		}
		position++
	}

	if useSpecial {
		password[position], err = randomChar(specialChars)
		if err != nil {
			return "", err
		}
		position++
	}

	// Preenche o restante da senha com caracteres aleatórios do pool
	for i := position; i < length; i++ {
		password[i], err = randomChar(charPool)
		if err != nil {
			return "", err
		}
	}

	// Embaralha os caracteres para maior aleatoriedade
	shuffle(password)

	return string(password), nil
}

// randomChar seleciona um caractere aleatório de uma string
func randomChar(chars string) (byte, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
	if err != nil {
		return 0, err
	}
	return chars[n.Int64()], nil
}

// shuffle embaralha os caracteres da senha
func shuffle(password []byte) {
	for i := len(password) - 1; i > 0; i-- {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		password[i], password[j.Int64()] = password[j.Int64()], password[i]
	}
}
