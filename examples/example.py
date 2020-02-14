
Paramiko
```python
  import hashlib, paramiko.agent
  data\_sha1 = haslib.sha1(recordToPublish).digest()
  agent = paramiko.agent.Agent()
  key = agent.keys\[0\]
  signature = key.sign\_ssh\_data(None, data\_sha1)
```
