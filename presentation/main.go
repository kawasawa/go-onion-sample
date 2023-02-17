package main

import (
	"context"
	"fmt"
	"go-onion-sample/domain/entity"
	"go-onion-sample/usecase"
	"os"
	"strconv"
)

func main() {
	err := execute()
	if err != nil {
		panic(err)
	}
}

func execute() error {
	ctx := context.Background()
	params := entity.ParseOSArgs(os.Args)

	service, err := di(ctx)
	if err != nil {
		return err
	}

	addResult, err := service.Add(ctx, params)
	if err != nil {
		return err
	}

	subResult, err := service.Subtract(ctx, params)
	if err != nil {
		return err
	}

	fmt.Println("add: " + strconv.Itoa(addResult))
	fmt.Println("sub: " + strconv.Itoa(subResult))
	return nil
}

func di(ctx context.Context) (calcService usecase.ICalcService, err error) {
	return usecase.NewCalcService(), nil
}
