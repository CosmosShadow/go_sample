package main

import "fmt"

func main() {
        /* 创建map */
        countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

        fmt.Println("原始地图")
        printMap(countryCapitalMap)

        /*删除元素*/ 
        delete(countryCapitalMap, "France")
        fmt.Println("\nFrance条目被删除")
        printMap(countryCapitalMap)
}

func printMap(theMap map[string] string) {
        for key := range theMap {
                fmt.Println(key, ": ", theMap[key])
        }
}