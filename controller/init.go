package controllers

import (
	"customized_lp/upgrade_server/upgrade"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	tempFile = "data.dat.temp"
	dataFile = "data.dat"
)

var (
	// andoidCheckers upgrade.CheckerList
	// iosCheckers    upgrade.CheckerList
	checkers upgrade.CheckerList
)

func init() {
	//init upgrade checker info
	checkers = readDataToCheckers(dataFile)
	if checkers == nil {
		checkers = upgrade.CheckerList{}
	}
}

func readDataToCheckers(data string) upgrade.CheckerList {
	bs, err := ioutil.ReadFile(data)
	if err != nil {
		fmt.Println("read file err: ", err)
		return nil
	}
	var savedCheckers upgrade.CheckerList
	err = json.Unmarshal(bs, &savedCheckers)
	if err != nil {
		return nil
	}

	return savedCheckers
}

func saveCheckers(list upgrade.CheckerList) error {
	err := saveToTempFile(list, tempFile)
	if err != nil {
		return err
	}
	return swapTempAndDataFile(dataFile, tempFile)
}

func saveToTempFile(list upgrade.CheckerList, temp string) error {
	bs, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(temp, bs, os.ModePerm)
}

func swapTempAndDataFile(data, temp string) error {
	return os.Rename(temp, data)
}
