package main
import(
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"

)
const(//DECLARING CONSTANTS
	markName = "GOLANG_ CLI_REMINDER"
	markValue = "1"
)
func main(){//USING os.Args for user input
	if len(os.Args) < 3{//IF USER INPUT IS LESS THAN 3 SHOW EXAMPLE
		fmt.Printf("Usage:%s <hh:mm> <message\n",os.Args[0])
		os.Exit(1)
	}
	now := time.Now() //Declaring a new variable for current time
	w := when.New(nil)
	w.Add(en.All...)//Adding for english
	w.Add(common.All...)
	t,err := w.Parse(os.Args[1],now)//Parsing the input
	if err != nil{
		fmt.Println(err)
		os.Exit(2)
	}
	if t == nil{
		fmt.Println("Unable to parse time")
		os.Exit(2)
	}
	if now.After(t.Time){
		fmt.Println("Set a future time")
		os.Exit(3)
	} 

	diff := t.Time.Sub(now)
	if(os.Getenv(markName) == markValue){
		time.Sleep(diff)
		err = beeep.Alert("Reminder",strings.Join(os.Args[2:],""),"assets/information.png")
		if err !=nil{
			fmt.Println(err)
			os.Exit(4)
		}
	}else{
		cmd := exec.Command(os.Args[0],os.Args[1:]...)
		cmd.Env = append(os.Environ(),fmt.Sprintf("%s=%s",markName,markValue))
	if err := cmd.Start(); err !=nil{
		fmt.Println(err)
		os.Exit(5)
	}
	fmt.Println("Reminder will be displayed after",diff.Round(time.Second))
	os.Exit(0)
	}

}