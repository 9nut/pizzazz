package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logFile := "testlogfile"
	port := "3001"
	if os.Getenv("HTTP_PLATFORM_PORT") != "" {
		logFile = "D:\\home\\site\\wwwroot\\testlogfile"
		port = os.Getenv("HTTP_PLATFORM_PORT")
	}

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
        <html>
            <body>
                <h1>Hello!</h1>
                <br />
		<p>You are authentic!</p>
                <pre>`)

		rf, _ := os.Open(logFile)
		defer rf.Close()
		scanner := bufio.NewScanner(rf)
		lineCount := 0
		for scanner.Scan() {
			lineStr := scanner.Text()
			fmt.Fprintf(w, lineStr)
			fmt.Fprintf(w, "<br />")
			lineCount++
		}

		fmt.Fprintf(w, "<br />")
		fmt.Fprintf(w, "Log Count: %v/1000", lineCount)
		fmt.Fprintf(w, "<br />")
		fmt.Fprintf(w, `
                </pre>
            </body>
        </html>`)

		if lineCount > 1000 {
			wf, _ := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
			defer wf.Close()
			w := bufio.NewWriter(wf)
			w.WriteString("")
			w.Flush()
		}
	})

	if err == nil {
		defer f.Close()
		log.SetOutput(f)
		log.Println("--->   UP @ " + port + "  <------")
	}

	http.ListenAndServe(":"+port, nil)
}

