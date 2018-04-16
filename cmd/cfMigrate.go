package cmd

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"net"
	"crypto/tls"
	"time"
	"log"
	"strconv"
	"regexp"
	// "context"
	"encoding/json"
	// "os"

	"github.com/spf13/cobra"
	"github.com/checkr/codeflow/server/plugins/codeflow"
	codeamp "github.com/codeamp/circuit/plugins/codeamp"
	codeamp_resolvers "github.com/codeamp/circuit/plugins/codeamp/resolvers"
	codeamp_plugins "github.com/codeamp/circuit/plugins"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/go-bongo/bongo"
	mgo "gopkg.in/mgo.v2"
	"github.com/spf13/viper"	
	"github.com/davecgh/go-spew/spew"
	// uuid "github.com/satori/go.uuid" 
)
var codeflowDB *bongo.Connection

// migrateCmd represents the migrate command
var cfMigrateCmd = &cobra.Command{
	Use:   "cfmigrate",
	Short: "Migrate Codeflow projects to CodeAmp",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[+] Codeflow to CodeAmp migration started.\n ---------------------------------------- \n")

		// init DB connection for codeflow
		fmt.Println("[*] Initializing Codeflow DB Connection")
		createCodeflowDBConnection()
		fmt.Println("[+] Successfully initialized Codeflow DB Connection")

		// init DB connection for codeamp
		fmt.Println("[*] Initializing CodeAmp Resolver")
		codeampDB, err := createCodeampDB()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("[+] Successfully initialized Codeamp Resolver")

		fmt.Println("[*] Initializing CodeAmp transistor...")		

		fmt.Println("[+] Successfully initialized CodeAmp transistor")		

		// adminContext := context.WithValue(context.Background(), "jwt", codeamp_resolvers.Claims{
		// 	UserID:      uuid.FromStringOrNil("codeamp").String(),
		// 	Email:       "codeamp",
		// 	Permissions: []string{"admin"},
		// })

		// TODO: Remove for production
		fmt.Println("[*] Cleaning Codeamp DB of all rows. REMOVE FOR PRODUCTION.")
		codeampDB.Unscoped().Delete(&codeamp_resolvers.Service{})
		codeampDB.Unscoped().Delete(&codeamp_resolvers.Secret{})
		codeampDB.Unscoped().Delete(&codeamp_resolvers.SecretValue{})
		codeampDB.Unscoped().Delete(&codeamp_resolvers.Project{})
		codeampDB.Unscoped().Delete(&codeamp_resolvers.ServiceSpec{})
		codeampDB.Unscoped().Delete(&codeamp_resolvers.Feature{})
		codeampDB.Unscoped().Delete(&codeamp_resolvers.Release{})
		fmt.Println("[+] Successfully cleaned Codeamp DB of all rows")
		
		// create CodeAmp project objects from codeflow
		projects := []codeflow.Project{}
		results := codeflowDB.Collection("projects").Find(bson.M{ })
		results.Query.All(&projects)

		reg, err := regexp.Compile("[^0-9]+")
		if err != nil {
			panic(err.Error())
		}

		// create service specs
		fmt.Println("[*] Porting service specs")
		codeflowServiceSpecs := []codeflow.ServiceSpec{}
		results = codeflowDB.Collection("serviceSpecs").Find(bson.M{})
		results.Query.All(&codeflowServiceSpecs)
		for _ , codeflowServiceSpec := range codeflowServiceSpecs {
			fmt.Println(fmt.Sprintf("[*] Transferring %s", codeflowServiceSpec.Name))
			codeampServiceSpec := codeamp_resolvers.ServiceSpec{
				Name: codeflowServiceSpec.Name,
				CpuRequest: reg.ReplaceAllString(codeflowServiceSpec.CpuRequest, ""),
				CpuLimit: reg.ReplaceAllString(codeflowServiceSpec.CpuLimit, ""),
				MemoryRequest: reg.ReplaceAllString(codeflowServiceSpec.MemoryRequest, ""),
				MemoryLimit: reg.ReplaceAllString(codeflowServiceSpec.MemoryLimit, ""),
				TerminationGracePeriod: strconv.Itoa(int(codeflowServiceSpec.TerminationGracePeriodSeconds)),
			}
			codeampDB.Create(&codeampServiceSpec)
			fmt.Println(fmt.Sprintf("[+] Successfully transferred %s", codeflowServiceSpec.Name))
		}
		fmt.Println("[+] Finished porting service spec \n\n")

		codeampUser := codeamp_resolvers.User{}
		if codeampDB.Where("email = ?", "kilgore@kilgore.trout").First(&codeampUser).RecordNotFound() {
			panic("Could not find CodeAmp user with email kilgore@kilgore.trout")
		}
		
		fmt.Println("[*] Porting projects")
		for _, project := range projects {
			fmt.Println(fmt.Sprintf("[*] Creating corresponding CodeAmp project for %s", project.Slug))
			codeampProject := codeamp_resolvers.Project{
				Name: project.Name,
				Slug: project.Slug,
				Repository: project.Repository,
				Secret: project.Secret,
				GitUrl: project.GitUrl,
				GitProtocol: project.GitProtocol,
				RsaPrivateKey: project.RsaPrivateKey,
				RsaPublicKey: project.RsaPublicKey,
			}
			codeampDB.Create(&codeampProject)			

			// fmt.Println("[*] Porting features")
			// // find the features tied to the project
			// codeflowFeatures := []codeflow.Feature{}
			// results = codeflowDB.Collection("features").Find(bson.M{ "projectId": bson.ObjectId(project.Id) })
			// results.Query.All(&codeflowFeatures)


			// for _, feature := range codeflowFeatures {
			// 	fmt.Println("[*] Porting feature ", feature.Hash)
			// 	// create codeamp feature
			// 	codeampFeature := codeamp_resolvers.Feature{
			// 		ProjectID: codeampProject.Model.ID,
			// 		Message: feature.Message,
			// 		User: feature.User,
			// 		Ref: feature.Ref,
			// 		ParentHash: feature.ParentHash,
			// 		Created: feature.Created,
			// 		Hash: feature.Hash,
			// 	}
			// 	codeampDB.Create(&codeampFeature)
			// }
			// fmt.Println("[+] Successfully ported features! \n")


			fmt.Println("[*] Porting environments...")
			// get envs in codeamp
			envs := []codeamp_resolvers.Environment{}
			codeampDB.Find(&envs)

			for _, env := range envs {
				fmt.Println(fmt.Sprintf("[*] Filling in environment %s", env.Key))


				fmt.Println("[*] Porting secrets...")
				// find and create the secrets tied to the project
				// secret := codeflow.Secret{}
				codeflowSecrets := []codeflow.Secret{}

				// bson.M{ "deleted": false, "projectId": project.Id } not working
				// so doing a manually-looped filter
				results = codeflowDB.Collection("secrets").Find(bson.M{ "projectId": bson.ObjectId(project.Id), "deleted": false })
				results.Query.All(&codeflowSecrets)

				codeampSecrets  := []codeamp_resolvers.Secret{}
				for _, secret := range codeflowSecrets {
					fmt.Println(fmt.Sprintf("[*] Creating secret %s", secret.Key))
					
					isSecret := false
					if string(secret.Type) == "protected-env" {
						isSecret = true
					}

					codeampSecret := codeamp_resolvers.Secret{
						Key: secret.Key,
						Scope: codeamp_resolvers.GetSecretScope("project"),
						EnvironmentID: env.Model.ID,
						IsSecret: isSecret,
						ProjectID: codeampProject.Model.ID,
						Type: codeamp_plugins.GetType(string(secret.Type)),
					}
					codeampDB.Create(&codeampSecret)

					codeampSecretValue := codeamp_resolvers.SecretValue{
						SecretID: codeampSecret.Model.ID,
						Value: secret.Value,
						UserID: codeampUser.Model.ID,
					}
					codeampDB.Create(&codeampSecretValue)
					codeampSecret.Value = codeampSecretValue
					codeampSecrets = append(codeampSecrets, codeampSecret)

					fmt.Println(fmt.Sprintf("[+] Successfully created Secret %s => %s", secret.Key, secret.Value))
				}
				fmt.Println("[+] Successfully ported secrets! \n\n")


				fmt.Println("[*] Porting services...")
				// find the services tied to the project
				codeflowServices := []codeflow.Service{}
				results = codeflowDB.Collection("services").Find(bson.M{ "projectId": bson.ObjectId(project.Id) })
				results.Query.All(&codeflowServices)
				codeampServices := []codeamp_resolvers.Service{}
				for _, codeflowService := range codeflowServices {
					if string(codeflowService.State) != "deleted" {
						fmt.Println("[*] Porting service ", codeflowService.Name, codeflowService.Id, codeflowService.SpecId)
						// get service spec
						codeflowServiceSpec := codeflow.ServiceSpec{}
						results = codeflowDB.Collection("serviceSpecs").Find(bson.M{ "_id": bson.ObjectId(codeflowService.SpecId) })
						results.Query.One(&codeflowServiceSpec)
	
						codeampServiceSpec := codeamp_resolvers.ServiceSpec{}
						if codeampDB.Where("name = ?", codeflowServiceSpec.Name).First(&codeampServiceSpec).RecordNotFound() {
							fmt.Println(fmt.Sprintf("[-] Could not find ServiceSpec %s in CodeAmp", codeflowServiceSpec.Name))
							continue
						}
	
						codeampServiceType := codeamp_plugins.GetType("general")
						if codeflowService.OneShot {
							codeampServiceType = codeamp_plugins.GetType("one-shot")
						}
	
						codeampService := codeamp_resolvers.Service{
							ProjectID: codeampProject.Model.ID,
							ServiceSpecID: codeampServiceSpec.Model.ID,
							Command: codeflowService.Command,
							EnvironmentID: env.Model.ID,
							Count: strconv.Itoa(codeflowService.Count),
							Type: codeampServiceType,
							Name: codeflowService.Name,
						}
						codeampDB.Create(&codeampService)
	
						// create ports arr
						codeampPorts := []codeamp_resolvers.ServicePort{}
						for _, codeflowPort := range codeflowService.Listeners {
							codeampPort := codeamp_resolvers.ServicePort{
								ServiceID: codeampService.Model.ID,
								Port: strconv.Itoa(codeflowPort.Port),
								Protocol: codeflowPort.Protocol,
							}
							codeampDB.Create(&codeampPort)
							codeampPorts = append(codeampPorts, codeampPort)
						}
						codeampService.Ports = codeampPorts
						codeampServices = append(codeampServices, codeampService)
					}
				}
				fmt.Println("[+] Succesfully ported services! \n")

				// create additional objects i.e. ProjectSettings, ProjectEnvironments
				fmt.Println("[*] Creating ProjectSettings... ", env, codeampProject.Slug)				
				projectSettings := codeamp_resolvers.ProjectSettings{
					EnvironmentID: env.Model.ID,
					ProjectID: codeampProject.Model.ID,
					GitBranch: "master",
					ContinuousDeploy: false,
				}
				codeampDB.Create(&projectSettings)
				fmt.Println("[+] Successfully created ProjectSettings")

				fmt.Println("[*] Creating ProjectEnvironment permission... ", env, codeampProject.Slug)
				projectEnvironment := codeamp_resolvers.ProjectEnvironment{
					EnvironmentID: env.Model.ID,
					ProjectID: codeampProject.Model.ID,
				}
				codeampDB.Create(&projectEnvironment)
				fmt.Println("[+] Successfully created ProjectEnvironment")	
				
				
				// Create project extensions
				fmt.Println("[*] Creating Project Extensions...")
				// Create DockerBuilder extension
				dockerBuilderDBExtension := codeamp_resolvers.Extension{}
				if codeampDB.Where("environment_id = ? and key = ?", env.Model.ID, "dockerbuilder").Find(&dockerBuilderDBExtension).RecordNotFound() {
					panic(err.Error())
				}

				dockerBuilderProjectExtension := codeamp_resolvers.ProjectExtension{
					ProjectID: codeampProject.Model.ID,
					ExtensionID: dockerBuilderDBExtension.Model.ID,
					State: codeamp_plugins.GetState("waiting"),
					StateMessage: "Migrated, click update to send an event.",
					Artifacts: postgres.Jsonb{[]byte("{}")},
					Config: dockerBuilderDBExtension.Config,
					CustomConfig: postgres.Jsonb{[]byte("{}")},
					EnvironmentID: env.Model.ID,
				}
				codeampDB.Create(&dockerBuilderProjectExtension)

				// get relevant information for project's corresponding load balancers in codeflow
				results = codeflowDB.Collection("extensions").Find(bson.M{ "projectId": bson.ObjectId(project.Id), "extension": "LoadBalancer" })
				codeflowLoadBalancers := []codeflow.LoadBalancer{}
				results.Query.All(&codeflowLoadBalancers)
				for _, codeflowLoadBalancer := range codeflowLoadBalancers {
					listenerPairs := []map[string]string{}
					for _, cfListenerPair := range codeflowLoadBalancer.ListenerPairs {
						listenerPairs = append(listenerPairs, map[string]string{
							"port": strconv.Itoa(cfListenerPair.Destination.Port),
							"containerPort": strconv.Itoa(cfListenerPair.Source.Port),
							"protocol": cfListenerPair.Destination.Protocol,
						})
					}

					codeflowService := codeflow.Service{}
					err = codeflowDB.Collection("services").FindById(bson.ObjectId(codeflowLoadBalancer.ServiceId), &codeflowService)
					if err != nil {
						panic(err.Error())
					}

					lbCustomConfig := map[string]interface{}{
						"name": codeflowLoadBalancer.Subdomain,
						"type": codeflowLoadBalancer.Type,
						"service": codeflowLoadBalancer.Name,
						"listener_pairs": listenerPairs,
					}
					marshaledLbCustomConfig, err := json.Marshal(lbCustomConfig)
					if err != nil {
						panic(err.Error())
					}
	
					// Create Kubernetes Deployments extension
					loadBalancersDBExtension := codeamp_resolvers.Extension{}
					if codeampDB.Where("environment_id = ? and key = ?", env.Model.ID, "kubernetesloadbalancers").Find(&loadBalancersDBExtension).RecordNotFound() {
						panic(err.Error())
					}				
					lbProjectExtension := codeamp_resolvers.ProjectExtension{
						ProjectID: codeampProject.Model.ID,
						ExtensionID: loadBalancersDBExtension.Model.ID,
						State: codeamp_plugins.GetState("waiting"),
						StateMessage: "Migrated, click update to send an event.",
						Artifacts: postgres.Jsonb{[]byte("{}")},
						Config: loadBalancersDBExtension.Config,
						CustomConfig: postgres.Jsonb{marshaledLbCustomConfig},
						EnvironmentID: env.Model.ID,
					}
					codeampDB.Create(&lbProjectExtension)			
				}

				// Create Kubernetes Deployments extension
				kubernetesDeploymentsDBExtension := codeamp_resolvers.Extension{}
				if codeampDB.Where("environment_id = ? and key = ?", env.Model.ID, "kubernetesdeployments").Find(&kubernetesDeploymentsDBExtension).RecordNotFound() {
					panic(err.Error())
				}		
				kubernetesProjectExtension := codeamp_resolvers.ProjectExtension{
					ProjectID: codeampProject.Model.ID,
					ExtensionID: kubernetesDeploymentsDBExtension.Model.ID,
					State: codeamp_plugins.GetState("waiting"),
					StateMessage: "Migrated, click update to send an event.",
					Artifacts: postgres.Jsonb{[]byte("{}")},
					Config: kubernetesDeploymentsDBExtension.Config,
					CustomConfig: postgres.Jsonb{[]byte("{}")},
					EnvironmentID: env.Model.ID,
				}
				codeampDB.Create(&kubernetesProjectExtension)						

				fmt.Println("[+] Successfully created project extensions\n\n")


				fmt.Println("[*] Porting Release...")
				// find and transform the most recent release tied to the project

				// marshaledCodeampServices, err := json.Marshal(codeampServices)
				// if err != nil {
				// 	panic(err.Error())
				// }

				// marshaledCodeampSecrets, err := json.Marshal(codeampSecrets)
				// if err != nil {
				// 	panic(err.Error())
				// }				

				codeampRelease := codeamp_resolvers.Release{
					ProjectID: codeampProject.Model.ID,
					EnvironmentID: env.Model.ID,
					UserID: codeampUser.Model.ID,
					State: codeamp_plugins.GetState("complete"),
					StateMessage: "migrated",
					// Services: postgres.Jsonb{marshaledCodeampServices},
					// Secrets: postgres.Jsonb{marshaledCodeampSecrets},
				}
				
				for {
					if codeflowDB.Session.Ping() == nil {							
						break
					}
					codeflowDB.Session.Refresh()
					time.Sleep(1)
				}

				results = codeflowDB.Collection("releases").Find(bson.M{ "projectId": bson.ObjectId(project.Id) })
				latestCodeflowRelease := codeflow.Release{}
				codeflowRelease := codeflow.Release{}
				for results.Next(&codeflowRelease) {
					if string(codeflowRelease.State) == "complete" && latestCodeflowRelease.Created.Unix() < codeflowRelease.Created.Unix() {
						latestCodeflowRelease = codeflowRelease
					} 
				}

				spew.Dump(latestCodeflowRelease.State, latestCodeflowRelease.Id.Hex())
				if string(latestCodeflowRelease.State) == "" || latestCodeflowRelease.Id.Hex() == "" {
					continue
				}
				
				if latestCodeflowRelease.Id.String() != "" {
					fmt.Println("[+] Found latest release! ", latestCodeflowRelease.Id, latestCodeflowRelease.HeadFeatureId, latestCodeflowRelease.TailFeatureId)					
					// head feature
					codeflowReleaseHeadFeature := codeflow.Feature{}

					results = codeflowDB.Collection("features").Find(bson.M{ "_id": bson.ObjectId(latestCodeflowRelease.HeadFeatureId) })
					results.Query.One(&codeflowReleaseHeadFeature)
					
					fmt.Println(codeflowReleaseHeadFeature.Message)

					codeampHeadFeature := codeamp_resolvers.Feature{
						ProjectID: codeampProject.Model.ID,
						Message: codeflowReleaseHeadFeature.Message,
						User: codeflowReleaseHeadFeature.User,
						Ref: codeflowReleaseHeadFeature.Ref,
						ParentHash: codeflowReleaseHeadFeature.ParentHash,
						Created: codeflowReleaseHeadFeature.Created,
						Hash: codeflowReleaseHeadFeature.Hash,
					}
					codeampDB.Create(&codeampHeadFeature)

					codeampRelease.HeadFeatureID = codeampHeadFeature.Model.ID				

					if latestCodeflowRelease.TailFeatureId != latestCodeflowRelease.HeadFeatureId {
						// tail feature
						codeflowReleaseTailFeature := codeflow.Feature{}
						results = codeflowDB.Collection("features").Find(bson.M{ "_id": bson.ObjectId(latestCodeflowRelease.TailFeatureId) })
						results.Query.One(&codeflowReleaseTailFeature)

						if codeflowReleaseTailFeature.Message == "" {
							spew.Dump(codeflowReleaseTailFeature)
							continue
						}

						fmt.Println(codeflowReleaseTailFeature.Message)
						codeampTailFeature := codeamp_resolvers.Feature{
							ProjectID: codeampProject.Model.ID,
							Message: codeflowReleaseTailFeature.Message,
							User: codeflowReleaseTailFeature.User,
							Ref: codeflowReleaseTailFeature.Ref,
							ParentHash: codeflowReleaseTailFeature.ParentHash,
							Created: codeflowReleaseTailFeature.Created,
							Hash: codeflowReleaseTailFeature.Hash,
						}
						codeampDB.Create(&codeampTailFeature)
						codeampRelease.TailFeatureID = codeampTailFeature.Model.ID
					} else {
						codeampRelease.TailFeatureID = codeampHeadFeature.Model.ID
					}

					codeampDB.Create(&codeampRelease)
					fmt.Println("[+] Successfully ported release \n")
				} else {
					fmt.Println("[.] No releases found.")
				}

				fmt.Println(fmt.Sprintf("Done filling objects in env %s", env.Key))				
			}		

			fmt.Println(fmt.Sprintf("[+] Successfully *fully* created %s for envs %s \n\n", project.Slug, envs))
		}

		fmt.Println("[+] Finished porting all projects!")
	},
}

func createCodeflowDBConnection() {
	var err error
	config := &bongo.Config{
		ConnectionString: viper.GetString("codeflow.mongodb.uri"),
		Database:         viper.GetString("codeflow.mongodb.database"),
	}

	if viper.GetBool("codeflow.mongodb.ssl") {
		if config.DialInfo, err = mgo.ParseURL(config.ConnectionString); err != nil {
			panic(fmt.Sprintf("cannot parse given URI %s due to error: %s", config.ConnectionString, err.Error()))
		}

		tlsConfig := &tls.Config{}
		config.DialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}

		config.DialInfo.Timeout = time.Second * viper.GetDuration("codeflow.mongodb.connection_timeout")
	}

	codeflowDB, err = bongo.Connect(config)
	if err != nil {
		log.Fatal(err)
	}

	// Try to reconnect if connection drops
	go func(session *mgo.Session) {
		var err error
		for {
			err = session.Ping()
			if err != nil {
				fmt.Println("Lost connection to MongoDB!!")
				session.Refresh()
				err = session.Ping()
				if err == nil {
					fmt.Println("Reconnect to MongoDB successful.")
				} else {
					panic("Reconnect to MongoDB failed!!")
				}
			}
			time.Sleep(time.Second * viper.GetDuration("codeflow.mongodb.health_check_interval"))
		}
	}(codeflowDB.Session)
}

func createCodeampDB() (resolver *gorm.DB, err error) {
	db, err := codeamp.NewDB(viper.GetString("codeamp.postgres.host"), viper.GetString("codeamp.postgres.port"), viper.GetString("codeamp.postgres.user"), viper.GetString("codeamp.postgres.dbname"), viper.GetString("codeamp.postgres.sslmode"), viper.GetString("codeamp.postgres.password"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func init() {
	RootCmd.AddCommand(cfMigrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}