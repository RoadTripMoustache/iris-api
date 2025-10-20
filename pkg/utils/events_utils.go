package utils

import "fmt"

func GenerateVoyageDetailSessionId(userId string, voyageId string) string {
	return fmt.Sprintf("%s-%s", userId, voyageId)
}
