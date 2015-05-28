package main
import "os"
import "fmt"
import "strconv"
import "container/list"
// import "time"

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage - {NumberOfBalls} [Minutes]")
		os.Exit(1)
	}	
	var ballnum, _ = strconv.Atoi(os.Args[1])
	var minutes int = -1
	if len(os.Args) == 3 {
		minutes, _ = strconv.Atoi(os.Args[2])
	}
	if ballnum => 27 && ballnum <= 127 {
		if minutes != -1 {
			m := list.New()
			for i := 1; i <= ballnum; i++{
				m.PushBack(i)	
			}
			getMinuteState(m, minutes)
		} else {
			m := make([]int, ballnum, 127)
			for i := 1; i <= ballnum; i++{
				m[i-1] = i
			}
			var days = getCycleLength(m)
			fmt.Printf("%d balls cycle after %d days.\n", ballnum, days)
		}
	}
}

func getCycleLength(m []int) int{
	var done bool = false
	var ball int = -1
	var tball int = -1

	singles := make([]int, 0, 4)
	fives := make([]int, 0, 11)
	hours := make([]int, 0, 11)
	var cycles int = 0

	for done != true {
		ball, m = m[0], m[1:]
		if len(singles) == 4 {
			for len(singles) != 0 {
				tball, singles = singles[len(singles)-1], singles[:len(singles)-1]
				m = append(m, tball)
			}
			if len(fives) == 11 {
				for len(fives) != 0 {
					tball, fives = fives[len(fives)-1], fives[:len(fives)-1]
					m = append(m, tball)
				}
				if len(hours) == 11 {
					//Return all the balls in the hour queue to the main
					//queue, then return the current ball to the queue	

					//Remove each element from the list, they are in-order, so 
					//this order will be preserved when the balls return to
					//the main tray
					for len(hours) != 0 {
						tball, hours = hours[len(hours)-1], hours[:len(hours)-1]
						m = append(m, tball)
					}
					m = append(m, ball)
					// var i int = 0
					// fmt.Printf("Length is %d  Capacity is %d ", len(m), cap(m))
					// for i = 0; i < len(m); i++{
					// 	fmt.Printf("%d ", m[i])	
					// }
					// fmt.Printf("\n")
					// time.Sleep(200 * time.Millisecond)
					cycles++
					if(checkCycle(m) == true){
						done = true
					}
				} else {
					hours = append(hours, ball)
				}
			} else {
				fives = append(fives, ball)
			}
		} else {
			singles = append(singles, ball)
		}
	}

	return cycles / 2.0
}


func checkCycle(m []int) bool{
	var i int = 1
	for i = 1; i <= len(m); i++ {
		if i != m[i-1] {
			return false
		}else{
			continue
		}
	}
	return true
}

func getMinuteState(m *list.List, minutes int){
	var counter int = 0
	var ball int = -1
	var tball int = -1

	singles := list.New()
	fives := list.New()
	hours := list.New()
	
	for counter < minutes {
		ball = m.Front().Value.(int)
		m.Remove(m.Front())

		//fmt.Printf("The value of ball is %d\n", ball)
		
		if singles.Len() == 4 {
			for singles.Len() != 0 {
				e := singles.Back()
				tmp := e.Prev()
				tball = e.Value.(int)
				singles.Remove(e)
				m.PushBack(tball)
				e = tmp
			}

			if fives.Len() == 11 {

				for fives.Len() != 0 {
					e := fives.Back()
					tmp := e.Prev()
					tball = e.Value.(int)
					fives.Remove(e)
					m.PushBack(tball)
					e = tmp
				}

				if hours.Len() == 11 {
					//Return all the balls in the hour queue to the main
					//queue, then return the current ball to the queue
					

					//Remove each element from the list, they are in-order, so 
					//this order will be preserved when the balls return to
					//the main tray

					for hours.Len() != 0 {
						e := hours.Back()
						tmp := e.Prev()
						tball = e.Value.(int)
						hours.Remove(e)
						m.PushBack(tball)
						e = tmp
					}
					m.PushBack(ball)
				
				} else {
					hours.PushBack(ball)
				}
			} else {
				fives.PushBack(ball)
			}
		} else {
			singles.PushBack(ball)
		}
		counter++
	}

	fmt.Printf("{\"Min\":[")
	if singles.Len() == 0 {
		fmt.Printf("],")
	} else {
		for e := singles.Front(); e != nil; e = e.Next() {
			if e.Next() == nil {
				fmt.Printf("%d],", e.Value.(int))
			} else {
				fmt.Printf("%d,", e.Value.(int))
			}
		}
	}
	fmt.Printf("\"FiveMin\":[")
	if fives.Len() == 0 {
		fmt.Printf("],")
	} else {
		for e := fives.Front(); e != nil; e = e.Next() {
			if e.Next() == nil {
				fmt.Printf("%d],", e.Value.(int))
			} else {
				fmt.Printf("%d,", e.Value.(int))
			}
		}
	}
	fmt.Printf("\"Hour\":[")
	if hours.Len() == 0 {
		fmt.Printf("],")
	} else {
		for e := hours.Front(); e != nil; e = e.Next() {
			if e.Next() == nil {
				fmt.Printf("%d],", e.Value.(int))
			} else {
				fmt.Printf("%d,", e.Value.(int))
			}
		}
	}
	fmt.Printf("\"Main\":[")
	if m.Len() == 0 {
		fmt.Printf("],")
	} else {
		for e := m.Front(); e != nil; e = e.Next() {
			if e.Next() == nil {
				fmt.Printf("%d]}\n", e.Value.(int))
			} else {
				fmt.Printf("%d,", e.Value.(int))
			}
		}
	}

	return
}
