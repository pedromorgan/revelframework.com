package controllers

import(
	"fmt"
	//"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

/*
These struct import the yaml jekyll format
There weare atht the _layouts/foo.html header but extracted
to /foo/_nav.yaml
 */

type NavPage struct {
	Label string ` yaml:"title" `
	Url string ` yaml:"url" `
}

type NavGroup struct {
	GroupTitle string ` yaml:"name" `
	Pages []NavPage ` yaml:"articles" `
}

type NavSection struct {
	Root string ` yaml:"root" `
	Name string ` yaml:"name" `
	SectionTitle string ` yaml:"section_title" `
	SubGroups[] NavGroup ` yaml:"nav" `
}



// This a hack to read yaml at top, of _layout/jekyll and return nav.yaml
// later we expect this to be a yaml in the directory.
// eg https://github.com/revel/revel.github.io/blob/master/_layouts/manual.html
func GetNav(section string) NavSection {

	nav_file := CLONES_DIR + "/revel.github.io/" + section + "/_nav.yaml"

	raw_bytes, err := ioutil.ReadFile(nav_file)
	if err != nil {
		fmt.Println("error, raw bytes", err)
	}

	var nav NavSection
	err = yaml.Unmarshal([]byte(raw_bytes), &nav)
	if err != nil {
		fmt.Println("error, yaml", err)
	}
	// scan lines until we hit <doctype
	// must be a better way.. am golang newview
	/*
	ret := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "<!DOCTYPE html>" {
			fmt.Println(ret)
			return ret
		}else {
			ret = ret + line + "\n"
		}

	}

	if err := scanner.Err(); err != nil {
		//log.Fatal(err)
	}
	*/

	//fmt.Println("nav=", nav)
	return nav
}
