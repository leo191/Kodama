package main


import (
	"bitbucket.org/leo191/kodama/orm"
	"fmt"
)


func main()  {
	
	cpu := orm.Probe{"check_cpu", "/usr/lib/probes/check-cpu.sh", 10, []orm.STATUS{orm.OK, orm.CRITICAL}}
	fmt.Println(cpu)
}