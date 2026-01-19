package validation

func ValidateTaskTitle(title string) error {
	if title == "" {
		return ErrTitleRequired
	}

	if len(title) > 255 {
		return ErrTitleTooLong
	}

	return nil
}
