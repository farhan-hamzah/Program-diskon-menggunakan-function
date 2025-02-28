package main
import "fmt"
func diskon(belanja float64, member bool)float64{
    var diskon float64
    if belanja > 100000 && !member{
        diskon = 0.9
    }else if diskon > 100000 && member{
        diskon = 0.95
    }else{
        diskon = 1
    }
    return(float64(belanja)*diskon)
}
func main(){
    var masukan, hasil float64
    var kartu bool
    fmt.Scan(&masukan, &kartu)
    hasil = diskon(masukan, kartu)
    fmt.Print(hasil)
}