#method 1: average value method
import numpy as np
n = 1_000_000
tmp = np.random.rand(n)
tmp = np.sqrt(1-tmp**2)
tmp = 4*tmp.mean()
print('The estimate result of average value method: ',tmp)

#method2: area method
import numpy as np
n = 1_000_000
tmp = np.random.rand(n,2)
tmp = np.square(tmp).sum(axis=1)
tmp = (tmp<1).sum() / n * 4
print('The estimate result of area method: ',tmp)
