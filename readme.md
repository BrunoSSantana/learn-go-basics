# Resumo para Estudo - Conceitos e Demonstrações em Go

## Pacote Principal e Importações
- **Pacote principal**: `package main`
- **Importações**: Uso do pacote `fmt` para entrada e saída.

---

## Constantes
- **Constantes de nível de pacote** são acessíveis por todas as funções do pacote:
  ```go
  const PackageLevelScope = "Package Level Scope"
  ```
E para torná-la pública para outros pacotes, basta usar a primeira letra maiúscula como está no exemplo acima.

---

## Função `main`
- Chama funções para demonstrar conceitos:
  - Variáveis e operadores
  - Ponteiros
  - Condicionais
  - Loops
  - Switch
  - Funções

---

## Conceitos e Demonstrações

### 1. **Variáveis e Operadores**
- Declaração com `var` (com ou sem valor inicial).
- Atribuição curta `:=` (não pode ser usada no nível de pacote).
- Reatribuição com `=`.

**Exemplo:**
```go
var oldValue string
shortDeclaration := "Olá, Mundo!"
shortDeclaration = "Novo Valor"
```

- Função `show` para exibir valores de qualquer tipo utilizado "interface vazia" que aceita qualquer tipo de dado, assim como "any" no TypeScript.
  ```go
  func show(data interface{}) {
      fmt.Println("Valor recebido:", data)
  }
  ```

---

### 2. **Ponteiros**
- Declaração e uso de ponteiros com os operadores `&` (endereço) e `*` (valor).

**Exemplo:**
```go
data := 10
pointer := &data
*pointer = 20
```

- Alteração de valor usando ponteiros:
  ```go
  // Função para alterar o valor de um inteiro usando um ponteiro
  func updateValueUsingPointer(data *int) int {
      // Atualiza o valor armazenado no endereço de memória
      *data = 20
      return *data
  }

  // Demonstrar o uso da função
  updatedValue := updateValueUsingPointer(&data)
  fmt.Println("Valor atualizado:", updatedValue)

  ```

---

### 3. **Condicionais**
- Uso do `if` com inicialização:
  ```go
  if data := 10; true {
      fmt.Println(data)
  }
  ```

---

### 4. **Loops**
- Tipos de `for`:
  - Clássico: `for i := 0; i < 5; i++`.
  - Com `range`: Itera sobre slices.
  - Infinito: `for { break }`.
  - Estilo `while`: `for i < 5`.

**Exemplo de Loop com `range`:**
```go
data := []int{1, 2, 3}
for index, value := range data {
    fmt.Printf("Índice: %v, Valor: %v\n", index, value)
}
```

---

### 5. **Switch**
- Avaliação de valores com `switch`:
  ```go
  switch data := 10; data {
  case 1:
      fmt.Println("Caso 1")
  default:
      fmt.Println("Caso padrão")
  }
  ```

---

### 6. **Funções**
#### Função Anônima
- Declaração e execução inline:
  ```go
  func() {
      fmt.Println("Função anônima")
  }()
  ```

#### Funções Nomeadas
- Com retorno simples:
  ```go
  func funcNamed(a, b int) int {
      return a * b
  }
  ```

- Com retorno nomeado:
  ```go
  func funcNamedReturn(a, b int) (result int) {
      result = a * b
      return
  }
  ```

- Retornos múltiplos:
  ```go
  func multipleReturn() (int, int) {
      return 1, 2
  }
  ```

#### Função Variádica
- Aceita múltiplos argumentos:
  ```go
  func variadicFunction(nums ...int) int {
      sum := 0
      for _, num := range nums {
          sum += num
      }
      return sum
  }
  ```

#### Função Retornando Função
- Uma função pode retornar outra função:
  ```go
  func outerFunction() func() int {
      x := 10
      return func() int {
          return x * x
      }
  }
  ```

---

## Organização e Estrutura
- Uso de modularização para separar conceitos em funções.
- Funções nomeadas e bem documentadas para reforçar o aprendizado.
- Exemplos práticos para cada conceito.

---

### Pontos de Foco para Estudo
1. Diferença entre `var` e `:=`.
2. Manipulação de ponteiros com `&` e `*`.
3. Diferentes formas de usar `for`.
4. Estrutura do `switch` em Go.
5. Funções variádicas e funções que retornam outras funções.

**Boa prática**: Documente e modularize seu código para facilitar o entendimento.