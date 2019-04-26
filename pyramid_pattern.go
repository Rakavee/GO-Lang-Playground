package main
import "fmt"

func main() {
    var n,i,j,k int
    fmt.Print("Enter number of rows:")
    fmt.Scan(&n)

    for i = 1; i <= n; i++ {
        for j = 1; j <= n-i; j++ {
            fmt.Print("  ")
        }
        for{
            k++
            fmt.Print(" P")
            if(k == 2*i-1){
                break
            }

        }
        k = 0
        fmt.Println("")
    }

}
