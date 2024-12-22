package userdata

type Query struct {
	UserData struct {
		User struct {
			Name string
			ID   int
		} `graphql:"user(id: $id)"`
	}
}
