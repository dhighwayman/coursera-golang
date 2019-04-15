package main

import (
  "fmt"
  "sync"
)

const philosophers = 5
const chopsticks = 5
const plateSize = 3

type ChopS struct {
  sync.Mutex
}

type Philo struct {
  id int
  leftCS, rightCS *ChopS
  remainingFood int
}

func (p Philo) eat(wg *sync.WaitGroup, host chan int) {
  defer wg.Done()
  for p.remainingFood != 0 {
    host <- 1
    p.leftCS.Lock()
    p.rightCS.Lock()
    fmt.Printf("starting to eat %v\n", p.id)
    p.remainingFood = p.remainingFood - 1
    fmt.Printf("philosophe %v eating\n", p.id)
    p.rightCS.Unlock()
    p.leftCS.Unlock()
    <- host
    fmt.Printf("finishing eating %v\n", p.id)
  }
}

func main() {
  var wg sync.WaitGroup
  var host = make(chan int, 1)
  CSticks := make([]*ChopS, chopsticks)
  for i := 0; i < chopsticks; i++ {
    CSticks[i] = new(ChopS)
  }
  philos := make([]*Philo, philosophers)
  for i := 1; i <= philosophers; i++ {
    minCs := i-1
    maxCs := i % philosophers
    if ( maxCs < minCs) {
      minCs, maxCs = maxCs, minCs
    }
    philos[i-1] = &Philo{i, CSticks[minCs], CSticks[maxCs], plateSize}
    fmt.Printf("Philosophe %v created, his lower chopstick is %v and the higher is %v\n", i, minCs, maxCs)
  }

  for i := 1; i <= philosophers; i++ {
    wg.Add(1)
    go philos[i-1].eat(&wg, host)
  }
  wg.Wait()

}
