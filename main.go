package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"

	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Elements struct {
	Name                          string      `json:"name"`
	Appearance                    string      `json:"appearance"`
	AtomicMass                    float64     `json:"atomic_mass"`
	Boil                          float64     `json:"boil"`
	Category                      string      `json:"category"`
	Color                         interface{} `json:"color"`
	Density                       float64     `json:"density"`
	DiscoveredBy                  string      `json:"discovered_by"`
	Melt                          float64     `json:"melt"`
	MolarHeat                     float64     `json:"molar_heat"`
	NamedBy                       string      `json:"named_by"`
	Number                        int         `json:"number"`
	Period                        int         `json:"period"`
	Phase                         string      `json:"phase"`
	Source                        string      `json:"source"`
	SpectralImg                   string      `json:"spectral_img"`
	Summary                       string      `json:"summary"`
	Symbol                        string      `json:"symbol"`
	Xpos                          int         `json:"xpos"`
	Ypos                          int         `json:"ypos"`
	Shells                        []int       `json:"shells"`
	ElectronConfiguration         string      `json:"electron_configuration"`
	ElectronConfigurationSemantic string      `json:"electron_configuration_semantic"`
	ElectronAffinity              float64     `json:"electron_affinity"`
	ElectronegativityPauling      float64     `json:"electronegativity_pauling"`
	IonizationEnergies            []float64   `json:"ionization_energies"`
	CpkHex                        string      `json:"cpk-hex"`
}

var element []Elements

// queryParam function

// all element route
func allElement(c *gin.Context) {
	c.JSON(http.StatusOK, element)
}

// element by name
func elementByName(c *gin.Context) {
	for _, item := range element {
		if item.Name == c.Param("name") {
			c.JSON(http.StatusOK, item)
		}
	}
}

// element by number
func elementByNumber(c *gin.Context) {
	for _, item := range element {
		if strconv.Itoa(item.Number) == c.Param("number") {
			c.JSON(http.StatusOK, item)
		}
	}
}

// element by symbol
func elementBySymbol(c *gin.Context) {
	for _, item := range element {
		if item.Symbol == c.Param("symbol") {
			c.JSON(http.StatusOK, item)
		}
	}
}

// element by phase
func elementByPhase(c *gin.Context) {
	for _, item := range element {
		if item.Phase == c.Param("phase") {
			c.JSON(http.StatusOK, item)
		}
	}
}

// random element
func randomElement(c *gin.Context) {
	min := 1
	max := 119
	c.JSON(http.StatusOK, element[rand.Intn(max-min)+min])
}
func main() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}
	file, err := os.ReadFile("element.json")
	if err != nil {
		panic("data not found!")
	}
	err = json.Unmarshal(file, &element)
	if err != nil {
		panic("data not found!")
	}
	r := gin.Default()
	r.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	r.GET("/", allElement)
	r.GET("/name/:name", elementByName)
	r.GET("/number/:number", elementByNumber)
	r.GET("/symbol/:symbol", elementBySymbol)
	r.GET("/phase/:phase", elementByPhase)
	r.GET("/random", randomElement)
	r.Run()
}
