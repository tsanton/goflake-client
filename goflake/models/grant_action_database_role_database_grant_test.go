package models_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	g "github.com/tsanton/goflake-client/goflake"
	i "github.com/tsanton/goflake-client/goflake/integration"
	a "github.com/tsanton/goflake-client/goflake/models/assets"
	ai "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	d "github.com/tsanton/goflake-client/goflake/models/describables"
	e "github.com/tsanton/goflake-client/goflake/models/entities"
	"github.com/tsanton/goflake-client/goflake/models/enums"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Test_grant_database_role_database_privilege(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	db := a.Database{
		Name:    "IGT_DATABASE_ROLES",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "SYSADMIN"},
	}
	databaseRole := a.DatabaseRole{
		Name:         "IGT_DEMO_ROLE",
		DatabaseName: db.Name,
		Comment:      "integration test goflake",
		Owner:        &a.Role{Name: "USERADMIN"},
	}
	privilege := a.GrantAction{
		Target:     &a.GrantActionDatabaseGrant[*a.DatabaseRole]{Principal: &databaseRole, DatabaseName: db.Name},
		Privileges: []enums.Privilege{enums.PrivilegeCreateDatabaseRole},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &databaseRole, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	grants, err := g.DescribeMany[*e.Grant](cli, &d.Grant{Principal: &d.DatabaseRole{Name: databaseRole.Name, DatabaseName: databaseRole.DatabaseName}})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Len(t, grants, 2) //Database roles are created with usage on database
	dbCreateRole, ok := lo.Find(grants, func(i *e.Grant) bool { return i.Privilege == enums.PrivilegeCreateDatabaseRole })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", dbCreateRole.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectDatabase, dbCreateRole.GrantedOn)
}

func Test_grant_database_role_database_privileges(t *testing.T) {
	/* Arrange */
	cli := i.Goflake()
	defer cli.Close()
	stack := u.Stack[ai.ISnowflakeAsset]{}
	defer g.DeleteAssets(cli, &stack)

	db := a.Database{
		Name:    "IGT_DEMO",
		Comment: "integration test goflake",
		Owner:   &a.Role{Name: "SYSADMIN"},
	}
	databaseRole := a.DatabaseRole{
		Name:         "IGT_DEMO_ROLE",
		DatabaseName: db.Name,
		Comment:      "integration test goflake",
		Owner:        &a.Role{Name: "USERADMIN"},
	}
	privilege := a.GrantAction{
		Target:     &a.GrantActionDatabaseGrant[*a.DatabaseRole]{Principal: &databaseRole, DatabaseName: db.Name},
		Privileges: []enums.Privilege{enums.PrivilegeCreateDatabaseRole, enums.PrivilegeMonitor},
	}

	/* Act */
	i.ErrorFailNow(t, g.RegisterAsset(cli, &db, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &databaseRole, &stack))
	i.ErrorFailNow(t, g.RegisterAsset(cli, &privilege, &stack))

	grants, err := g.DescribeMany[*e.Grant](cli, &d.Grant{Principal: &d.DatabaseRole{Name: databaseRole.Name, DatabaseName: databaseRole.DatabaseName}})

	/* Assert */
	i.ErrorFailNow(t, err)
	assert.Len(t, grants, 3) //Database roles are created with usage on database

	dbCreateRole, ok := lo.Find(grants, func(i *e.Grant) bool { return i.Privilege == enums.PrivilegeCreateDatabaseRole })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", dbCreateRole.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectDatabase, dbCreateRole.GrantedOn)

	dbMonitor, ok := lo.Find(grants, func(i *e.Grant) bool { return i.Privilege == enums.PrivilegeMonitor })
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", dbMonitor.GrantedBy)
	assert.Equal(t, enums.SnowflakeObjectDatabase, dbMonitor.GrantedOn)
}