// Conway's Game of Life Imprementation
package main

import (
  "bytes"
  "fmt"
  "math/rand"
  "time"
)

type World struct {
  s     [][]bool
  w, h  int
}

func NewWorld(w, h int) (World) {
  s := make([][]bool, h)
  for i := range s {
    s[i] = make([]bool, w)
  }

  return World{s: s, w: w, h: h}
}

func (world *World) InitWorld() {
  rand.Seed(time.Now().UTC().UnixNano())

  for h:= 0; h < world.h; h++ {
    for w:= 0; w < world.w; w++ {
      if rand.Int31n(2) == 0 {
        world.s[h][w] = false
      } else {
        world.s[h][w] = true
      }
    }
  }
}

func (world *World) Tick() {
  evolution := NewWorld(world.w, world.h)

  for h := 0; h < world.h; h++ {
    for w := 0; w < world.w; w++ {
      evolution.s[h][w] = world.CellNextState(h, w)
    }
  }

  *world = evolution
}

func (world *World) CellNextState(h, w int) (bool) {
  alive := 0
  for i := -1; i <= 1; i++ {
    for j := -1; j <= 1; j++ {
      if (i != 0 || j !=0) && world.CellAlive(h+i, w+j)  {
        alive++
      }
    }
  }
  return alive == 3 || alive == 2 && world.s[h][w]
}

func (world *World) CellAlive(h, w int) bool {
  if (h < 0 || h >= world.h || w < 0 || w >= world.w) {
    return false
  } else {
    return world.s[h][w]
  }
}

func (world *World) String() string {
  var buf bytes.Buffer
  for h := 0; h < world.h; h++ {
    for w := 0; w < world.w; w++ {
      b := byte(' ')
      if world.CellAlive(h, w) {
        b = '*'
      }
      buf.WriteByte(b)
    }
    buf.WriteByte('\n')
  }
  return buf.String()
}

func main() {
  world := NewWorld(40, 15)
  world.InitWorld()

  for i := 0; i < 300; i++ {
    world.Tick()
    fmt.Println(world.String())
    time.Sleep(time.Second / 30)
  }
}
