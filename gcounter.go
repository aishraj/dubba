package dubba

import "fmt"
import "reflect"

type CRDT interface {
	Merge(CRDT) *CRDT
}

type GCounter struct {
	data map[string]int
	CRDT
}

type emptyStruct struct{}

const defaultNodeName = "1"

//NewGCounter initializes and returns a new G-Counter
func NewGCounter() *GCounter {
	data := make(map[string]int)
	gCounter := GCounter{data: data}
	return &gCounter
}

func NewDataGCounter(d map[string]int) *GCounter {
	return &GCounter{data: d}
}

//Increment increases the value of the GCounter of the default node by 1
func (counter *GCounter) Increment() error {
	return counter.IncrementNode(defaultNodeName, 1)
}

//IncrementNode increases the value of a node by given delta
func (counter *GCounter) IncrementNode(nodeName string, delta int) error {
	if delta < 0 {
		return fmt.Errorf("Can't decrement a GCounter")
	}
	if _, ok := counter.data[nodeName]; ok {
		counter.data[nodeName] += delta
	} else {
		counter.data[nodeName] = delta
	}
	return nil
}

//Value returns the current value of G-Counter
func (counter *GCounter) Value() int {
	totalSum := 0
	for _, v := range counter.data {
		totalSum += v
	}
	return totalSum
}

//IsEqualTo returns true if this G-Counter equals the GCounter passed as the argument
func (counter *GCounter) IsEqualTo(c GCounter) bool {
	return reflect.DeepEqual(counter.data, c.data)
}

//Merge merges two G-Counters
func (counter *GCounter) Merge(another GCounter) *GCounter {
	keySet := make(map[string]emptyStruct, 0)
	retMap := make(map[string]int)
	for key := range counter.data {

		keySet[key] = emptyStruct{}
	}
	for key := range another.data {

		keySet[key] = emptyStruct{}
	}
	for key := range keySet {
		counts := make([]int, 0)
		if _, ok := counter.data[key]; ok {
			counts = append(counts, counter.data[key])
		}
		if _, ok := another.data[key]; ok {
			counts = append(counts, another.data[key])
		}
		maxVal := getMax(counts)
		retMap[key] = maxVal
	}

	return NewDataGCounter(retMap)
}

//getMax returns the the max element from a slice of +ve int
func getMax(c []int) int {
	prev := 0 //since +ve int assumption holds true
	for _, item := range c {
		if item > prev {
			prev = item
		}
	}
	return prev
}
