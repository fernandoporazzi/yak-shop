package cmd

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fernandoporazzi/yak-shop/app/entity"
	"github.com/fernandoporazzi/yak-shop/app/service"
	"github.com/spf13/cobra"
)

var getDataCmd = &cobra.Command{
	Use:   "get-data",
	Short: "Returns data according to a given XML and elapsed days",
	Long:  "Returns stock given a XML file as input and a the number of elapsed days",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")

		if err != nil {
			fmt.Println("Error getting flag `file`")
			panic(err)
		}

		days, err := cmd.Flags().GetInt64("days")
		if err != nil {
			fmt.Println("Error getting flag `days`")
			panic(err)
		}

		xmlFile, err := os.Open(file)
		if err != nil {
			fmt.Println("Error")
			panic(err)
		}
		defer xmlFile.Close()

		byteValue, err := ioutil.ReadAll(xmlFile)
		if err != nil {
			fmt.Println("Error creating byteValue")
			panic(err)
		}

		var herd entity.Herd

		err = xml.Unmarshal(byteValue, &herd)
		if err != nil {
			fmt.Println("Error unmarshalling xml")
			panic(err)
		}

		// in case we want to use a database
		// repository := repository.NewStockRepository(herd)
		stockService := service.NewStockService(herd)

		fmt.Println(stockService.GetMilkByDays(days))
		fmt.Println(stockService.GetSkinByDays(days))
	},
}
