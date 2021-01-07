package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"school-supply-list/database"
	"strconv"
)

type supplyItem struct {
	Id       int    `json:"item_id"`
	Supply   string `json:"item_name"`
	Desc     string `json:"item_desc"`
	Category string `json:"item_category"`
}

func createSupply(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var supply supplyItem
		err := json.NewDecoder(c.Request.Body).Decode(&supply)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}
		row := db.Db.QueryRow(`INSERT INTO supply_item (supply_name, supply_desc) 
		  VALUES ($1, $2) RETURNING id`, supply.Supply, supply.Desc)
		err = row.Scan(&supply.Id)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, supply)
	}
}

func getSupply(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		supply := supplyItem{
			Id: id,
		}

		row := db.Db.QueryRow(`SELECT supply_name, supply_desc FROM supply_item 
											WHERE supply_item.id=$1`, id)
		err = row.Scan(&supply.Supply, &supply.Desc)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, supply)
	}
}

func getAllSupplies(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var supplies []supplyItem
		rows, err := db.Db.Query(`SELECT id, supply_name, supply_desc FROM supply_item`)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}
		for rows.Next() {
			var supply supplyItem
			err = rows.Scan(&supply.Supply, &supply.Desc, &supply.Desc)
			supplies = append(supplies, supply)
		}
		c.JSON(200, supplies)
	}
}

func updateSupply(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		var supply supplyItem
		err = json.NewDecoder(c.Request.Body).Decode(&supply)

		supply.Id = id

		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}
		row := db.Db.QueryRow(`UPDATE supply_item SET supply_name=$1, supply_desc=$2
	   		where id=$3 returning id, supply_name, supply_desc`, supply.Supply, supply.Desc)
		//Scan the latest changes into the supply struct
		err = row.Scan(&supply.Id, &supply.Supply, supply.Desc)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, supply)
	}
}

func deleteSupply(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		row := db.Db.QueryRow(`DELETE FROM supply_item where id=$1`, id)
		if row.Err() != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, nil)
	}
}
