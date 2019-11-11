package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gookit/color"
)

func main() {
	var err error
	var client = &http.Client{}

	file, err := os.Open("./email.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator
		email := scanner.Text()

		request, err := http.NewRequest("GET", "https://twitter.com/users/email_available?email="+email, nil)

		if err != nil {
			log.Fatal(err)
		}

		response, err := client.Do(request)

		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(response.Body)

		json.NewDecoder(response.Body)

		keyVal := make(map[string]string)
		json.Unmarshal(body, &keyVal)

		if keyVal["msg"] == "Tersedia!" {
			color.New(color.FgGreen, color.BgBlack).Println("Selamat, Email anda : "+email, "Dapat Didaftarkan.", "Message : ", keyVal["msg"])
			createFile(email, keyVal["msg"])
		} else {
			color.New(color.FgRed, color.BgBlack).Println("Email : "+email, keyVal["msg"])
			createFile(email, keyVal["msg"])
		}

	}

}

func createFile(email string, indicator string) {
	var newEmail = email + "\n"
	if indicator == "Tersedia!" {
		f, err := os.OpenFile("notRegistered.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte(newEmail)); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	} else {
		f, err := os.OpenFile("registered.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte(newEmail)); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
	
  //this is my change

}
