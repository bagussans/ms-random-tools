package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"

	"github.com/bagussans/ms-support-golang/utils"
)

type Tools struct{}
type CalcStudentAvgScorePayload struct {
	Type string                      `json:"type"`
	Data []CalcStudentAvgScoreHolder `json:"data"`
}
type ImageEditorPayload struct {
	Base64 string                    `json:"base64"`
	Filter *string                   `json:"filter"`
	Resize *ImageEditorResizePayload `json:"resize"`
}

type ImageEditorResizePayload struct {
	Width  int `json:"width"`
	Heigth int `json:"height"`
}
type CalcStudentAvgScoreHolder struct {
	Name       string  `json:"name"`
	AvgDesired float64 `json:"avgDesired"`
	AmountMany int     `json:"amountMany"`
	Min        int     `json:"min"`
	Max        int     `json:"max"`
}
type CalcStudentAvgScoreResp struct {
	Name      string `json:"name"`
	ScoreList []int  `json:"scoreList"`
	AvgScore  int    `json:"avgScore"`
	Error     string `json:"error,omitempty"`
}

func (t *Tools) CalcStudentAvgScore(w http.ResponseWriter, r *http.Request) {
	var calcStudentAvgScorePayload CalcStudentAvgScorePayload
	err := json.NewDecoder(r.Body).Decode(&calcStudentAvgScorePayload)
	if err != nil {
		resp, err := utils.NewErrorResponse(err, http.StatusBadRequest)
		utils.SendJSON(w, resp, err)
		return
	}

	finalResult := []CalcStudentAvgScoreResp{}
	for _, student := range calcStudentAvgScorePayload.Data {
		loopCounter := 0
		keeploop := true
		for {
			tempValue := 0
			finalListScore := []int{}
			loopCounter += 1

			if !keeploop {
				break
			}

			for i := 0; i < student.AmountMany; i++ {
				val := utils.RandomIntFromInterval(student.Min, student.Max)
				finalListScore = append(finalListScore, val)
				tempValue += val
			}

			tempValue = tempValue / (student.AmountMany)

			if float64(tempValue) >= student.AvgDesired-0.9 && float64(tempValue) <= student.AvgDesired+0.9 {
				keeploop = false
				finalResult = append(finalResult, CalcStudentAvgScoreResp{
					Name:      student.Name,
					ScoreList: finalListScore,
					AvgScore:  tempValue,
					Error:     "",
				})
			}

			if loopCounter >= 1500 {
				keeploop = false
				finalResult = append(finalResult, CalcStudentAvgScoreResp{
					Name:      student.Name,
					ScoreList: []int{},
					AvgScore:  0,
					Error:     "Error too many loops",
				})
			}
		}
	}

	if calcStudentAvgScorePayload.Type == "text" {
		resp, err := utils.NewSuccessResponse("success", finalResult, http.StatusOK)
		utils.SendJSON(w, resp, err)
	} else {
		header := []string{}
		row := [][]string{}
		filename := "student_score_avg_" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".csv"

		// create header, loop only once
		for _, s := range finalResult {

			if len(s.ScoreList) > 0 {
				header = append(header, "Name")
				for i := range s.ScoreList {
					header = append(header, "Score_"+strconv.Itoa(i+1))
				}
				header = append(header, "Average Score")
				break
			}

		}

		//create row
		for _, s := range finalResult {

			value := []string{}
			value = append(value, s.Name)

			// strArr := make([]string, len(s.ScoreList))
			for _, v := range s.ScoreList {
				value = append(value, strconv.Itoa(v))
			}

			value = append(value, strconv.Itoa(s.AvgScore))
			row = append(row, value)
		}

		csvTemp := utils.GenerateCSVBase64(
			filename,
			header,
			row,
		)

		constructResp := map[string]string{
			"filename": filename,
			"file":     csvTemp,
		}

		resp, err := utils.NewSuccessResponse("success", constructResp, http.StatusOK)
		utils.SendJSON(w, resp, err)
	}
}

func (t *Tools) ImageEditor(w http.ResponseWriter, r *http.Request) {
	var imageEditorPayload ImageEditorPayload
	err := json.NewDecoder(r.Body).Decode(&imageEditorPayload)
	if err != nil {
		resp, _ := utils.NewErrorResponse(err, http.StatusBadRequest)
		utils.SendJSON(w, resp, err)
		return
	}

	base64String := imageEditorPayload.Base64
	fmt.Printf("Received Base64 prefix: %.50s...\n", base64String)
	fmt.Printf("Received Base64 length: %d\n", len(base64String))

	// Strip data URI prefix
	if strings.HasPrefix(base64String, "data:") {
		parts := strings.SplitN(base64String, ",", 2)
		if len(parts) == 2 {
			base64String = parts[1]
		}
	}

	// Normalize (remove line breaks)
	base64String = strings.ReplaceAll(base64String, "\n", "")
	base64String = strings.ReplaceAll(base64String, "\r", "")
	base64String = strings.TrimSpace(base64String)

	// Decode Base64
	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)

	if err != nil {
		log.Printf("Error decoding base64: %v", err)
		resp, _ := utils.NewErrorResponse(err, http.StatusBadRequest)
		utils.SendJSON(w, resp, err)
		return
	}

	// Decode image
	imgReader := bytes.NewReader(decodedBytes)
	img, format, err := image.Decode(imgReader)
	if err != nil {
		log.Printf("Error decoding image: %v", err)
		resp, _ := utils.NewErrorResponse(err, http.StatusBadRequest)
		utils.SendJSON(w, resp, err)
		return
	}

	var editedImg image.Image
	var resizedImage image.Image
	var filteredImage image.Image

	//RESIZE
	if imageEditorPayload.Resize != nil {
		resizedImage = utils.ResizeUtil(img, imageEditorPayload.Resize.Width, imageEditorPayload.Resize.Heigth)
		editedImg = resizedImage
	} else {
		editedImg = img
	}

	//FILTER !
	if imageEditorPayload.Filter != nil {
		if *imageEditorPayload.Filter == "Grayscale" {
			filteredImage = utils.GrayscaleUtil(editedImg)
			editedImg = filteredImage
		}
	}

	var buf bytes.Buffer
	_ = png.Encode(&buf, editedImg)

	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	// Optional: Add data URI prefix so the frontend can directly use it as <img src="...">
	dataURI := "data:image/" + format + ";base64," + base64Str
	constructResp := map[string]string{
		"fileBase64": dataURI,
	}
	resp, err := utils.NewSuccessResponse("success", constructResp, http.StatusOK)
	utils.SendJSON(w, resp, err)

	// // Optionally, save the image to a file
	// outputFile, err := os.Create("decoded_image." + format)
	// if err != nil {
	// 	log.Fatalf("Error creating output file: %v", err)
	// 	resp, err := utils.NewErrorResponse(err, http.StatusBadRequest)
	// 	utils.SendJSON(w, resp, err)
	// 	return
	// }
	// defer outputFile.Close()

	// switch format {
	// case "jpeg":
	// 	err = os.WriteFile(outputFile.Name(), decodedBytes, 0644) // For JPEG, raw bytes are often sufficient
	// case "png":
	// 	// For PNG, you might need to encode the image.Image object
	// 	// if you want to ensure proper compression or metadata handling.
	// 	// For simplicity, we're writing the raw decoded bytes here.
	// 	err = os.WriteFile(outputFile.Name(), decodedBytes, 0644)
	// default:
	// 	log.Printf("Unsupported format for direct writing: %s. Consider encoding the image.Image object.", format)
	// 	// For other formats, you would typically use an encoder from the respective package, e.g., png.Encode
	// }

	// if err != nil {
	// 	log.Fatalf("Error writing image to file: %v", err)
	// 	resp, err := utils.NewErrorResponse(err, http.StatusBadRequest)
	// 	utils.SendJSON(w, resp, err)
	// 	return
	// }

	// fmt.Printf("Image successfully decoded and saved as %s\n", outputFile.Name())
}
