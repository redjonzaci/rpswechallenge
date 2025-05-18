package main

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, HEAD, OPTIONS, PATCH, POST, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/calculate", func(c *gin.Context) {
		var calculateRequest struct {
			PackSizes []int `json:"packSizes"`
			Amount    int   `json:"amount"`
		}

		if err := c.ShouldBindJSON(&calculateRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		orderPacks, err := CalculateOrderPacks(calculateRequest.PackSizes, calculateRequest.Amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"orderPacks": orderPacks})
	})

	return r
}

// CalculateOrderPacks calculates the optimal pack distribution for an order based on the following rules:
// 1. Only whole packs can be sent
// 2. Send the least amount of items to fulfill the order (priority)
// 3. Send as few packs as possible
//
// Parameters:
//   - packSizes: Slice of available pack sizes
//   - amountOrdered: Total amount ordered
//
// Returns:
//   - A map of pack sizes to their counts
//   - An error if no solution exists
func CalculateOrderPacks(packSizes []int, amountOrdered int) (map[int]int, error) {
	// Create a copy of packSizes and sort it in descending order
	sortedSizes := make([]int, len(packSizes))
	copy(sortedSizes, packSizes)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedSizes)))

	// Find the largest pack size to determine the maximum possible amount
	maxPackSize := sortedSizes[0]
	maxPossible := amountOrdered + maxPackSize

	// Step 1: Find the minimum valid amount >= amountOrdered that can be made
	// This is to satisfy rule #2: send out the least amount of items

	// Initialize an array to track which amounts can be made
	// canMake[i] = true if amount i can be made using available pack sizes
	canMake := make([]bool, maxPossible+1)
	canMake[0] = true

	// Calculate which amounts can be made using the available pack sizes
	for _, size := range packSizes {
		for i := size; i <= maxPossible; i++ {
			if canMake[i-size] {
				canMake[i] = true
			}
		}
	}

	// Find minimum amount >= amountOrdered that can be made
	minValidAmount := amountOrdered
	for minValidAmount <= maxPossible && !canMake[minValidAmount] {
		minValidAmount++
	}

	// If no valid amount found
	if minValidAmount > maxPossible {
		return nil, fmt.Errorf("no solution exists for the given pack sizes and amount")
	}

	// Step 2: Find minimum number of packs to make minValidAmount
	// This satisfies rule #3: send out as few packs as possible

	// Initialize array to track minimum packs needed for each amount
	minPacks := make([]int, maxPossible+1)
	for i := range minPacks {
		minPacks[i] = maxPossible + 1 // Use a value larger than possible pack count as "infinity"
	}
	minPacks[0] = 0

	// Track which pack was used for each amount (for reconstruction)
	packUsed := make([]int, maxPossible+1)
	for i := range packUsed {
		packUsed[i] = -1
	}

	// Calculate minimum packs needed for each amount
	for i := 1; i <= maxPossible; i++ {
		for _, size := range packSizes {
			if i >= size && minPacks[i-size] != maxPossible+1 {
				if minPacks[i-size]+1 < minPacks[i] {
					minPacks[i] = minPacks[i-size] + 1
					packUsed[i] = size
				}
			}
		}
	}

	// Step 3: Reconstruct the solution
	result := make(map[int]int)
	for _, size := range packSizes {
		result[size] = 0
	}

	remaining := minValidAmount
	for remaining > 0 {
		size := packUsed[remaining]
		if size == -1 {
			// This should not happen if our DP solution is correct
			return nil, fmt.Errorf("internal error: no valid pack size found during reconstruction")
		}
		result[size]++
		remaining -= size
	}

	return result, nil
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
