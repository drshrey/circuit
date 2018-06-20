package graphql_resolver_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/codeamp/circuit/plugins/codeamp/db"
	graphql_resolver "github.com/codeamp/circuit/plugins/codeamp/graphql"
	graphql "github.com/graph-gophers/graphql-go"

	"github.com/codeamp/circuit/plugins/codeamp"
	"github.com/codeamp/circuit/plugins/codeamp/model"
	"github.com/codeamp/circuit/test"
	uuid "github.com/satori/go.uuid"

	log "github.com/codeamp/logger"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/assert"
)

type UserTestSuite struct {
	suite.Suite
	Resolver     *graphql_resolver.Resolver
	UserResolver *graphql_resolver.UserResolver

	cleanupUserIDs []uuid.UUID
	createdUserID  uuid.UUID
}

func (suite *UserTestSuite) SetupTest() {
	migrators := []interface{}{
		&model.User{},
		&model.UserPermission{},
	}

	db, err := test.SetupResolverTest(migrators)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
		return
	}

	_ = codeamp.CodeAmp{}
	_ = &graphql_resolver.Resolver{DB: db, Events: nil, Redis: nil}

	suite.Resolver = &graphql_resolver.Resolver{DB: db}
	suite.UserResolver = &graphql_resolver.UserResolver{DBUserResolver: &db_resolver.UserResolver{DB: db}}
}

/* Test successful project permissions update */
func (ts *UserTestSuite) Test1GormCreateUser() {
	user := model.User{
		Email:    fmt.Sprintf("test%d@example.com", time.Now().Unix()),
		Password: "example",
	}

	res := ts.UserResolver.DBUserResolver.DB.Create(&user)
	if res.Error != nil {
		assert.FailNow(ts.T(), res.Error.Error())
	}

	ts.createdUserID = user.ID
	log.Debug("Created User with ID ", user.ID)
}

func (ts *UserTestSuite) Test2GQLBGetUser() {
	ctx := context.WithValue(context.Background(), "jwt", model.Claims{
		UserID:      "foo",
		Email:       "foo@gmail.com",
		Permissions: []string{"admin"},
	})

	var usr struct {
		ID *graphql.ID
	}
	graphqlID := graphql.ID(ts.createdUserID.String())
	usr.ID = &graphqlID

	_, err := ts.Resolver.User(ctx, &usr)
	if err != nil {
		ts.T().Log(ts.createdUserID.String())
		assert.FailNow(ts.T(), err.Error())
	}
}

func (ts *UserTestSuite) Test3GormDeleteUser() {
	user := model.User{
		Model: model.Model{ID: ts.createdUserID},
	}

	res := ts.UserResolver.DBUserResolver.DB.Delete(&user)
	if res.Error != nil {
		assert.FailNow(ts.T(), res.Error.Error())
	}
}

func (ts *UserTestSuite) Test4GormCreate5Users() {
	for x := 0; x < 5; x++ {
		user := model.User{
			Email:    fmt.Sprintf("test%d@example.com", time.Now().Unix()),
			Password: "example",
		}

		res := ts.UserResolver.DBUserResolver.DB.Create(&user)
		if res.Error != nil {
			assert.FailNow(ts.T(), res.Error.Error())
		}

		ts.cleanupUserIDs = append(ts.cleanupUserIDs, user.ID)
	}
}

func (ts *UserTestSuite) Test5GQLBGet5Users() {
	ctx := context.WithValue(context.Background(), "jwt", model.Claims{
		UserID:      "foo",
		Email:       "foo@gmail.com",
		Permissions: []string{"admin"},
	})

	var usr struct {
		ID *graphql.ID
	}
	graphqlID := graphql.ID(ts.createdUserID.String())
	usr.ID = &graphqlID

	res, err := ts.Resolver.Users(ctx)
	if err != nil {
		ts.T().Log(ts.createdUserID.String())
		assert.FailNow(ts.T(), err.Error())
	}

	assert.True(ts.T(), len(res) >= 5)
}

func TearDownTest(ts *UserTestSuite) {
	ts.UserResolver.DBUserResolver.DB.Unscoped().Delete(&model.User{Model: model.Model{ID: ts.createdUserID}})

	for _, i := range ts.cleanupUserIDs {
		ts.UserResolver.DBUserResolver.DB.Unscoped().Delete(&model.User{Model: model.Model{ID: i}})
	}
}

func TestSuiteUserResolver(t *testing.T) {
	ts := new(UserTestSuite)
	suite.Run(t, ts)

	TearDownTest(ts)
}
