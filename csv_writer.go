package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"afakih", "fajduwani", "dewangga"})
	_ = writer.Write([]string{"laras", "ayu", "hyein"})
	_ = writer.Write([]string{"seunghee", "liz", "binnie"})

	writer.Flush()
}
