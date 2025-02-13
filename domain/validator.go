package domain

	type Validator struct {
		errors map[string]string
	}

	func NewValidator() *Validator {
		return &Validator{errors: make(map[string]string)}
	}

	func (v *Validator) MustBeLongerThan(field, value string, high int) bool {

		if _, ok := v.errors[field]; ok {
			return false
		}

		if value == "" {
			 return true
		}

		if len(value) < high {
			v.errors[field] = ErrNotLongerEnough{
				field: field,
				amount: high,
			}.Error()

			return false
		}

		return true

	}

	func (v *Validator) MustBeNotEmpty(field, value string) bool {
		if _, ok := v.errors[field]; ok {
			return false
		}

		if value == "" {
			v.errors[field] = ErrIsRequired{field: field}.Error()
			return false
		}
		return true
	}

	func (v *Validator) MustBeValidEmail(field, email string) bool {
		if _, ok := v.errors[field]; ok {
			return false
		}

		if !emailRegexp.MatchString(email) {
			v.errors[field] = ErrEmailBadFormat.Error()
			return false
		}
		return true
	}

	func (v *Validator) IsValid() bool {
		return len(v.errors) == 0
	}