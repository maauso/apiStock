package derefstring

// DerefString convert pointer string to string
func DerefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}
