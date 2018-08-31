package main

import "io/ioutil"
import "errors"

type BrainGoFuck struct {
	source string
	memory []uint8
	carretePos uint
}

func (this *BrainGoFuck) ReadFile(fileName string) error {
	b, err := ioutil.ReadFile(fileName)
    if err != nil {
        return errors.New("Cannot open file")
    }
    this.source = string(b)
    return nil
}

func (this BrainGoFuck) PrintSource() {
	println(this.source[0])
}

func (this *BrainGoFuck) Step() {

}

func NewBrainGoFuck () BrainGoFuck {
	var memoryCells uint = 30000
	var carretePos uint = 0
	return BrainGoFuck{"", make([]uint8, memoryCells), carretePos}
}

func main() {
	var interpretator BrainGoFuck = NewBrainGoFuck() 
	interpretator.ReadFile("hello.bf")
	interpretator.PrintSource()
}