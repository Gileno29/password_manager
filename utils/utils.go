package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"golang.org/x/crypto/bcrypt"
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
