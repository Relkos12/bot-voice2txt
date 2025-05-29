package entity

import "os"

//ConvertedText is text from voice message
type ConvertedText struct {
	ChatID float64
	File   *os.File
}
