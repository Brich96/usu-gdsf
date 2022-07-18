package db

import (
	"encoding/json"

	"github.com/jak103/usu-gdsf/models"
	"github.com/jak103/usu-gdsf/log"
)

func CreateGamesFromJson() []models.Game {
	result := []models.Game{}
	err := json.Unmarshal([]byte(JSON_SEED_DATA), &result)

	if err != nil {
		log.Warn("Unable to seed example data.")
		return nil
	}

	return result
}

const JSON_SEED_DATA = `[
  {
    "name": "Star Wars 1",
    "author": "George Lucas and Disney",
    "creationDate": "5/25/1977",
    "version": "1.0.1"
  },
  {
    "name": "Star Wars 2",
    "author": "George Lucas and Disney",
    "creationDate": "5/26/1977",
    "version": "1.0.2"
  },
  {
    "name": "Star Wars 3",
    "author": "George Lucas and Disney",
    "creationDate": "5/27/1977",
    "version": "1.0.3"
  },
  {
    "name": "Star Wars 4",
    "author": "George Lucas and Disney",
    "creationDate": "5/28/1977",
    "version": "1.0.4"
  },
  {
    "name": "Star Wars 5",
    "author": "George Lucas and Disney",
    "creationDate": "5/29/1977",
    "version": "1.0.5"
  },
  {
    "name": "Star Wars 6",
    "author": "George Lucas and Disney",
    "creationDate": "5/30/1977",
    "version": "1.0.6"
  },
  {
    "name": "Star Wars 7",
    "author": "George Lucas and Disney",
    "creationDate": "5/31/1977",
    "version": "1.0.7"
  },
  {
    "name": "Star Wars 8",
    "author": "George Lucas and Disney",
    "creationDate": "6/1/1977",
    "version": "1.0.8"
  },
  {
    "name": "Star Wars 9",
    "author": "George Lucas and Disney",
    "creationDate": "6/2/1977",
    "version": "1.0.9"
  },
  {
    "name": "Star Wars 10",
    "author": "George Lucas and Disney",
    "creationDate": "6/3/1977",
    "version": "1.0.10"
  },
  {
    "name": "Star Wars 11",
    "author": "George Lucas and Disney",
    "creationDate": "6/4/1977",
    "version": "1.0.11"
  },
  {
    "name": "Star Wars 12",
    "author": "George Lucas and Disney",
    "creationDate": "6/5/1977",
    "version": "1.0.12"
  },
  {
    "name": "Star Wars 13",
    "author": "George Lucas and Disney",
    "creationDate": "6/6/1977",
    "version": "1.0.13"
  },
  {
    "name": "Star Wars 14",
    "author": "George Lucas and Disney",
    "creationDate": "6/7/1977",
    "version": "1.0.14"
  },
  {
    "name": "Star Wars 15",
    "author": "George Lucas and Disney",
    "creationDate": "6/8/1977",
    "version": "1.0.15"
  },
  {
    "name": "Star Wars 16",
    "author": "George Lucas and Disney",
    "creationDate": "6/9/1977",
    "version": "1.0.16"
  },
  {
    "name": "Star Wars 17",
    "author": "George Lucas and Disney",
    "creationDate": "6/10/1977",
    "version": "1.0.17"
  },
  {
    "name": "Star Wars 18",
    "author": "George Lucas and Disney",
    "creationDate": "6/11/1977",
    "version": "1.0.18"
  },
  {
    "name": "Star Wars 19",
    "author": "George Lucas and Disney",
    "creationDate": "6/12/1977",
    "version": "1.0.19"
  },
  {
    "name": "Star Wars 20",
    "author": "George Lucas and Disney",
    "creationDate": "6/13/1977",
    "version": "1.0.20"
  },
  {
    "name": "Star Wars 21",
    "author": "George Lucas and Disney",
    "creationDate": "6/14/1977",
    "version": "1.0.21"
  },
  {
    "name": "Star Wars 22",
    "author": "George Lucas and Disney",
    "creationDate": "6/15/1977",
    "version": "1.0.22"
  },
  {
    "name": "Star Wars 23",
    "author": "George Lucas and Disney",
    "creationDate": "6/16/1977",
    "version": "1.0.23"
  },
  {
    "name": "Star Wars 24",
    "author": "George Lucas and Disney",
    "creationDate": "6/17/1977",
    "version": "1.0.24"
  },
  {
    "name": "Star Wars 25",
    "author": "George Lucas and Disney",
    "creationDate": "6/18/1977",
    "version": "1.0.25"
  },
  {
    "name": "Star Wars 26",
    "author": "George Lucas and Disney",
    "creationDate": "6/19/1977",
    "version": "1.0.26"
  },
  {
    "name": "Star Wars 27",
    "author": "George Lucas and Disney",
    "creationDate": "6/20/1977",
    "version": "1.0.27"
  },
  {
    "name": "Star Wars 28",
    "author": "George Lucas and Disney",
    "creationDate": "6/21/1977",
    "version": "1.0.28"
  },
  {
    "name": "Star Wars 29",
    "author": "George Lucas and Disney",
    "creationDate": "6/22/1977",
    "version": "1.0.29"
  },
  {
    "name": "Star Wars 30",
    "author": "George Lucas and Disney",
    "creationDate": "6/23/1977",
    "version": "1.0.30"
  },
  {
    "name": "Star Wars 31",
    "author": "George Lucas and Disney",
    "creationDate": "6/24/1977",
    "version": "1.0.31"
  },
  {
    "name": "Star Wars 32",
    "author": "George Lucas and Disney",
    "creationDate": "6/25/1977",
    "version": "1.0.32"
  },
  {
    "name": "Star Wars 33",
    "author": "George Lucas and Disney",
    "creationDate": "6/26/1977",
    "version": "1.0.33"
  },
  {
    "name": "Star Wars 34",
    "author": "George Lucas and Disney",
    "creationDate": "6/27/1977",
    "version": "1.0.34"
  },
  {
    "name": "Star Wars 35",
    "author": "George Lucas and Disney",
    "creationDate": "6/28/1977",
    "version": "1.0.35"
  },
  {
    "name": "Star Wars 36",
    "author": "George Lucas and Disney",
    "creationDate": "6/29/1977",
    "version": "1.0.36"
  },
  {
    "name": "Star Wars 37",
    "author": "George Lucas and Disney",
    "creationDate": "6/30/1977",
    "version": "1.0.37"
  },
  {
    "name": "Star Wars 38",
    "author": "George Lucas and Disney",
    "creationDate": "7/1/1977",
    "version": "1.0.38"
  },
  {
    "name": "Star Wars 39",
    "author": "George Lucas and Disney",
    "creationDate": "7/2/1977",
    "version": "1.0.39"
  },
  {
    "name": "Star Wars 40",
    "author": "George Lucas and Disney",
    "creationDate": "7/3/1977",
    "version": "1.0.40"
  },
  {
    "name": "Star Wars 41",
    "author": "George Lucas and Disney",
    "creationDate": "7/4/1977",
    "version": "1.0.41"
  },
  {
    "name": "Star Wars 42",
    "author": "George Lucas and Disney",
    "creationDate": "7/5/1977",
    "version": "1.0.42"
  },
  {
    "name": "Star Wars 43",
    "author": "George Lucas and Disney",
    "creationDate": "7/6/1977",
    "version": "1.0.43"
  },
  {
    "name": "Star Wars 44",
    "author": "George Lucas and Disney",
    "creationDate": "7/7/1977",
    "version": "1.0.44"
  },
  {
    "name": "Star Wars 45",
    "author": "George Lucas and Disney",
    "creationDate": "7/8/1977",
    "version": "1.0.45"
  },
  {
    "name": "Star Wars 46",
    "author": "George Lucas and Disney",
    "creationDate": "7/9/1977",
    "version": "1.0.46"
  },
  {
    "name": "Star Wars 47",
    "author": "George Lucas and Disney",
    "creationDate": "7/10/1977",
    "version": "1.0.47"
  },
  {
    "name": "Star Wars 48",
    "author": "George Lucas and Disney",
    "creationDate": "7/11/1977",
    "version": "1.0.48"
  },
  {
    "name": "Star Wars 49",
    "author": "George Lucas and Disney",
    "creationDate": "7/12/1977",
    "version": "1.0.49"
  },
  {
    "name": "Star Wars 50",
    "author": "George Lucas and Disney",
    "creationDate": "7/13/1977",
    "version": "1.0.50"
  },
  {
    "name": "Star Wars 51",
    "author": "George Lucas and Disney",
    "creationDate": "7/14/1977",
    "version": "1.0.51"
  },
  {
    "name": "Star Wars 52",
    "author": "George Lucas and Disney",
    "creationDate": "7/15/1977",
    "version": "1.0.52"
  },
  {
    "name": "Star Wars 53",
    "author": "George Lucas and Disney",
    "creationDate": "7/16/1977",
    "version": "1.0.53"
  },
  {
    "name": "Star Wars 54",
    "author": "George Lucas and Disney",
    "creationDate": "7/17/1977",
    "version": "1.0.54"
  },
  {
    "name": "Star Wars 55",
    "author": "George Lucas and Disney",
    "creationDate": "7/18/1977",
    "version": "1.0.55"
  },
  {
    "name": "Star Wars 56",
    "author": "George Lucas and Disney",
    "creationDate": "7/19/1977",
    "version": "1.0.56"
  },
  {
    "name": "Star Wars 57",
    "author": "George Lucas and Disney",
    "creationDate": "7/20/1977",
    "version": "1.0.57"
  },
  {
    "name": "Star Wars 58",
    "author": "George Lucas and Disney",
    "creationDate": "7/21/1977",
    "version": "1.0.58"
  },
  {
    "name": "Star Wars 59",
    "author": "George Lucas and Disney",
    "creationDate": "7/22/1977",
    "version": "1.0.59"
  },
  {
    "name": "Star Wars 60",
    "author": "George Lucas and Disney",
    "creationDate": "7/23/1977",
    "version": "1.0.60"
  },
  {
    "name": "Star Wars 61",
    "author": "George Lucas and Disney",
    "creationDate": "7/24/1977",
    "version": "1.0.61"
  },
  {
    "name": "Star Wars 62",
    "author": "George Lucas and Disney",
    "creationDate": "7/25/1977",
    "version": "1.0.62"
  },
  {
    "name": "Star Wars 63",
    "author": "George Lucas and Disney",
    "creationDate": "7/26/1977",
    "version": "1.0.63"
  },
  {
    "name": "Star Wars 64",
    "author": "George Lucas and Disney",
    "creationDate": "7/27/1977",
    "version": "1.0.64"
  },
  {
    "name": "Star Wars 65",
    "author": "George Lucas and Disney",
    "creationDate": "7/28/1977",
    "version": "1.0.65"
  },
  {
    "name": "Star Wars 66",
    "author": "George Lucas and Disney",
    "creationDate": "7/29/1977",
    "version": "1.0.66"
  },
  {
    "name": "Star Wars 67",
    "author": "George Lucas and Disney",
    "creationDate": "7/30/1977",
    "version": "1.0.67"
  },
  {
    "name": "Star Wars 68",
    "author": "George Lucas and Disney",
    "creationDate": "7/31/1977",
    "version": "1.0.68"
  },
  {
    "name": "Star Wars 69",
    "author": "George Lucas and Disney",
    "creationDate": "8/1/1977",
    "version": "1.0.69"
  },
  {
    "name": "Star Wars 70",
    "author": "George Lucas and Disney",
    "creationDate": "8/2/1977",
    "version": "1.0.70"
  },
  {
    "name": "Star Wars 71",
    "author": "George Lucas and Disney",
    "creationDate": "8/3/1977",
    "version": "1.0.71"
  },
  {
    "name": "Star Wars 72",
    "author": "George Lucas and Disney",
    "creationDate": "8/4/1977",
    "version": "1.0.72"
  },
  {
    "name": "Star Wars 73",
    "author": "George Lucas and Disney",
    "creationDate": "8/5/1977",
    "version": "1.0.73"
  },
  {
    "name": "Star Wars 74",
    "author": "George Lucas and Disney",
    "creationDate": "8/6/1977",
    "version": "1.0.74"
  },
  {
    "name": "Star Wars 75",
    "author": "George Lucas and Disney",
    "creationDate": "8/7/1977",
    "version": "1.0.75"
  },
  {
    "name": "Star Wars 76",
    "author": "George Lucas and Disney",
    "creationDate": "8/8/1977",
    "version": "1.0.76"
  },
  {
    "name": "Star Wars 77",
    "author": "George Lucas and Disney",
    "creationDate": "8/9/1977",
    "version": "1.0.77"
  },
  {
    "name": "Star Wars 78",
    "author": "George Lucas and Disney",
    "creationDate": "8/10/1977",
    "version": "1.0.78"
  },
  {
    "name": "Star Wars 79",
    "author": "George Lucas and Disney",
    "creationDate": "8/11/1977",
    "version": "1.0.79"
  },
  {
    "name": "Star Wars 80",
    "author": "George Lucas and Disney",
    "creationDate": "8/12/1977",
    "version": "1.0.80"
  },
  {
    "name": "Star Wars 81",
    "author": "George Lucas and Disney",
    "creationDate": "8/13/1977",
    "version": "1.0.81"
  },
  {
    "name": "Star Wars 82",
    "author": "George Lucas and Disney",
    "creationDate": "8/14/1977",
    "version": "1.0.82"
  },
  {
    "name": "Star Wars 83",
    "author": "George Lucas and Disney",
    "creationDate": "8/15/1977",
    "version": "1.0.83"
  },
  {
    "name": "Star Wars 84",
    "author": "George Lucas and Disney",
    "creationDate": "8/16/1977",
    "version": "1.0.84"
  },
  {
    "name": "Star Wars 85",
    "author": "George Lucas and Disney",
    "creationDate": "8/17/1977",
    "version": "1.0.85"
  },
  {
    "name": "Star Wars 86",
    "author": "George Lucas and Disney",
    "creationDate": "8/18/1977",
    "version": "1.0.86"
  },
  {
    "name": "Star Wars 87",
    "author": "George Lucas and Disney",
    "creationDate": "8/19/1977",
    "version": "1.0.87"
  },
  {
    "name": "Star Wars 88",
    "author": "George Lucas and Disney",
    "creationDate": "8/20/1977",
    "version": "1.0.88"
  },
  {
    "name": "Star Wars 89",
    "author": "George Lucas and Disney",
    "creationDate": "8/21/1977",
    "version": "1.0.89"
  },
  {
    "name": "Star Wars 90",
    "author": "George Lucas and Disney",
    "creationDate": "8/22/1977",
    "version": "1.0.90"
  },
  {
    "name": "Star Wars 91",
    "author": "George Lucas and Disney",
    "creationDate": "8/23/1977",
    "version": "1.0.91"
  },
  {
    "name": "Star Wars 92",
    "author": "George Lucas and Disney",
    "creationDate": "8/24/1977",
    "version": "1.0.92"
  },
  {
    "name": "Star Wars 93",
    "author": "George Lucas and Disney",
    "creationDate": "8/25/1977",
    "version": "1.0.93"
  },
  {
    "name": "Star Wars 94",
    "author": "George Lucas and Disney",
    "creationDate": "8/26/1977",
    "version": "1.0.94"
  },
  {
    "name": "Star Wars 95",
    "author": "George Lucas and Disney",
    "creationDate": "8/27/1977",
    "version": "1.0.95"
  },
  {
    "name": "Star Wars 96",
    "author": "George Lucas and Disney",
    "creationDate": "8/28/1977",
    "version": "1.0.96"
  },
  {
    "name": "Star Wars 97",
    "author": "George Lucas and Disney",
    "creationDate": "8/29/1977",
    "version": "1.0.97"
  },
  {
    "name": "Star Wars 98",
    "author": "George Lucas and Disney",
    "creationDate": "8/30/1977",
    "version": "1.0.98"
  },
  {
    "name": "Star Wars 99",
    "author": "George Lucas and Disney",
    "creationDate": "8/31/1977",
    "version": "1.0.99"
  },
  {
    "name": "Star Wars 100",
    "author": "George Lucas and Disney",
    "creationDate": "9/1/1977",
    "version": "1.0.100"
  }
]`