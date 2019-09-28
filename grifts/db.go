package grifts

import (
	"fmt"

	"github.com/arschles/gifm-site/models"
	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		screencasts := []models.Screencast{
			{
				EpisodeNum: 1,
				Title:      "This is episode 1!",
				Intro:      "Here's an intro",
				Markdown:   "# Here's some markdown",
			},
			{
				EpisodeNum: 1,
				Title:      "This is episode 2!",
				Intro:      "Here's an intro",
				Markdown:   "# Here's some markdown",
			},
			{
				EpisodeNum: 1,
				Title:      "This is episode 3!",
				Intro:      "Here's an intro",
				Markdown:   "# Here's some markdown",
			},
		}
		for _, screencast := range screencasts {
			if _, err := models.DB.ValidateAndCreate(&screencast); err != nil {
				return err
			}
		}
		fmt.Println("Seeded the DB with", len(screencasts), "screencast(s)")
		return nil
	})

})
