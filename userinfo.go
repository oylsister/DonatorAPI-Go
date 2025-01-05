package main

import (
    "database/sql"
    "net/http"
	"github.com/gin-gonic/gin"
)

// Json can be rename to something else
type UserInfo struct {
    UserID     		int       `json:"id"`
    UserAuth   		string    `json:"auth"`
    DonateTier   	string    `json:"donateTier"`
    ExpireTime 		sql.NullString    `json:"expireTime"`
}

func getUserInfo(c * gin.Context, db *sql.DB) {
	steamAuth := c.Param("steamAuth")

	// if steamAuth is not empty or null
	if(steamAuth != "") {
		var user UserInfo
		err := db.QueryRow("SELECT user_id, user_auth, donate_tier, expire_time FROM userinfo WHERE user_auth = ?", steamAuth).Scan(&user.UserID, &user.UserAuth, &user.DonateTier, &user.ExpireTime)
		
		if(err != nil) {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, user)
	} else {
		rows, err := db.Query("SELECT user_id, user_auth, donate_tier, expire_time FROM userinfo")

		// nil mean success I guess
		if (err != nil) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) 
				return 
		} 
		defer rows.Close()

		var users []UserInfo
		for rows.Next() {
			var user UserInfo
			if err := rows.Scan(&user.UserID, &user.UserAuth, &user.DonateTier, &user.ExpireTime); err != nil { 
				c.JSON(http.StatusInternalServerError, gin.H {"error": err.Error()}) 
				return 
			}

			users = append(users, user)
		}
		c.JSON(http.StatusOK, users)
	}
}