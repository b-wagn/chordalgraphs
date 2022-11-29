import subprocess
import matplotlib.pyplot as plt


seed = 7290832189
num_instances = 10

tests = ["recognition", "lexbfs", "validitycheck", "clique", "coloring", "cliquecover and indset"]
vertices = range(10,2000,100)
times = []
for i in range(len(tests)):
	times.append([])

for v in vertices:
	print("Running with " + str(v) + " vertices.")
	output = subprocess.check_output(["./algo", str(seed), str(v), str(num_instances)]) 
	outputs = output.split("\n")
	for i in range(len(tests)):
		times[i].append(float(outputs[i]))


for i in range(len(tests)):
	plt.plot(vertices, times[i])
	plt.xlabel('Number of Vertices')
	plt.ylabel('Average Execution time in milliseconds')
	plt.title('Performance of algorithm ' + tests[i])
	plt.legend()
	plt.show()

