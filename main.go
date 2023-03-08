package main

import (
	"fmt"
	"math"
)

// define a point struct to represent individual points with x and y coordinates
type point struct {
    x, y int
}

// define the points as global variables
var (
    a = point{2, 2}
    b = point{0, 5}
    c = point{8, 10}
    d = point{2, 9}
    e = point{3, 5}
    f = point{1, 10}
    g = point{7, 14}
)

func main() {
    // define slice of all points to be included in the path
    points := []point{a, b, c, d, e, f, g}

    // calculate the distance between every pair of points
    distances := make([][]float64, len(points))
    for i := range distances {
        distances[i] = make([]float64, len(points))
        for j := range distances[i] {
            if i == j {
                distances[i][j] = 0
                continue
            }
            distances[i][j] = math.Sqrt(float64((points[i].x-points[j].x)*(points[i].x-points[j].x) + (points[i].y-points[j].y)*(points[i].y-points[j].y)))
        }
    }

    // initialize the path as starting at A
    path := []int{0}
    // create a map to keep track of unvisited points
    unvisited := make(map[int]bool)
    // add all points except A to the unvisited map
    for i := 1; i < len(points); i++ {
        unvisited[i] = true
    }

    // while there are still unvisited points
    for len(unvisited) > 0 {
        // find the closest unvisited point to the last visited point
        bestDist := math.Inf(1)
        bestNext := -1
        for next := range unvisited {
            dist := distances[path[len(path)-1]][next]
            if dist < bestDist {
                bestDist = dist
                bestNext = next
            }
        }

        // add the closest unvisited point to the path and mark it as visited
        path = append(path, bestNext)
        delete(unvisited, bestNext)
    }

    // print out the final result
    fmt.Print("The most optimal path starting from A and going through every point is: [")
    for _, p := range path {
        // translate the point index to its corresponding letter for readability
        switch p {
        case 0:
            fmt.Print("A ")
        case 1:
            fmt.Print("B ")
        case 2:
            fmt.Print("C ")
        case 3:
            fmt.Print("D ")
        case 4:
            fmt.Print("E ")
        case 5:
            fmt.Print("F ")
        case 6:
            fmt.Print("G ")
        default:
            fmt.Print("? ")
        }
    }
    fmt.Println("]")
}
