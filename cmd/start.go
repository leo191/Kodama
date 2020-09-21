/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"io/ioutil"
	"bitbucket.org/leo191/kodama/utils"
	"github.com/spf13/cobra"
	"time"
	"github.com/elliotchance/sshtunnel"
)

var logger = utils.NewLogger("test.log")



// func runProbe(client *ssh.Client, source string){
// 	sess, err := client.NewSession()
// 	logger.Log(2, source, err)
// 	script, _ := os.OpenFile(source, os.O_RDWR, 0755)
// 	sess.Stdin = script
// 	cmd := "bash -s --"
// 	out, err := sess.CombinedOutput(cmd)
// 	logger.Log(3, string(out), err)
// 	defer sess.Close()
// 	defer script.Close()
// }


func readHosts(*cobra.Command, []string){

	tunnel := sshtunnel.NewSSHTunnel(
		// User and host of tunnel server, it will default to port 22
		// if not specified.
		"ec2-user@jumpbox.us-east-1.mydomain.com",
	 
		// Pick ONE of the following authentication methods:
		sshtunnel.PrivateKeyFile("/home/leo/Downloads/awstest.pem"), // 1. private key
	 
		// The destination host and port of the actual server.
		"dqrsdfdssdfx.us-east-1.redshift.amazonaws.com:5439",
		
		// The local port you want to bind the remote port to.
		// Specifying "0" will lead to a random port.
		"8443",
	 )
	 
	// f, err := ioutil.ReadFile("hosts.txt")
	// logger.Log(3, string(f), err)
	// client, _, err := connectToHost("ubuntu", string(f)+":22")
	// // for i := 0; i<10; i++ {
	// for i:=0;i<1000;i++ {
	// 	go runProbe(client, "process")
	// 	// go runProbe(client, "os")
	// }
	// // go runProbe(client, "process")
	// // go runProbe(client, "os")
	// 	// go runProbe(client, "os")
	// // }
	// time.Sleep(time.Second * 2)
	// defer client.Close()
}

// func connectToHost(user, hostport string) (*ssh.Client, *ssh.Session, error){
// 	sshConfig := &ssh.ClientConfig{
// 		User: user,
// 		Auth: []ssh.AuthMethod{
// 			publicKey("/home/leo/Downloads/awstest.pem"),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 		}
// 	client, err := ssh.Dial("tcp", hostport, sshConfig)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	session, err := client.NewSession()
// 	if err != nil {
// 		client.Close()
// 		return nil, nil, err
// 	}
// 	return client, session, nil
// }

// func publicKey(path string) ssh.AuthMethod {
// 	key, err := ioutil.ReadFile(path)
// 	if err != nil {
// 	 panic(err)
// 	}
// 	signer, err := ssh.ParsePrivateKey(key)
// 	if err != nil {
// 	 panic(err)
// 	}
// 	return ssh.PublicKeys(signer)
//    }
 //   



// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start kodama",
	Run: readHosts,
}

func init() {
	rootCmd.AddCommand(startCmd)

	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().StringP("toggle", "t", "duck", "Help message for toggle")
}
