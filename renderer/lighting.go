package renderer

import (
	"fmt"

	"github.com/PlayBlockiro/CraftWorks/utils"
)

type LightData struct {
	PosX  float64 `json:"pos_x"`
	PosY  float64 `json:"pos_y"`
	PosZ  float64 `json:"pos_z"`
	Level int     `json:"level"` // Light intensity 0-15
}

func ProcessLighting(lightMap []LightData) {
	for _, light := range lightMap {
		utils.LogInfo(
			fmt.Sprintf("Light at (%.2f, %.2f, %.2f) Level: %d", light.PosX, light.PosY, light.PosZ, light.Level),
		)
	}
}
