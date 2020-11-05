package yandex

import (
	"testing"
	"time"
)

func TestSendPoint(t *testing.T) {

	// тестовый сервер
	serverURL := "http://tst.extjams.maps.yandex.net/mtr_collect/1.x/"

	client := New(clid)

	item := models.GNSSData{
		EventTime: time.Now().UTC().Unix(),
		Lat:       44.56,
		Lon:       33.78,
		Speed:     45,
		Heading:   240,
	}

	err := client.SendPoint(item)
	if err != nil {
		t.Error(err)
	}
}
