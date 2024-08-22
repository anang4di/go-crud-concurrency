package user

type UserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func FormatUser(user User) UserFormatter {
	return UserFormatter(user)
}

func FormatUsers(users []User) []UserFormatter {
	usersFormatter := []UserFormatter{}

	for _, user := range users {
		formatter := FormatUser(user)
		usersFormatter = append(usersFormatter, formatter)
	}

	return usersFormatter
}
