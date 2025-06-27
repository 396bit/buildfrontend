# buildfrontend

A simple go:generate tool.

I looks for a **frontend** folder. If found an the folder contains a package.json, it calls
```sh
cd frontend
bun run build
```


TODO make the js tool (bun) configurable
