package orm


// STATUS stands for exit status of check
type STATUS int8

const (
	// OK for success
	OK STATUS = 0
	// WARNING for not OK
	WARNING = 1
	// CRITICAL for failure
	CRITICAL = 2

)

// Instance represents data model of instances or object of host under monitoring.
type Instance struct {
	InstanceName string //`json:”instancename,omitempty”`
	InstanceIP string 
	probe []Probe 
}


// Probe represent data model for checks.
type Probe struct {
    Name string
    Location string
	Interval int
	History []STATUS
}