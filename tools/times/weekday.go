package times

var weekdayText = map[string]string {
	"Monday": "星期一",
	"Tuesday": "星期二",
	"Wednesday": "星期三",
	"Thursday": "星期四",
	"Friday": "星期五",
	"Saturday": "星期六",
	"Sunday": "星期日",
}

func WeekdayText(weekday string) string {
	return weekdayText[weekday]
}
