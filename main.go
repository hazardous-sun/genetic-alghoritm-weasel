package main

import (
	"fmt"
	"math/rand"
)

const (
	Destiny          = "METHINKS IT IS LIKE A WEASEL"
	DestinyLen       = len(Destiny)
	ChromosomesCount = 100
	MutationFactor   = 4
	Prize            = 100
)

func main() {
	adaptabilityIndex := [ChromosomesCount]int{}
	newChromosomes := initPopulation()
	chromosomes := initPopulation()
	cycleCount := 0
	bestIndex := 0
	for adaptabilityIndex[bestIndex] < DestinyLen*Prize {
		checkAdaptability(&chromosomes, &adaptabilityIndex)
		bestIndex = bestMatch(&adaptabilityIndex)
		printBestMatch(cycleCount, bestIndex, &chromosomes, &adaptabilityIndex)
		crossOver(&chromosomes, &newChromosomes, &adaptabilityIndex)
		mutate(&chromosomes)
		cycleCount++
	}
}

func initPopulation() [ChromosomesCount][DestinyLen]rune {
	arr := [ChromosomesCount][DestinyLen]rune{}
	for i := 0; i < ChromosomesCount; i++ {
		for j := 0; j < DestinyLen; j++ {
			arr[i][j] = 'Z'
		}
	}
	return arr
}

func checkAdaptability(chromosomes *[ChromosomesCount][DestinyLen]rune, adaptIndex *[ChromosomesCount]int) {
	for i := 0; i < ChromosomesCount; i++ {
		index := 1
		for j := 0; j < DestinyLen; j++ {
			if rune(Destiny[j]) == chromosomes[i][j] {
				index += Prize
			}
		}
		adaptIndex[i] = index
	}
}

func bestMatch(adaptIndex *[ChromosomesCount]int) int {
	var bestIndex int
	best := 0
	for i := 0; i < ChromosomesCount; i++ {
		if adaptIndex[i] > best {
			best = adaptIndex[i]
			bestIndex = i
		}
	}
	return bestIndex
}

func printBestMatch(cycle int, index int, chromosomes *[ChromosomesCount][DestinyLen]rune, adaptIndex *[ChromosomesCount]int) {
	fmt.Printf("Cycle %d - %d - : '", cycle, adaptIndex[index]-1)
	for i := 0; i < DestinyLen; i++ {
		fmt.Printf("%c", chromosomes[index][i])
	}
	fmt.Println("'")
}

func crossOver(chrom *[ChromosomesCount][DestinyLen]rune, newChrom *[ChromosomesCount][DestinyLen]rune, adaptIndex *[ChromosomesCount]int) {
	sum := sumIndices(adaptIndex)
	parents := [2]int{}
	indices := [2]int{}
	crossingOverPoint := -1
	for i := 0; i < (ChromosomesCount / 2); i++ {
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
			newChrom[i+(ChromosomesCount/2)][j] = chrom[indices[1]][j]
		}

		for j := crossingOverPoint; j < DestinyLen; j++ {
			newChrom[i][j] = chrom[indices[1]][j]
			newChrom[i+(ChromosomesCount/2)][j] = chrom[indices[0]][j]
		}
	}

	for i := 0; i < ChromosomesCount; i++ {
		for j := 0; j < DestinyLen; j++ {
			chrom[i][j] = newChrom[i][j]
		}
	}
}

func sumIndices(adaptIndex *[ChromosomesCount]int) int {
	sum := 0
	for i := 0; i < ChromosomesCount; i++ {
		sum += adaptIndex[i]
	}
	return sum
}

func mutate(chrom *[ChromosomesCount][DestinyLen]rune) {
	mutationCount := rand.Intn(MutationFactor)
	for i := 0; i < mutationCount; i++ {
		chosenChrom := rand.Intn(ChromosomesCount)
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
