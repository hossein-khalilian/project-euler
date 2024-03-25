package main

import (
    "fmt"
)

func count_sundays() int{
    months := []string {"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
    days := []string {"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
    month_days := make(map[string] int, 0)
    month_days["sep"] = 30
    month_days["apr"] = 30
    month_days["jun"] = 30
    month_days["nov"] = 30

    month_days["jan"] = 31
    month_days["mar"] = 31
    month_days["may"] = 31
    month_days["jul"] = 31
    month_days["aug"] = 31
    month_days["oct"] = 31
    month_days["dec"] = 31

    month_days["feb_rain"] = 28
    month_days["feb_leap"] = 29

    year := 1900
    month := 0
    index := 0
    count := 0    
    for {
        // fmt.Println(months[month % 12], year)
        // fmt.Println(month % 12)
        for i := range month_days[months[month % 12]]{
            i += 1
            if year >= 1901{
                if i == 1 && days[(index)%7] == "sunday"{
                    count += 1
                }
            }
            index += 1
        }

        month += 1
        if month % 12 == 0{
            year += 1
            if year % 4 == 0 && year % 100 != 0{
                month_days["feb"] = month_days["feb_leap"]
                
            } else {
                month_days["feb"] = month_days["feb_rain"]
            }
            if year % 400 == 0{
                month_days["feb"] = month_days["feb_leap"]
            }
            
            if year > 2000{
                break
            }
        }
    }

    return count
}

func main(){
    fmt.Println(count_sundays())
}