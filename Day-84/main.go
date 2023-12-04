package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/hex"
    "fmt"
)

func decrypt(encryptedString string, keyString string) (decryptedString string) {
    key, _ := hex.DecodeString(keyString)
    ciphertext, _ := hex.DecodeString(encryptedString)

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err.Error())
    }

    if len(ciphertext) < aes.BlockSize {
        panic("ciphertext too short")
    }
    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)

    // XORKeyStream can work in-place if the two arguments are the same.
    stream.XORKeyStream(ciphertext, ciphertext)

    return fmt.Sprintf("%s", ciphertext)
}



func main() {
    key := "6368616e676520746869732070617373" // 這個 key 必須是 32 位元組長度的十六進位字串

    encrypted := "3095212d200a4d0265662ad73aeb7115cac2f88e4a7b8a939705ca3f69486232111eef86eed5826f3ff7ba6d5f4fb096fe5de605"
    decrypted := decrypt(encrypted, key)
    fmt.Println("解密後: ", decrypted)
}
