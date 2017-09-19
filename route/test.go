package main

// import (
//     "github.com/phbai/api/logic"
//     "github.com/bitly/go-simplejson"
//     "fmt"
// )

// func main() {
//     finished := make(chan bool)

//     go GetChannel()
//     <- finished
//     // fmt.Println(res)
// }

// func GetChannel() *simplejson.Json {
//     url := "Movie/GetChannel"
//     data := make(map[string]string)
//     result := make(chan *simplejson.Json)

//     logic.SendRequest(url, data, func(j *simplejson.Json) {
//         fmt.Println(j);
//         // finished <- true
//         result <- j
//     })
//     fmt.Println(result);
//     return <- result
// }