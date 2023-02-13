package split

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/HiDoYa/rent-cat/app/models"
	"github.com/gin-gonic/gin"
)

// Controller holds implementations of API endpoints
type Controller struct {}

// GetSplit gets expense split percentage
// @Summary gets expense split percentage
// @Schemes
// @Description gets expense split percentage
// @Tags split
// @Accept json
// @Produce json
// @Success 200 {object} models.Split
// @Router /split/:year/:month [get]
func (cont Controller) GetSplit(c *gin.Context) {
	yearStr := c.Param("year")
	monthStr := c.Param("month")
	
	strconv.Atoi(monthStr)
	strconv.Atoi(yearStr)

	// TODO Get from database

	split := models.Split{
		MyPercentage: 0.88,
		HerPercentage: 0.12,
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

	var myPercentage float32
	var herPercentage float32

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

	// TODO Store in database

	c.JSON(http.StatusOK, gin.H{
		"split": split,
	})
}