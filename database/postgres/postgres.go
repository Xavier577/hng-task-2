package postgres

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

var client *sqlx.DB

type PgConnectCfg struct {
	Host     string `json:"host"`
	PORT     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

func parseCfg(pgCfg *PgConnectCfg) string {
	var cfgMap map[string]any

	cfgJson, _ := json.Marshal(&pgCfg)

	_ = json.Unmarshal(cfgJson, &cfgMap)

	var cfgPairs []string

	for key, val := range cfgMap {
		if val != nil && val != "" {
			cfgPairs = append(cfgPairs, fmt.Sprintf("%s=%v", key, val))
		}
	}

	return strings.Join(cfgPairs, " ")
}

func Connect(pgCfg *PgConnectCfg) error {
	var connectionError error

	dataSource := parseCfg(pgCfg)

	log.Println(dataSource)

	client, connectionError = sqlx.Connect("postgres", dataSource)

	return connectionError
}

func Client() *sqlx.DB {
	return client
}
