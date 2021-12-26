package db

// This function will return a Snowflake ID
// to be used as primary key for tables that require it
func GenerateId() int64 {
	return snowflakeNode.Generate().Int64()
}
