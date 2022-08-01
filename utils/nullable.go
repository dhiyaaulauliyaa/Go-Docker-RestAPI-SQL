package util

import (
	"database/sql"

	"gopkg.in/guregu/null.v4"
)

func Int32ToNullable(data sql.NullInt32) null.Int {
	var result null.Int
	if data.Valid {
		result = null.Int{
			NullInt64: sql.NullInt64{
				Int64: int64(data.Int32),
				Valid: data.Valid,
			},
		}
	}

	return result
}

func NullableToInt32(data null.Int) sql.NullInt32 {
	return sql.NullInt32{
		Int32: int32(data.Int64),
		Valid: data.Valid,
	}
}

func StringToNullable(data sql.NullString) null.String {
	var result null.String
	if data.Valid {
		result = null.NewString(
			data.String,
			data.Valid,
		)
	}

	return result
}

func NullableToString(data null.String) sql.NullString {
	return sql.NullString{
		String: data.String,
		Valid:  data.Valid,
	}
}
