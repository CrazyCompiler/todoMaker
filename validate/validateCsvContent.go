package validate

func IsValidPriority(priority string) bool{
	return priority=="High" || priority == "Medium" || priority == "Low"
}

