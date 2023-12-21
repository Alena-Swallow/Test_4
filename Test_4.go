go
package main

import "fmt"

// MyStruct структура
type MyStruct struct {
    id   int
    name string
}

// NewMyStruct статический метод для создания экземпляра MyStruct
func NewMyStruct(name string, id int) MyStruct {
    return MyStruct{id, name}
}

// Name метод для получения имени структуры
func (ms MyStruct) Name() string {
    return ms.name
}

// Id метод для получения идентификатора структуры
func (ms MyStruct) Id() int {
    return ms.id
}

func main() {
    myStruct := NewMyStruct("John", 123)
    fmt.Println("Name:", myStruct.Name())
    fmt.Println("Id:", myStruct.Id())
}
