package NewEmployee

import "fmt"

type Employee26 struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee26)LeavesRemaining()  {
	fmt.Printf("\n %s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}