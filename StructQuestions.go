
// ##########################
// #       STRUCTS          #
// ##########################

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	fmt.Println("hello")

	// Create an instance of Person struct
	p := Person{
		Name:    "Abhishek",
		Age:     36,
		Address: "mahagun mywoods",
	}
	fmt.Printf("Struct is %+v\n", p)

	// Accessing struct fields
	p.Name = "Abhishek Shukla"
	fmt.Println(p.Address)

	// Update struct field via a function
	UpdateStruct(&p, "somu")

	// Pointer to struct
	p1 := &p
	p1.Age = 24
	fmt.Printf("Struct pointer is %+v\n", p1)

	// Copying a struct to create a new instance
	p2 := *p1
	fmt.Printf("p1 Struct is %p\n", &p1)
	fmt.Printf("p2 Struct is %p\n", &p2)

	// Nested struct example
	type Member struct {
		Name   string
		Age    int
		Person *Person
	}
	m1 := Member{
		Name: "tick",
		Age:  12,
		Person: &Person{
			Name:    "Abhi",
			Age:     12,
			Address: "mahagun",
		},
	}
	fmt.Printf("Nested Struct is %+v\n", m1)

	// Struct tags with validation
	type User struct {
		ID           int    `json:"-"` // private field
		Name         string `json:"name" validate:"required, len=5"`
		Email        string `json:"email" validate:"required,email"`
		EnrollmentNo int32  `json:"enrollmentno" validate:"required,gte=232323"`
		Address      string `json:"address,omitempty"`
	}
	user := User{
		ID:           12121,
		Name:         "Abhishek",
		EnrollmentNo: 98989898,
		Address:      "mahagun",
	}
	fmt.Printf("User Struct with tags: %+v\n", user)

	// Marshal struct to JSON
	bt, err := json.Marshal(user)
	if err == nil {
		newD := string(bt)
		fmt.Println("JSON representation:", newD)
	}

	// Unmarshal JSON to struct
	var newUser User
	jsonD := `{"id":12122,"name":"Abhishek","enrollmentno":11111,"address":"mahagun"}`
	err1 := json.Unmarshal([]byte(jsonD), &newUser)
	if err1 == nil {
		fmt.Println("Unmarshalled Struct is", newUser)
	}

	// Accessing struct tags using reflection
	ReadTags(&user)

	// Anonymous struct
	var st = struct {
		Name string
		Age  int
	}{
		"hello",
		23,
	}
	fmt.Println("Anonymous struct is", st)

	// Convert struct to map using reflection
	fmt.Println("Struct as map:", ConvertStructToMap(&p1))
}

// ##########################
// #   HELPER FUNCTIONS     #
// ##########################

// Convert struct to map using reflection
func ConvertStructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		result[typ.Field(i).Name] = val.Field(i).Interface()
	}
	return result
}

// Update struct fields in a function
func UpdateStruct(p *Person, newName string) *Person {
	p.Name = newName
	fmt.Printf("Updated Struct is %+v\n", p)
	return p
}

// Compare two structs
func CompareStructs() {
	p1 := Person{
		Name:    "hello",
		Age:     32,
		Address: "mahagun",
	}
	p2 := Person{
		Name:    "hello",
		Age:     32,
		Address: "mahagun",
	}
	if p1 == p2 {
		fmt.Println("Structs are equal")
	} else {
		fmt.Println("Structs are different")
	}
}

// Access struct tags using reflection
func ReadTags(data interface{}) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fmt.Println("Field:", field.Name)
		fmt.Println("JSON Tag:", field.Tag.Get("json"))
		fmt.Println("Validate Tag:", field.Tag.Get("validate"))
	}
}

