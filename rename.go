package main

import ("fmt"
		"os"
		"io/ioutil"
		"math/rand"
		"time"
)

func main() {

	fmt.Print("This is the filename shuffle application. Please remember to specify your folder path as an argument.\n")
	fmt.Print("For example \"C:/Users/users/Desktop/folder\"\n")
	fmt.Printf("Listing all files in the directory %s.\n",os.Args[1])

	//Now Lets Get the Files in the Directory
	files, _ := ioutil.ReadDir(os.Args[1])
	fmt.Printf("The folder contains %d items\n", len(files))
	oldstring := ""
	for _, f := range files {
		//fmt.Printf("%s\n",RandomString(5))
		newString := RandomString(5)
		
		for {
			if oldstring == newString {
				newString = RandomString(5)
			} else { 
				oldstring = newString
				break
			}
		}

		fmt.Printf("Processing File %s with OldString: %s and NewString %s\n", f.Name(), oldstring, newString)
		newName := newString + "-" + f.Name()
		os.Rename(os.Args[1]+"/"+f.Name(), os.Args[1] + "/" + newName)
        fmt.Printf("File Renamed to %s\n", newName)
    }
}


func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}