package main
import "fmt"

func binarySearch(list []int, key int){

    low := 0
    high := len(list) - 1

    for low <= high{
        mid := (low + high) / 2

        if list[mid] < key {
            low = mid + 1
        }else{
            high = mid - 1
        }
    }

    if low == len(list) || list[low] != key {
        fmt.Printf("%d not found in the list.",key)
        return
    }

    fmt.Printf("%d found at %dnd position.",key,low+1)
}


func main(){
    list := []int{3, 34, 56, 78, 133, 500, 679, 1092}
    fmt.Print("The list is:")
    fmt.Println(list)
    key := 7
    binarySearch(list, key)
}
