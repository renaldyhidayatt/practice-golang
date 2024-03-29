package simplecipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type ShiftCipher struct {
	length int
	shift  []int
}

func NewCaesar() Cipher {
	panic("Please implement the NewCaesar function")
}

func NewShift(distance int) Cipher {
	panic("Please implement the NewShift function")
}

func (c shift) Encode(input string) string {
	panic("Please implement the Encode function")
}

func (c shift) Decode(input string) string {
	panic("Please implement the Decode function")
}

func NewVigenere(key string) Cipher {
	panic("Please implement the NewVigenere function")
}

func (v vigenere) Encode(input string) string {
	panic("Please implement the Encode function")
}

func (v vigenere) Decode(input string) string {
	panic("Please implement the Decode function")
}
