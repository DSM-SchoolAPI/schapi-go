package schapi

import (
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

type Menu struct {
	Breakfast []string `json:"breakfast"`
	Lunch     []string `json:"lunch"`
	Dinner    []string `json:"dinner"`
}

func NewMenu() *Menu {
	return &Menu{
		Breakfast: make([]string, 0),
		Lunch: make([]string, 0),
		Dinner: make([]string, 0),
	}
}

var URL = "http://%s/sts_sci_md00_001.do?schulCode=%s&schulCrseScCode=%d&schulKndScCode=0%d&schYm=%d%02d"
var TIMINGS = [3]string{"조식", "중식", "석식"}

func (api SchoolAPI) GetMonthlyMenuURL(year int, month int) string {
	return fmt.Sprintf(URL, api.Region, api.Code, api.Kind, api.Kind, year, month)
}

func (api SchoolAPI) GetMenuData(text string) map[int][]string {
	menu := make(map[int][]string)

	for i, _ := range TIMINGS {
		menu[i] = make([]string, 0)
	}

	re, _ := regexp.Compile("[가-힣]+\\([가-힣]+\\)|[가-힣]+")
	chunks := re.FindAllString(text, -1)
	timing := 0

	for _, s := range chunks {
		if matched, _ := regexp.MatchString("[조중석]식", s); matched {
			for i, t := range TIMINGS {
				if s == t {
					timing = i
					break
				}
			}
		} else {
			menu[timing] = append(menu[timing], s)
		}
	}

	return menu
}

func (api SchoolAPI) GetMenuItem(data map[int][]string) Menu {
	breakfast := data[0]
	if len(breakfast) == 0 {
		breakfast = nil
	}
	
	lunch := data[1]
	if len(lunch) == 0 {
		lunch = nil
	}

	dinner := data[2]
	if len(dinner) == 0 {
		dinner = nil
	}

	return Menu {
		Breakfast: breakfast,
		Lunch: lunch,
		Dinner: dinner,
	}
}

func (api SchoolAPI) GetMonthlyMenus(year int, month int) map[int]Menu {
	menus := make(map[int]Menu, 0)
	url := api.GetMonthlyMenuURL(year, month)
	doc, _ := goquery.NewDocument(url)
	
	items := doc.Find(".tbl_type3.tbl_calendar td").FilterFunction(func(_ int, s *goquery.Selection) bool {
		return s.Text() != " "
	})

	items.Each(func(i int, s *goquery.Selection) {
		data := api.GetMenuData(s.Text())
		menu := api.GetMenuItem(data)

		menus[i + 1] = menu
	})

	return menus
}
