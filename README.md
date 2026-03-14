# Линтер для проверки лог-записей

---

Линтер для Go, совместимый с golangci-lint, который будет анализировать лог-записи в коде и проверять их соответствие установленным правилам.

## 1. Установка

```bash
git clone https://github.com/Denis-Mukhametshin-74/selectel-linter.git
cd selectel-linter
go build -o selectel-linter.exe ./cmd/selectel-linter
```

## 2. Использование как standalone линтер

### Запуск на текущем проекте

```bash
./selectel-linter.exe ./...
```

### Запуск на конкретных файлах

```bash
./selectel-linter.exe file.go
```

### Запуск на пакете

```bash
./selectel-linter.exe github.com/username/project/...
```

## 3. Интеграция с golangci-lint

### 3.1 Сборка плагина (на Linux/macOS)

```bash
cd selectel-linter
go build -buildmode=plugin -o selectel-linter.so ./pkg/golangci
```

### 3.2 Настройка .golangci.yml

Создайте файл `.golangci.yml` в корне вашего проекта:

```yaml
version: "2"
linters:
  enable:
    - selectel-linter
  settings:
    custom:
      selectel-linter:
        path: /полный/путь/к/selectel-linter.so
        original-name: selectellinter
```

### 3.3 Запуск golangci-lint

```bash
# Запуск только нашего линтера
golangci-lint run --enable=selectel-linter ./...

# Запуск со всеми линтерами
golangci-lint run
```

## 4. Проверка работоспособности (на Windows)

Создайте тестовый файл `test.go`:

```go
package main

import "log/slog"

func main() {
    // Должны вызвать ошибки
    slog.Info("Starting server")
    slog.Info("запуск сервера")
    slog.Info("starting server!")
    slog.Info("user password: 123")
}
```

Запустите:

```bash
./selectel-linter.exe test.go
```

Ожидаемый вывод:
```
test.go:7:2: лог-сообщения должны начинаться со строчной буквы
test.go:8:2: лог-сообщения должны быть только на английском языке
test.go:8:2: лог-сообщения не должны содержать спецсимволы или эмодзи
test.go:9:2: лог-сообщения не должны содержать спецсимволы или эмодзи
test.go:10:2: лог-сообщения не должны содержать спецсимволы или эмодзи
test.go:10:2: лог-сообщения не должны содержать потенциально чувствительные данные
```
