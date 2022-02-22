package integration_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/m-rcd/go-rest-api/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
		var response models.JsonBookResponse
		json.Unmarshal(body, &response)

		Expect(response.Type).To(Equal("success"))
		Expect(response.StatusCode).To(Equal(200))
		Expect(response.Data[0].Title).To(Equal("hello"))
		Expect(response.Data[0].Author).To(Equal("writer"))
		Expect(response.Data[0].Content).To(Equal("iexist"))
	})

	It("gets a book", func() {
		c := http.Client{}
		resp, err := c.Get("http://localhost:10000/book/48")
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		var response models.Book
		json.Unmarshal(body, &response)

		Expect(response.ID).To(Equal("48"))
		Expect(response.Title).To(Equal("hello"))
		Expect(response.Author).To(Equal("writer"))
		Expect(response.Content).To(Equal("iexist"))
	})

	It("gets all book", func() {
		c := http.Client{}
		resp, err := c.Get("http://localhost:10000/books/")
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		var response []models.Book
		json.Unmarshal(body, &response)

		Expect(len(response)).NotTo(Equal(0))

	})

	It("updates book", func() {
		c := http.Client{}
		patchData := bytes.NewBuffer([]byte(`{"title":"uu", "author": "jj", "content":"iexddist"}`))
		req, err := http.NewRequest("PATCH", "http://localhost:10000/book/49", patchData)
		Expect(err).NotTo(HaveOccurred())
		resp, _ := c.Do(req)
		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		defer req.Body.Close()
		var response models.JsonBookResponse
		json.Unmarshal(body, &response)
		Expect(response.Type).To(Equal("success"))
		Expect(response.StatusCode).To(Equal(200))
		Expect(response.Data[0].Title).To(Equal("uu"))
		Expect(response.Data[0].Author).To(Equal("jj"))
		Expect(response.Data[0].Content).To(Equal("iexddist"))
	})

	It("deletes a book", func() {
		c := http.Client{}
		req, err := http.NewRequest("DELETE", "http://localhost:10000/book/44", bytes.NewBuffer([]byte{}))
		Expect(err).NotTo(HaveOccurred())
		resp, _ := c.Do(req)
		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()
		Expect(err).NotTo(HaveOccurred())
		var response models.JsonBookResponse
		json.Unmarshal(body, &response)
		Expect(response.Type).To(Equal("success"))
		Expect(response.StatusCode).To(Equal(200))
		Expect(response.Message).To(Equal("The book was successfully deleted"))

	})
})
