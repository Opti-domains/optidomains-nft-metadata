package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

var TOWN_DOMAINS_PREFIX string
var TOWN_DOMAINS_SUFFIX string

func main() {
	r := gin.Default()

	TOWN_DOMAINS_PREFIX_BYTES, err := os.ReadFile("./images/domains/town/TEMPLATE_PREFIX")

	if err != nil {
		panic(err)
	}

	TOWN_DOMAINS_SUFFIX_BYTES, err := os.ReadFile("./images/domains/town/TEMPLATE_SUFFIX")

	if err != nil {
		panic(err)
	}

	TOWN_DOMAINS_PREFIX = string(TOWN_DOMAINS_PREFIX_BYTES)
	TOWN_DOMAINS_SUFFIX = string(TOWN_DOMAINS_SUFFIX_BYTES)

	r.GET("/collection/domains/town", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":          "Bored Town Domains",
			"description":   ".town Domains by Opti.Domains X Bored Town for Bored Town holders",
			"image":         "https://metadata.opti.domains/images/domains/town.png",
			"external_link": "https://town.opti.domains",
		})
	})

	r.GET("/collection/domains/op", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":          "Opti.Domains",
			"description":   ".op Domains by Opti.Domains for Optimism and OP stack ecosystems",
			"image":         "https://metadata.opti.domains/images/domains/op.png",
			"external_link": "https://opti.domains",
		})
	})

	r.GET("/token/domains/:id", func(c *gin.Context) {
		id := c.Param("id")

		if strings.HasPrefix(id, "town") {
			name, err := getDomainNameFromId(strings.TrimPrefix(id, "town"))

			if err != nil {
				if err.Error() == "not found" {
					c.AbortWithStatus(http.StatusNotFound)
				} else {
					c.AbortWithStatus(http.StatusBadRequest)
				}
			}

			c.JSON(http.StatusOK, gin.H{
				"name":         name,
				"description":  ".town Domains by Opti.Domains X Bored Town for Bored Town holders",
				"image":        "https://metadata.opti.domains/images/domains/town/" + name + ".svg",
				"external_url": "https://town.opti.domains",
			})
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})

	// Render images
	r.StaticFile("/images/domains/town.png", "images/domains/town/avatar.png")

	r.GET("/images/domains/town/:name", func(c *gin.Context) {
		name := c.Param("name")

		if !strings.HasSuffix(name, ".svg") {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		name = strings.TrimSuffix(name, ".svg")

		if len(name) > 22 {
			var prefix string

			if len(name) <= 28 {
				prefix = strings.ReplaceAll(TOWN_DOMAINS_PREFIX, "font-size: 360px;", "font-size: 280px;")
			} else if len(name) <= 40 {
				prefix = strings.ReplaceAll(TOWN_DOMAINS_PREFIX, "font-size: 360px;", "font-size: 200px;")
			} else {
				name = name[:35] + "....town"
				prefix = strings.ReplaceAll(TOWN_DOMAINS_PREFIX, "font-size: 360px;", "font-size: 180px;")
			}

			c.Render(http.StatusOK, render.Data{
				ContentType: "image/svg+xml",
				Data:        []byte(prefix + name + TOWN_DOMAINS_SUFFIX),
			})
		} else {
			c.Render(http.StatusOK, render.Data{
				ContentType: "image/svg+xml",
				Data:        []byte(TOWN_DOMAINS_PREFIX + name + TOWN_DOMAINS_SUFFIX),
			})
		}
	})

	r.Run("0.0.0.0:1888")
}
