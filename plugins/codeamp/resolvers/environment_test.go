package codeamp_resolvers_test

import (
	"fmt"
	"log"
	"testing"

	resolvers "github.com/codeamp/circuit/plugins/codeamp/resolvers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	graphql "github.com/neelance/graphql-go"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EnvironmentTestSuite struct {
	suite.Suite
	Resolver *resolvers.Resolver
}

func (suite *EnvironmentTestSuite) SetupTest() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("../../../configs/circuit.test.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		viper.GetString("plugins.codeamp.postgres.host"),
		viper.GetString("plugins.codeamp.postgres.port"),
		viper.GetString("plugins.codeamp.postgres.user"),
		viper.GetString("plugins.codeamp.postgres.dbname"),
		viper.GetString("plugins.codeamp.postgres.sslmode"),
		viper.GetString("plugins.codeamp.postgres.password"),
	))
	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(
		&resolvers.Environment{},
	)
	suite.Resolver = &resolvers.Resolver{DB: db}
}

/* Test successful env. creation */
func (suite *EnvironmentTestSuite) TestCreateEnvironment() {
	envInput := resolvers.EnvironmentInput{
		Name:  "test",
		Key:   "foo",
		Color: "color",
	}

	envResolver, err := suite.Resolver.CreateEnvironment(nil, &struct{ Environment *resolvers.EnvironmentInput }{Environment: &envInput})
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(suite.T(), envResolver.Name(), "test")
	assert.Equal(suite.T(), envResolver.Key(), "foo")
	assert.Equal(suite.T(), envResolver.Color(), "color")
	assert.NotEqual(suite.T(), envResolver.Color(), "wrongcolor")

	suite.TearDownTest([]string{envResolver.Model.ID.String()})
}

/* Test successful env. update */
func (suite *EnvironmentTestSuite) TestUpdateEnvironment() {
	envInput := resolvers.EnvironmentInput{
		Name:  "test",
		Key:   "foo",
		Color: "color",
	}

	envResolver, err := suite.Resolver.CreateEnvironment(nil, &struct{ Environment *resolvers.EnvironmentInput }{Environment: &envInput})
	if err != nil {
		log.Fatal(err.Error())
	}

	// update environment's name with same id
	envId := envResolver.Model.ID.String()
	envInput.ID = &envId
	envInput.Color = "red"
	envInput.Name = "test2"
	// key SHOULD be ignored
	envInput.Key = "diffkey"

	updateEnvResolver, err := suite.Resolver.UpdateEnvironment(nil, &struct{ Environment *resolvers.EnvironmentInput }{Environment: &envInput})
	if err != nil {
		log.Fatal(err.Error())
	}

	assert.Equal(suite.T(), updateEnvResolver.ID(), graphql.ID(envResolver.Model.ID.String()))
	assert.Equal(suite.T(), updateEnvResolver.Name(), "test2")
	assert.Equal(suite.T(), updateEnvResolver.Color(), "red")
	assert.Equal(suite.T(), updateEnvResolver.Key(), "foo")
	assert.NotEqual(suite.T(), updateEnvResolver.Name(), "diffkey")

	suite.TearDownTest([]string{updateEnvResolver.Model.ID.String()})
}

func (suite *EnvironmentTestSuite) TearDownTest(ids []string) {
	for _, id := range ids {
		suite.Resolver.DB.Where("id = ?", id).Delete(&resolvers.Environment{})
	}
}

func TestEnvironmentTestSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentTestSuite))
}