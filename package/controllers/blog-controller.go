package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Threx-code/go-blog/package/models"
	"github.com/Threx-code/go-blog/package/utils"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	allBlogs := models.Index()
	response, _ := json.Marshal(allBlogs)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogId := vars["id"]
	ID, err := strconv.ParseInt(blogId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	singleBlog, _ := models.Read(ID)
	response, _ := json.Marshal(singleBlog)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogId := vars["id"]
	ID, err := strconv.ParseInt(blogId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	singleBlog := models.Destroy(ID)
	response, _ := json.Marshal(singleBlog)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write(response)
}

func Store(w http.ResponseWriter, r *http.Request) {
	CreateBlog := &models.Blogs{}
	utils.ParseBody(r, CreateBlog)
	blog := CreateBlog.Store()
	response, _ := json.Marshal(blog)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func Update(w http.ResponseWriter, r *http.Request) {
	updateBlog := &models.Blogs{}
	utils.ParseBody(r, updateBlog)
	vars := mux.Vars(r)
	blogId := vars["id"]
	ID, err := strconv.ParseInt(blogId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	singleBlog, db := models.Read(ID)

	if updateBlog.Title != "" {
		singleBlog.Title = updateBlog.Title
	}
	if updateBlog.Author != "" {
		singleBlog.Author = updateBlog.Author
	}
	if updateBlog.Content != "" {
		singleBlog.Content = updateBlog.Content
	}

	db.Save(singleBlog)

	response, _ := json.Marshal(singleBlog)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
