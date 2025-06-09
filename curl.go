package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"encoding/base64"
)

// makeRequest realiza una petición HTTP.
// headersStr: String en formato "Header1: Valor1\nHeader2: Valor2".
// Si response es NULL, devuelve un nuevo *C.char (debe liberarse con Free).
func makeRequest(method string, url *C.char, body *C.char, headersStr *C.char) *C.char {
	goUrl := C.GoString(url)
	var goBody []byte
	if body != nil {
		goBody = []byte(C.GoString(body))
	}
	req, err := http.NewRequest(method, goUrl, bytes.NewReader(goBody))
	if err != nil {
		return nil
	}
	if headersStr != nil {
		headers := C.GoString(headersStr)
		for _, line := range strings.Split(headers, "\n") {
			if line == "" {
				continue
			}
			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				continue // Formato inválido, ignorar
			}
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			req.Header.Add(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	return C.CString(string(respBody))
}

//export Header
func Header(key *C.char, value *C.char) *C.char {
	gokey := C.GoString(key)
	govalue := C.GoString(value)
	header := gokey+": " + govalue
	return C.CString(header)
}

//export HeaderAuthToken
func HeaderAuthToken(token *C.char) *C.char {
	goToken := C.GoString(token)
	header := "Authorization: Bearer " + goToken
	return C.CString(header)
}

//export HeaderAuthBasic
func HeaderAuthBasic(user *C.char, pass *C.char) *C.char {
	goUser := C.GoString(user)
	goPass := C.GoString(pass)
	authString := goUser + ":" + goPass
	encoded := base64.StdEncoding.EncodeToString([]byte(authString))
	header := "Authorization: Basic " + encoded
	return C.CString(header)
}

//export Get
func Get(url *C.char, headers *C.char, body *C.char) *C.char {
	return makeRequest("GET", url, body, headers)
}

//export Post
func Post(url *C.char, headers *C.char, body *C.char) *C.char {
	return makeRequest("POST", url, body, headers)
}

//export Put
func Put(url *C.char, headers *C.char, body *C.char) *C.char {
	return makeRequest("PUT", url, body, headers)
}

//export Patch
func Patch(url *C.char, headers *C.char, body *C.char) *C.char {
	return makeRequest("PATCH", url, body, headers)
}

//export Delete
func Delete(url *C.char, headers *C.char, body *C.char) *C.char {
	return makeRequest("DELETE", url, body, headers)
}

//export Head
func Head(url *C.char, headers *C.char, body *C.char) *C.char {
	return makeRequest("HEAD", url, body, headers)
}

//export Options
func Options(url *C.char, headers *C.char, body *C.char) *C.char {
	return makeRequest("OPTIONS", url, body, headers)
}


func main() {}