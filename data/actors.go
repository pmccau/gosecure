package data

type Location struct {
	Address 		string
	ZipCode 		string
	City			string
	State 			string
	Country 		string
	Name 			string
	Pin 			string
	IsArmed			bool
}

type Person struct {
	Name 			string
	Email 			string
	NotifyByEmail	bool
	Phone 			string
	NotifyByPhone   bool
}

// NewLocation is a constructor for location
func NewLocation(Address string, City string, State string,
	Country string, ZipCode string, Name string) *Location {
	return &Location {
		Address: Address,
		City: City,
		State: State,
		Country: Country,
		ZipCode: ZipCode,
		Name: Name,
		IsArmed: false,
	}
}

// NewPerson is a constructor for Person
func NewPerson(Name string, Email string, Phone string) *Person {
	return &Person {
		Name: Name,
		Email: Email,
		Phone: Phone,
		NotifyByEmail: true,
		NotifyByPhone: false,
	}
}