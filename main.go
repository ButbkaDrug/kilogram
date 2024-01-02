package main

import(
    "github.com/butbkadrug/kilogram/cmd"
    _ "github.com/butbkadrug/kilogram/cmd/get"
    _ "github.com/butbkadrug/kilogram/cmd/send"
)
func main() {
	cmd.Execute()
}
