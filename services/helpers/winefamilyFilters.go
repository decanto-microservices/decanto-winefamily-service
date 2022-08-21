package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func BuildFiltersFromContext(c *gin.Context) (filters *bson.M, err error) {
	winetypesString := c.Query("winetypeIds")
	winetypeIdsString := strings.Split(winetypesString, ",")

	var winetypeIds []int64
	for _, element := range winetypeIdsString {
		id, err := strconv.ParseInt(element, 10, 64)

		if err != nil {
			break
		}

		winetypeIds = append(winetypeIds, id)
	}

	winecolorsString := c.Query("winecolorIds")
	winecolorIdsString := strings.Split(winecolorsString, ",")

	var winecolorIds []int64
	for _, element := range winecolorIdsString {
		id, err := strconv.ParseInt(element, 10, 64)

		if err != nil {
			return nil, err
		}

		winecolorIds = append(winecolorIds, id)
	}

	structure, err := strconv.ParseFloat(c.Query("structure"), 64)
	structureDelta, err := strconv.ParseFloat(c.Query("structureD"), 64)

	softness, err := strconv.ParseFloat(c.Query("softness"), 64)
	softnessDelta, err := strconv.ParseFloat(c.Query("softnessD"), 64)

	hardness, err := strconv.ParseFloat(c.Query("hardness"), 64)
	hardnessDelta, err := strconv.ParseFloat(c.Query("hardnessD"), 64)

	sweetness, err := strconv.ParseFloat(c.Query("sweetness"), 64)
	sweetnessDelta, err := strconv.ParseFloat(c.Query("sweetnessD"), 64)

	foodSx, err := strconv.ParseFloat(c.Query("foodSx"), 64)
	foodSxDelta, err := strconv.ParseFloat(c.Query("foodSxD"), 64)

	foodDx, err := strconv.ParseFloat(c.Query("foodDx"), 64)
	foodDxDelta, err := strconv.ParseFloat(c.Query("foodDxD"), 64)

	if err != nil {
		return nil, err
	}

	filters = &bson.M{
		"winetypeId":         bson.M{"$in": winetypeIds},
		"winecolorId":        bson.M{"$in": winecolorIds},
		"_wfBody":            bson.M{"$gte": structure - structureDelta, "$lte": structure + structureDelta},
		"_wfHard":            bson.M{"$gte": softness - softnessDelta, "$lte": softness + softnessDelta},
		"_wfSoft":            bson.M{"$gte": hardness - hardnessDelta, "$lte": hardness + hardnessDelta},
		"winefamilyDolcezza": bson.M{"$gte": sweetness - sweetnessDelta, "$lte": sweetness + sweetnessDelta},
		"_wfCurveSx":         bson.M{"$gte": foodSx - foodSxDelta, "$lte": foodSx + foodSxDelta},
		"_wfCurveDx":         bson.M{"$gte": foodDx - foodDxDelta, "$lte": foodDx + foodDxDelta},
	}

	fmt.Println(filters)

	return
}
