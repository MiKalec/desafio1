package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println(err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	responseBodyBytes, err := io.ReadAll(res.Body)
	cotacao := string(responseBodyBytes)
	clean := strings.ReplaceAll(cotacao, "\"", "")
	clean = strings.ReplaceAll(clean, "\n", "")
	clean = strings.ReplaceAll(clean, "\r", "")

	if cotacao != "" {
		toFile := fmt.Sprintf("Dólar: {%s}", clean)
		_, err = file.WriteString(toFile)
		if err != nil {
			log.Println(err)
			return
		}
	}

}
