# Synchronization primitives

- [Mutex 1 (Atomics)](./mutex-on-atomic)
- [Mutex 2 (Channels)](./mutex-on-channel)

---

## Base

### Slice & array

## Concurrency patterns

### Fan-in pattern

Написать функцию, которая сливает данные из нескольких каналов в один

```go
func fanin(sources ...<-chan int) <-chan int {
    // CODE
}
```

[Solution](./fan-in/)

### Fan-out pattern

Написать функцию, распределяет входные данные по нескольким каналам в соответствии с предикатом

```go
func fanout(source <-chan int, predicate func(int) bool) [2]<-chan int {
    // CODE
}
```

[Solution](./fan-out/)

### Done channel

Реализовать механизм, при котором `caller` завершает асинхронный `worker` и продолжает выполнение только после его полного завершения

[Solution](./done-ch)

### Pipeline pattern

Реализовать функцию, которая читает значения из канала, применяет к ним некоторый `func apply()` и шлет в другой канал

[Solution](./pipeline)

### Worker pool

Модернизировать предыдущий пример, добавив пул размера `N`.

[Solution](./workerpool)

### Semaphore

Реализовать `WaitGroup` на семафоре

### Slow function

Пусть некоторый внешний сервис имеет API:

```go
func slow() (int, error) {
    // Something slow
}
```

Наш сервис готов ждать результат `N` секунд или вернуть ошибку по таймауту.
Написать обертку для данной функции, реализующую данный функционал

```go
func wrapWithTimeout(/* TODO */) /* TODO */ {
    /* TODO */
}
```

Изменить реализацию обертки с учетом того, что если функция API вернула ошибку, но у нас еще есть время подождать, то мы можем сделать повторный запрос (Функция API идемпотента)

```go
func wrapWithTimeoutAndRetry(/* TODO */) /* TODO */ {
    /* TODO */
}
```

## Interview Qs

### Fetch URls

#### Base

Написать функцию, которая опрашивает `url`-ы из списка и опрашивает
