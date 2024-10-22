package modc

import "github.com/SulaimanSayyed21/go-lib-examp/modb"

// sayHelloC calls modb.SayHelloB adds its own message

func SayHelloC() string {
	return modb.SayHelloB() + "add Hello from Module C"
}
