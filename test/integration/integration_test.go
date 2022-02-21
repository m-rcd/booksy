package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Book `json:"data"`
	Message string `json:"message"`
}

var _ = Describe("Integration", func() {
	It("homepage", func() {
		c := http.Client{}
		resp, err := c.Get("http://localhost:10000/")
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Eventually(body).Should(ContainSubstring("Welcome to the HomePage!"))
	})

	It("creates new book", func() {
		c := http.Client{}
		postData := bytes.NewBuffer([]byte(`{"title":"hello", "author": "writer", "content":"iexist"}`))
		resp, err := c.Post("http://localhost:10000/book", "application/json", postData)
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		var response JsonResponse
		json.Unmarshal(body, &response)
		fmt.Println(response)

		Expect(response.Message).To(ContainSubstring("success"))
	})
})
