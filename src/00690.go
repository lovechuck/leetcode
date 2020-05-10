package src

/*690. 员工的重要性*/

/**
* Definition for Employee.
* type Employee struct {
*     Id int
*     Importance int
*     Subordinates []int
* }
 */

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	hash := make(map[int]*Employee)
	for _, employee := range employees {
		hash[employee.Id] = employee
	}
	sum := 0
	sum = dfs(hash[id], hash)
	return sum
}

func dfs(employee *Employee, hash map[int]*Employee) int {
	sum := 0
	sum += employee.Importance
	for _, subordinate := range employee.Subordinates {
		sum += dfs(hash[subordinate], hash)
	}
	return sum
}
