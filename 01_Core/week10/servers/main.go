package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)

}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		//Scan each line sent in the header
		ln := scanner.Text()
		fmt.Println(ln)
		//If we are on the first header, send to mux to route
		if i == 0 {
			mux(conn, ln)
		}

		if ln == "" {
			//If we get to blank line, the headers have completed
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	method := strings.Fields(ln)[0]
	url := strings.Fields(ln)[1]
	fmt.Println("Method is " + method + ". URL is " + url)

	if url == "/" {
		index(conn)
	} else if url == "/about" {
		about(conn)
	} else if url == "/products" {
		products(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	  <head>
		<meta charset="utf-8">
		<title>Untitled</title>
		<meta name="description" content="This is an example of a meta description.">
	  </head>
	  <body>
	  <a href="/">Go home </a><br>
	  <a href="/about">Find out more about us</a><br>
	  <a href="/products">Search our products</a>
	  </body>`

	//Add the necessary headers
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func about(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	  <head>
		<meta charset="utf-8">
		<title>Untitled</title>
		<meta name="description" content="This is an example of a meta description.">
	  </head>
	  <body>
	  <p>Welcome to the about page!</p>
	  </body>`

	//Add the necessary headers
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func products(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	  <head>
		<meta charset="utf-8">
		<title>Untitled</title>
		<meta name="description" content="This is an example of a meta description.">
	  </head>
	  <body>
	  <p>Welcome to the products page!</p>
	  </body>`

	//Add the necessary headers
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
