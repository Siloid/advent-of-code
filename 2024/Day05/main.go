package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func parseInput(path string) ([]string, []string) {
    file, _ := os.Open(path)
    scanner := bufio.NewScanner(file)

    var orderingRules []string
    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            // Found input file divider betweens rules and updates
            break
        } else {
            orderingRules = append(orderingRules, line)
        }
    }

    var pageUpdates []string
    for scanner.Scan() {
        line := scanner.Text()
        pageUpdates = append(pageUpdates, line)
    }
    file.Close()

    return orderingRules, pageUpdates
}

func createRegexs(rules []string) []*regexp.Regexp {
    // Creates a set of "reverse rule" regexs that when matched means the rule was broken
	ruleRegexs := make([]*regexp.Regexp, 0)

	for _, rule := range rules {
        pages := strings.Split(rule, "|")
        beforeNum := pages[0]
        afterNum := pages[1]

        ruleRegexStr := "(^|,)" + afterNum + "," + "(\\d+,)*" + beforeNum + "(,|$)"
		regex, _ := regexp.Compile(ruleRegexStr)
		ruleRegexs = append(ruleRegexs, regex)
	}
    return ruleRegexs
}

func isValidUpdate(update string, ruleRegexs []*regexp.Regexp) bool {
	for _, rule := range ruleRegexs {
		if rule.MatchString(update) {
            return false
        }
	}
    return true
}

func separateUpdates(pageUpdates []string, orderingRules []string) ([]string, []string) {
    // Separates valid vs invalid updates
    ruleRegexs := createRegexs(orderingRules)

    var validUpdates []string
    var invalidUpdates []string
    for _, update := range pageUpdates {
        if isValidUpdate(update, ruleRegexs) {
            validUpdates = append(validUpdates, update)
        } else {
            invalidUpdates = append(invalidUpdates, update)
        }
    }
    return validUpdates, invalidUpdates
}

func getUpdateMiddleValue(update string) int {
    pages := strings.Split(update, ",")
    middleValueStr := pages[len(pages) / 2]
    value, _ := strconv.Atoi((middleValueStr))
    return value
}

func sumUpdates(pageUpdates []string) int {
    updateTotal := 0
    for _, update := range pageUpdates {
        updateTotal += getUpdateMiddleValue(update)
    }
    return updateTotal
}

func createRulesMap(rules []string) map[string][]string {
    rulesMap := make(map[string][]string)
    for _, rule := range rules {
        pages := strings.Split(rule, "|")
        beforeNum := pages[0]
        afterNum := pages[1]
        afterSlice, inMap := rulesMap[beforeNum]
        if inMap {
            rulesMap[beforeNum] = append(afterSlice, afterNum)
        } else {
            rulesMap[beforeNum] = []string{afterNum}
        }
    }
    return rulesMap
}

func correctUpdate(update string, rulesMap map[string][]string ) string {
    // Always attempt to put a page as late as possible, as rules only
    // define what pages must come before others
    pages := strings.Split(update, ",")
    var correctedUpdate []string
    for _, nextPage := range pages {
        inserted := false
        afterPagesSlice, inMap := rulesMap[nextPage]
        if !inMap {
            // No rules apply, put it at the end 
            correctedUpdate = slices.Insert(correctedUpdate, len(correctedUpdate), nextPage)
            continue
        }
        for index, page := range correctedUpdate {
            if slices.Contains(afterPagesSlice, page) {
                correctedUpdate = slices.Insert(correctedUpdate, index, nextPage)
                inserted = true
                break
            }
        }
        if !inserted {
            // Doesn't have to come before any existing pages, put it at the end
            correctedUpdate = slices.Insert(correctedUpdate, len(correctedUpdate), nextPage)
        }
    }

    var correctedUpdateString string
    for i, page := range correctedUpdate {
        if i == 0 {
            correctedUpdateString = page
        } else {
            correctedUpdateString = correctedUpdateString + "," + page
        }
    }
    return correctedUpdateString
}

func fixInvalidUpdates(updates []string, rules []string) []string {
    rulesMap := createRulesMap(rules)
    var correctedUpdates []string
    for _, update := range updates {
        correctedUpdate := correctUpdate(update, rulesMap)
        correctedUpdates = append(correctedUpdates, correctedUpdate)
    }
    return correctedUpdates
}

func main() {
    orderingRules, pageUpdates := parseInput("./input.txt")
    validUpdates, invalidUpdates := separateUpdates(pageUpdates, orderingRules)
    
    // Part1
    total := sumUpdates(validUpdates)
    fmt.Printf("(Part 1) - valid update sum: %d\n", total)

    // Part2
    correctedUpdates := fixInvalidUpdates(invalidUpdates, orderingRules)
    total = sumUpdates(correctedUpdates)
    fmt.Printf("(Part 2) - corrected invalid update sum: %d\n", total)
}