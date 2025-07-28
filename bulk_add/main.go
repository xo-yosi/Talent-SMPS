package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Student struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
}

type ResponseData struct {
	StudentID int    `json:"student_id"`
	Name      string `json:"name"`
}

func main() {
	// Open CSV file
	file, err := os.Open("students.csv")
	if err != nil {
		fmt.Println("❌ Error opening CSV:", err)
		return
	}
	defer file.Close()

	// Read all records
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("❌ Error reading CSV:", err)
		return
	}

	if len(records) <= 1 {
		fmt.Println("⚠️ CSV has no data rows (only header?)")
		return
	}

	// Output file to store successful registrations
	outFile, err := os.Create("registered_students.txt")
	if err != nil {
		fmt.Println("❌ Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// API endpoint and auth token
	url := "https://talent-smps.onrender.com/api/v1/student-register"
	authToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTM3MzMwODEsImlhdCI6MTc1MzY0NjY4MSwicm9sZSI6ImFkbWluIiwidXNlcl9pZCI6Ijc0YWZkNWJjLTRmYTItNDg3MS04MzRhLTdhYjliOTM0Nzg3OSJ9.lF41lYfTfd7pX3D-mko3BV-U0pxaGhjMUP4qQjdOP1M" // 🔐 Replace this

	client := &http.Client{}

	// Loop through CSV records
	for i, record := range records[1:] {
		if len(record) < 4 {
			fmt.Printf("⚠️ Row %d skipped: not enough fields → %+v\n", i+2, record)
			continue
		}

		age, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Printf("⚠️ Row %d skipped: invalid age: %v\n", i+2, err)
			continue
		}

		student := Student{
			Name:        record[0],
			Age:         age,
			Gender:      record[2],
			PhoneNumber: record[3],
		}

		// Convert to JSON
		body, err := json.Marshal(student)
		if err != nil {
			fmt.Printf("❌ Row %d: Failed to marshal JSON: %v\n", i+2, err)
			continue
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
		if err != nil {
			fmt.Printf("❌ Row %d: Failed to create request: %v\n", i+2, err)
			continue
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", authToken)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("❌ Row %d: Request failed: %v\n", i+2, err)
			continue
		}

		var res ResponseData
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			fmt.Printf("❌ Row %d: Failed to decode response: %v\n", i+2, err)
			resp.Body.Close()
			continue
		}
		resp.Body.Close()

		// Log successful registration
		// Log successful registration
		line := fmt.Sprintf("%s - %d\n", student.Name, res.StudentID)
		if _, err := outFile.WriteString(line); err != nil {
			fmt.Printf("❌ Row %d: Failed to write to file: %v\n", i+2, err)
		} else {
			fmt.Printf("✅ Registered %s → ID %d\n", student.Name, res.StudentID)
		}

		// 2-second delay between requests
		time.Sleep(2 * time.Second)
	}
}
