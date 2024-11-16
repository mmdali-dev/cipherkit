package rsa

import "io/ioutil"

func LoadKeyFromFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}

func SaveKeyToFile(filename string, keyString string) error {
	return ioutil.WriteFile(filename, []byte(keyString), 0600)
}
