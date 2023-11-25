package irepositorys

var Name = "Nieraldo Lima"

type PostalCode struct {
	District   string `json:"district"`
	PostalCode string `json:"postalcode"`
	City       string `json:"city"`
	Street     string `json:"street"`
	State      string `json:"state"`
}

type Subscriber struct {
	Last_update int64      `json:"last_update"`
	Name        string     `json:"name"`
	Phone       string     `json:"phone"`
	PortNumber  string     `json:"portnumber"`
	Session_id  string     `json:"session_id"`
	Status      bool       `json:"status"`
	Cep         PostalCode `json:"postalcode" opcional:"true"`
}

func NewSubscriberWithCep(
	last_update int64,
	name string,
	phone string,
	portnumber string,
	session_id string,
	status bool,
	cep PostalCode,
) *Subscriber {
	return &Subscriber{
		last_update,
		name,
		phone,
		portnumber,
		session_id,
		status,
		cep,
	}
}

func NewSubscriberWithoutCep(
	last_update int64,
	name string,
	phone string,
	portnumber string,
	session_id string,
	status bool) *Subscriber {
	return &Subscriber{
		last_update,
		name,
		phone,
		portnumber,
		session_id,
		status,
		PostalCode{},
	}
}

type ISubscribersService interface {
	GetAll() (*Subscriber, error)
}

func GetAllSubscribers(i ISubscribersService) (*Subscriber, error) {
	return i.GetAll()
}
