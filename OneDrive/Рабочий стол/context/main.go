package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second *10)//withTimeout Это дедлайн< то есть мы даем какому либо процессу время на выполнения.Если функция не успевает то мы его блокируем и выводим соответствующий текст 
	parse(ctx)
}

func parse(ctx context.Context){
	for{
		select{
		case <- time.After(time.Second * 4):
			fmt.Println("parsing completed")
			return
		case <- ctx.Done():
			fmt.Println("deadline exceded")
			return
		}
	}
}