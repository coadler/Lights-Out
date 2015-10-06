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

const yellow = "http://cadler.co/wp-content/uploads/2015/09/yellow-e1442941699604.jpg"
const black = "http://cadler.co/wp-content/uploads/2015/09/black-e1442941689285.jpg"

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
	   		fmt.Fprintf(w, tail)
   		}

   		if length > 1 {
   			switch urlToTest[0:1] {
   				default:
   				case "1":
   					switch urlToTest[1:2] {
   						default:
   						case "1":
   							SwitchColor(&l11)
   							SwitchColor(&l12)
   							SwitchColor(&l21)
		   				case "2":
		   					SwitchColor(&l11)
		   					SwitchColor(&l12)
		   					SwitchColor(&l13)
		   					SwitchColor(&l22)
		   				case "3":
		   					SwitchColor(&l12)
		   					SwitchColor(&l13)
		   					SwitchColor(&l14)
		   					SwitchColor(&l23)
		   				case "4":
		   					SwitchColor(&l13)
		   					SwitchColor(&l14)
		   					SwitchColor(&l15)
		   					SwitchColor(&l24)
		   				case "5":
		   					SwitchColor(&l14)
		   					SwitchColor(&l15)
		   					SwitchColor(&l25)
   					}
   				case "2":
   					switch urlToTest[1:2] {
   						default:
   						case "1":
   							SwitchColor(&l11)
   							SwitchColor(&l21)
   							SwitchColor(&l31)
   							SwitchColor(&l22)
		   				case "2":
		   					SwitchColor(&l12)
		   					SwitchColor(&l21)
		   					SwitchColor(&l22)
		   					SwitchColor(&l23)
		   					SwitchColor(&l32)
		   				case "3":
		   					SwitchColor(&l13)
		   					SwitchColor(&l22)
		   					SwitchColor(&l23)
		   					SwitchColor(&l24)
		   					SwitchColor(&l33)
		   				case "4":
		   					SwitchColor(&l14)
		   					SwitchColor(&l23)
		   					SwitchColor(&l24)
		   					SwitchColor(&l25)
		   					SwitchColor(&l34)
		   				case "5":
		   					SwitchColor(&l24)
		   					SwitchColor(&l25)
		   					SwitchColor(&l15)
		   					SwitchColor(&l35)
   					}
   				case "3":
   					switch urlToTest[1:2] {
   						default:
   						case "1":
   							SwitchColor(&l21)
   							SwitchColor(&l31)
   							SwitchColor(&l32)
   							SwitchColor(&l41)
		   				case "2":
		   					SwitchColor(&l22)
		   					SwitchColor(&l31)
		   					SwitchColor(&l32)
		   					SwitchColor(&l33)
		   					SwitchColor(&l42)
		   				case "3":
		   					SwitchColor(&l23)
		   					SwitchColor(&l32)
		   					SwitchColor(&l33)
		   					SwitchColor(&l34)
		   					SwitchColor(&l43)
		   				case "4":
		   					SwitchColor(&l24)
		   					SwitchColor(&l33)
		   					SwitchColor(&l34)
		   					SwitchColor(&l35)
		   					SwitchColor(&l44)
		   				case "5":
		   					SwitchColor(&l34)
		   					SwitchColor(&l35)
		   					SwitchColor(&l25)
		   					SwitchColor(&l45)
   					}
   				case "4":
   					switch urlToTest[1:2] {
   						default:
   						case "1":
   							SwitchColor(&l31)
   							SwitchColor(&l41)
   							SwitchColor(&l42)
   							SwitchColor(&l51)
		   				case "2":
		   					SwitchColor(&l32)
		   					SwitchColor(&l41)
		   					SwitchColor(&l42)
		   					SwitchColor(&l43)
		   					SwitchColor(&l52)
		   				case "3":
		   					SwitchColor(&l33)
		   					SwitchColor(&l42)
		   					SwitchColor(&l43)
		   					SwitchColor(&l44)
		   					SwitchColor(&l53)
		   				case "4":
		   					SwitchColor(&l34)
		   					SwitchColor(&l43)
		   					SwitchColor(&l44)
		   					SwitchColor(&l45)
		   					SwitchColor(&l54)
		   				case "5":
		   					SwitchColor(&l44)
		   					SwitchColor(&l45)
		   					SwitchColor(&l35)
		   					SwitchColor(&l55)
   					}
   				case "5":
   					switch urlToTest[1:2] {
   						default:
   						case "1":
   							SwitchColor(&l41)
   							SwitchColor(&l51)
   							SwitchColor(&l52)
		   				case "2":
		   					SwitchColor(&l42)
		   					SwitchColor(&l51)
		   					SwitchColor(&l52)
		   					SwitchColor(&l53)
		   				case "3":
		   					SwitchColor(&l43)
		   					SwitchColor(&l52)
		   					SwitchColor(&l53)
		   					SwitchColor(&l54)
		   				case "4":
		   					SwitchColor(&l44)
		   					SwitchColor(&l53)
		   					SwitchColor(&l54)
		   					SwitchColor(&l55)
		   				case "5":
		   					SwitchColor(&l45)
		   					SwitchColor(&l54)
		   					SwitchColor(&l55)
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
	   		fmt.Fprintf(w, tail)
        if urlToTest != "reset" {
   			    fmt.Fprintf(w, "You just pressed square %v-%v", url[1:2], url[2:3])
        }
   		}
    }
    if err != nil {
        panic(err)
    }
}

func SwitchColor(color *string) {
	if *color == yellow {
		*color = black
	} else {
		*color = yellow
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
    //http.ListenAndServe(":8080", nil)
}

// cd /Users/colinadler/GitHub/GoApp

/*
		<div class="floated_img">
			<a href="11"><img src="{{.L1}}" alt="square" /></a>
			<a href="12"><img src="{{.L2}}" alt="square" /></a>
			<a href="13"><img src="{{.L3}}" alt="square" /></a>
			<a href="14"><img src="{{.L4}}" alt="square" /></a>
			<a href="15"><img src="{{.L5}}" alt="square" /></a>
		</div>
		<div class="floated_img">
			<a href="21"><img src="{{.On}}" alt="square" /></a>
			<a href="22"><img src="{{.On}}" alt="square" /></a>
			<a href="23"><img src="{{.On}}" alt="square" /></a>
			<a href="24"><img src="{{.On}}" alt="square" /></a>
			<a href="25"><img src="{{.On}}" alt="square" /></a>
		</div>
		<div class="floated_img">
			<a href="31"><img src="{{.On}}" alt="square" /></a>
			<a href="32"><img src="{{.On}}" alt="square" /></a>
			<a href="33"><img src="{{.On}}" alt="square" /></a>
			<a href="34"><img src="{{.On}}" alt="square" /></a>
			<a href="35"><img src="{{.On}}" alt="square" /></a>
		</div>
		<div class="floated_img">
			<a href="41"><img src="{{.On}}" alt="square" /></a>
			<a href="42"><img src="{{.On}}" alt="square" /></a>
			<a href="43"><img src="{{.On}}" alt="square" /></a>
			<a href="44"><img src="{{.On}}" alt="square" /></a>
			<a href="45"><img src="{{.On}}" alt="square" /></a>
		</div>
		<div class="floated_img">
			<a href="51"><img src="{{.On}}" alt="square" /></a>
			<a href="52"><img src="{{.On}}" alt="square" /></a>
			<a href="53"><img src="{{.On}}" alt="square" /></a>
			<a href="54"><img src="{{.On}}" alt="square" /></a>
			<a href="55"><img src="{{.On}}" alt="square" /></a>
		</div>
*/

// goapp deploy -application lights-outg app.yaml
// export PATH=/usr/local/go/src/sdk/go_appengine:$PATH
