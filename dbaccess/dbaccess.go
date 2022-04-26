package dbaccess

import (
	"fmt"

	"github.com/Jarod-CG/Prototipo_Libros/structs"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var (
	driver *neo4j.Driver
)

func getDriver() *neo4j.Driver {
	if driver == nil {
		auth := neo4j.BasicAuth(structs.Config.Username, structs.Config.Password, "")
		driverI, err := neo4j.NewDriver(structs.Config.Uri, auth)
		if err != nil {
			panic(err)
		}
		driver = &driverI
	}
	return driver
}

func TestNeo4j() {
	// Aura requires you to use "neo4j+s" protocol
	// (You need to replace your connection details, username and password)
	auth := neo4j.BasicAuth(structs.Config.Username, structs.Config.Password, "")
	// You typically have one driver instance for the entire application. The
	// driver maintains a pool of database connections to be used by the sessions.
	// The driver is thread safe.
	driver, err := neo4j.NewDriver(structs.Config.Uri, auth)
	if err != nil {
		panic(err)
	}
	// Don't forget to close the driver connection when you are finished with it
	defer driver.Close()
	// Create a session to run transactions in. Sessions are lightweight to
	// create and use. Sessions are NOT thread safe.
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	// WriteTransaction retries the operation in case of transient errors by
	// invoking the anonymous function multiple times until it succeeds.
	records, err := session.WriteTransaction(
		func(tx neo4j.Transaction) (interface{}, error) {
			// To learn more about the Cypher syntax, see https://neo4j.com/docs/cypher-manual/current/
			// The Reference Card is also a good resource for keywords https://neo4j.com/docs/cypher-refcard/current/
			createRelationshipBetweenPeopleQuery := `
				MERGE (p1:Person { name: $person1_name })
				MERGE (p2:Person { name: $person2_name })
				MERGE (p1)-[:KNOWS]->(p2)
				RETURN p1, p2`
			result, err := tx.Run(createRelationshipBetweenPeopleQuery, map[string]interface{}{
				"person1_name": "Alice",
				"person2_name": "David",
			})
			if err != nil {
				// Return the error received from driver here to indicate rollback,
				// the error is analyzed by the driver to determine if it should try again.
				return nil, err
			}
			// Collects all records and commits the transaction (as long as
			// Collect doesn't return an error).
			// Beware that Collect will buffer the records in memory.
			return result.Collect()
		})
	if err != nil {
		panic(err)
	}
	for _, record := range records.([]*neo4j.Record) {
		firstPerson := record.Values[0].(neo4j.Node)
		fmt.Printf("First: '%s'\n", firstPerson.Props["name"].(string))
		secondPerson := record.Values[1].(neo4j.Node)
		fmt.Printf("Second: '%s'\n", secondPerson.Props["name"].(string))
	}

	// Now read the created persons. By using ReadTransaction method a connection
	// to a read replica can be used which reduces load on writer nodes in cluster.
	_, err = session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		// Code within this function might be invoked more than once in case of
		// transient errors.
		readPersonByName := `
			MATCH (p:Person)
			WHERE p.name = $person_name
			RETURN p.name AS name`
		result, err := tx.Run(readPersonByName, map[string]interface{}{
			"person_name": "Alice",
		})
		if err != nil {
			return nil, err
		}
		// Iterate over the result within the transaction instead of using
		// Collect (just to show how it looks...). Result.Next returns true
		// while a record could be retrieved, in case of error result.Err()
		// will return the error.
		for result.Next() {
			fmt.Printf("Person name: '%s' \n", result.Record().Values[0].(string))
		}
		// Again, return any error back to driver to indicate rollback and
		// retry in case of transient error.
		return nil, result.Err()
	})
	if err != nil {
		panic(err)
	}
}

func CreateAuthor(authorName, birth string) {
	dr := *getDriver()
	session := dr.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(
		func(tx neo4j.Transaction) (interface{}, error) {
			query := `
				CREATE (a:Author 
					{Name: $name, DateOfBirth: $birth}
					) 
				RETURN a.Name AS Name, a.DateOfBirth AS DateOfBirth`

			result, err := tx.Run(query, map[string]interface{}{
				"name":  authorName,
				"birth": birth,
			})
			if err != nil {
				return nil, err
			}
			return result.Collect()
		})
	if err != nil {
		panic(err)
	}
}

func CreateBook(authorName, bookName, topic, date string, price float64, quantity int) {
	dr := *getDriver()
	session := dr.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(
		func(tx neo4j.Transaction) (interface{}, error) {
			query := `
				MERGE (a:Author {Name: $authorName}) 
				MERGE (b:Book {Name: $bookName, Topic: $topic, Date: $date, Price: $price, Quantity: $quantity})
				MERGE (a)-[:WRITE]->(b)
				`

			result, err := tx.Run(query, map[string]interface{}{
				"authorName": authorName,
				"bookName":   bookName,
				"topic":      topic,
				"date":       date,
				"price":      price,
				"quantity":   quantity,
			})
			if err != nil {
				return nil, err
			}
			return result.Collect()
		})
	if err != nil {
		panic(err)
	}
}

func CreateReader(readerName string) {
	dr := *getDriver()
	session := dr.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(
		func(tx neo4j.Transaction) (interface{}, error) {
			query := `
				CREATE (r:Reader 
					{Name: $readerName}
				) 
				`
			result, err := tx.Run(query, map[string]interface{}{
				"readerName": readerName,
			})
			if err != nil {
				return nil, err
			}
			return result.Collect()
		})
	if err != nil {
		panic(err)
	}
}

func OrderBook(readerName, bookName string) {
	dr := *getDriver()
	session := dr.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(
		func(tx neo4j.Transaction) (interface{}, error) {
			query := `
				MATCH (b:Book {Name: $bookName})
				MATCH (r:Reader {Name: $readerName})
				CREATE (r)-[:ORDER {Date: date()}]->(b)
				`
			result, err := tx.Run(query, map[string]interface{}{
				"readerName": readerName,
				"bookName":   bookName})
			if err != nil {
				return nil, err
			}
			return result.Collect()
		})
	if err != nil {
		panic(err)
	}
}

func BuyBook(readerName, bookName string) {
	dr := *getDriver()
	session := dr.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(
		func(tx neo4j.Transaction) (interface{}, error) {
			query := `
			MATCH (b:Book{Name: $bookName}) 
			MATCH (r:Reader{Name:$readerName})
			WHERE b.Quantity > 0 
			CREATE (r)-[:BUY{Date:date()}]->(b) 
			SET b.Quantity = b.Quantity - 1
			`
			result, err := tx.Run(query, map[string]interface{}{
				"readerName": readerName,
				"bookName":   bookName})
			if err != nil {
				return nil, err
			}
			return result.Collect()
		})
	if err != nil {
		panic(err)
	}
}

func WatchAuthor(readerName, authorName string) {
	dr := *getDriver()
	session := dr.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(
		func(tx neo4j.Transaction) (interface{}, error) {
			query := `
			MATCH (r:Reader{Name:$readerName})
			MATCH (a:Author{Name: $authorName}) 
			CREATE (r)-[:WATCH{Date:date()}]->(a)
			`
			result, err := tx.Run(query, map[string]interface{}{
				"readerName": readerName,
				"authorName": authorName})
			if err != nil {
				return nil, err
			}
			return result.Collect()
		})
	if err != nil {
		panic(err)
	}
}

func WatchBook(readerName, bookName string) {
	dr := *getDriver()
	session := dr.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(
		func(tx neo4j.Transaction) (interface{}, error) {
			query := `
			MATCH (r:Reader{Name:$readerName})
			MATCH (b:Book{Name: $bookName}) 
			CREATE (r)-[:WATCH{Date:date()}]->(b)
			`
			result, err := tx.Run(query, map[string]interface{}{
				"readerName": readerName,
				"bookName":   bookName})
			if err != nil {
				return nil, err
			}
			return result.Collect()
		})
	if err != nil {
		panic(err)
	}
}

func CleanDB() {
	dr := *getDriver()
	session := dr.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(
		func(tx neo4j.Transaction) (interface{}, error) {
			query := `
			MATCH (r:Reader{Name:$readerName})
			MATCH (b:Book{Name: $bookName}) 
			CREATE (r)-[:WATCH{Date:date()}]->(b)
			`
			result, err := tx.Run(query, map[string]interface{}{})
			if err != nil {
				return nil, err
			}
			return result.Collect()
		})
	if err != nil {
		panic(err)
	}
}
