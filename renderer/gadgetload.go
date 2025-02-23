package renderer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/PlayBlockiro/CraftWorks/config"
	"github.com/PlayBlockiro/CraftWorks/utils"
)

type GadgetData struct {
	Color  string  `json:"color"`
	PosX   float64 `json:"pos_x"`
	PosY   float64 `json:"pos_y"`
	PosZ   float64 `json:"pos_z"`
	Scale  float64 `json:"scale"`
	Rotate float64 `json:"rotate"`
}

func LoadGadget(gadgetID string, cfg *config.Config) error {
	url := fmt.Sprintf("http://api.blockiro.mindity.net/gadnet/%s.glb", gadgetID)
	gadgetPath := filepath.Join("blocks", fmt.Sprintf("%s.glb", gadgetID))

	// Download the gadget model if not already present
	if _, err := os.Stat(gadgetPath); os.IsNotExist(err) {
		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("failed to fetch gadget model: %v", err)
		}
		defer resp.Body.Close()

		out, err := os.Create(gadgetPath)
		if err != nil {
			return fmt.Errorf("failed to create gadget file: %v", err)
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return fmt.Errorf("failed to save gadget model: %v", err)
		}

		utils.LogInfo(fmt.Sprintf("Gadget %s downloaded successfully", gadgetID))
	}

	gadgetData := GadgetData{}
	gadgetConfigPath := filepath.Join("config", "main.json")
	data, err := os.ReadFile(gadgetConfigPath)
	if err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}

	err = json.Unmarshal(data, &gadgetData)
	if err != nil {
		return fmt.Errorf("failed to parse gadget data: %v", err)
	}

	utils.LogInfo(fmt.Sprintf("Applying transformations to %s: %+v", gadgetID, gadgetData))
	return nil
}
