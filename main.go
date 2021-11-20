package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"os/exec"
	"net/url"
)

//made by JustFoxxo so have fun!

type A struct {
	A map[string]string
}

type JSON struct {
	Li []A
}

func getAuthor(text string, mod *string) string {
	regex := regexp.MustCompile(`\((\w+)\)$`)
	author := regex.FindAllString(text, 1)
	*mod = strings.ReplaceAll(text, author[0], "")

	author[0] = strings.ReplaceAll(author[0], "(", "")
	author[0] = strings.ReplaceAll(author[0], ")", "")

	return author[0]
}

func writeFileSpace(modlist string, file string) {
	jsonFile, _ := os.ReadFile(file)
	modlistWrite := fmt.Sprintf("%v\n%v", string(jsonFile), modlist)
	os.WriteFile(file, []byte(modlistWrite), 0777)
}

func scan(s *string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*s = scanner.Text()
}

func main() {
	var jsonUnmarshamal JSON
	jsonFile, _ := os.ReadFile("modlist.json")
	json.Unmarshal(jsonFile, &jsonUnmarshamal)

	for i := range jsonUnmarshamal.Li {
		var git string
		var short string

		text := jsonUnmarshamal.Li[i].A["#text"]
		href := jsonUnmarshamal.Li[i].A["@href"]
		author := getAuthor(text, &text)
		
		cmd := exec.Command("bash", "-c", fmt.Sprintf("firefox-nightly --new-tab %v", href))
		cmd.Run()

		fmt.Println(href)
		fmt.Printf("GitHub >")
		scan(&git)
		urlText := url.QueryEscape(text)
		command := fmt.Sprintf("firefox-nightly --new-tab \"https://www.curseforge.com/minecraft/mc-mods/search?&search=%v\"", urlText)
		cmd = exec.Command("bash", "-c", command)
		cmd.Run()
		fmt.Printf("Short  >")
		scan(&short)

		modlist := fmt.Sprintf("|%v|%v|%v|%v|%v|", text, author, href, git, short)

		writeFileSpace(modlist, "modlist.txt")
		fmt.Println(modlist)
	}
}
