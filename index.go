//CombarRobot Drive Train Calculator based on AskAron at http://runamok.tech/AskAaron/optimum.html

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	step1()
}

func step1() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Weight Supported by the Wheel")
	scanner.Scan()
	weightSupportedByWheel, nmb := strconv.ParseFloat(scanner.Text(), 64)
	FrictionCoeffecient := 0.8 //default value for rubber tire and wood/steel surface; change if needed
	if nmb == nil {
		maxTireForce := weightSupportedByWheel * FrictionCoeffecient
		step2(maxTireForce)
	} else {
		step1()
	}
}

func step2(maxTireForce float64) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Tire Radius")
	scanner.Scan()
	tireRadius, nmb := strconv.ParseFloat(scanner.Text(), 64)
	if nmb == nil {
		torqForMaxForce := maxTireForce * tireRadius
		step3(torqForMaxForce)
	} else {
		step2(maxTireForce)
	}
}

func step3(torqForMaxForce float64) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Motor Stall Torque")
	scanner.Scan()
	motorStallTorque, nmb := strconv.ParseFloat(scanner.Text(), 64)
	if nmb == nil {
		gearReductionNeeded := torqForMaxForce / motorStallTorque
		step4(gearReductionNeeded)
	} else {
		step3(torqForMaxForce)
	}
}

func step4(gearReductionNeeded float64) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Amount Times Torque Needed (recomended 1.5 to 2)")
	scanner.Scan()
	timesTorque, nmb := strconv.ParseFloat(scanner.Text(), 64)
	if nmb == nil {
		desirableRangeOfGearReduction := gearReductionNeeded * timesTorque
		fmt.Println("Desirable Range of Gear Reduction: ")
		fmt.Println(desirableRangeOfGearReduction)
	} else {
		step4(gearReductionNeeded)
	}
}
