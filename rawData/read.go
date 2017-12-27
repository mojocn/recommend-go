package rawData

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	. "github.com/skelterjohn/go.matrix"
	"errors"
)


func readCsvFile(csvPath string) [][]string {
	f, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // this needs to be after the err check
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return lines
}
func ReadScoreCsvToMatrix()(rowSlice []int,colSlice []int,mx *DenseMatrix) {
	csvPath := "/Users/ericzhou/go/src/github.com/mojocn/recommend-go/rawData/ratings.csv"
	lines := readCsvFile(csvPath)
	colSlice = make([]int, 0)
	rowSlice = make([]int, 0)

	for _, line := range lines {
		//userId
		if userId, err := strconv.Atoi(line[0]);err == nil {
			rowSlice = appendIfMissing(rowSlice,userId)
		}else {
			log.Fatal(err)
		}


		//movie_id
		if movieId, err := strconv.Atoi(line[1]); err == nil {
			colSlice = appendIfMissing(colSlice,movieId)
		}else {
			log.Fatal(err)
		}
	}
	//创建矩阵
	//填充
	mat := Zeros(len(rowSlice), len(colSlice))
	for _, line := range lines {
		//计算行号
		userId, _ := strconv.Atoi(line[0])
		if userId == 0 {
			continue
		}
		rowIdx,_ := indexOfValueInSlice(rowSlice,userId)


		//计算col 序号
		movieId, _ := strconv.Atoi(line[1])
		colIdx,_ := indexOfValueInSlice(colSlice,movieId)

		//填充数据
		score,_ := strconv.ParseFloat(line[2],64)


		mat.Set(rowIdx, colIdx, score)
	}

	return rowSlice,colSlice,mat
}

func ReadMovies(){

}

func appendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func indexOfValueInSlice(slice []int,value int) (int, error) {
	for index, ele := range slice {
		if ele == value {
			return index,nil
		}
	}
	return 0, errors.New("can't find value of in slice",)
}