package table

type TableInfo struct {
	TableName string   `json:"table_name"`
	Columns   []string `json:"columns"`
}

var talentTable = map[string]map[string]string{
	"talents": {
		"id":     "int",
		"name":   "string",
		"height": "float",
		"weight": "float",
		"age":    "int",
	},
}

func getKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func getValues(m map[string]string) []string {
	values := make([]string, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func GetTableInfo() TableInfo {

	return TableInfo{
		TableName: "talents",
		Columns:   getKeys(talentTable["talents"]),
	}
}
