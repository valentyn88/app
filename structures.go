package app

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    bool   `json:"gender"`
	BirthDate int    `json:"birth_date"`
	Age       int    `json:"age"`
	Visits    Visits `json:"-"`
}

type UserService interface {
	Get(id int) (*User, error)
	Update(id int, values map[string]interface{}) error
	Create(u *User) error
	Visits() (Visits, error)
}

type Location struct {
	ID       int
	Place    string
	Country  string
	City     string
	Distance int
	Visits   Visits
}

type LocationService interface {
	Get(id int) (*Location, error)
	Update(l *Location) error
	Create(l *Location) error
	Avg(id int) (int, error)
}

type Visit struct {
	ID        int
	Location  int
	User      int
	VisitedAt int
	Mark      int
}

type VisitService interface {
	Get(id int) (*Visit, error)
	Update(v *Visit) error
	Create(v *Visit) error
}

type Visits []*Visit
