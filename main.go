package main

const (
	DESTINY           = "METHINKS IT IS LIKE A WEASEL"
	DestinyLen        = len(DESTINY)
	ChromossomesCount = 100
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
		printBestMatch(cycleCount, bestIndex, &chromossomes)
		crossOver(&adaptabilityIndex, &chromossomes, &newChromosomes)
		mutate()
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
		index := 0
		for j := 0; j < DestinyLen; j++ {
			if string(DESTINY[j]) == chromossomes[i][j] {
				index++
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

func printBestMatch(cycle int, index int, chromossomes *[ChromossomesCount][DestinyLen]string) {}

func crossOver(adaptIndex *[ChromossomesCount]int, chrom *[ChromossomesCount][DestinyLen]string, newChrom *[ChromossomesCount][DestinyLen]string) {
}

func sumIndices(adaptIndex *[ChromossomesCount]int) int {
	return 0
}

func mutate() {

}
