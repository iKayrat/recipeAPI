package utils

import (
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	db "github.com/iKairat/RecipeAPI/internal/db/sqlc"
	"github.com/lib/pq"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const Migration = `CREATE TABLE "recipes" (
	"id" BIGSERIAL PRIMARY KEY,
	"name" VARCHAR(255) NOT NULL,
	"description" TEXT NOT NULL,
	"ingredients" TEXT[] NOT NULL,
	"steps" JSONB NOT NULL,
	"total_time" SMALLINT NOT NULL,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
  );`

func InitPopulation() {
	log.Println("start init population..")

	args, err := Populate()
	if err != nil {
		log.Fatal(err)
		return
	}

	dbsource := os.Getenv("DBSOURCE")
	if dbsource == "" {
		dbsource = "postgresql://root:kaak@localhost:5432/recipe?sslmode=disable"
	}

	conn, err := sql.Open("postgres", dbsource)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Query(Migration)
	if err != nil {
		log.Println("failed to run migrations", err.Error())
		return
	}

	count := 0
	row := conn.QueryRow("SELECT COUNT(*) FROM recipes")
	err = row.Scan(&count)
	if err != nil {
		log.Println(err)
	}

	if count >= 99 {
		log.Println("Init Proccess canceled")
		return
	}

	tx, err := conn.Begin()
	if err != nil {
		log.Fatal("tx begin", err)
		return
	}

	// Prepare the SQL statement for bulk insert
	stmt, err := tx.Prepare("INSERT INTO recipes (id, name, description, ingredients, steps, total_time, created_at, updated_at) VALUES ($1, $2, $3, $4, $5,$6,$7,$8)")
	if err != nil {
		log.Fatal("prepare:", err)
	}
	defer stmt.Close()

	stepMaps := make([]map[string]int, 3)
	sMp1 := make(map[string]int, 1)
	sMp2 := make(map[string]int, 1)
	sMp3 := make(map[string]int, 1)
	s := RandomString(steps)
	i := int(RandomInt(10, 50))
	s1 := RandomString(steps)
	s2 := RandomString(steps)
	i1 := int(RandomInt(10, 50))
	i2 := int(RandomInt(10, 50))
	sMp1[s] = int(i)
	sMp2[s1] = int(i1)
	sMp2[s2] = int(i2)
	stepMaps[0] = sMp1
	stepMaps[1] = sMp2
	stepMaps[2] = sMp3
	// res := fmt.Sprintf("{%v,%v,%v}", stepMaps[0], stepMaps[1], stepMaps[2])

	totaltime := 0
	for _, v := range stepMaps {
		for _, m := range v {
			totaltime += m
		}
	}
	args[i].TotalTime = int16(totaltime)

	stepJson, _ := json.Marshal(stepMaps)

	// Insert each recipe in the transaction
	for _, recipe := range args {
		_, err := stmt.Exec(recipe.ID, recipe.Name, recipe.Description, pq.Array(RandomSlice(ingredients, 3)), stepJson, totaltime, recipe.CreatedAt, recipe.UpdatedAt)
		if err != nil {
			log.Println(err)
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Population Done")
}

var (
	names       = []string{"Homemade pizza", "Twice-baked potatoes", "Chicken pot pie (easier than you think!)", "Homemade corn dogs", "Shepherd's pie", "Calzones", "French bread pizza", "Enchiladas", "Egg rolls (homemade)", "Chicken parmesan", "Fried chicken (baked or fried)", "Chicken cordon bleu", "Breaded chicken ", "Chicken nuggets ", "Chicken wings", "Cheesy bacon chicken", "Dijon chicken", "Chicken fingers", "Chicken and broccoli casserole", "Orange chicken", "Cilantro-lime chicken", "Lemon chicken"}
	steps       = []string{"Broiling", "Dry Heat Cooking", "Baking", "Roasting", "Poaching", "Simmering", "Boiling", "Steaming", "Stewing", "AlBaste", "Blanch", "Brunoise", "Caramelize", "Chiffonade", "Clarify", "Cure", "Deglaze", "Dredge", "Emulsify", "Fillet", "Flambe", "Fold", "Julienne", "Meuniere", "Parboil", "Reduce", "Scald", "Spatchcock"}
	ingredients = []string{"Olive oil", "All purpose flour", "Butter", "Chicken", "Sugar", "Olive oil", "All purpose flour", "Butter", "Chicken", "Sugar", "Salt", "Egg", "Rice", "Vegetable oil", "Pork", "Beef", "Cheese", "Garlic", "Orange", "Turkey", "Onion", "Corn", "Whole milk", "Mayonnaise", "Chiles", "Almonds", "Bacon", "Mushrooms", "Coconut", "Beets", "Strawberries", "Fennel", "Lamb", "Apple", "Shrimp"}

	description = []string{
		"pasta that is firm and slightly undercooked",
		"pour juices or liquid fat over meat while it cooks",
		"scald food in b`oiling water for a quick moment and then place it in cold water to stop the cooking process",
		"cut foods in to a 1/8 size dice",
		"heat sugars until they are browned",
		"roll up leafy greens or herbs and cut into long, thin slices",
		"melt butter and separate the solids from the butterfat",
		"preserve foods by adding salt and drawing out moisture",
		"dissolve browned food residue in a hot pan with liquid",
		"coat moist foods in a dry ingredient, like flour",
		"blend two liquids like oil and water",
		"cut a portion of meat or fish",
		"cover a food in a flammable liquid, like brandy or rum, and light it briefly on fire",
		"incorporate an ingredient with a careful motion that retains air",
		"cut foods into long thin strips",
		"method of cooking, usually used with fish, in which the food is lightly dusted with flour and sauteed in butter",
		"precook foods by boiling for a short time",
		"thicken a liquid mixture by boiling or simmering, causing moisture to evaporate",
		"heat a liquid just to the boiling point",
		"split open a whole chicken or turkey for easy grilling",
	}
)

func RandomString(list []string) string {
	random := rand.Intn(len(list))

	return list[random]
}

func RandomSlice(list []string, size int) []string {
	// Shuffle the list in-place
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	// Take the first 'size' elements as the random slice
	return list[:size]
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func Populate() ([]db.Recipe, error) {

	args := make([]db.Recipe, 100)

	for i := 1; i < len(args); i++ {

		name := RandomString(names)
		des := RandomString(description)
		str := RandomSlice(ingredients, 3)

		args[i].ID = int64(i)
		args[i].Name = name
		args[i].Description = des
		args[i].Ingredients = append(args[i].Ingredients, str...)
		args[i].Steps = nil
		args[i].CreatedAt = time.Now()
		args[i].UpdatedAt = time.Now()

	}

	return args, nil

}
