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
	newChromosomes := initPopulation()
	chromossomes := initPopulation()
	cycleCount := 0
	bestIndex := 0
	for adaptabilityIndex[bestIndex] < DestinyLen+1 {
		checkAdaptability(&chromossomes, &adaptabilityIndex)
		bestIndex = bestMatch(&adaptabilityIndex)
		printBestMatch(cycleCount, bestIndex, &chromossomes, &adaptabilityIndex)
		crossOver(&chromossomes, &newChromosomes, &adaptabilityIndex)
		mutate(&chromossomes)
		cycleCount++
	}
}

func initPopulation() [ChromossomesCount][DestinyLen]rune {
	arr := [ChromossomesCount][DestinyLen]rune{}
	for i := 0; i < ChromossomesCount; i++ {
		for j := 0; j < DestinyLen; j++ {
			arr[i][j] = 'Z'
		}
	}
	return arr
}

func checkAdaptability(chromossomes *[ChromossomesCount][DestinyLen]rune, adaptIndex *[ChromossomesCount]int) {
	for i := 0; i < ChromossomesCount; i++ {
		index := 1
		for j := 0; j < DestinyLen; j++ {
			if rune(Destiny[j]) == chromossomes[i][j] {
				index += 1
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

func printBestMatch(cycle int, index int, chromossomes *[ChromossomesCount][DestinyLen]rune, adaptIndex *[ChromossomesCount]int) {
	fmt.Printf("Cycle %d - %d - : '", cycle, adaptIndex[index]-1)
	for i := 0; i < DestinyLen; i++ {
		fmt.Printf("%c", chromossomes[index][i])
	}
	fmt.Println("'")
}

func crossOver(chrom *[ChromossomesCount][DestinyLen]rune, newChrom *[ChromossomesCount][DestinyLen]rune, adaptIndex *[ChromossomesCount]int) {
	sum := sumIndices(adaptIndex)
	parents := [2]int{}
	indices := [2]int{}
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

		// Roulette for first parent
		indices[0] = 0
		for parents[0] > 0 {
			parents[0] -= adaptIndex[indices[0]]
			indices[0]++
		}
		indices[0]--

		if indices[0] < 0 {
			indices[0] = 0
		}

		// Roulette for second parent
		indices[1] = 0
		for parents[1] > 0 {
			parents[1] -= adaptIndex[indices[1]]
			indices[1]++
		}
		indices[1]--

		if indices[1] < 0 {
			indices[1] = 0
		}

		crossingOverPoint = rand.Intn(8) + 2

		for j := 0; j < crossingOverPoint; j++ {
			newChrom[i][j] = chrom[indices[0]][j]
			newChrom[i+(ChromossomesCount/2)][j] = chrom[indices[1]][j]
		}

		for j := crossingOverPoint; j < DestinyLen; j++ {
			newChrom[i][j] = chrom[indices[1]][j]
			newChrom[i+(ChromossomesCount/2)][j] = chrom[indices[0]][j]
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

func mutate(chrom *[ChromossomesCount][DestinyLen]rune) {
	mutationCount := rand.Intn(MutationFactor)
	for i := 0; i < mutationCount; i++ {
		chosenChrom := rand.Intn(ChromossomesCount)
		mutationPoint := rand.Intn(DestinyLen)
		options :=
			[]rune{
				'A', 'B', 'C', 'D',
				'E', 'F', 'G', 'H',
				'I', 'J', 'K', 'L',
				'M', 'N', 'O', 'P',
				'Q', 'R', 'S', 'T',
				'U', 'V', 'X', 'W',
				'Y', 'Z', ' ',
			}
		pos := rand.Intn(len(options))
		temp := options[pos]
		chrom[chosenChrom][mutationPoint] = temp
	}
}
