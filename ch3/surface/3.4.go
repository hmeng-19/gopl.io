// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"io"
	"log"
	"net/http"
	"math"
	"fmt"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	/*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})
	*/
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		var cellsF int
		var widthF int
		var heightF int
		var err error
		param_cells := r.FormValue("cells")
		if param_cells != "" {
			cellsF, err = strconv.Atoi(param_cells)
			if err != nil {
				log.Print(err)
			}
		} else {
			cellsF = cells
		}

		param_width := r.FormValue("width")
		if param_width != "" {
			widthF, err = strconv.Atoi(param_width)
			if err != nil {
				log.Print(err)
			}
		} else {
			widthF = width
		}

		param_height := r.FormValue("height")
		if param_height != "" {
			heightF, err = strconv.Atoi(param_height)
			if err != nil {
				log.Print(err)
			}
		} else {
			heightF = height
		}

		SVGCreate(w, cellsF, widthF, heightF)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func SVGCreate(out io.Writer, cells int, width int, height int) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height)
			bx, by := corner(i, j, width, height)
			cx, cy := corner(i, j+1, width, height)
			dx, dy := corner(i+1, j+1, width, height)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j, width, height int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

