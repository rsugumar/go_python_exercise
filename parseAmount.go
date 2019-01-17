package main

import "C"

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
)

type Configuration struct {
	SupportedCurrencySymbol []string `json:"supportedCurrencySymbol"`
	SupportedRange          []string `json:"supportedRange"`
}

func parseConfigFile(fileName string) (Configuration, error) {
	configuration := Configuration{}
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("error: ", err)
		return configuration, errors.New("Error occurred opening file")
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&configuration); err != nil {
		fmt.Println("error: ", err)
		return configuration, errors.New("Error occurred opening file")
	}
	//fmt.Println(configuration)
	return configuration, nil
}

func getCurrencyAmount(currencyVal string, currencyStr []string) (float64, error) {
	var currency string
	var no float64
	fmt.Sscanf(currencyVal, "%s %f", &currency, &no)

	sort.Strings(currencyStr)
	found := sort.SearchStrings(currencyStr, currency)
	if found < len(currencyStr) && currencyStr[found] == currency {
		return no, nil
	} else {
		e := fmt.Sprintf("Invalid currency found: %s; Expected Currency: %s",
			currency, currencyStr)
		return 0, errors.New(e)
	}
}

func getRangeValue(amount float64, rangeVals []string) (float64, float64, error) {
	var startRange, endRange float64
	isInRange := false
	for i := range rangeVals {
		fmt.Sscanf(rangeVals[i], "%f-%f", &startRange, &endRange)
		if amount >= startRange && amount <= endRange {
			isInRange = true
			break
		}
	}

	if !isInRange {
		return -1, -1, errors.New("Not in range")
	} else {
		return startRange, endRange, nil
	}
}

//export processInput
func processInput(configFileName string, amount string) {
	config, err := parseConfigFile(configFileName)
	if err != nil {
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
		fmt.Println("Config file - Error found. Valid JSON config expected")
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
		return
	}

	no, err := getCurrencyAmount(amount, config.SupportedCurrencySymbol)
	if err != nil {
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
		fmt.Println("getCurrencyAmount() failed! Invalid input: ", amount)
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
		return
	}
	start, end, err := getRangeValue(no, config.SupportedRange)
	fmt.Println("------------------------------------------------------------------------------------------------------------------")
	fmt.Println("*******")
	fmt.Println("RESULT: ")
	fmt.Println("*******")
	if err != nil {
		var rangeStr string
		for i := range config.SupportedRange {
			rangeStr += config.SupportedRange[i] + ", "
		}
		fmt.Println(fmt.Sprintf("%s NOT in the given range: %s", amount, rangeStr))
	} else {
		fmt.Println(fmt.Sprintf("%s falls WITHIN range %f - %f", amount, start, end))
	}
	fmt.Println("------------------------------------------------------------------------------------------------------------------")
}

func main() {
}
