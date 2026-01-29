package validation

func ValidateTaskId(id int) error {
	if id == 0 {
		return ErrIdRequired
	}

	return nil
}
