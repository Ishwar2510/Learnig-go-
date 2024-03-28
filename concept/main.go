package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	// wap1

	// to demostrate working of a channel  and gotoutine
	// the channel can send unli=mied file
	// c := make(chan int)
	// demon1(40, c)

	// wap2

	//demostrate the working of map with where te map is used by different go routine
	// to change it value .
	// wap2()

	// wap3
	// impliment linked list
	// linkedlist()

	// wap 4
	// impliment Queue
	// queue()

	// wap5
	// impliment set
	// set()

	// wap6
	// impliment tree data structure
	// tree()

}

// ******************   start  wap2   *************************
// use of mutex .
// a struct is required because we want to
// combine  both the map and mutex as one entity

type myControlledMap struct {
	mutex sync.Mutex
	myMap map[string]int
}

// define all the method u need
func (m *myControlledMap) insert(key string, value int) {

	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.myMap[key] = value

}
func (m *myControlledMap) delete(key string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.myMap, key)

}
func (m *myControlledMap) get(key string) (int, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if key, ok := m.myMap[key]; !ok {
		return key, errors.New(" key not available")
	} else {
		return key, nil
	}

}

func wap2() {
	myMapWithMutex := myControlledMap{myMap: make(map[string]int)}
	// here we will use waitgroup to keep track of all the go process .
	var wg sync.WaitGroup
	for i := 1; i < 30; i++ {
		myMapWithMutex.insert(strconv.Itoa(i), i)
	}

	for i := 1; i < 30; i++ {
		wg.Add(1)
		if i%5 == 0 {
			// since each function are to be used as go
			go func(i int) {
				myMapWithMutex.insert(strconv.Itoa(i), i*5)
				wg.Done()
			}(i) //  as javascript go also has closure, if we donot  pass i then the go  will refer i =30
		} else if i%2 == 0 {
			go func(i int) {
				input := i
				if result, err := myMapWithMutex.get(strconv.Itoa(i)); err != nil {
					fmt.Println(err, input)
				} else {
					fmt.Println("the data for", i, "is ", result)
				}
				wg.Done()
			}(i)
		} else {
			go func(i int) {
				myMapWithMutex.delete(strconv.Itoa(i))
				wg.Done()
			}(i)
		}

	}
	wg.Wait()
	for key, value := range myMapWithMutex.myMap {
		fmt.Println(key, "--->", value)
	}

}

// **************************   end wap2  *****************************

// *************************   start wap1 ******************************
func demon1(times int, c chan int) {
	// start a go routine that will do some work and send the response over the channel
	go func() {
		for i := 1; i < times; i++ {
			time.Sleep(1 * time.Second)
			c <- i
		}
		close(c)
	}()
	// infinte loop runs till the chaneel is not closed
	for {
		data, isOpen := <-c
		if !isOpen {
			fmt.Println("channel is closed")
			return
		}
		fmt.Println(data)

	}
}

// *********************** end wap 2 ****************************************

// *********************** start wap 3 ****************************

func linkedlist() {
	linkedList := NewLinkedList()
	for _, value := range []int{1, 2, 3, 4, 5, 6, 7} {
		linkedList.Insert(value)
	}
	linkedList.PrintList()
	linkedList.ReverseList()
	linkedList.PrintList()
	linkedList.RemoveFirst()
	linkedList.RemoveLast()
	linkedList.PrintList()
	fmt.Println(linkedList.GetMiddle())

}

// ************************** end wap 3 *****************************

// *************************** start wap4 ****************************

func queue() {
	myqueue := NewQueue()
	for _, value := range []int{2, 3, 4, 5, 1, 13, 4, 45} {
		myqueue.Add(value)
	}
	myqueue.PrintQueue()
	myqueue.Remove()
	myqueue.PrintQueue()
}

// ************************** end wap 4 *****************************

// ************************** start wap 5 *****************************

func set() {
	set := newSet()

	for _, value := range []int{1, 2, 3, 4, 4, 5, 6, 1, 2, 3, 4, 5} {
		set.Add(value)
	}
	set.Printset()
	set.Remove(3)
	set.Remove(40)
	set.Printset()

	set2 := newSet()
	for _, value := range []string{"hi", "how", "how", "are", "you", "all", "you"} {
		set2.Add(value)
	}
	set2.Printset()

}

// ************************** end wap 5 *****************************

// ************************** end wap 6*****************************

func tree() {
	tree := NewTree()
	for _, value := range []int{20, 10, 2, 3, 110, 56, 34, 1, 5, 6, 200, 300, 500, 600, 700, 800, 900, 120} {
		tree.Insert(value)
	}
	tree.InOrder()
	fmt.Println()
	tree.Preorder()
	fmt.Println()
	tree.Postorder()
	fmt.Println()
	tree.Leftview()
}

// ************************** start wap 6 *****************************
