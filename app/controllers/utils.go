package controllers

import (
	"fmt"
	"os"
	"bufio"
	//"//io/ioutil"
	"html/template"
	//"path/filepath"
	"strings"

	//"github.com/revel/revel"
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v2"
	"github.com/pksunkara/pygments"
)

type YamlJekyllPage struct {
	Title string ` yaml:"title" `
	Layout string ` yaml:"layout" ` // not used in revel
}

type PageData struct {
	Title string
	HTML template.HTML
}

// a mardown file has some yaml at the top  which contains
// title, so this scans line by line.. (TODO jekyll tag replace)
// ---
// title: foo
// layout: manual
// ---
// TODO , this is real slow and needs a regex expert
func ReadMarkdownPage( section, page string) PageData {

	var pd PageData

	md_file_path := CLONES_DIR + "/revel.github.io/" + section + "/" + page + ".md"

	file, err := os.Open(md_file_path)
	if err != nil {
		fmt.Println("error md", err)
	}
	defer file.Close()

	yaml_bounds := "---"
	yaml_str := ""
	body_str := ""
	found_yaml_start := false
	in_yaml := false // we always expect yaml ??
	in_code := false
	code_str := ""
	lexer := ""
	scanner := bufio.NewScanner(file)
	for  scanner.Scan() {
		line := scanner.Text()
		if line == yaml_bounds {
			if found_yaml_start == false {
				in_yaml = true
				found_yaml_start = true
			} else {
				in_yaml = false
			}
		} else {
			if in_yaml {
				yaml_str += line + "\n"
			} else {
				// TODO need a regex for "{%highlight foo %}"
				if len(line) > 2 && line[0:2] == "{%" && strings.Contains(line, "endhighlight")  == false && strings.Contains(line, "highlight")  {
					//fmt.Println("GOT CODE=" , line)
					xline := line
					xline = strings.Replace(xline, "{%", "", 1)
					xline = strings.Replace(xline, "%}", "", 1)
					xline = strings.Replace(xline, "highlight", "", 1)
					xline = strings.TrimSpace(xline)
					//fmt.Println("GOT CODE=" , line, xline)
					lexer = xline
					//if line == "{% highlight go %}" {
					//body_str += "``` go\n"
					body_str += "\n"
					code_str = ""
					in_code = true

					//}
				} else if len(line) > 2 && line[0:2] == "{%" && strings.Contains(line, "endhighlight")  {
					//fmt.Println("END CODE=" , line)
					//body_str += "##########" + code_str + "###########"
					hi_str := pygments.Highlight(code_str, lexer, "html", "utf-8")
					//fmt.Println("hi=", hi_str)
					body_str += string(hi_str)
					//body_str += "```\n"
					body_str += "\n"
					in_code = false

				} else {
					if in_code {
						code_str += line + "\n"
					} else {
						body_str += line+"\n"
					}

				}
			}
		}

	}
	if err := scanner.Err(); err != nil {
		//log.Fatal(err)
	}


	// parse yaml header bit
	var yami YamlJekyllPage
	err = yaml.Unmarshal([]byte(yaml_str), &yami)
	if err != nil {
		fmt.Println("error md", err)
	}
	pd.Title = yami.Title

	//fmt.Println("===", pd.Title, yami)

	// convert markdown
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS

	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	//htmlFlags |= blackfriday.HTML_GITHUB_BLOCKCODE
	//htmlFlags |= blackfriday.HTML_TOC

	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")
	output := blackfriday.Markdown([]byte(body_str), renderer, extensions)

	pd.HTML = template.HTML(output)

	//fmt.Println("yamiii", yami)

	return pd

}
