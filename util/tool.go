package util

import "math/rand"
import "strings"

func RandomUserId() string {
    arr := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
    var str []string

    for i := 0; i < 32; i++ {
        index := rand.Intn(len(arr) - 1)
        str = append(str, arr[index])
    }

    result := strings.Join(str, "")
    return result[0:8] + "-" + result[8:12] + "-" + result[12:16] + "-" + result[16:20] + "-" + result[20:32];
}































