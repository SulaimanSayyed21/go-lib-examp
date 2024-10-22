package modb

import "github.com/SulaimanSayyed21/go-lib-examp/moda" 

// sayHelloB calls moda.SayHelloA adds its own message

func SayHelloB() string {
	return moda.SayHelloA() + "add Hello from Module B"
}
