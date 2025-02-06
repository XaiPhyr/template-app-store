package utils

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		SMTP   SMTPConf     `yaml:"smtp"`
		Server ServerConfig `yaml:"server"`
	}

	ServerConfig struct {
		JwtKey string `yaml:"jwt_key"`
	}

	SMTPConf struct {
		Host   string `yaml:"host"`
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
		Port   int    `yaml:"port"`
		Sender string `yaml:"sender"`
	}
)

func InitConfig() Config {
	var cfg Config
	env := os.Getenv("APP_ENVIRONMENT")

	rootFile := fmt.Sprintf("./conf/config-%s.yml", env)
	f, err := os.Open(rootFile)
	if err != nil {
		log.Printf("Error: %s", err)
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Printf("Error: %s", err)
	}

	return cfg
}

func Paginator(q url.Values) (limit, offset int, keyword, qExt string, sortBy []string) {
	limit, _ = strconv.Atoi(q.Get("size"))
	page, _ := strconv.Atoi(q.Get("page"))
	sorts := q.Get("sort")
	keyword = q.Get("q")
	qExt = q.Get("qExt")

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	offset = (page - 1) * limit

	if len(sorts) > 0 {
		for _, sort := range strings.Split(sorts, ",") {
			if strings.Contains(sort, "-") {
				sortBy = append(sortBy, strings.ReplaceAll(sort, "-", "")+" DESC")
			} else {
				sortBy = append(sortBy, sort+" ASC")
			}
		}
	}

	return
}
