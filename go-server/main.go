package main


import (
    "fmt"
    "errors"
)


func main() {
    // var intNum int = 234;
    // fmt.Println("Integer number", intNum);
    
    // var floatNum float64 = 234.3;
    // fmt.Println("Float number", floatNum);
    
    // var1, var2 := 2,4
    // fmt.Println(var1, var2);
//    var myString string = "hello world" 
//    printString(myString)

    var myInt int = 234
    var myInt2 int = 45

    var answer, err =  divideNum(myInt, myInt2)

    if err != nil {
        fmt.Printf("%v\n", err)
    }else {
        fmt.Printf("result of the integer division is %v", answer)
    }

}

func printString(s string) {
    fmt.Println( s);
}


func divideNum(a int, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    var result int = a / b
    return result, nil
}
