package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
)

func main() {
	fmt.Println("hellow buda")
	go func() {
		url := "http://localhost:5001/api/user"
		callPost(url)
		callGet(url)
		callPut(url)
		callGet(url)
		callDelete(url)
		callGet(url)
	}()
	httpCurd()
}

func callGet(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("failed to fetch the request")
	}

	respBody, _ := io.ReadAll(resp.Body)

	fmt.Println("from GET request", string(respBody))
}

func callPost(url string) {
	reqBody, _ := json.Marshal(map[string]any{
		"username": "test",
		"name":     "actual test",
		"id":       123455,
	})

	buffBody := bytes.NewBuffer(reqBody)
	resp, err := http.Post(url, "application/json", buffBody)

	if err != nil {
		fmt.Println("failed to fetch the request")
	}

	respBody, _ := io.ReadAll(resp.Body)

	fmt.Println("from POST request", string(respBody))
}
func callPut(url string) {
	reqBody, _ := json.Marshal(map[string]any{
		"username": "test",
		"name":     "actually a new test",
		"id":       123455,
	})

	buffBody := bytes.NewBuffer(reqBody)
	newUrl := fmt.Sprintf("%s?username=test", url)
	req, err := http.NewRequest(http.MethodPut, newUrl, buffBody)

	if err != nil {
		fmt.Println("failed to do the request")
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("failed to fetch the request")
	}

	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	fmt.Println("from PUT request", string(respBody))
}

func callDelete(url string) {
	newUrl := fmt.Sprintf("%s?username=test", url)
	req, err := http.NewRequest(http.MethodDelete, newUrl, nil)

	if err != nil {
		fmt.Println("failed to do the request")
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("failed to fetch the request")
	}

	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	fmt.Println("from DELETE request", string(respBody))
}

func httpCurd() {
	type User struct {
		Username       string `json:"username"`
		Name           string `json:"name"`
		IdentityNumber int    `json:"id"`
	}
	users := []*User{}

	errRes := map[string]any{
		"error": "something went wrong",
	}

	http.HandleFunc("GET /api/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonUsers := map[string]any{
			"users": users,
		}
		data, err := json.Marshal(jsonUsers)
		if err != nil {
			data, _ := json.Marshal(errRes)
			w.Write(data)
		}
		w.Write(data)

	})
	http.HandleFunc("POST /api/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			data, _ := json.Marshal(errRes)
			w.Write(data)
		}
		typedBody := User{}
		err = json.Unmarshal(body, &typedBody)
		if err != nil {
			data, _ := json.Marshal(errRes)
			w.Write(data)
			return
		}
		users = append(users, &typedBody)
		w.Write(body)

	})
	http.HandleFunc("PUT /api/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		query := r.URL.Query()
		username := query.Get("username")
		if username == "" {
			errRes := map[string]any{
				"error": "something went wrong",
			}
			data, _ := json.Marshal(errRes)
			w.Write(data)
			return

		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			data, _ := json.Marshal(errRes)
			w.Write(data)
			return
		}

		index := slices.IndexFunc(users, func(ele *User) bool {
			return (*ele).Username == username
		})

		if index != -1 {
			users = slices.Delete(users, index, 1)
			typedBody := User{}
			err := json.Unmarshal(body, &typedBody)
			if err != nil {
				data, _ := json.Marshal(errRes)
				w.Write(data)
			}
			users = slices.Insert(users, index, &typedBody)
		}
		w.Write(body)

	})
	http.HandleFunc("DELETE /api/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		query := r.URL.Query()
		username := query.Get("username")
		if username == "" {
			errRes := map[string]any{
				"error": "something went wrong",
			}
			data, _ := json.Marshal(errRes)
			w.Write(data)
			return

		}

		index := slices.IndexFunc(users, func(ele *User) bool {
			return (*ele).Username == username
		})

		if index == -1 {
			errRes := map[string]any{
				"error": "User not found",
			}
			data, _ := json.Marshal(errRes)
			w.Write(data)
			return

		}
		user := users[index]
		users = slices.Delete(users, index, 1)
		jsonData, _ := json.Marshal(map[string]any{
			"message":        "success",
			"deleted_record": user,
		})
		w.Write(jsonData)

	})

	http.ListenAndServe(":5001", nil)
}
