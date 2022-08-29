package service

import (
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bibishkind/shtrafovnet-test-task/models"
)

const searchURL = "https://www.rusprofile.ru/search"

func (srvc *service) GetCompanyByINN(inn string) (*models.Company, error) {
	resp, err := srvc.httpClient.Get(fmt.Sprintf("%s?query=%s", searchURL, inn))
	if err != nil {
		return nil, err
	}

	body := resp.Body
	defer resp.Body.Close()

	return parseHTMLOfCompany(body)
}

func parseHTMLOfCompany(body io.Reader) (*models.Company, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	var company models.Company
	company.Owner = &models.Owner{}

	doc.Find(".company-row").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			company.INN = s.Find("#clip_inn").Text()
			company.KPP = s.Find("#clip_kpp").Text()
		}
		if i == 4 {
			text := s.Find(".company-info__text").Text()
			data := strings.Split(text, " ")
			owner := &models.Owner{
				Firstname:  data[1],
				Middlename: data[2],
				Lastname:   data[0],
			}
			company.Owner = owner
		}
	})

	return &company, nil
}
