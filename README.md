# A simple program to extract a palette from a picture. 

Image manipulation in Golang.

The task to extract a color palette from a picture is always arbitrary, and the approach of this program is to: blur -> resize -> kmeans cluster.

## Dependency
github.com/esimov/stackblur-go
github.com/nfnt/resize
github.com/muesli/kmeans
github.com/parnurzeal/gorequest

## Usage
To install

```bash
go get github.com/yiwenwang9702/gocolors
```

Results can be refined by accessing the Colormind API (http://colormind.io/). According to the website, their algorithm is a pretrained GAN that refines or completes a palette.

An example is available under the Examples directory.

Colormind's GAN seems to work pretty well. However, the results generated by example_1.go do seem better than those generated by directly uploading the picture to the Colormind website. The difference might come from the different methods to sample colors from a picture.

Again, this is only a matter of opinion.

Might add a feature to sort the palette in some way.

Ongoing.
