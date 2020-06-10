package gtcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"os/user"
)

func CreateRandomyKey() (key []byte, err error) {
	key = make([]byte, 32)
	_, err = io.ReadFull(rand.Reader, key)
	return
}

func CreateFixedKey() (key []byte, err error) {
	var usr *user.User
	if usr, err = user.Current(); err != nil {
		return nil, err
	}
	var mac string
	if mac, err = GetMacAddr(); err != nil {
		return nil, err
	}
	pp := fmt.Sprintf("%s-%s-ü432$6ätç(&]*£à", usr.Gid, mac)
	x := sha256.Sum256([]byte(pp))
	key = x[:]
	return
}

func CreatePasswordKey(password string) (key []byte, err error) {
	x := sha256.Sum256([]byte(password + "äjfa(&%çXog$$äü!èDS"))
	key = x[:]
	return
}

func Encrypt(data, key []byte) (ret []byte, err error) {
	var gcm cipher.AEAD
	if gcm, err = createGcm(key); err != nil {
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}
	ret = gcm.Seal(nonce, nonce, data, nil)
	ret = []byte(base64.StdEncoding.EncodeToString(ret))
	return
}

func Decrypt(data, key []byte) (ret []byte, err error) {
	data, err = base64.StdEncoding.DecodeString(string(data))
	var gcm cipher.AEAD
	if gcm, err = createGcm(key); err != nil {
		return
	}
	nonceSize := gcm.NonceSize()
	nonce, sealed := data[:nonceSize], data[nonceSize:]
	if ret, err = gcm.Open(nil, nonce, sealed, nil); err != nil {
		return
	}
	return ret, nil
}

func createGcm(key []byte) (gcm cipher.AEAD, err error) {
	var c cipher.Block
	if c, err = aes.NewCipher(key); err != nil {
		return
	}
	gcm, err = cipher.NewGCM(c)
	return
}

func GetMacAddr() (string, error) {
	// Credits to mattn from stackoverflow
	ifas, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var as string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as += a
		}
	}
	return as, nil
}
