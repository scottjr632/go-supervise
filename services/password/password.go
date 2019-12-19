package password

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

const defaultFileName = "creds.txt"
const defaultAdminUser = "admin"

func generateRandomASCIIString(length int) (string, error) {
	result := ""
	for {
		if len(result) >= length {
			return result, nil
		}
		num, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		// Make sure that the number/byte/letter is inside
		// the range of printable ASCII characters (excluding space and DEL)
		if n > 32 && n < 127 {
			result += string(n)
		}
	}
}

func SaveBasicPassword() {
	f, err := os.Create(defaultFileName)
	must(err)
	defer f.Close()

	pass, err := generateRandomASCIIString(32)
	must(err)

	_, err = f.WriteString(pass)
	must(err)

	fmt.Printf(" Default password\n%s\n", pass)
}

func GetBasicPassword() (string, error) {
	data, err := ioutil.ReadFile(defaultFileName)
	return string(data), err
}

func GetBasicUser() string {
	return defaultAdminUser
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
