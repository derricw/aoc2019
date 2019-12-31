package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	//"strconv"
	//"strings"
)

func extractLayers(in []byte, width, height int) [][]byte {
	layerSize := width * height
	layerCount := len(in) / layerSize
	layers := [][]byte{}
	for i := 0; i < layerCount; i++ {
		layers = append(layers, in[i*layerSize:(i+1)*layerSize])
	}
	return layers
}

func zProjection(layers [][]byte, width, height int) string {
	zProj := ""
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			zProj += func(layers [][]byte, row, col int) string {
				for layer := 0; layer < len(layers); layer++ {
					pix := layers[layer][row*width+col]
					if pix != byte('2') {
						return string(pix)
					}
				}
				return "?"
			}(layers, row, col)
		}
	}
	return zProj
}

func solveP1(layers [][]byte) {
	zeros := []int{}
	ones := []int{}
	twos := []int{}
	for _, layer := range layers {
		zerosCount := 0
		onesCount := 0
		twosCount := 0
		for _, pix := range layer {
			if pix == byte('0') {
				zerosCount++
			} else if pix == byte('1') {
				onesCount++
			} else if pix == byte('2') {
				twosCount++
			}
		}
		zeros = append(zeros, zerosCount)
		ones = append(ones, onesCount)
		twos = append(twos, twosCount)
	}
	minZeros := 0
	minZerosLayer := 0
	for i, count := range zeros {
		if i == 0 {
			minZeros = count
		} else if count < minZeros {
			minZeros = count
			minZerosLayer = i
		}
	}
	output := ones[minZerosLayer] * twos[minZerosLayer]
	fmt.Printf("Pt1 Answer: %v\n", output)
}

func solveP2(layers [][]byte, width, height int) {
	zProj := zProjection(layers, width, height)
	fmt.Printf("Pt2 Answer:\n")
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			pix := zProj[row*width+col]
			if pix == byte('1') {
				fmt.Printf("▓")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func readInput(in io.Reader) []byte {
	buf, _ := ioutil.ReadAll(in)
	return buf
}

func main() {
	width := 25
	height := 6
	in := readInput(os.Stdin)
	layers := extractLayers(in, width, height)

	solveP1(layers)
	solveP2(layers, width, height)
}
