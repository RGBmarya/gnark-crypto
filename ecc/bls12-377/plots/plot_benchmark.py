import matplotlib.pyplot as plt
import re
import numpy as np

def parse_log_cores(file_path):
    points = []
    ns_op = []
    cores = []

    with open(file_path, 'r') as file:
        for line in file:
            # get each element from the line
            match = re.match(r'BenchmarkMultiExpG1/(\d+)_points-(\d+)\s+\d+\s+(\d+) ns/op', line)
            if match:
                num_points = int(match.group(1))  
                nanoseconds = int(match.group(3)) 
                core_count = int(match.group(2))
                
                points.append(num_points)
                ns_op.append(nanoseconds)
                cores.append(core_count)
    
    return points, ns_op, cores


def plot_benchmark_cores(file_path_og, file_path_modified):
    points_og, ns_op_og, cores_og = parse_log_cores(file_path_og)
    points_mod, ns_op_mod, cores_mod = parse_log_cores(file_path_modified)
    
    # num of unique points
    unique_points = sorted(set(points_og).union(points_mod))
    
    for num_points in unique_points:
        
        points_og_filtered = [p for p in points_og if p == num_points]
        ns_op_og_filtered = [ns for p, ns in zip(points_og, ns_op_og) if p == num_points]
        cores_og_filtered = [cores for p, cores in zip(points_og, cores_og) if p == num_points]
        
        points_mod_filtered = [p for p in points_mod if p == num_points]
        ns_op_mod_filtered = [ns for p, ns in zip(points_mod, ns_op_mod) if p == num_points]
        cores_mod_filtered = [cores for p, cores in zip(points_mod, cores_mod) if p == num_points]
        
        if points_og_filtered and points_mod_filtered:
            plt.figure(figsize=(12, 6))
            plt.plot(cores_og_filtered, ns_op_og_filtered, marker='o', linestyle='-', color='k', label='baseline')
            plt.plot(cores_mod_filtered, ns_op_mod_filtered, marker='o', linestyle='-', color='r', label='initial SIMD')
            
            plt.xscale('linear')
            plt.xticks(cores_og_filtered, fontsize=12)
            plt.yticks(fontsize=12)
            plt.yscale('log')
            plt.xlabel('Number of Cores',fontsize=15)
            plt.ylabel('Nanoseconds per Operation (ns/op)', fontsize=15)
            plt.title(f'Benchmark Results for MultiExp - 2^{int(np.log2(num_points))} points', fontsize=15)
            plt.legend(bbox_to_anchor=(1, 0.8), loc='upper right', ncol=1)
            plt.grid(True, which="both", ls="--", linewidth=0.5)
            #plt.savefig(f'simd_2^{int(np.log2(num_points))}.png')
            plt.show()


#plot_benchmark_cores('./gnarkcrypto.log', './modified.log', './purego.log')
plot_benchmark_cores('./{yourpath}.log', './{yourpath}.log')

