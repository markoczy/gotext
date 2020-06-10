package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os/user"
	"path"
	"time"

	"github.com/markoczy/gotext/common"
)

const (
	keyFile = "pk.enc"
)

func CreateRandomyKey() (key []byte, err error) {
	key = make([]byte, 32)
	_, err = io.ReadFull(rand.Reader, key)
	return
}

func CreateSecondaryKey() (key []byte, err error) {
	var usr *user.User
	if usr, err = user.Current(); err != nil {
		return nil, err
	}
	ts := time.Now().Truncate(time.Hour * 24).Unix()
	pp := fmt.Sprintf("%s-%d", usr.Gid, ts)
	x := sha256.Sum256([]byte(pp))
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
	return
}

func Decrypt(data, key []byte) (ret []byte, err error) {
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

func LoadPrimaryKey(password string) ([]byte, error) {
	keyDir, err := common.InitKeyDir()
	if err != nil {
		return nil, err
	}
	keyPath := path.Join(keyDir, keyFile)
	if !common.FileExists(keyPath) {
		if err = savePrimaryKey(password); err != nil {
			return nil, err
		}
	}

	pkEnc, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	pk, err := Decrypt(pkEnc, []byte(password))
	return pk, nil
}

func savePrimaryKey(password string) error {
	keyDir, err := common.InitKeyDir()
	if err != nil {
		return err
	}
	keyPath := path.Join(keyDir, keyFile)

	if common.FileExists(keyPath) {
		return fmt.Errorf("Key File already exists!")
	}

	pk, err := CreateRandomyKey()
	if err != nil {
		return err
	}

	pkEnc, err := Encrypt(pk, []byte{})
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(keyPath, pkEnc, 0666)
	return err
}

func createGcm(key []byte) (gcm cipher.AEAD, err error) {
	var c cipher.Block
	if c, err = aes.NewCipher(key); err != nil {
		return
	}
	gcm, err = cipher.NewGCM(c)
	return
}
