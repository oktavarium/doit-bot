package dbo

func stringFromPointer(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func stringToPointer(s string) *string {
	return &s
}

func boolFromPointer(b *bool) bool {
	if b == nil {
		return false
	}

	return *b
}
