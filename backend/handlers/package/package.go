package pkg

import (
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	resty "github.com/go-resty/resty/v2"
	"github.com/google/go-github/v28/github"
)

func Routes(r *gin.RouterGroup) {
	r.GET("package/issues/:name", getPackageIssues)
}

func getPackageIssues(c *gin.Context) {
	client := resty.New()
	response, _ := client.R().
		SetResult(&NpmPackageResponse{}).
		Get("https://api.npms.io/v2/package/" + c.Param("name"))

	npmResponse := response.Result().(*NpmPackageResponse)

	err := npmResponse.Validate()

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	repoOwner, err := npmResponse.GetRepoOwner()
	repoName, err := npmResponse.GetRepoName()

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	githubClient := github.NewClient(nil) //add token
	issues, _, err := githubClient.Issues.ListByRepo(context.Background(), repoOwner, repoName, nil)

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, issues)
}

func (r *NpmPackageResponse) GetRepoOwner() (string, error) {
	splittedLink := strings.Split(*r.Collected.Metadata.Links.Homepage, "/")

	if len(splittedLink) < 2 {
		return "", errors.New("Failed to get repo owner")
	}

	return splittedLink[3], nil
}

func (r *NpmPackageResponse) GetRepoName() (string, error) {
	splittedLink := strings.Split(*r.Collected.Metadata.Links.Homepage, "/")

	if len(splittedLink) < 3 {
		return "", errors.New("Failed to get repo name")
	}

	return splittedLink[4], nil
}

func (r *NpmPackageResponse) Validate() error {
	if r.Collected == nil || r.Collected.Metadata == nil || r.Collected.Metadata.Links == nil || r.Collected.Metadata.Links.Homepage == nil {
		return errors.New("Unknown package")
	}

	if !strings.Contains(*r.Collected.Metadata.Links.Homepage, "github") {
		return errors.New("Not a github repo")
	}

	return nil
}

type NpmPackageResponse struct {
	Collected *NpmPackageCollectedResponse `json:"collected"`
}

type NpmPackageCollectedResponse struct {
	Metadata *NpmPackageMetadataResponse `json:"metadata,omitempty"`
}

type NpmPackageMetadataResponse struct {
	Links *NpmPackageLinksResponse `json:"links,omitempty"`
}

type NpmPackageLinksResponse struct {
	Homepage *string `json:"homepage,omitempty"`
}
