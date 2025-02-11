package main

import (
	"log"

	"github.com/go-pdf/fpdf"
)

func addHeader(pdf *fpdf.Fpdf) {
	
	logoWidth := 30.0
	logoHeight := 10.0
	pdf.ImageOptions("ASI.png", 10, 10, logoWidth, logoHeight, false, fpdf.ImageOptions{}, 0, "")

	
	pdf.SetXY(10+logoWidth+5, 10) 
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, logoHeight, "Amit Prakash", "", 0, "L", false, 0, "")

	
	pdf.SetY(10 + logoHeight + 10) 

	
	pdf.SetLineWidth(0.5)
	pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
	pdf.Ln(10)
}

func addFooter(pdf *fpdf.Fpdf) {
	currentY := pdf.GetY()
	
	pdf.SetY(-15)
	pdf.SetFont("Arial", "I", 10)
	pdf.CellFormat(0, 10, "© 2025 amitprakash.me.", "", 0, "C", false, 0, "")
	n
	pdf.SetY(currentY)
}

func main() {
	
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(10, 10, 10)

	
	pdf.SetFooterFunc(func() {
		addFooter(pdf)
	})

	
	pdf.SetAutoPageBreak(true, 20)

	pdf.AddPage()
	addHeader(pdf)

	
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Terms and Conditions", "", 1, "C", false, 0, "")
	pdf.Ln(10)

	
	sections := []struct {
		title, content string
	}{
		{"TERMS", "This is the first section content. It contains detailed information about topic one.\nMore details can span multiple lines and pages as needed."},
		{"User Licence", "This is the second section content. More details about topic two go here.\nThis text is long enough to trigger page breaks automatically."},
		{"Enjoy Madi", "This is the third section content. Additional details for topic three.\nMore lines to ensure multiple pages are created dynamically."},
		{"TERMS", "This is the first section content. It contains detailed information about topic one.\nMore details can span multiple lines and pages as needed."},
		{"User Licence", "This is the second section content. More details about topic two go here.\nThis text is long enough to trigger page breaks automatically."},
		{"Enjoy Madi", "This is the third section content. Additional details for topic three.\nMore lines to ensure multiple pages are created dynamically."},
		{"TERMS", "This is the first section content. It contains detailed information about topic one.\nMore details can span multiple lines and pages as needed."},
		{"User Licence", "This is the second section content. More details about topic two go here.\nThis text is long enough to trigger page breaks automatically."},
		{"Enjoy Madi", "This is the third section content. Additional details for topic three.\nMore lines to ensure multiple pages are created dynamically."},
		{"TERMS", "This is the first section content. It contains detailed information about topic one.\nMore details can span multiple lines and pages as needed."},
		{"User Licence", "This is the second section content. More details about topic two go here.\nThis text is long enough to trigger page breaks automatically."},
		{"Enjoy Madi", "This is the third section content. Additional details for topic three.\nMore lines to ensure multiple pages are created dynamically."},
	}

	for _, section := range sections {
		pdf.SetFont("Arial", "B", 14)
		pdf.CellFormat(0, 10, section.title, "", 1, "L", false, 0, "")
		pdf.Ln(5)

		pdf.SetFont("Arial", "", 12)
		contentLines := pdf.SplitLines([]byte(section.content), 190)
		for _, line := range contentLines {
			pdf.CellFormat(0, 10, string(line), "", 1, "L", false, 0, "")
		}
		pdf.Ln(5)
	}

	
	err := pdf.OutputFileAndClose("ou11tput.pdf")
	if err != nil {
		log.Fatalf("Failed to create PDF: %v", err)
	}

	log.Println("PDF file successfully created: output.pdf")
}
