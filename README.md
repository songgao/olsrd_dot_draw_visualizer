# olsrd Dot Draw Visualizer
A simple HTTP server for displaying real-time topology information retrieved from `dot_draw` plugin of [olsrd][olsrd] routing daemon. The topology is updated on the fly every 200 ms (can be changed in `html/script.js`).

[olsrd]: http://www.olsr.org/

## Usage
```bash
git clone git://github.com/songgao/olsrd_dot_draw_visualizer.git
cd olsrd_dot_draw_visualizer
go run main.go # make sure you have Go installed
```
Now, go to `http://localhost:8080/` to view the graph changing in real-time. Change the port number in `main.go` if needed.

* The binary `olsrd_dot_draw_visualizer` in source directory is for x86_64 on Linux

## Credit
- viz.js ([https://github.com/mdaines/viz.js/](https://github.com/mdaines/viz.js/))
- Go ([http://golang.org/](http://golang.org/))
- jquery ([http://jquery.com/](http://jquery.com/))
