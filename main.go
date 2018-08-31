package main

import "io/ioutil"
import "errors"

type BrainGoFuck struct {
	source string
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

func NewBrainGoFuck () BrainGoFuck {
	return BrainGoFuck{""}
}

func main() {
	var interpretator BrainGoFuck = NewBrainGoFuck() 
	interpretator.ReadFile("hello.bf")
	interpretator.PrintSource()
}