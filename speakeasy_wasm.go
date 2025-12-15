//go:build js && wasm

package speakeasy

func getPassword() (password string, err error) {
	return "", errors.New("password input not supported in wasm")
}
