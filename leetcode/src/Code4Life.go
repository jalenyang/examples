package main

import (
	"fmt"
	"strings"
	"os"
)

/**
 * Bring data on patient samples from the diagnosis machine to the laboratory with enough molecules to produce medicine!
 **/

type Robot struct {
	target    string
	storage   []int
	carrying  []Sample
	expertise []int
	score     int
	eta       int
	cloud     [] Sample
}

type Sample struct {
	sampleId      int
	carriedBy     int
	rank          int
	health        int
	cost          []int
	diagnosed     bool
	ready         bool
	expertiseGain string
}

func goEmpty() {
	fmt.Println("GOTO", "EMPTY")
}

func goTo(module string) {
	fmt.Println("GOTO", module)
}

func connect(data interface{}) {
	fmt.Println("CONNECT", data)
}

func sum(expertise []int) int {
	return expertise[0] + expertise[1] + expertise[2] + expertise[3] + expertise[4]
}

func getLabSample(robot Robot) (Sample, bool) {
	storage := robot.storage
	carrying := robot.carrying
	expertise := robot.expertise

	for i := 0; i < len(carrying); i++ {
		need_more := false
		for j := 0; j < 5; j++ {
			if storage[j]+expertise[j] < carrying[i].cost[j] {
				need_more = true
				break
			}
		}
		if need_more && i != len(carrying)-1 {
			continue
		} else if need_more && i == len(carrying)-1 {
			return carrying[i], false
		}
		if !need_more {
			return carrying[i], true
		}
	}

	return Sample{sampleId: -1}, false
}

func isPossbibleToLab(storage [] int, carrying []Sample, expertise [] int, availables []int) bool {
	count := 0
	for index, sample := range carrying {
		for i := 0; i < 5; i++ {
			if storage[i]+expertise[i]+availables[i] < sample.cost[i] {
				carrying[index].ready = false
				count++
				break
			}
		}
	}
	if count > 0 {
		return false
	}
	return true
}

func getNeededModleCule(robots [] Robot, availables []int) string {
	storage := robots[0].storage
	carrying := robots[0].carrying
	expertise := robots[0].expertise
	need_modlecule := ""

	if !isPossbibleToLab(storage, carrying, expertise, availables) {
		return "none"
	}

	//sort the samples by rank asc
	var sortCarrying []Sample
	sortCarrying = sortByRank(carrying)
	for _, sample := range sortCarrying {
		if !sample.ready {
			continue
		}
		for i := 0; i < 5; i++ {
			if storage[i]+expertise[i] < sample.cost[i] {
				if checkAvailable(availables, string("ABCDE"[i])) {
					need_modlecule = string("ABCDE"[i])
				}
			}
		}
		if need_modlecule == "" {
			for index, mol := range []string{"A", "B", "C", "D", "E"} {
				storage[index] -= sample.cost[index] - expertise[index]
				if sample.expertiseGain == mol {
					expertise[index] += 1
				}
			}
		}
	}
	return need_modlecule
}

func sortByRank(carrying []Sample) []Sample {
	var sortCarrying []Sample
	for index, sample := range carrying {
		if len(sortCarrying) == 0 {
			sortCarrying = append(sortCarrying, sample)
		} else {
			if sample.rank < carrying[index-1].health {
				sortCarrying = append([]Sample{sample}, sortCarrying...)
			} else {
				sortCarrying = append(sortCarrying, sample)
			}
		}
	}
	return sortCarrying
}

func getPossibleModleCule(robots [] Robot, availables []int) string {
	storage := robots[0].storage
	carrying := robots[0].carrying
	expertise := robots[0].expertise
	need_modleCule := ""
	var sortCarrying []Sample
	sortCarrying = sortByRank(carrying)

	for _, carrying := range sortCarrying {
		for i := 0; i < 5; i++ {
			if storage[i]+expertise[i] < carrying.cost[i] {
				if checkAvailable(availables, string("ABCDE"[i])) {
					need_modleCule = string("ABCDE"[i])
				}
			}
		}
	}
	if need_modleCule == "" {
		need_modleCule = getBlockModleCule(robots[1].storage, robots[1].expertise, robots[1].carrying, availables)
	}
	return need_modleCule
}

func getBlockModleCule(storage [] int, expertise [] int, carrying [] Sample, availables []int) string {
	need_modlecule := ""

	var sortCarrying []Sample
	sortCarrying = sortByRank(carrying)

	for _, sample := range sortCarrying {
		if !sample.diagnosed {
			continue
		}
		for i := 0; i < 5; i++ {
			if storage[i]+expertise[i] < sample.cost[i] {
				if checkAvailable(availables, string("ABCDE"[i])) {
					need_modlecule = string("ABCDE"[i])
					break
				}
			}
		}
	}
	return need_modlecule
}

func getAvailableOne(robots []Robot, samples [] Sample, availables []int) int {
	sampleId := -1
	for _, sample := range samples {
		if readyTolab(robots[0].storage, robots[0].expertise, sample) {
			sampleId = sample.sampleId
			break
		}
	}
	return sampleId
}

func retryToGetModlecule(robots []Robot, samples [] Sample, availables []int) int {
	sampleId := -1
	for _, sample := range samples {
		if possibleReadyToLab(robots, availables, sample) {
			sampleId = sample.sampleId
			break
		}
	}
	return sampleId
}

func readyTolab(storage [] int, expertise [] int, sampe Sample) bool {
	for j := 0; j < 5; j++ {
		if storage[j]+expertise[j] < sampe.cost[j] {
			return false
		}
	}
	return true
}

func possibleReadyToLab(robots [] Robot, availables []int, sample Sample) bool {
	storage0 := robots[0].storage
	storage1 := robots[1].storage
	//carrying0 := robots[0].carrying
	carrying1 := robots[1].carrying
	expertise0 := robots[0].expertise
	expertise1 := robots[1].expertise
	target1 := robots[1].target
	cost1 := make([]int, 5)

	for _, sample := range carrying1 {
		for i := 0; i < 5; i++ {
			cost1[i] += sample.cost[i] - expertise1[i]
		}
	}

	if target1 == "LABORATORY" {
		for j := 0; j < 5; j++ {
			if storage0[j]+expertise0[j]+availables[j]+storage1[j] < sample.cost[j] {
				return false
			}
		}
	} else {
		for j := 0; j < 5; j++ {
			if storage0[j]+expertise0[j]+availables[j] < sample.cost[j] {
				return false
			}
		}
	}
	return true
}

func checkAvailable(availables [] int, molecule string) bool {
	if availables[strings.Index("ABCDE", molecule)] > 0 {
		return true
	}
	return false
}

func sumCost(carrying []Sample) int {
	cost := 0
	for _, sample := range carrying {
		cost += sum(sample.cost)
	}
	return cost
}

func connectRank(expertise []int, carrying []Sample, storage []int, availables []int, otherTarget string, otherCarrying []Sample) int {
	if sum(storage) == 10 {
		return 1
	}
	if sum(expertise) < 1 {
		return 1
	}
	factor := sum(expertise) + 10 - sum(storage)-sumCost(carrying)
	fmt.Fprintln(os.Stderr, "sum(expertise):", sum(expertise))
	fmt.Fprintln(os.Stderr, "10 - sum(storage):", 10-sum(storage))
	if factor < 5 {
		return 1
	} else if factor < 10{
		return 2
	} else if factor < 18 {
		return 3
	}
	return 1
}

func connectSample(expertise []int, carrying []Sample, storage []int, availables []int, otherTarget string, otherCarrying []Sample) {
	if len(carrying) < 3 {
		rankId := connectRank(expertise, carrying, storage, availables, otherTarget, otherCarrying)
		connect(rankId)
	} else {
		goTo("DIAGNOSIS")
	}
}

func main() {
	var projectCount int
	fmt.Scan(&projectCount)

	for i := 0; i < projectCount; i++ {
		var a, b, c, d, e int
		fmt.Scan(&a, &b, &c, &d, &e)
	}

	for {
		var robot Robot
		var robots []Robot
		for i := 0; i < 2; i++ {
			var target string
			var eta, score, storageA, storageB, storageC, storageD, storageE, expertiseA, expertiseB, expertiseC, expertiseD, expertiseE int
			fmt.Scan(&target, &eta, &score, &storageA, &storageB, &storageC, &storageD, &storageE, &expertiseA, &expertiseB, &expertiseC, &expertiseD, &expertiseE)
			robot.target = target
			robot.score = score
			robot.storage = []int{storageA, storageB, storageC, storageD, storageE}
			robot.expertise = []int{expertiseA, expertiseB, expertiseC, expertiseD, expertiseE}
			robot.eta = eta
			robots = append(robots, robot)
		}

		var availableA, availableB, availableC, availableD, availableE int
		fmt.Scan(&availableA, &availableB, &availableC, &availableD, &availableE)

		availables := []int{availableA, availableB, availableC, availableD, availableE}

		var sampleCount int
		fmt.Scan(&sampleCount)
		var sample Sample
		var samples []Sample

		for i := 0; i < sampleCount; i++ {
			var sampleId, carriedBy, rank int
			var expertiseGain string
			var health, costA, costB, costC, costD, costE int
			fmt.Scan(&sampleId, &carriedBy, &rank, &expertiseGain, &health, &costA, &costB, &costC, &costD, &costE)
			sample.sampleId = sampleId
			sample.carriedBy = carriedBy
			sample.rank = rank
			sample.health = health
			sample.cost = []int{costA, costB, costC, costD, costE}
			sample.diagnosed = health != -1
			sample.expertiseGain = expertiseGain
			sample.ready = true
			samples = append(samples, sample)
			if carriedBy > -1 {
				robots[carriedBy].carrying = append(robots[carriedBy].carrying, sample)
			} else if sample.diagnosed {
				robots[0].cloud = append(robots[0].cloud, sample)
			}
		}

		myRobot := robots[0]
		storage := myRobot.storage
		allStorage := sum(storage)
		expertise := myRobot.expertise
		carrying := myRobot.carrying
		eta := myRobot.eta
		target := myRobot.target

		otherRobot := robots[1]
		//otherStorage := otherRobot.storage
		//otherExpertise := otherRobot.expertise
		otherCarrying := otherRobot.carrying
		//otherEta := otherRobot.eta
		otherTarget := otherRobot.target

		switch target {
		case "START_POS":
			goTo("SAMPLES")
		case "SAMPLES":
			if eta > 0 {
				goEmpty()
			} else {
				connectSample(expertise, carrying, storage, availables, otherTarget, otherCarrying)
			}
		case "DIAGNOSIS":
			size := len(carrying)
			if eta > 0 {
				goEmpty()
			} else {
				if size > 0 && !carrying[0].diagnosed {
					connect(carrying[0].sampleId)
				} else if size > 1 && !carrying[1].diagnosed {
					connect(carrying[1].sampleId)
				} else if size > 2 && !carrying[2].diagnosed {
					connect(carrying[2].sampleId)
				} else if size > 0 && carrying[0].diagnosed && !possibleReadyToLab(robots, availables, carrying[0]) {
					connect(carrying[0].sampleId)
				} else if size > 1 && carrying[1].diagnosed && !possibleReadyToLab(robots, availables, carrying[1]) {
					connect(carrying[0].sampleId)
				} else if size > 2 && carrying[2].diagnosed && !possibleReadyToLab(robots, availables, carrying[2]) {
					connect(carrying[0].sampleId)
				} else if size < 3 && len(myRobot.cloud) > 0 {
					sampleId := getAvailableOne(robots, myRobot.cloud, availables)
					if sampleId != -1 {
						connect(sampleId)
					} else if size > 0 && allStorage == 10 {
						goTo("LABORATORY")
					} else if size > 0 && allStorage < 10 {
						goTo("MOLECULES")
					} else {
						goTo("SAMPLES")
					}
				} else if size > 0 && allStorage == 10 {
					goTo("LABORATORY")
				} else if size > 0 && allStorage < 10 {
					goTo("MOLECULES")
				} else {
					goTo("SAMPLES")
				}
			}
		case "MOLECULES":
			if eta > 0 {
				goEmpty()
			} else {
				molecule := getNeededModleCule(robots, availables)
				molecule2 := getPossibleModleCule(robots, availables)
				if molecule == "none" {
					if allStorage < 10 && molecule2 != "" {
						connect(molecule2)
					} else if len(carrying) < 3 {
						goTo("SAMPLES")
					} else if allStorage < 10 {
						blockMolecule := getBlockModleCule(robots[1].storage, robots[1].expertise, robots[1].carrying, availables)
						if blockMolecule != "" {
							connect(molecule)
						} else {
							goTo("DIAGNOSIS")
						}
					} else {
						goTo("DIAGNOSIS")
					}
				} else if molecule != "" && allStorage < 10 {
					connect(molecule)
				} else if allStorage < 10 {
					molecule = getBlockModleCule(robots[1].storage, robots[1].expertise, robots[1].carrying, availables)
					if molecule != "" {
						connect(molecule)
					} else {
						goTo("LABORATORY")
					}
				} else {
					goTo("LABORATORY")
				}

			}

		case "LABORATORY":
			if eta > 0 {
				goEmpty()
			} else {
				sample, ready := getLabSample(myRobot)
				if sample.sampleId == -1 {
					goTo("SAMPLES")
				} else if !ready {
					if allStorage == 10 {
						if len(carrying) < 3 {
							goTo("SAMPLES")
						} else {
							goTo("DIAGNOSIS")
						}
					} else {
						sampleId := retryToGetModlecule(robots, myRobot.carrying, availables)
						cloudId := getAvailableOne(robots, myRobot.cloud, availables)
						if sampleId != -1 && allStorage < 10 {
							goTo("MOLECULES")
						} else if len(carrying) == 3 {
							goTo("DIAGNOSIS")
						} else if len(myRobot.cloud) > 0 && cloudId != -1 {
							goTo("DIAGNOSIS")
						} else if (robots[1].target == "DIAGNOSIS" || robots[1].target == "MOLECULES") && allStorage < 10 {
							goTo("MOLECULES")
						} else {
							goTo("SAMPLES")
						}
					}
				} else {
					connect(sample.sampleId)
				}
			}
		}

	}
}
