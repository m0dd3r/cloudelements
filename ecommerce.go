package h2go

var ecommerce_base_path = "/hubs/ecommerce/"
var ecommerce_paths = map[string]string{
	"customers": ecommerce_base_path + "customers",
}

type Address struct {
	Address1     string
	Address2     string
	City         string
	Company      string
	Id           int
	Province     string
	Phone        string
	Zip          string
	ProvinceCode string
	CountryCode  string
	CountryName  string
	LastName     string
	FirstName    string
	Country      string
}

type Customer struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string
	Id               int
	Tags             string
	Note             string
	LastOrderId      string
	UpdatedAt        string
	LastOrderName    string
	TotalSpent       string
	State            string
	CreatedAt        string
	VerifiedEmail    bool
	AcceptsMarketing bool
	OrdersCount      int
	DefaultAddress   Address
	Addresses        []Address
}

func (s Session) GetEcommerceCustomers(query string, page, page_size int) ([]*Customer, error) {
	//
	// Struct to hold error response
	//
	errors := struct {
		Message string
	}{}

	url := BASE_URL + ecommerce_paths["customers"]

	res := []*Customer{}

	resp, err := s.Get(url, nil, &res, errors)
	if resp.Status() == 200 {
		return res, nil
	} else {
		return nil, err
	}
}
