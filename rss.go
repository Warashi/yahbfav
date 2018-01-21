package yahbfav

import (
	"encoding/xml"
	"fmt"
	"time"
)

const day = 24 * time.Hour

type Date struct {
	time.Time
}

func (dt *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	parse, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return err
	}
	*dt = Date{parse}
	return nil
}

func (d Date) String() string {
	s := time.Now().Sub(d.Time)
	if s > day {
		return d.Format("01/02")
	}
	if s > time.Hour {
		return fmt.Sprintf("%.0f時間前", s.Hours())
	}
	if s > time.Minute {
		return fmt.Sprintf("%.0f分前", s.Minutes())
	}
	if s > time.Second {
		return fmt.Sprintf("%.0f秒前", s.Seconds())
	}
	return "now"
}

type RSS struct {
	Items []struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Creator     string `xml:"creator"`
		Date        Date   `xml:"date"`
	} `xml:"item"`
}
