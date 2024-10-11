package enum

type Name string

const (
	ZhangSan Name = "ZhangSan"
	LiShi    Name = "LiShi"
	Other    Name = "Other"
)

func convert(name Name) string {
	switch name {
	case ZhangSan:
		return "San Zhang"
	case LiShi:
		return "Shi Li"
	case Other:
	default:
	}

	return "invalid"
}
