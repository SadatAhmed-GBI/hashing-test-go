/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc"
	hashConstant := "f9c977ea-b423-4b3e-9952-d38cec720b70"
	hashValue := fmt.Sprintf("%s%s", hashConstant, data)

	// calculate SHA-256 hash of the input
	hash := sha256.Sum256([]byte(hashValue))

	sEnc := b64.StdEncoding.EncodeToString(hash[:])
	fmt.Println(sEnc)
}
