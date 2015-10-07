package main

import (
    "fmt"
    "net/http"
    //"log"
    "html/template"
    "math/rand"
    "time"
    "unicode/utf8"
)

type State struct {
	L1, L2, L3, L4, L5, N1, N2, N3, N4, N5 string
}

var yellow string = "http://cadler.co/wp-content/uploads/2015/09/yellow-e1442941699604.jpg"
var black string = "http://cadler.co/wp-content/uploads/2015/09/black-e1442941689285.jpg"
var green string = "http://cadler.co/wp-content/uploads/2015/10/green-e1443791954640.png"

var (
	l11 string = RandColor()
	l12 string = RandColor()
	l13 string = RandColor()
	l14 string = RandColor()
	l15 string = RandColor()
	)
var (
	l21 string = RandColor()
	l22 string = RandColor()
	l23 string = RandColor()
	l24 string = RandColor()
	l25 string = RandColor()
	)
var (
	l31 string = RandColor()
	l32 string = RandColor()
	l33 string = RandColor()
	l34 string = RandColor()
	l35 string = RandColor()
	)
var (
	l41 string = RandColor()
	l42 string = RandColor()
	l43 string = RandColor()
	l44 string = RandColor()
	l45 string = RandColor()
	)
var (
	l51 string = RandColor()
	l52 string = RandColor()
	l53 string = RandColor()
	l54 string = RandColor()
	l55 string = RandColor()
	)



const head = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>
		Lights
	</title>
</head>
	<body>
`

const body = `
		<div class="floated_img">
			<a href="{{.N1}}"><img src="{{.L1}}" alt="square" /></a>
			<a href="{{.N2}}"><img src="{{.L2}}" alt="square" /></a>
			<a href="{{.N3}}"><img src="{{.L3}}" alt="square" /></a>
			<a href="{{.N4}}"><img src="{{.L4}}" alt="square" /></a>
			<a href="{{.N5}}"><img src="{{.L5}}" alt="square" /></a>
		</div>
`

const solveHTML = `<input type=button onClick="parent.location='solve'" value='Solve row'>`

const tail = `
	</body>
</html>
`

func Index(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("lights").Parse(body)
    url := r.URL.Path
    urlToTest := url[1:]
    length := utf8.RuneCountInString(urlToTest)
    if err == nil {
    	if length == 0 {
	    	fmt.Fprintf(w, head)
	   		lights1 := State{l11, l12, l13, l14, l15, "11", "12", "13", "14", "15"}
	   		lights2 := State{l21, l22, l23, l24, l25, "21", "22", "23", "24", "25"}
	   		lights3 := State{l31, l32, l33, l34, l35, "31", "32", "33", "34", "35"}
	   		lights4 := State{l41, l42, l43, l44, l45, "41", "42", "43", "44", "45"}
	   		lights5 := State{l51, l52, l53, l54, l55, "51", "52", "53", "54", "55"}
	   		tmpl.Execute(w, lights1)
	   		tmpl.Execute(w, lights2)
	   		tmpl.Execute(w, lights3)
	   		tmpl.Execute(w, lights4)
	   		tmpl.Execute(w, lights5)
        fmt.Fprintf(w, solveHTML)
	   		fmt.Fprintf(w, tail)
   		}

   		if length > 1 && urlToTest != "solve" {
   			switch urlToTest[0:1] {
   				default:
   				case "1":
   					switch urlToTest[1:2] {
   						default:
   						case "1":
   							SwitchColor(&l11, true)
   							SwitchColor(&l12, false)
   							SwitchColor(&l21, false)
		   				case "2":
		   					SwitchColor(&l11, false)
		   					SwitchColor(&l12, true)
		   					SwitchColor(&l13, false)
		   					SwitchColor(&l22, false)
		   				case "3":
		   					SwitchColor(&l12, false)
		   					SwitchColor(&l13, true)
		   					SwitchColor(&l14, false)
		   					SwitchColor(&l23, false)
		   				case "4":
		   					SwitchColor(&l13, false)
		   					SwitchColor(&l14, true)
		   					SwitchColor(&l15, false)
		   					SwitchColor(&l24, false)
		   				case "5":
		   					SwitchColor(&l14, false)
		   					SwitchColor(&l15, true)
		   					SwitchColor(&l25, false)
   					}
   				case "2":
   					switch urlToTest[1:2] {
   						default:
   						case "1":
   							SwitchColor(&l11, false)
   							SwitchColor(&l21, true)
   							SwitchColor(&l31, false)
   							SwitchColor(&l22, false)
		   				case "2":
		   					SwitchColor(&l12, false)
		   					SwitchColor(&l21, false)
		   					SwitchColor(&l22, true)
		   					SwitchColor(&l23, false)
		   					SwitchColor(&l32, false)
		   				case "3":
		   					SwitchColor(&l13, false)
		   					SwitchColor(&l22, false)
		   					SwitchColor(&l23, true)
		   					SwitchColor(&l24, false)
		   					SwitchColor(&l33, false)
		   				case "4":
		   					SwitchColor(&l14, false)
		   					SwitchColor(&l23, false)
		   					SwitchColor(&l24, true)
		   					SwitchColor(&l25, false)
		   					SwitchColor(&l34, false)
		   				case "5":
		   					SwitchColor(&l24, false)
		   					SwitchColor(&l25, true)
		   					SwitchColor(&l15, false)
		   					SwitchColor(&l35, false)
   					}
   				case "3":
   					switch urlToTest[1:2] {
   						default:
   						case "1":
   							SwitchColor(&l21, false)
   							SwitchColor(&l31, true)
   							SwitchColor(&l32, false)
   							SwitchColor(&l41, false)
		   				case "2":
		   					SwitchColor(&l22, false)
		   					SwitchColor(&l31, false)
		   					SwitchColor(&l32, true)
		   					SwitchColor(&l33, false)
		   					SwitchColor(&l42, false)
		   				case "3":
		   					SwitchColor(&l23, false)
		   					SwitchColor(&l32, false)
		   					SwitchColor(&l33, true)
		   					SwitchColor(&l34, false)
		   					SwitchColor(&l43, false)
		   				case "4":
		   					SwitchColor(&l24, false)
		   					SwitchColor(&l33, false)
		   					SwitchColor(&l34, true)
		   					SwitchColor(&l35, false)
		   					SwitchColor(&l44, false)
		   				case "5":
		   					SwitchColor(&l34, false)
		   					SwitchColor(&l35, true)
		   					SwitchColor(&l25, false)
		   					SwitchColor(&l45, false)
   					}
   				case "4":
   					switch urlToTest[1:2] {
   						default:
   						case "1":
   							SwitchColor(&l31, false)
   							SwitchColor(&l41, true)
   							SwitchColor(&l42, false)
   							SwitchColor(&l51, false)
		   				case "2":
		   					SwitchColor(&l32, false)
		   					SwitchColor(&l41, false)
		   					SwitchColor(&l42, true)
		   					SwitchColor(&l43, false)
		   					SwitchColor(&l52, false)
		   				case "3":
		   					SwitchColor(&l33, false)
		   					SwitchColor(&l42, false)
		   					SwitchColor(&l43, true)
		   					SwitchColor(&l44, false)
		   					SwitchColor(&l53, false)
		   				case "4":
		   					SwitchColor(&l34, false)
		   					SwitchColor(&l43, false)
		   					SwitchColor(&l44, true)
		   					SwitchColor(&l45, false)
		   					SwitchColor(&l54, false)
		   				case "5":
		   					SwitchColor(&l44, false)
		   					SwitchColor(&l45, true)
		   					SwitchColor(&l35, false)
		   					SwitchColor(&l55, false)
   					}
   				case "5":
   					switch urlToTest[1:2] {
   						default:
   						case "1":
   							SwitchColor(&l41, false)
   							SwitchColor(&l51, true)
   							SwitchColor(&l52, false)
		   				case "2":
		   					SwitchColor(&l42, false)
		   					SwitchColor(&l51, false)
		   					SwitchColor(&l52, true)
		   					SwitchColor(&l53, false)
		   				case "3":
		   					SwitchColor(&l43, false)
		   					SwitchColor(&l52, false)
		   					SwitchColor(&l53, true)
		   					SwitchColor(&l54, false)
		   				case "4":
		   					SwitchColor(&l44, false)
		   					SwitchColor(&l53, false)
		   					SwitchColor(&l54, true)
		   					SwitchColor(&l55, false)
		   				case "5":
		   					SwitchColor(&l45, false)
		   					SwitchColor(&l54, false)
		   					SwitchColor(&l55, true)
   					}
   			}
	    	fmt.Fprintf(w, head)
	   		lights1 := State{l11, l12, l13, l14, l15, "11", "12", "13", "14", "15"}
	   		lights2 := State{l21, l22, l23, l24, l25, "21", "22", "23", "24", "25"}
	   		lights3 := State{l31, l32, l33, l34, l35, "31", "32", "33", "34", "35"}
	   		lights4 := State{l41, l42, l43, l44, l45, "41", "42", "43", "44", "45"}
	   		lights5 := State{l51, l52, l53, l54, l55, "51", "52", "53", "54", "55"}
	   		tmpl.Execute(w, lights1)
	   		tmpl.Execute(w, lights2)
	   		tmpl.Execute(w, lights3)
	   		tmpl.Execute(w, lights4)
	   		tmpl.Execute(w, lights5)
   			fmt.Fprintf(w, "You just pressed square %v-%v", url[1:2], url[2:3])
        fmt.Fprintf(w, solveHTML)
	   		fmt.Fprintf(w, tail)
   		}
      if urlToTest == "solve" {
        Solver()
        fmt.Fprintf(w, head)
        lights1 := State{l11, l12, l13, l14, l15, "11", "12", "13", "14", "15"}
        lights2 := State{l21, l22, l23, l24, l25, "21", "22", "23", "24", "25"}
        lights3 := State{l31, l32, l33, l34, l35, "31", "32", "33", "34", "35"}
        lights4 := State{l41, l42, l43, l44, l45, "41", "42", "43", "44", "45"}
        lights5 := State{l51, l52, l53, l54, l55, "51", "52", "53", "54", "55"}
        tmpl.Execute(w, lights1)
        tmpl.Execute(w, lights2)
        tmpl.Execute(w, lights3)
        tmpl.Execute(w, lights4)
        tmpl.Execute(w, lights5)
        fmt.Fprintf(w, solveHTML)
        fmt.Fprintf(w, tail)
      }
    }
    if err != nil {
        panic(err)
    }
}

func SwitchColor(color *string, switchGreen bool) {
	if *color == yellow {
		*color = black
	} else if *color == black {
		*color = yellow
	} else {
		if switchGreen {
			if *color == green {
        *color = black
      }
		}
	}
}

func RandColor() (string){
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	colorNumber := r.Intn(2) + 1;
	var toReturn string

	switch (colorNumber) {
		default: 	toReturn = yellow //"http://cadler.co/wp-content/uploads/2015/09/yellow-e1442941699604.jpg"
		case 1:		toReturn = yellow //"http://cadler.co/wp-content/uploads/2015/09/yellow-e1442941699604.jpg"
		case 2:		toReturn = black  //"http://cadler.co/wp-content/uploads/2015/09/black-e1442941689285.jpg"
	}
	return toReturn
}

func FindRow() int{
  row1 := []string{l11, l12, l13, l14, l15}
  row2 := []string{l21, l22, l23, l24, l25}
  row3 := []string{l31, l32, l33, l34, l35}
  row4 := []string{l41, l42, l43, l44, l45}
  row5 := []string{l51, l52, l53, l54, l55}
  var toReturn int
  yellowCount1, yellowCount2, yellowCount3, yellowCount4, yellowCount5 := 0, 0, 0, 0, 0
    for i := 0; i < 5; i++ {
      if row1[i] == yellow {
        toReturn = 1
        fmt.Println("FindRow = 1")
        yellowCount1++
      }
      if i == 4 {
        continue RowLoop2
      }
    }
  RowLoop2:
    if yellowCount1 == 0 {
      for i := 0; i < 5; i++ {
        if row2[i] == yellow {
          toReturn = 2
          fmt.Println("FindRow = 2")
          yellowCount2++
        }
      }
    }
  RowLoop3:
    for i := 0; i < 5; i++ {
      if row3[i] == yellow {
        toReturn = 3
        fmt.Println("FindRow = 3")
        break
      }
  }
  RowLoop4:
    for i := 0; i < 5; i++ {
      if row4[i] == yellow {
        toReturn = 4
        fmt.Println("FindRow = 4")
        break
      }
  }
  RowLoop5:
    for i := 0; i < 5; i++ {
      if row5[i] == yellow {
        toReturn = 5
        fmt.Println("FindRow = 5")
        break
      }
  }

  return toReturn
}

func Solver() {
  row := FindRow()
  row1 := []string{l11, l12, l13, l14, l15}
  row2 := []string{l21, l22, l23, l24, l25}
  row3 := []string{l31, l32, l33, l34, l35}
  row4 := []string{l41, l42, l43, l44, l45}
  row5 := []string{l51, l52, l53, l54, l55}
  l12, l14, l15 := &l12, &l14, &l15
  l21, l22, l23, l24, l25 := &l21, &l22, &l23, &l24, &l25
  l31, l32, l33, l34, l35 := &l31, &l32, &l33, &l34, &l35
  l41, l42, l43, l44, l45 := &l41, &l42, &l43, &l44, &l45
  l51, l52, l53, l54, l55 := &l51, &l52, &l53, &l54, &l55
  //yellow := &yellow
	if row == 1 {
		for i := 0; i < 5; i++ {
			if row1[i] == yellow {
				switch i {
					default:
						fmt.Println("failed in row 1")
					case 0:
						*l21 = green
					case 1:
						*l22 = green
					case 2:
						*l23 = green
					case 3:
						*l24 = green
					case 4:
						*l25 = green
				}
			}
		}
	}
  if row == 2 {
		for i := 0; i < 5; i++ {
			if row2[i] == yellow {
				switch i {
					default:
						fmt.Println("failed in row 2")
					case 0:
						*l31 = green
					case 1:
						*l32 = green
					case 2:
						*l33 = green
					case 3:
						*l34 = green
					case 4:
						*l35 = green
				}
			}
		}
	}
  if row == 3 {
		for i := 0; i < 5; i++ {
			if row3[i] == yellow {
				switch i {
					default:
						fmt.Println("failed in row 3")
					case 0:
						*l41 = green
					case 1:
						*l42 = green
					case 2:
						*l43 = green
					case 3:
						*l44 = green
					case 4:
						*l45 = green
				}
			}
		}
	}
  if row == 4 {
		for i := 0; i < 5; i++ {
			if row4[i] == yellow {
				switch i {
					default:
						fmt.Println("failed in row 4")
					case 0:
						*l51 = green
					case 1:
						*l52 = green
					case 2:
						*l53 = green
					case 3:
						*l54 = green
					case 4:
						*l55 = green
				}
			}
		}
	}
  if row == 5 {
		for i := 0; i < 5; i++ {
			if row5[i] == yellow {
        if i == 0 {
          *l14 = green
          *l15 = green
          break
        }
        if i == 1 {
          *l12 = green
          *l15 = green
          break
        }
        if i == 2 {
          *l14 = green
          break
        }
        if (i == 4) || (i == 5) {
          fmt.Println("CANNOT BE SOLVED")
        }
			}
		}
	}
  if (row != 1) && (row != 2) && (row != 3) && (row != 4) && (row != 5) {
    fmt.Println("the problem is the row finder")
  }
}

func main() {
    http.HandleFunc("/", Index)
    http.HandleFunc("/.*", Index)
    http.ListenAndServe(":8080", nil)
}

// cd /Users/colinadler/GitHub/GoApp

// goapp deploy -application lights-game app.yaml
// export PATH=/usr/local/go/src/sdk/go_appengine:$PATH
