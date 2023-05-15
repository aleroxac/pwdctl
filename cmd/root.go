/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
)

var letter int
var number int

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")

func RandLetters(num_letters int) string {
	rand_rune := make([]rune, num_letters)
	for i := range rand_rune {
		rand_rune[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(rand_rune)
}

func RandNumbers(num_numbers int) string {
	rand_number := make([]rune, num_numbers)
	for i := range rand_number {
		rand_number[i] = numberRunes[rand.Intn(10)]
	}
	return string(rand_number)
}

func GeneratePassword(c *cobra.Command, args []string) error {
	rand_letter := RandLetters(letter)
	rand_number := RandNumbers(number)
	temp_rand_pass := rand_letter + rand_number

	perm := rand.Perm(len(temp_rand_pass))
	shuffledRunes := make([]rune, len(temp_rand_pass))

	for i, randIndex := range perm {
		shuffledRunes[i] = []rune(temp_rand_pass)[randIndex]
	}

	fmt.Printf("%s\n", string(shuffledRunes))
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "pwdctl [option]",
	Short: "Command line tool to manage passwords.",
	RunE:  GeneratePassword,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&letter, "letters", "l", 0, "Include at least one letter in the password")
	rootCmd.Flags().IntVarP(&number, "numbers", "n", 0, "Include at least one number in the password")
}
