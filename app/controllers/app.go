package controllers

import (
	"fmt"
	"os"
	"bufio"
	//"//io/ioutil"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/revel/revel"
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v2"
	"github.com/pksunkara/pygments"
)

var CLONES_DIR = os.Getenv("GOPATH") + "/src/github.com/pedromorgan/revel-www/externals"

type CurrPage struct {
	//Title string
	Version string
	SectionUrl string
	SectionTitle string
	PageUrl string
	PageTitle string
	//Version string
	Lang string
}

//var Site *SiteStruct

func GetCurrPage(section, section_title, version, lang, page string) CurrPage {

	s := CurrPage{SectionUrl: section, SectionTitle: section_title, PageUrl: page, Version: version, Lang: lang}
	return s
}

type App struct {
	*revel.Controller
}

func (c App) IndexPage() revel.Result {
	return c.Render()
}

// Only allow spiders on prod site
func (c App) RobotsTxt() revel.Result {

	s := "User-agent: *\n"
	if revel.Config.BoolDefault("site.robots", false)  == false {
		s += "Disallow: /\n"
	}
	s += "\n"

	return c.RenderText(s)
}


func (c App) ManualPage(ver, lang, page string) revel.Result {

	site_section := "manual"

	cPage := GetCurrPage(site_section, "Manual", ver, lang, page)


	nav := GetNav(site_section)
	c.RenderArgs["nav"] = nav


	page_no_ext := page
	if filepath.Ext(page) == ".html" { // wtf includes the .
		page_no_ext = page[0: len(page) - 5]
	}

	// the file path to the jekyll markdown file
	//md_file_path := CLONES_DIR + "/revel.github.io/manual/" + page_no_ext + ".md"

	// read the file
	// TODO check  it exists or catch error
	//raw_md_bytes, err := ioutil.ReadFile(md_file_path)
	//if err != nil {
	//	fmt.Println("errrr", err)
	//}

	// convert markdown to html
	//page_html_bytes := blackfriday.MarkdownCommon(raw_md_bytes)
	//fmt.Println("html", string(page_html_bytes))

	// use template.HTML to "unespae" encoding.. ie proper html not &lt;escaped
	pdata := ReadMarkdownPage(site_section, page_no_ext)
	c.RenderArgs["page_content"] = pdata.HTML
	//c.RenderArgs["md"] = md


	cPage.PageTitle = pdata.Title
	c.RenderArgs["cPage"] = cPage

	return c.Render()
}

func (c App) GithubPage() revel.Result {
	return c.Render()
}


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
