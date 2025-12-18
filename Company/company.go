package Company

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	Name   string
	Salary float64
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Full-time: %s, Salary: %.2f", f.Name, f.Salary)
}

type PartTimeEmployee struct {
	Name      string
	HourlyPay float64
	Hours     int
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Part-time: %s, Pay: %.2f x %d", p.Name, p.HourlyPay, p.Hours)
}

type Company struct {
	Employees map[uint64]Employee
}

func NewCompany() *Company {
	return &Company{
		Employees: make(map[uint64]Employee),
	}
}

func (c *Company) AddEmployee(id uint64, emp Employee) {
	c.Employees[id] = emp
}

func (c *Company) ListEmployees() []string {
	var list []string
	for _, emp := range c.Employees {
		list = append(list, emp.GetDetails())
	}
	return list
}
