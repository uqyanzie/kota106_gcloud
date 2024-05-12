package sqlgen

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/uqyanzie/kota106_gcloud/model"
)

func GenerateSql(request model.Request) string {
	var conditions []sq.Sqlizer

	if request.GetConnection() == model.AND {
		conditions = sq.And{}
	} else {
		conditions = sq.Or{}
	}

	for _, condition := range request.GetConditions() {
		switch condition.GetOperator() {
		case model.EQUAL:
			conditions = append(conditions, sq.Eq{condition.GetColumn(): condition.GetValue()})
		case model.NOT_EQUAL:
			conditions = append(conditions, sq.NotEq{condition.GetColumn(): condition.GetValue()})
		case model.GREATER_THAN:
			conditions = append(conditions, sq.Gt{condition.GetColumn(): condition.GetValue()})
		case model.LESS_THAN:
			conditions = append(conditions, sq.Lt{condition.GetColumn(): condition.GetValue()})
		case model.GREATER_THAN_OR_EQUAL:
			conditions = append(conditions, sq.GtOrEq{condition.GetColumn(): condition.GetValue()})
		case model.LESS_THAN_OR_EQUAL:
			conditions = append(conditions, sq.LtOrEq{condition.GetColumn(): condition.GetValue()})
		}
	}

	// Generate SQL
	sql, _, _ := sq.Select(request.GetProjection()...).From("talents").Where(conditions).ToSql()
	response := sql

	return response
}
