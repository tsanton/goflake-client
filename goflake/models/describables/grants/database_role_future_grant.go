package rolegrants

import (
	"fmt"

	d "github.com/tsanton/goflake-client/goflake/models/describables"
)

var (
	_ d.ISnowflakeDescribable = &DatabaseRoleFutureGrant{}
)

type DatabaseRoleFutureGrant struct {
	RoleName     string
	DatabaseName string
}

Error here to easily back to TODO: SHOW FUTURE GRANTS TO DATABASE ROLE not yet implemented
func (r *DatabaseRoleFutureGrant) GetDescribeStatement() string {
	return fmt.Sprintf(`
with show_future_grants_to_database_role as procedure(role_name varchar, database_name varchar)
	returns variant not null
	language python
	runtime_version = '3.8'
	packages = ('snowflake-snowpark-python')
	handler = 'show_future_grants_to_database_role_py'
as '
def show_future_grants_to_database_role_py(snowpark_session, role_name_py:str, database_name_py:str):
	res = []
	for row in snowpark_session.sql(f"SHOW FUTURE GRANTS TO DATABASE ROLE {database_name_py.upper()}.{role_name_py.upper()}").to_local_iterator():
		res.append(row.as_dict())
	return {
		"role_name": role_name_py,
		"grants": res
	}
'
call show_future_grants_to_database_role('%[1]s', '%[2]s');
	`, r.RoleName, r.DatabaseName)
}

func (*DatabaseRoleFutureGrant) IsProcedure() bool {
	return true
}
