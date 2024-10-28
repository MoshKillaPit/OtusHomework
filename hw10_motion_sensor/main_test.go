package main

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func Test_randomNum(t *testing.T) {
	type args struct {
		randomNumbers chan int
		limit         int
		duration      time.Duration
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "100",
			args: args{
				randomNumbers: make(chan int, 100),
				limit:         100,
				duration:      1 * time.Second,
			},
			want: 100,
		},
		{
			name: "Time",
			args: args{
				randomNumbers: make(chan int, 100),
				limit:         100,
				duration:      1 * time.Nanosecond,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				randomNum(tt.args.randomNumbers, tt.args.limit, tt.args.duration)
			}()

			// Небольшая задержка для ожидания закрытия канала
			time.Sleep(10 * time.Millisecond)

			count := 0
			for range tt.args.randomNumbers {
				count++
			}
			wg.Wait()
			if count >= tt.want {
				t.Errorf("randomNum() = %v, want less than %v", count, tt.want)
			}
		})
	}
}

func Test_refactor(t *testing.T) {
	type args struct {
		randomNumbers chan int
		resultChan    chan Result
	}
	tests := []struct {
		name           string
		args           args
		inputNumbers   []int
		expectedResult []Result
	}{
		{
			name: "Multiple groups of 10",
			args: args{
				randomNumbers: make(chan int, 20),
				resultChan:    make(chan Result, 2),
			},
			inputNumbers: []int{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10, // Группа 1, среднее 5
				11, 12, 13, 14, 15, 16, 17, 18, 19, 20, // Группа 2, среднее 15
			},
			expectedResult: []Result{
				{GroupNumber: 1, Average: 5},
				{GroupNumber: 2, Average: 15},
			},
		},
		{
			name: "Less than 10 numbers",
			args: args{
				randomNumbers: make(chan int, 5),
				resultChan:    make(chan Result, 1),
			},
			inputNumbers: []int{3, 5, 7}, // Остаток, среднее 5
			expectedResult: []Result{
				{GroupNumber: 1, Average: 5},
			},
		},
		{
			name: "More than 10 but not a multiple of 10",
			args: args{
				randomNumbers: make(chan int, 15),
				resultChan:    make(chan Result, 2),
			},
			inputNumbers: []int{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10, // Группа 1, среднее 5
				11, 12, 13, // Остаток, среднее 12
			},
			expectedResult: []Result{
				{GroupNumber: 1, Average: 5},
				{GroupNumber: 2, Average: 12},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				for _, num := range tt.inputNumbers {
					tt.args.randomNumbers <- num
				}
				close(tt.args.randomNumbers)
			}()

			go refactor(tt.args.randomNumbers, tt.args.resultChan)

			var results []Result
			for res := range tt.args.resultChan {
				results = append(results, res)
			}

			if !reflect.DeepEqual(results, tt.expectedResult) {
				t.Errorf("refactor() = %v, want %v", results, tt.expectedResult)
			}
		})
	}
}
