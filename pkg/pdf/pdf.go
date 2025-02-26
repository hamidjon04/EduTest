package pdf

import (
	"edutest/pkg/model"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

func CreateTestTemplate(template model.CreatePdf) (string, error) {
	// PDF yaratish
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 14)

	// Talaba ma'lumotlari
	pdf.Cell(0, 10, fmt.Sprintf("Student ID: %s", template.StudentId))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Name: %s %s", template.Name, template.Lastname))
	pdf.Ln(15)

	// Savollar boshi
	for i, question := range template.Questions {
		// 1-savoldan oldin Subject1 ni markazga joylash
		if i == 0 {
			pdf.SetFont("Arial", "B", 16)
			pdf.CellFormat(0, 10, template.Subject1, "", 1, "C", false, 0, "")
			pdf.Ln(5)
			pdf.SetFont("Arial", "", 14)
		}

		// 31-savoldan oldin Subject2 ni markazga joylash
		if i == 30 {
			pdf.SetFont("Arial", "B", 16)
			pdf.CellFormat(0, 10, template.Subject2, "", 1, "C", false, 0, "")
			pdf.Ln(5)
			pdf.SetFont("Arial", "", 14)
		}

		// Savol
		pdf.SetFont("Arial", "B", 12)
		pdf.MultiCell(0, 7, fmt.Sprintf("%d. %s", i+1, question.QuestionText), "", "L", false)
		pdf.SetFont("Arial", "", 12)
		pdf.Ln(3)

		// Variantlar
		pdf.Cell(0, 7, fmt.Sprintf("A) %s", question.Options.A))
		pdf.Ln(7)
		pdf.Cell(0, 7, fmt.Sprintf("B) %s", question.Options.B))
		pdf.Ln(7)
		pdf.Cell(0, 7, fmt.Sprintf("C) %s", question.Options.C))
		pdf.Ln(7)
		pdf.Cell(0, 7, fmt.Sprintf("D) %s", question.Options.D))
		pdf.Ln(15)
	}

	// Faylni saqlash katalogi
	dir := "./storage/pdfs"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	// Fayl nomi
	filePath := filepath.Join(dir, fmt.Sprintf("%s.pdf", template.TemplateId))
	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to save PDF: %w", err)
	}

	return filePath, nil
}
