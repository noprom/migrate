// +build clickhouse

package cli

import (
	_ "github.com/noprom/migrate/v4/database/clickhouse"
	_ "github.com/ClickHouse/clickhouse-go"
)
