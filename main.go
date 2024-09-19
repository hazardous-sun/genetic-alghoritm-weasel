package main

import (
	"fmt"
	"math/rand"
)

const (
	Destiny           = "METHINKS IT IS LIKE A WEASEL"
	DestinyLen        = len(Destiny)
	ChromossomesCount = 100
	MutationFactor    = 4
)

func main() {
	adaptabilityIndex := [ChromossomesCount]int{}
	newChromosomes := [ChromossomesCount][DestinyLen]string{}
	chromossomes := initPopulation()
	cycleCount := 0
	bestIndex := 0
	for adaptabilityIndex[bestIndex] < DestinyLen {
		checkAdaptability(&chromossomes, &adaptabilityIndex)
		bestIndex = bestMatch(&adaptabilityIndex)
		printBestMatch(cycleCount, bestIndex, &chromossomes, &adaptabilityIndex)
		crossOver(&chromossomes, &newChromosomes, &adaptabilityIndex)
		mutate(&chromossomes)
		cycleCount++
	}
}

func initPopulation() [ChromossomesCount][DestinyLen]string {
	arr := [ChromossomesCount][DestinyLen]string{}
	for i := 0; i < ChromossomesCount; i++ {
		for j := 0; j < DestinyLen; j++ {
			arr[i][j] = "Z"
		}
	}
	return arr
}

func checkAdaptability(chromossomes *[ChromossomesCount][DestinyLen]string, adaptIndex *[ChromossomesCount]int) {
	for i := 0; i < ChromossomesCount; i++ {
		index := 50
		for j := 0; j < DestinyLen; j++ {
			if string(Destiny[j]) == chromossomes[i][j] {
				index += 50
			}
		}
		adaptIndex[i] = index
	}
}

func bestMatch(adaptIndex *[ChromossomesCount]int) int {
	var bestIndex int
	best := 0
	for i := 0; i < ChromossomesCount; i++ {
		if adaptIndex[i] > best {
			best = adaptIndex[i]
			bestIndex = i
		}
	}
	return bestIndex
}

func printBestMatch(cycle int, index int, chromossomes *[ChromossomesCount][DestinyLen]string, adaptIndex *[ChromossomesCount]int) {
	fmt.Printf("Cycle %d - %d - : '", cycle, adaptIndex[index])
	for i := 0; i < DestinyLen; i++ {
		fmt.Printf("%s", chromossomes[index][i])
	}
	fmt.Println("'")
}

func crossOver(chrom *[ChromossomesCount][DestinyLen]string, newChrom *[ChromossomesCount][DestinyLen]string, adaptIndex *[ChromossomesCount]int) {
	sum := sumIndices(adaptIndex)
	parents := [2]int{}
	parentsIndices := [2]int{}
	crossingOverPoint := -1
	for i := 0; i < (ChromossomesCount / 2); i++ {
		parents[0] = rand.Intn(sum)
		parents[1] = rand.Intn(sum)

		if parents[0] == parents[1] {
			if parents[1] < 49 {
				parents[1]++
			} else {
				parents[1]--
			}
		}

		parentsIndices[0] = 0
		for parents[0] > 0 {
			parents[0] -= adaptIndex[parentsIndices[0]]
			parentsIndices[0]++
		}
		parentsIndices[0]--

		if parentsIndices[0] < 0 {
			parentsIndices[0] = 0
		}

		parentsIndices[1] = 0
		for parents[1] > 0 {
			parents[1] -= adaptIndex[parentsIndices[1]]
			parentsIndices[1]++
		}
		parentsIndices[1]--

		if parentsIndices[1] < 0 {
			parentsIndices[1] = 0
		}

		crossingOverPoint = rand.Intn(8) + 2

		for j := 0; j < crossingOverPoint; j++ {
			newChrom[i][j] = chrom[parentsIndices[0]][j]
			newChrom[i+(ChromossomesCount/2)][j] = chrom[parentsIndices[1]][j]
		}

		for j := 0; j < crossingOverPoint; j++ {
			newChrom[i][j] = chrom[parentsIndices[1]][j]
			newChrom[i+(ChromossomesCount/2)][j] = chrom[parentsIndices[0]][j]
		}
	}

	for i := 0; i < ChromossomesCount; i++ {
		for j := 0; j < DestinyLen; j++ {
			chrom[i][j] = newChrom[i][j]
		}
	}
}

func sumIndices(adaptIndex *[ChromossomesCount]int) int {
	sum := 0
	for i := 0; i < ChromossomesCount; i++ {
		sum += adaptIndex[i]
	}
	return sum
}

func mutate(chrom *[ChromossomesCount][DestinyLen]string) {
	mutationCount := rand.Intn(MutationFactor)
	for i := 0; i < mutationCount; i++ {
		chosenChrom := rand.Intn(ChromossomesCount)
		mutationPoint := rand.Intn(DestinyLen)
		options := "ABCDEFGHIJKLMNOPQRSTUVXWYZ _"
		pos := rand.Intn(len(options) - 1)
		temp := options[pos : pos+1]
		chrom[chosenChrom][mutationPoint] = temp
	}
}
