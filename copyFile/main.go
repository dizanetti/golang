package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {

	var jdk Jdk

	configJdkFile := open(configFile)
	defer configJdkFile.Close()

	config, errConfig := ioutil.ReadAll(configJdkFile)
	checkError(errConfig)

	json.Unmarshal([]byte(config), &jdk)

	for i := 0; i < len(jdk.JdkVersion); i++ {
		if jdk.JdkVersion[i].Version == os.Args[1] {
			copy(jdk.JdkVersion[i].Java, jdk.JavaHome+javaExe)
			copy(jdk.JdkVersion[i].JavaC, jdk.JavaHome+javacExe)
			copy(jdk.JdkVersion[i].JavaW, jdk.JavaHome+javawExe)

			break
		}
	}
}
