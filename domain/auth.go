package domain
	import "regexp"

	var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	type RegisterPayload struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
		Username        string `json:"username"`
	}

	func (r *RegisterPayload) IsValid() (bool, map[string]string) {
		v := NewValidator()

		v.MustBeNotEmpty("email", r.Email)
		v.MustBeValidEmail("email", r.Email)

		v.MustBeNotEmpty("password", r.Password)
		v.MustBeLongerThan("password", r.Password, 6)

		v.MustBeNotEmpty("confirmPassword", r.ConfirmPassword)

		v.MustBeNotEmpty("username", r.Username)
		v.MustBeLongerThan("username", r.Username, 3)

		return v.IsValid(), v.errors
	}

	func (d *Domain) Register(payload RegisterPayload) (*User, error) {
		userExist, _ := d.DB.UserRepo.GetByEmail(payload.Email)
		if userExist != nil {
			return nil, ErrUserWithEmailAlreadyExist
		}

		userExist, _ = d.DB.UserRepo.GetByUsername(payload.Username)
		if userExist != nil {
			return nil, ErrUserWithUsernameAlreadyExist
		}

		password, err := d.setPassword(payload.Password)
		if err != nil {
			return nil, err
		}

		data := &User{
			Username: payload.Username,
			Email:    payload.Email,
			Password: *password,
		}

		user, err := d.DB.UserRepo.Create(data)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	func (d *Domain) setPassword(password string) (*string, error) {
		return nil, nil
	}