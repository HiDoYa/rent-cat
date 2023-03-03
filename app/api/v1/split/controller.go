package split

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/HiDoYa/rent-cat/app/models"
	"github.com/HiDoYa/rent-cat/app/svc/database"
	"github.com/gin-gonic/gin"
)

// Controller holds implementations of API endpoints
type Controller struct {}

// GetLatestSplit gets latest expense split percentage
// @Summary gets latest expense split percentage
// @Schemes
// @Description gets latest expense split percentage
// @Tags split
// @Accept json
// @Produce json
// @Success 200 {object} models.Split
// @Router /split/latest [get]
func (cont Controller) GetLatestSplit(c *gin.Context) {
	client, err := database.MakeDb()
	if err != nil {
	}

	split, err := client.SelectLatestSplit()
	if err != nil {
	}

	c.JSON(http.StatusOK, gin.H{
		"split": split,
	})
}


// PostSplit adds expense split percentage
// @Summary adds expense split percentage
// @Schemes
// @Description adds expense split percentage
// @Param Split body models.SplitSpecifier true "Split Descriptor"
// @Tags split
// @Accept json
// @Produce json
// @Success 200 {object} models.Split
// @Router /split [post]
func (cont Controller) PostSplit(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
	}

	var body models.SplitSpecifier
	err = json.Unmarshal(jsonData, &body)
	if err != nil {
	}

	var myPercentage, herPercentage float32

	if body.MyNetIncome != 0 && body.HerNetIncome != 0 {

	} else if body.MyPercentage != 0 && body.HerPercentage != 0 {
		myPercentage = body.MyPercentage
		herPercentage = body.HerPercentage
	} else {
		// No valid data passed
		return
	}

	currentTime := time.Now()
	currentTime.Year()
	currentTime.Month()

	split := models.Split{
		MyPercentage: myPercentage,
		HerPercentage: herPercentage,
	}

	client, err := database.MakeDb()
	if err != nil {
	}

	_, err = client.InsertSplit(split)
	if err != nil {
	}

	c.JSON(http.StatusOK, gin.H{
		"split": split,
	})
}