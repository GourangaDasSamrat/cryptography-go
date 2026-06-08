package main

func crypt(textCh, keyCh <-chan byte, result chan<- byte) {
	defer close(result)
	for textByte := range textCh {
		keyByte := <-keyCh
		result <- textByte ^ keyByte
	}
}
