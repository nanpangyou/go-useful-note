package std_byte

import (
	"bytes"
	"fmt"
)

func TestByteBuffer() {
	testBuffer()

	testBufferWrite()
}

func testBuffer() {
	var b bytes.Buffer
	fmt.Printf("b: %v\n", b)

	b1 := bytes.NewBufferString("hello xxx")
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b1: %T\n", b1)
	b2 := bytes.NewBuffer([]byte("hello xxx"))
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("b2: %T\n", b2)
}

func testBufferWrite() {
	var b bytes.Buffer
	n, _ := b.WriteString("hello")
	fmt.Printf("n: %v\n", n)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b.String(): %v\n", b.String())
}
