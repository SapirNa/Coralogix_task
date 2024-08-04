package main

import (
	"simple-app/utils"
)

func main() {
	R := utils.Read("input.csv").FilterRows(2, "Sheryl").GetColumn(12).Sum_column()
	utils.Write(R.GetOutput(), "output.csv")
}
