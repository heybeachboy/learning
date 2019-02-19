package Encrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"runtime"
	"log"
)

func DesCbc_Encrypt(plainText, key []byte) ([]byte) {
	newKey := getKey(key)
	//1. 指定使用des算法,创建并返回一个使用DES算法的cipher.Block接口。
	block, err := des.NewCipher(newKey)
	if err != nil {
		panic(err)
	}

	//2.对明文进行分组填充处理
	paddingText := PKCS5Padding(plainText, block.BlockSize())

	//3.指定使用哪种分组模式 返回一个密码分组链接模式的、底层用b加密的BlockMode接口，初始向量iv的长度必须等于b的块尺寸。
	iv := []byte("wumansgy") //初始化向量
	blockMode := cipher.NewCBCEncrypter(block, iv)

	//4.加密dst src
	cipherText := make([]byte, len(paddingText))
	// 将填充好的paddingText传入进去, 加密的结果 : ciherText
	blockMode.CryptBlocks(cipherText, paddingText)
	return cipherText
}

func DesCbc_Decrypt(cipherText ,key []byte)([]byte) {
	newKey := getKey(key)
	//1.指定解密算法des
	block, err := des.NewCipher(newKey)
	if err!=nil{
		panic(err)
	}

	//2.指定使用哪种分组模式进行解密 返回一个密码分组链接模式的、底层用b解密的BlockMode接口，初始向量iv必须和加密时使用的iv相同。
	iv:=[]byte("wumansgy")   //初始化向量
	blockMode := cipher.NewCBCDecrypter(block, iv)

	//3. 解密
	plainText := make([]byte,len(cipherText))
	blockMode.CryptBlocks(plainText,cipherText)

	//5. 删除填充的内容
	//删除之前防止出现用户输入两次密钥不一样，引起panic,所以做一个错误处理
	//防止用户传的密钥不正确导致panic,这里恢复程序并打印错误
	defer func(){
		if err:=recover();err!=nil{
			switch err.(type){
			case runtime.Error:
				log.Println("runtime err:",err,"请检查密钥是否正确")
			default:
				log.Println("error:",err)
			}
		}
	}()
	unPaddingText := PKCS5UnPadding(plainText)
	return unPaddingText
}

func PKCS5UnPadding(plainText []byte)[]byte{
	//0. 获取总长度
	length := len(plainText)
	// 1.获取最后一个字节
	number:= int(plainText[length-1])
	//2. 删除最后一个字节数
	return plainText[:length-number]
}

// 使用pks5的方式填充
func PKCS5Padding(plainText []byte, blockSize int) []byte {
	// 1. 计算最后一个分组缺多少个字节
	padding := blockSize - (len(plainText) % blockSize)
	// 2. 创建一个大小为padding的切片, 每个字节的值为padding
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 3. 将padText添加到原始数据的后边, 将最后一个分组缺少的字节数补齐
	newText := append(plainText, padText...)
	return newText
}

/**
 *处理密钥KEY的问题
 */
func getKey(key []byte) ([]byte) {
	length := len(key)
	if length == 0 {
		return []byte("zhoulelel")
	}

	if length > 0 && length < 8 {
		key = append(key, bytes.Repeat([]byte{0}, 8-length)...)
		return key
	}

	return key[:8]
}
