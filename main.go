package main

import (
	"bufio"
	"estiam/dictionary"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/arriqaaq/flashdb"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create the dictionary
	dico := dictionary.New()

	// Add the database
	config := &flashdb.Config{Path: "./db", EvictionInterval: 10}
	db, err := flashdb.New(config)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Initialize the router
	router := gin.Default()

	// Start the CLI for CRUD operations
	go cli(dico)

	// Route to retrieve the list of words and their definitions
	router.GET("/list", func(c *gin.Context) {
		err := db.View(func(tx *flashdb.Tx) error {
			val, err := tx.Get("java")
			if err != nil {
				return err
			}
			c.JSON(http.StatusOK, gin.H{
				"mot":        "java",
				"définition": val,
			})
			return nil
		})

		if err != nil {
			log.Println("Error retrieving word from flashDB:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving word"})
			return
		}
	})

	// Route to retrieve the definition of a word
	router.GET("/word/:name", func(c *gin.Context) {
		err := db.View(func(tx *flashdb.Tx) error {
			word := c.Params.ByName("name")
			val, err := tx.Get(word)
			if err != nil {
				return err
			}
			c.JSON(http.StatusOK, gin.H{
				"mot":        word,
				"définition": val,
			})
			return nil
		})

		if err != nil {
			log.Println("Error retrieving word from flashDB:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving word"})
			return
		}
	})

	// Route to add a word and its definition
	router.POST("/word", func(c *gin.Context) {
		mot := c.PostForm("mot")
		definition := c.PostForm("definition")

		err := db.Update(func(tx *flashdb.Tx) error {
			err := tx.Set(mot, definition)
			return err
		})

		if err != nil {
			log.Println("Error adding word to flashDB:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding word"})
			return
		}

		dico.Add(mot, definition)

		c.JSON(http.StatusOK, gin.H{"message": "Word added successfully"})
	})

	// Route to delete a word
	router.DELETE("/delete/:name", func(c *gin.Context) {
		word := c.Params.ByName("name")

		err := db.Update(func(tx *flashdb.Tx) error {
			tx.Delete(word)
			if err != nil {
				return err
			}
			c.JSON(http.StatusOK, gin.H{
				"msg": "Word deleted",
			})
			return nil
		})

		if err != nil {
			log.Println("Error deleting word from flashDB:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting word"})
			return
		}
	})

	// Route to update a word
	router.POST("/update/:name", func(c *gin.Context) {
		word := c.Params.ByName("name")
		definition := c.PostForm("definition")

		err := db.Update(func(tx *flashdb.Tx) error {
			tx.Set(word, definition)
			if err != nil {
				return err
			}
			c.JSON(http.StatusOK, gin.H{
				"msg": "Word updated",
			})
			return nil
		})

		if err != nil {
			log.Println("Error updating word in flashDB:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating word"})
			return
		}
	})

	router.Run()
}

func cli(dico *dictionary.Dictionary) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter ADD, DEF, REMOVE, LIST, EXIT: ")
		arg, _ := reader.ReadString('\n')

		word := strings.TrimSpace(arg)

		switch word {
		case "ADD":
			actionAdd(dico, reader)
		case "DEF":
			actionDefine(dico, reader)
		case "REMOVE":
			actionRemove(dico, reader)
		case "LIST":
			actionList(dico)
		case "EXIT":
			return
		}
	}
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter the word: ")
	mot, _ := reader.ReadString('\n')
	mot = strings.TrimSpace(mot)

	fmt.Print("Enter the definition: ")
	definition, _ := reader.ReadString('\n')
	definition = strings.TrimSpace(definition)

	d.Add(mot, definition)
}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter the word: ")
	mot, _ := reader.ReadString('\n')
	mot = strings.TrimSpace(mot)

	entrie, _ := d.Get(mot)
	fmt.Println("  Definition:", entrie)
}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter the word: ")
	mot, _ := reader.ReadString('\n')
	mot = strings.TrimSpace(mot)

	d.Remove(mot)
}

func actionList(d *dictionary.Dictionary) {
	words, entries := d.List()

	fmt.Println("Words in the dictionary:")

	for _, word := range words {
		entry := entries[word]

		fmt.Println("- Word:", word)
		fmt.Println("- Definition:", entry)
	}
}
