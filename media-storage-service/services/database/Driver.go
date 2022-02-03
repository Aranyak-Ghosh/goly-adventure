package database

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

func NewDbDriver() (neo4j.Driver, error) {
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "test", ""))
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func CloseDriver(driver neo4j.Driver) {
	driver.Close()
}
