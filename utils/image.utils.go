package utils

import (
	"fmt"

	"github.com/google/uuid"
)
func EncodeImageImage(img_name string) string{
	newName := fmt.Sprintf("%s_%s.png",img_name,uuid.New().String());
	return newName;
}

