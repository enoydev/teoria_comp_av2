# Comparativo Merge Sort: Go vs. Python

Este projeto simples demonstra e compara o desempenho do algoritmo Merge Sort (recursivo) implementado em Go e Python.

### Propósito


O objetivo é observar e quantificar as diferenças de desempenho (tempo de execução, média e desvio padrão) entre as implementações do Merge Sort em linguagens com características de execução distintas (Go compilado vs. Python interpretado), para diferentes tamanhos e tipos de dados.

### Algoritmo


Apenas a versão recursiva do Merge Sort é utilizada para uma comparação justa entre as duas linguagens, garantindo que o mesmo padrão algorítmico esteja sendo testado.


- Complexidade de Tempo (Big-Theta): \(\Theta(N \log N)\)

- Complexidade de Espaço (Big-Theta): \(\Theta(N)\)

### Como Executar

##### 1. Preparação (Go)
Certifique-se de ter Go instalado (versão 1.18+ para suporte a genéricos).
```
# Crie o diretório do projeto Go
mkdir go-merge-sort-comparison
cd go-merge-sort-comparison
go mod init go-merge-sort-comparison

# Salve o código Go fornecido (main.go) neste diretório.
```



##### 2. Preparação (Python)
Certifique-se de ter Python instalado (versão 3.x).

```
# Crie o diretório do projeto Python
mkdir python-merge-sort-comparison
cd python-merge-sort-comparison

# Salve o código Python fornecido (merge_sort.py) neste diretório.
```

##### 3. Execução e Coleta de Dados
Execute cada script e redirecione a saída para um arquivo para análise:

Go
```
go run main.go > go_results.txt
```

Python
```
python merge_sort.py > python_results.txt
```

4. Análise
Compare os arquivos go_results.txt e python_results.txt. Ambos os arquivos contêm tabelas com:


- Tamanho N: O número de elementos na lista.

- Tipo Entrada: Tipo de ordenação inicial da lista (aleatório, ordenado, inverso).

- Média (ns): Tempo médio de execução em nanosegundos (30 execuções).

- Desvio Padrão (ns): Desvio padrão dos tempos de execução.

- Ordenado?: Confirmação se a lista resultante está corretamente ordenada.
