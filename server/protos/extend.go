package protos

var (
	actionMap = map[string]string{
		"A0000": ActionsEnum_name[0],
		"A0001": ActionsEnum_name[1],
		"A1001": ActionsEnum_name[2],
		"A1002": ActionsEnum_name[3],
		"A1003": ActionsEnum_name[4],
		"A2000": ActionsEnum_name[5],
		"A2001": ActionsEnum_name[6],
		"A3000": ActionsEnum_name[7],
		"A3001": ActionsEnum_name[8],
	}
)

func ActionCodeToIndex(code string) string {
	if val, ok := actionMap[code]; ok {
		return val
	} else {
		return ""
	}
}
