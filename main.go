package main

import ("fmt"; 
"github.com/dhavalraj007/helloSubmodule/hello";
"github.com/dhavalraj007/helloSubmodule/Bye");

func main() {
	fmt.Println(hello.Hello("dhaval"));
	fmt.Println(Bye.Bye("dhaval"));
}
