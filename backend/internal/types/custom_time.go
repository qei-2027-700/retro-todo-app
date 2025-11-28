package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// CustomTime は ISO8601 形式（マイクロ秒なし）でJSONをシリアライズする time.Time のラッパー
type CustomTime time.Time

// MarshalJSON は JSON エンコード時に ISO8601 形式（マイクロ秒なし）にフォーマット
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	t := time.Time(ct)
	if t.IsZero() {
		return []byte("null"), nil
	}
	// ISO8601 形式（マイクロ秒なし）: 2025-11-28T15:19:49Z
	formatted := t.UTC().Format("2006-01-02T15:04:05Z")
	return []byte(fmt.Sprintf(`"%s"`, formatted)), nil
}

// UnmarshalJSON は JSON デコード時に時刻をパース
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	str := string(data)
	if len(str) > 2 {
		str = str[1 : len(str)-1] // クォートを除去
	}

	// 複数のフォーマットを試行
	formats := []string{
		"2006-01-02T15:04:05Z",
		time.RFC3339,
		time.RFC3339Nano,
	}

	var t time.Time
	var err error
	for _, format := range formats {
		t, err = time.Parse(format, str)
		if err == nil {
			*ct = CustomTime(t)
			return nil
		}
	}

	return err
}

// Scan は database/sql のスキャン時に使用
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		*ct = CustomTime(time.Time{})
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*ct = CustomTime(v)
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into CustomTime", value)
	}
}

// Value は database/sql の値変換時に使用
func (ct CustomTime) Value() (driver.Value, error) {
	return time.Time(ct), nil
}
