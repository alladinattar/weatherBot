package tests

import (
	"github.com/tgBot/models"
	"testing"
)

func TestGetImage(t *testing.T) {
	wthr := models.Weather{}
	wthr.Weather = append(wthr.Weather)
	if wthr.GetImage() != "./images/clouds.jpg" {
		t.Error("Invalid image")
	}

}
