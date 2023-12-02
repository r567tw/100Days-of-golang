package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "io"
)

func encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
    key, _ := hex.DecodeString(keyString)
    plaintext := []byte(stringToEncrypt)

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err.Error())
    }

    secret := make([]byte, aes.BlockSize+len(plaintext))
    iv := secret[:aes.BlockSize]
    if _, err = io.ReadFull(rand.Reader, iv); err != nil {
        panic(err.Error())
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(secret[aes.BlockSize:], plaintext)

    return fmt.Sprintf("%x", secret)
}

func main() {
    key := "6368616e676520746869732070617373" // 這個 key 必須是 32 位元組長度的十六進位字串
    plaintext := "您好，這是一段測試文字！"

    fmt.Println("原始文字: ", plaintext)

    encrypted := encrypt(plaintext, key)
    fmt.Println("加密後: ", encrypted)
}
