package main

import (
	"log"

	"github.com/go-pdf/fpdf"
)

func main() {

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(10, 10, 10)
	pdf.AddPage()

	
	logoWidth := 30.0
	pdf.ImageOptions("no.png", 10, 10, logoWidth, 0, false, fpdf.ImageOptions{}, 0, "")


	pdf.SetXY(10+logoWidth+15, 10) 
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 10, "No Text", "", 0, "L", false, 0, "")


	pdf.SetY(40)

	
	pdf.SetLineWidth(0.5)
	pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
	pdf.Ln(10)


	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Terms and Conditions", "", 1, "C", false, 0, "")
	pdf.Ln(10)

	
	sections := []struct {
		title, content string
	}{
		{"TERMS", "There is no need of nay terms and conditions you can do anything here man, I am with you 100%"},
		{"User License", "Your Driving License is enough of YOU :P."},
		{"Enjoy Madi", "Don't read boring stuff please"},
	}

	for _, section := range sections {
		pdf.SetFont("Arial", "B", 14)
		pdf.CellFormat(0, 10, section.title, "", 1, "L", false, 0, "")
		pdf.Ln(5)

		pdf.SetFont("Arial", "", 12)
		pdf.MultiCell(0, 10, section.content, "", "L", false)
		pdf.Ln(5)
	}

	
	err := pdf.OutputFileAndClose("ASI TERMS and Conditions.pdf")
	if err != nil {
		log.Fatalf("Failed to create PDF: %v", err)
	}

	log.Println("PDF file successfully created: NO TERMS and Conditions.pdf")
}
