/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import(
    "github.com/butbkadrug/kilogram/cmd"
    _ "github.com/butbkadrug/kilogram/cmd/get"
    _ "github.com/butbkadrug/kilogram/cmd/new"
    _ "github.com/butbkadrug/kilogram/cmd/send"
)
func main() {
	cmd.Execute()
}
