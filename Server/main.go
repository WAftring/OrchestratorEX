package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
)

var page = `
<html>
<body>
	<ul>
		<li><a href='/'>Home</a></li>
		<li><a href='/api'>API</a></li>
	</ul>
	<pre>
	{{.Content}}
	</pre>
</body>
</html>`

type Data struct {
	Content string
}

func main() {

	log.Println("Server starting")
	var content = "Home Page"
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		log.Println("Request for /")
		RenderPage(w, content)
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		log.Println("Request for /api")
		s := API("bing.com")
		s += API("api.local")
		RenderPage(w, s)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/api", h2)
	log.Println("Listening on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func RenderPage(w http.ResponseWriter, content string) {

	data := &Data{Content: content}
	tmpl := template.New("page")
	tmpl, err := tmpl.Parse(page)
	if err != nil {
		log.Println("Failed to parse template with ", err)
		io.WriteString(w, "Failed to parse template")
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Failed to execute template with ", err)
		io.WriteString(w, "Failed to execute template")
	}

}

func API(endpoint string) (res string) {
	var logstr string
	log.Println("Looking up ", endpoint)
	resp, err := net.LookupHost(endpoint)
	if err != nil {
		logstr = fmt.Sprintf("\nDNS Lookup for %s: FAILED\n\tERR: %s", endpoint, err)
		return logstr
	} else {
		dns_string := ""
		for index, value := range resp {
			if index == 0 {
				dns_string = fmt.Sprintf("[%d] %s\n", index, value)
			} else {
				dns_string = fmt.Sprintf("%s\t[%d] %s\n", dns_string, index, value)
			}
		}
		logstr = fmt.Sprintf("DNS Lookup for %s: SUCCESS\n\t%s", endpoint, dns_string)
	}

	log.Println(logstr)
	// NOTE: We have some weird rendering in the html
	// This extra linebreak fixes it - wiaftrin
	body := fmt.Sprintf("\n%s", logstr)
	log.Println("Testing TCP Connectivity to ", endpoint)
	var dialstr = fmt.Sprintf("%s:80", endpoint)
	socket, err := net.Dial("tcp4", dialstr)
	defer socket.Close()
	if err != nil {
		logstr = fmt.Sprintf("Dial %s:80: FAILED \n\tERR: %s", endpoint, err)
	} else {
		logstr = fmt.Sprintf("Dial %s:80: SUCCESS\n\tLocal Addr:\t%s\n\tRemote Addr:\t%s\n",
			endpoint,
			socket.LocalAddr().String(),
			socket.RemoteAddr().String())
	}

	log.Println(logstr)
	body = fmt.Sprintf("%s\n%s",
		body,
		logstr)
	return body

}
