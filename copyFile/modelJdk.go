package main

type Jdk struct {
	JavaHome   string    `json:"javaHome"`
	JdkVersion []Version `json:"jdk"`
}

type Version struct {
	Version string `json:"version"`
	Java    string `json:"java"`
	JavaC   string `json:"javac"`
	JavaW   string `json:"javaw"`
}
