package main

import (
        "econding/json"
        "fmt"
)

type person struct {
        First string
        Last string
        Ltk bool
        Items []string
}

func main() {
        p1 := person{
                First: "James",
                Last: "Bond",
                LTK: true,
                Items: []string{
                        "martini",
                        "walther PPK",
                        "Astin Martin",
                },
        }
        bs, err := json.Marshal(p1)
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(p1)
        fmt.Println("-------")
        fmt.Println(string(bs))
}
