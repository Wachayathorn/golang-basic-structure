package client

import (
	"encoding/json"
	"net/http"

	bookhttpinfo "github.com/wachayathorn/golang-basic-structure/pkg/handler/book/httpinfo"
	"github.com/wachayathorn/golang-basic-structure/pkg/model"
	bookrepository "github.com/wachayathorn/golang-basic-structure/pkg/repository/book"
)

func (c *Client) AddBook(book model.Book) (bookrepository.Book, error) {
	var result bookrepository.Book

	apiPath := bookhttpinfo.GroupV1 + bookhttpinfo.AddBook
	res, err := c.Utils.DoRequest(c.Url, http.MethodPost, apiPath, nil)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, nil
}
