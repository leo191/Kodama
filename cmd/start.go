/*
Copyright © 2020 SUBHADIP BANERJEE subhadipbanerjee527@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"bitbucket.org/leo191/kodama/utils"
	"github.com/spf13/cobra"
	"context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"bitbucket.org/leo191/kodama/orm"
)

var logger = utils.NewLogger("kodama.log")






func readConfig(cmd *cobra.Command, args []string){
	//
	// part where it reads config and decides backend
	//
	addr := "localhost"
	clientOptions := options.Client().ApplyURI("mongodb://"+ addr +":27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	logger.Log(3, "Connected with "+ addr, err)
	
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	
	logger.Log(3, "Pingged "+ addr, err)

	// cpu := orm.Probe{"check_cpu", "/usr/lib/probes/check-cpu.sh", 10, []orm.STATUS{orm.OK, orm.CRITICAL}}
	collection := client.Database("Test").Collection("Probe")
	filter := bson.D{{"name", "check_cpu"}}
	update := bson.D{
		{"$inc", bson.D{
			{"interval", 1},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	logger.Log(2, "Updated ", err)
	fmt.Printf("%v   %v", updateResult.MatchedCount, updateResult.ModifiedCount)
	var cpucheck orm.Probe
	err = collection.FindOne(context.TODO(), filter).Decode(&cpucheck)
	logger.Log(3, "Found ", err)
	fmt.Println(cpucheck)
	
	

}

// Authenticate mongodb
func authDB() {} 

// Fetch host settings from backend(mongodb).
func fetchSettings() {}

// Fork separate goroutines for every check with intervel as context.WithTimeout.
func spawnKodama() {} 



// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start kodama",
	Run: readConfig,
}

func init() {
	rootCmd.AddCommand(startCmd)

	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:


	// Declare all flags like log path, mongoport, host, password, probes path.
	startCmd.Flags().StringP("toggle", "t", "duck", "Help message for toggle")
}
