package supplies

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"school-supply-list/database"
	"strconv"
)

type SupplyItem struct {
	ID       int            `json:"id"`
	ListID   int            `json:"list_id"`
	Supply   string         `json:"supply"`
	Desc     string         `json:"desc"`
	Category sql.NullString `json:"item_category"`
}

func CreateSupply(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var supply SupplyItem
		err := json.NewDecoder(c.Request.Body).Decode(&supply)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}
		row := db.Db.QueryRow(`INSERT INTO supply_item (list_id, supply_name, supply_desc) 
		  VALUES ($1, $2, $3) RETURNING id`, supply.ListID, supply.Supply, supply.Desc)
		err = row.Scan(&supply.ID)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, supply)
	}
}

func GetSupply(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		supply := SupplyItem{
			ID: -1,
		}
		row := db.Db.QueryRow(`SELECT id, list_id, supply_name, supply_desc FROM supply_item 
											WHERE supply_item.id=$1`, id)
		err = row.Scan(&supply.ID, &supply.ListID, &supply.Supply, &supply.Desc)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}
		if supply.ID == -1 {
			c.AbortWithStatusJSON(404, "Supply does not exist")
		}
		c.JSON(200, supply)
	}
}

func GetAllSupplies(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var supplies []SupplyItem
		rows, err := db.Db.Query(`SELECT id, list_id, supply_name, supply_desc FROM supply_item`)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}
		for rows.Next() {
			var supply SupplyItem
			err = rows.Scan(&supply.ID, &supply.ListID, &supply.Supply, &supply.Desc)
			supplies = append(supplies, supply)
		}
		c.JSON(200, supplies)
	}
}

func UpdateSupply(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid id. Must be an integer.")
			return
		}
		var supply SupplyItem
		err = json.NewDecoder(c.Request.Body).Decode(&supply)

		supply.ID = id

		if err != nil {
			c.AbortWithStatusJSON(400, "Invalid request.")
			return
		}
		row := db.Db.QueryRow(`UPDATE supply_item SET supply_name=$1, supply_desc=$2
	   		where id=$3 returning id, supply_name, supply_desc`, supply.Supply, supply.Desc)
		//Scan the latest changes into the supply struct
		err = row.Scan(&supply.ID, &supply.Supply, supply.Desc)
		if err != nil {
			database.CheckDBErr(err.(*pq.Error), c)
			return
		}

		c.JSON(200, supply)
	}
}

func DeleteSupply(db *database.DB) gin.HandlerFunc {
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
