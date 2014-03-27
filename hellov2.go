//Hello GO!!!
/*
package main

import 
(
"fmt"
"os"
"strings"
)

func main()
{//take care with spaces and lines, this code will not compile. why?. cause the spaces and white lines
who:="World"
if len(os.Args)>1
	{
	who=strings.Join(os.Args[1:]," ")
	}

	fmt.Println("Hello",who)
}*/
//but this, without spaces, compiles correctly =)

package main
import ( 
"fmt"//provides functions for formatting text/read formatted text
"os"//plataform independent operation system
"strings"//functions for manipulating strings
)
func main() {
	who := "world!" // := declare and start a variable
		if len(os.Args) > 1 { //number of lines added in the command line 0 or more than 0
 			who = strings.Join(os.Args[1:], " ") 
		}

	fmt.Println("hello", who) 
}




