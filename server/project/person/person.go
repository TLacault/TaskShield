package person

type Role string

const (
	HOST                Role = "Host"
	COWORKER_READ       Role = "Coworker-R"
	COWORKER_READ_WRITE Role = "Coworker-RW"
)

type Person struct {
	firstName string
	lastName  string
	role      Role
}

/* ********************************************************************************************* */

func New(firstName string, lastName string, role Role) Person {
	person := Person{
		firstName: firstName,
		lastName:  lastName,
		role:      role,
	}
	return person
}

/* ********************************************************************************************* */

func (p *Person) GetFirstName() string { return p.firstName }
func (p *Person) GetLastName() string  { return p.lastName }
func (p *Person) GetRole() string      { return string(p.role) }

/* ********************************************************************************************* */

func (p *Person) ToString() string {
	return string(p.role) + ":" + p.firstName + "_" + p.lastName
}

func (p *Person) ToStringFormat() string {
	return "[" + string(p.role) + "] " + p.firstName + " " + p.lastName
}
