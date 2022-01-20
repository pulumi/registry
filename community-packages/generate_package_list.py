import json

import json

# Opening JSON file
f = open('package-list.json')

data = json.load(f)

print(json.dumps(data))

# Closing file
f.close()
