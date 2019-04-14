# A simple program to extract a palette from a picture. 

Image manipulation in Golang.

The task to extract a color palette from a picture is always arbitrary, and the current approach is to: blur -> resize -> kmeans cluster.

To install

```bash
go get github.com/yiwenwang9702/gocolors
```

An example is available under the Examples directory.

Might add a feature to sort the palette in some way.

Might use the Colormind API (http://colormind.io/) to refine the results.

Ongoing.
