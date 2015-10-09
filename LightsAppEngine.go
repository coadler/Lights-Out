package lights

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
//var green string = "http://cadler.co/wp-content/uploads/2015/10/green-e1443791954640.png"
var blackGreen string = "http://cadler.co/wp-content/uploads/2015/10/blackGreen.png"
var yellowGreen string = "http://cadler.co/wp-content/uploads/2015/10/yellowGreen.png"

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

const resetHTML = `<input type=button onClick="parent.location='reset'" value='Reset puzzle'>`

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
      if urlToTest == "reset" {
        Reset()
      }
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
        fmt.Fprintf(w, resetHTML)
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
        if urlToTest != "reset" {
   			    fmt.Fprintf(w, "You just pressed square %v-%v </br>", url[1:2], url[2:3])
        }
        fmt.Fprintf(w, solveHTML)
        fmt.Fprintf(w, resetHTML)
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
        fmt.Fprintf(w, resetHTML)
        fmt.Fprintf(w, tail)
      }
    }
    if err != nil {
        panic(err)
    }
}

func SwitchColor(color *string, switchGreen bool) {
  //newColor := color
    switch *color {
      default:
        fmt.Println("Failed to switch color")
        //fmt.Println(newColor)
        fmt.Println("Color to change:", color)
        fmt.Println("Yellow:", yellow)
        fmt.Println("Black:", black)
        fmt.Println("YellowGreen:", yellowGreen)
        fmt.Println("BlackGreen:", blackGreen, "\n")
      case yellow:
        *color = black
      case black:
        *color = yellow
      case yellowGreen:
        if switchGreen {
          *color = black
        } else {
          *color = blackGreen
        }
      case blackGreen:
        if switchGreen {
          *color = yellow
        } else {
          *color = yellowGreen
        }
    }
  /*
	if *color == yellow {
		*color = black
	} else if *color == black {
		*color = yellow
	} else if *color == yellowGreen {
    *color = blackGreen
	} else if *color == blackGreen {
    *color = yellowGreen
  } else if (*color == yellowGreen) && (switchGreen == true) {
    *color = black
  } else if (*color == blackGreen) && (switchGreen == true) {
    *color = yellow
  } else {
    fmt.Println("Failed to switch color")
  }
  */
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
  row1 := []string{l15, l14, l13, l12, l11}
  row2 := []string{l25, l24, l23, l22, l21}
  row3 := []string{l35, l34, l33, l32, l31}
  row4 := []string{l45, l44, l43, l42, l41}
  row5 := []string{l55, l54, l53, l52, l51}
  var toReturn int
  yellowCount1, yellowCount2, yellowCount3, yellowCount4, yellowCount5 := 0, 0, 0, 0, 0
  for i := 4; i >= 0; i-- {
    if row1[i] == yellow {
      yellowCount1++
      break
    }
  }
  for i := 4; i >= 0; i-- {
    if row2[i] == yellow {
      yellowCount2++
      break
    }
  }
  for i := 4; i >= 0; i-- {
    if row3[i] == yellow {
      yellowCount3++
      break
    }
  }
  for i := 4; i >= 0; i-- {
    if row4[i] == yellow {
      yellowCount4++
      break
    }
  }
  for i := 4; i >= 0; i-- {
    if row5[i] == yellow {
      yellowCount5++
      break
    }
  }
  if yellowCount1 != 0 {
    toReturn = 1
  } else {
    if yellowCount2 != 0 {
      toReturn = 2
    } else {
      if yellowCount3 != 0 {
        toReturn = 3
      } else {
        if yellowCount4 != 0 {
          toReturn = 4
        } else {
          toReturn = 5
        }
      }
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
	if row == 1 {
		for i := 0; i < 5; i++ {
			if row1[i] == yellow {
				switch i {
					default:
						fmt.Println("failed in row 1")
					case 0:
            if *l21 == yellow {
              *l21 = yellowGreen
            } else {
              *l21 = blackGreen
            }
					case 1:
            if *l22 == yellow {
              *l22 = yellowGreen
            } else {
              *l22 = blackGreen
            }
					case 2:
            if *l23 == yellow {
              *l23 = yellowGreen
            } else {
              *l23 = blackGreen
            }
					case 3:
            if *l24 == yellow {
              *l24 = yellowGreen
            } else {
              *l24 = blackGreen
            }
					case 4:
            if *l25 == yellow {
              *l25 = yellowGreen
            } else {
              *l25 = blackGreen
            }
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
            if *l31 == yellow {
              *l31 = yellowGreen
            } else {
              *l31 = blackGreen
            }
					case 1:
            if *l32 == yellow {
              *l32 = yellowGreen
            } else {
              *l32 = blackGreen
            }
					case 2:
            if *l33 == yellow {
              *l33 = yellowGreen
            } else {
              *l33 = blackGreen
            }
					case 3:
            if *l34 == yellow {
              *l34 = yellowGreen
            } else {
              *l34 = blackGreen
            }
					case 4:
            if *l35 == yellow {
              *l35 = yellowGreen
            } else {
              *l35 = blackGreen
            }
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
            if *l41 == yellow {
              *l41 = yellowGreen
            } else {
              *l41 = blackGreen
            }
					case 1:
            if *l42 == yellow {
              *l42 = yellowGreen
            } else {
              *l42 = blackGreen
            }
					case 2:
            if *l43 == yellow {
              *l43 = yellowGreen
            } else {
              *l43 = blackGreen
            }
					case 3:
            if *l44 == yellow {
              *l44 = yellowGreen
            } else {
              *l44 = blackGreen
            }
					case 4:
            if *l45 == yellow {
              *l45 = yellowGreen
            } else {
              *l45 = blackGreen
            }
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
            if *l51 == yellow {
              *l51 = yellowGreen
            } else {
              *l51 = blackGreen
            }
					case 1:
            if *l52 == yellow {
              *l52 = yellowGreen
            } else {
              *l52 = blackGreen
            }
					case 2:
            if *l53 == yellow {
              *l53 = yellowGreen
            } else {
              *l53 = blackGreen
            }
					case 3:
            if *l54 == yellow {
              *l54 = yellowGreen
            } else {
              *l54 = blackGreen
            }
					case 4:
            if *l55 == yellow {
              *l55 = yellowGreen
            } else {
              *l55 = blackGreen
            }
				}
			}
		}
	}
  if row == 5 {
		for i := 0; i < 5; i++ {
			if row5[i] == yellow {
        if i == 0 {
          if *l14 == yellow {
            *l14 = yellowGreen
          } else {
            *l14 = blackGreen
          }
          if *l15 == yellow {
            *l15 = yellowGreen
          } else {
            *l15 = blackGreen
          }
          break
        }
        if i == 1 {
          if *l12 == yellow {
            *l12 = yellowGreen
          } else {
            *l12 = blackGreen
          }
          if *l15 == yellow {
            *l15 = yellowGreen
          } else {
            *l15 = blackGreen
          }
          break
        }
        if i == 2 {
          if *l14 == yellow {
            *l14 = yellowGreen
          } else {
            *l14 = blackGreen
          }
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

func RandColorP(block *string) {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	colorNumber := r.Intn(2) + 1;

	switch (colorNumber) {
		default: 	*block = yellow //"http://cadler.co/wp-content/uploads/2015/09/yellow-e1442941699604.jpg"
		case 1:		*block = yellow //"http://cadler.co/wp-content/uploads/2015/09/yellow-e1442941699604.jpg"
		case 2:		*block = black  //"http://cadler.co/wp-content/uploads/2015/09/black-e1442941689285.jpg"
	}
}

func Reset() {
  	RandColorP(&l11)
    RandColorP(&l12)
    RandColorP(&l13)
    RandColorP(&l14)
    RandColorP(&l15)
    RandColorP(&l21)
    RandColorP(&l22)
    RandColorP(&l23)
    RandColorP(&l24)
    RandColorP(&l25)
    RandColorP(&l31)
    RandColorP(&l32)
    RandColorP(&l33)
    RandColorP(&l34)
    RandColorP(&l35)
    RandColorP(&l41)
    RandColorP(&l42)
    RandColorP(&l43)
    RandColorP(&l44)
    RandColorP(&l45)
    RandColorP(&l51)
    RandColorP(&l52)
    RandColorP(&l53)
    RandColorP(&l54)
    RandColorP(&l55)
}

func init() {
    http.HandleFunc("/", Index)
    http.HandleFunc("/.*", Index)
}

// cd /Users/colinadler/GitHub/GoApp

// goapp deploy -application lights-outg app.yaml
// export PATH=/usr/local/go/src/sdk/go_appengine:$PATH
