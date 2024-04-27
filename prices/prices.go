package prices

import (
	"fmt"

	"githhub.com/devphasex/tax-calculator/conversion"
	"githhub.com/devphasex/tax-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64            `json:"tax_rate"`
	InputPrices       []float64          `json:"input_prices"`
	TaxIncludedPrices map[string]string  `json:"tax_included_prices"`
	IOManager         iomanager.IOManger `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}
	priceInterestMap := make(map[string]string, len(job.InputPrices))

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		key := fmt.Sprintf("%.2f", price)
		priceInterestMap[key] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = priceInterestMap
	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iomd iomanager.IOManger, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   iomd,
	}
}
