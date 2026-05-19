package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", homeHandler)

	fmt.Printf("Kitchen Display System showcase is running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = fmt.Fprint(w, `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Kitchen Display System Showcase</title>
</head>
<body>
  <h1>Kitchen Display System (KDS)</h1>
  <p>An overview of our Kitchen Display System app, including demos, key features, and bug tracking by label and priority.</p>

  <section>
    <h2>Overview</h2>
    <p>KDS helps kitchen teams manage incoming orders, track preparation status, and improve service speed and order accuracy.</p>
  </section>

  <section>
    <h2>Demos</h2>
    <ul>
      <li>Live kitchen queue board demo</li>
      <li>Order bump and recall workflow demo</li>
      <li>Kitchen performance metrics dashboard demo</li>
    </ul>
  </section>

  <section>
    <h2>Key Features</h2>
    <ul>
      <li>Real-time order updates</li>
      <li>Station-based filtering</li>
      <li>Priority highlighting for urgent orders</li>
      <li>Service-time tracking and analytics</li>
    </ul>
  </section>

  <section>
    <h2>Bugs by Label and Priority</h2>
    <table border="1" cellpadding="6" cellspacing="0">
      <thead>
        <tr>
          <th>Label</th>
          <th>Priority</th>
          <th>Open Bugs</th>
        </tr>
      </thead>
      <tbody>
        <tr><td>UI</td><td>P1</td><td>4</td></tr>
        <tr><td>Backend</td><td>P0</td><td>2</td></tr>
        <tr><td>Integrations</td><td>P2</td><td>3</td></tr>
        <tr><td>Performance</td><td>P1</td><td>1</td></tr>
      </tbody>
    </table>
  </section>
</body>
</html>`)
}
