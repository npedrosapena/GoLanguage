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
"fmt"
"os"
"strings"
)
func main() {
	who := "World!" 
		if len(os.Args) > 1 { 
 			who = strings.Join(os.Args[1:], " ") 
		}

	fmt.Println("Hello", who) 
}




