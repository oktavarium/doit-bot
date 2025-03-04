package dbo

import "go.mongodb.org/mongo-driver/bson/primitive"

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

func objectIDsToString(ids []primitive.ObjectID) []string {
	result := make([]string, 0, len(ids))
	for _, id := range ids {
		result = append(result, id.Hex())
	}
	return result
}
