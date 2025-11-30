import random
import time
import math

def merge_sort(collection: list) -> list:
    """Pure Python implementation of merge sort algorithm - OPTIMIZED"""
    
    def merge(left: list, right: list) -> list:
        result = []
        i = j = 0  # Usar índices em vez de pop(0)
        
        # Merge usando índices (O(n) em vez de O(n²))
        while i < len(left) and j < len(right):
            if left[i] <= right[j]:
                result.append(left[i])
                i += 1
            else:
                result.append(right[j])
                j += 1
        
        # Adicionar elementos restantes
        result.extend(left[i:])
        result.extend(right[j:])
        return result
    
    if len(collection) <= 1:
        return collection
    
    mid_index = len(collection) // 2
    return merge(merge_sort(collection[:mid_index]), merge_sort(collection[mid_index:]))


# --- Data Generation Functions ---
def generate_random_list(size: int) -> list:
    return [random.randint(0, size * 10) for _ in range(size)]

def generate_sorted_list(size: int) -> list:
    return list(range(size))

def generate_reverse_sorted_list(size: int) -> list:
    return list(range(size - 1, -1, -1))

def is_list_sorted(lst: list) -> bool:
    return all(lst[i] <= lst[i+1] for i in range(len(lst)-1)) if lst else True


def calculate_mean_and_std_dev(times: list) -> tuple[float, float]:
    """Calcula média e desvio padrão de tempos em nanosegundos"""
    if not times:
        return 0.0, 0.0
    
    mean = sum(times) / len(times)
    
    if len(times) < 2:
        return mean, 0.0
    
    # Desvio padrão amostral
    variance = sum((t - mean) ** 2 for t in times) / (len(times) - 1)
    std_dev = math.sqrt(variance)
    return mean, std_dev


# --- Comparação de Testes ---
def run_comparison_tests():
    N_VALUES = [10, 100, 1000, 5000, 10000, 50000, 100000, 200000, 500000]
    INPUT_TYPES = {
        "aleatorio": generate_random_list,
        "ordenado":  generate_sorted_list,
        "inverso":   generate_reverse_sorted_list,
    }
    NUM_TRIALS = 30
    
    print("### Python Merge Sort (Recursivo Otimizado) - Tempos de Execução")
    print(f"{'Tamanho N':<10} | {'Tipo Entrada':<15} | {'Média (ns)':<15} | {'Desvio Padrão (ns)':<19} | {'Ordenado?'}")
    print("-" * 85)
    
    for n in N_VALUES:
        for input_name, generator in INPUT_TYPES.items():
            execution_times = []
            sorted_flag = False
            
            # Gera dados uma vez para todos os trials (mais justo)
            test_data = generator(n)
            
            for i in range(NUM_TRIALS):
                data_copy = test_data.copy()  # Cópia rasa (suficiente para ints)
                
                start = time.perf_counter_ns()
                sorted_data = merge_sort(data_copy)
                end = time.perf_counter_ns()
                
                execution_times.append(end - start)
                
                if i == 0:
                    sorted_flag = is_list_sorted(sorted_data)
            
            mean, std_dev = calculate_mean_and_std_dev(execution_times)
            print(f"{n:<10} | {input_name:<15} | {mean:<15.2f} | {std_dev:<19.2f} | {sorted_flag}")
        
        print("-" * 85)


if __name__ == "__main__":
    random.seed(42)  # Seed fixo para reproducibilidade
    run_comparison_tests()