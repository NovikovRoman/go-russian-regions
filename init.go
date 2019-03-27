package go_russian_regions

import (
	"fmt"
	"regexp"
)

func (c *Client) init() {
	c.regions = []*Region{
		{Name: "Республика Адыгея", Code: "01", reString: `(^|[^а-яё])адыге.+`},
		{Name: "Республика Алтай", Code: "04", reString: `(^|[^а-яё])алта[йяюе](\s|[^а-яё]|$)`},
		{Name: "Республика Башкортостан", Code: "02", reString: `(^|[^а-яё])башкортостан(.{1}|)(\s|[^а-яё]|$)`},
		{Name: "Республика Бурятия", Code: "03", reString: `(^|[^а-яё])буряти[ия](\s|[^а-яё]|$)`},
		{Name: "Республика Дагестан", Code: "05", reString: `(^|[^а-яё])Дагестан(.{1}|)(\s|[^а-яё]|$)`},
		{Name: "Республика Ингушетия", Code: "06", reString: `(^|[^а-яё])Ингушети[ия](\s|[^а-яё]|$)`},
		{Name: "Кабардино-Балкарская республика", Code: "07", reString: `(^|[^а-яё])Кабардино\-Балкар.+`},
		{Name: "Республика Калмыкия", Code: "08", reString: `(^|[^а-яё])Калмык.+`},
		{Name: "Карачаево-Черкесская республика", Code: "09", reString: `(^|[^а-яё])Карачаево\-Черке.+`},
		{Name: "Республика Карелия", Code: "10", reString: `(^|[^а-яё])Карели.+`},
		{Name: "Республика Коми", Code: "11", reString: `(^|[^а-яё])Коми(\s|[^а-яё]|$)`},
		{Name: "Республика Крым", Code: "91", reString: `(^|[^а-яё])Крым`},
		{Name: "Республика Марий Эл", Code: "12", reString: `(^|[^а-яё])Марий(\s|\-)Эл`},
		{Name: "Республика Мордовия", Code: "13", reString: `(^|[^а-яё])Мордов.+`},
		{Name: "Республика Саха (Якутия)", Code: "14", reString: `([\s\(\)]|)(Саха|Якути[яи])(\s|[^а-яё]|$)`},
		{Name: "Республика Северная Осетия — Алания", Code: "15", reString: `(^|[^а-яё])(осети|алани)[яи](\s|[^а-яё]|$)`},
		{Name: "Республика Татарстан", Code: "16", reString: `(^|[^а-яё])(Татарстан.{0,1}|Татари[яи])(\s|[^а-яё]|$)`},
		{Name: "Республика Тыва", Code: "17", reString: `(^|[^а-яё])Тыв[ае](\s|[^а-яё]|$)`},
		{Name: "Удмуртская республика", Code: "18", reString: `(^|[^а-яё])Удмурт.+`},
		{Name: "Республика Хакасия", Code: "19", reString: `(^|[^а-яё])Хакаси[ия](\s|[^а-яё]|$)`},
		{Name: "Чеченская республика", Code: "20", reString: `(^|[^а-яё])(Чечен|Чечн).+`},
		{Name: "Чувашская республика", Code: "21", reString: `(^|[^а-яё])Чуваш.+`},
		{Name: "Алтайский край", Code: "22", reString: `(^|[^а-яё])Алтайс.+`},
		{Name: "Забайкальский край", Code: "75", reString: `(^|[^а-яё])Забайк.+`},
		{Name: "Камчатский край", Code: "41", reString: `(^|[^а-яё])Камчат.+`},
		{Name: "Краснодарский край", Code: "23", reString: `(^|[^а-яё])Краснода.+`},
		{Name: "Красноярский край", Code: "24", reString: `(^|[^а-яё])Краснояр.+`},
		{Name: "Пермский край", Code: "59", reString: `(^|[^а-яё])Пермс.+`},
		{Name: "Приморский край", Code: "25", reString: `(^|[^а-яё])Примор.+`},
		{Name: "Ставропольский край", Code: "26", reString: `(^|[^а-яё])Ставропол.+`},
		{Name: "Хабаровский край", Code: "27", reString: `(^|[^а-яё])Хабаров.+`},
		{Name: "Амурская область", Code: "28", reString: `(^|[^а-яё])Амурс.+`},
		{Name: "Архангельская область", Code: "29", reString: `(^|[^а-яё])Архангель.+`},
		{Name: "Астраханская область", Code: "30", reString: `(^|[^а-яё])Астраханс.+`},
		{Name: "Белгородская область", Code: "31", reString: `(^|[^а-яё])Белгородс.+`},
		{Name: "Брянская область", Code: "32", reString: `(^|[^а-яё])Брянск.+`},
		{Name: "Владимирская область", Code: "33", reString: `(^|[^а-яё])Владимирс.+`},
		{Name: "Волгоградская область", Code: "34", reString: `(^|[^а-яё])Волгоградс.+`},
		{Name: "Вологодская область", Code: "35", reString: `(^|[^а-яё])Вологодс.+`},
		{Name: "Воронежская область", Code: "36", reString: `(^|[^а-яё])Воронежс.+`},
		{Name: "Ивановская область", Code: "37", reString: `(^|[^а-яё])Ивановс.+`},
		{Name: "Иркутская область", Code: "38", reString: `(^|[^а-яё])Иркутс.+`},
		{Name: "Калининградская область", Code: "39", reString: `(^|[^а-яё])Калининградс.+`},
		{Name: "Калужская область", Code: "40", reString: `(^|[^а-яё])Калужс.+`},
		{Name: "Кемеровская область", Code: "42", reString: `(^|[^а-яё])Кемеровс.+`},
		{Name: "Кировская область", Code: "43", reString: `(^|[^а-яё])Кировс.+`},
		{Name: "Костромская область", Code: "44", reString: `(^|[^а-яё])Костромс.+`},
		{Name: "Курганская область", Code: "45", reString: `(^|[^а-яё])Курганс.+`},
		{Name: "Курская область", Code: "46", reString: `(^|[^а-яё])Курск.+`},
		{Name: "Ленинградская область", Code: "47", reString: `(^|[^а-яё])Ленинградс.+`},
		{Name: "Липецкая область", Code: "48", reString: `(^|[^а-яё])Липецк.+`},
		{Name: "Магаданская область", Code: "49", reString: `(^|[^а-яё])Магаданс.+`},
		{Name: "Московская область", Code: "50", reString: `(^|[^а-яё])Московс.+`},
		{Name: "Мурманская область", Code: "51", reString: `(^|[^а-яё])Мурманск.+`},
		{Name: "Нижегородская область", Code: "52", reString: `(^|[^а-яё])Нижегородс.+`},
		{Name: "Новгородская область", Code: "53", reString: `(^|[^а-яё])Новгородс.+`},
		{Name: "Новосибирская область", Code: "54", reString: `(^|[^а-яё])Новосибирск.+`},
		{Name: "Омская область", Code: "55", reString: `(^|[^а-яё])Омск.+`},
		{Name: "Оренбургская область", Code: "56", reString: `(^|[^а-яё])Оренбургс.+`},
		{Name: "Орловская область", Code: "57", reString: `(^|[^а-яё])Орловс.+`},
		{Name: "Пензенская область", Code: "58", reString: `(^|[^а-яё])Пензенс.+`},
		{Name: "Псковская область", Code: "60", reString: `(^|[^а-яё])Псковс.+`},
		{Name: "Ростовская область", Code: "61", reString: `(^|[^а-яё])Ростовс.+`},
		{Name: "Рязанская область", Code: "62", reString: `(^|[^а-яё])Рязанс.+`},
		{Name: "Самарская область", Code: "63", reString: `(^|[^а-яё])Самарс.+`},
		{Name: "Саратовская область", Code: "64", reString: `(^|[^а-яё])Саратовс.+`},
		{Name: "Сахалинская область", Code: "65", reString: `(^|[^а-яё])Сахалинс.+`},
		{Name: "Свердловская область", Code: "66", reString: `(^|[^а-яё])Свердловс.+`},
		{Name: "Смоленская область", Code: "67", reString: `(^|[^а-яё])Смоленс.+`},
		{Name: "Тамбовская область", Code: "68", reString: `(^|[^а-яё])Тамбовс.+`},
		{Name: "Тверская область", Code: "69", reString: `(^|[^а-яё])Тверс.+`},
		{Name: "Томская область", Code: "70", reString: `(^|[^а-яё])Томск.+`},
		{Name: "Тульская область", Code: "71", reString: `(^|[^а-яё])Тульс.+`},
		{Name: "Тюменская область", Code: "72", reString: `(^|[^а-яё])Тюменс.+`},
		{Name: "Ульяновская область", Code: "73", reString: `(^|[^а-яё])Ульяновс.+`},
		{Name: "Челябинская область", Code: "74", reString: `(^|[^а-яё])Челябинс.+`},
		{Name: "Ярославская область", Code: "76", reString: `(^|[^а-яё])Ярославс.+`},
		{Name: "Москва", Code: "77", reString: `(^|[^а-яё])Москв[аеыуо](й|\s|[^а-яё]|$)`},
		{Name: "Санкт-Петербург", Code: "78", reString: `(^|[^а-яё])(Санкт[-\s])*Петербур.+`},
		{Name: "Севастополь", Code: "92", reString: `(^|[^а-яё])Севастопол.+`},
		{Name: "Еврейская автономная область", Code: "79", reString: `(^|[^а-яё])Еврейс.+`},
		{Name: "Ненецкий автономный округ", Code: "83", reString: `(^|[^а-яё\-])Ненецк.+`},
		{Name: "Ханты-Мансийский автономный округ - Югра", Code: "86",
			reString: `(^|[^а-яё\-])(Ханты\-Мансийс.+|Югр.+|ХМАО([^а-яё]|$))`},
		{Name: "Чукотский автономный округ", Code: "87", reString: `(^|[^а-яё])чукот.+`},
		{Name: "Ямало-Ненецкий автономный округ", Code: "89", reString: `(^|[^а-яё])(Ямало\-Ненецк.+|ЯНАО([^а-яё]|$))`},
	}
	for i := range c.regions {
		if c.regions[i].reString != "" {
			c.regions[i].re = regexp.MustCompile(fmt.Sprintf("(?i)%s", c.regions[i].reString))
		}
	}
}