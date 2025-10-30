# Concurrency

[TOC]

## Проблемы многопоточного программирования

### Data race

`Data race` - несинхронизированный досутуп к ячейке памяти двумя или более потоками, где как минимум один из потоков осуществляет доступ на запись.

[Пример `data race`](./data-race/)

```go
func main() {
	n := 0
	wg := sync.WaitGroup{}
	for range 1000 {
		wg.Go(func() {
			n++
		})
	}
	wg.Wait()
	fmt.Println(n) // 1000 or 994 or ...
}
```

### Race condition

`Race condition` - ситуация, при которой результат исполнения программы зависит от порядка планирования потоков.

[Пример `race condition`](./race-condition/)

```go
func main() {
	wg := sync.WaitGroup{}
	for i := range 10 {
		wg.Go(func() {
			fmt.Println(i)
		})
	}
	wg.Wait()
}
```

`Data race` не тождественно равен `race condition`. Эти баги могут быть причиной друг друга, а могут возникать отдельно, как в примерах выше.

## Проблемы многопоточной синхронизации
