// microservice.go
package main

import (
	"html/template"
	"net/http"

	pb "grpc/proto" // Import the generated code

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var tp1 *template.Template

func main() {
	http.HandleFunc("/sayhello", func(w http.ResponseWriter, r *http.Request) {
		// Handle the HTTP request to call the gRPC server
		conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			http.Error(w, "Failed to connect server", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		c := pb.NewMyServiceClient(conn)
		name := r.URL.Query().Get("name")
		response, err := c.SayHello(r.Context(), &pb.HelloRequest{Name: name})
		if err != nil {
			// log.Fatalf("Failed to connect to gRPC server: %v", err)
			http.Error(w, "Failed to call gRPC server", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(response.Message))

		// Render the HTML page
		tmpl, err := template.New("hello").Parse(response.Message)
		if err != nil {
			http.Error(w, "Failed to parse template", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Failed to render HTML", http.StatusInternalServerError)
		}
	})

	h2 := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	}

	// Serve the HTML GUI
	http.HandleFunc("/", h2)

	http.ListenAndServe(":8080", nil)

}
