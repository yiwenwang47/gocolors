# A simple program to extract a palette from a picture. 

Image manipulation in Golang.

The task to extract a color palette from a picture is always arbitrary, and the approach of this program is to: blur -> resize -> kmeans cluster.

## Dependency
https://github.com/esimov/stackblur-go

https://github.com/nfnt/resize

https://github.com/muesli/kmeans

https://github.com/parnurzeal/gorequest

## Usage
To install

```bash
go get github.com/yiwenwang9702/gocolors
```

Results can be refined by accessing the Colormind API (http://colormind.io/). According to the website, their algorithm is a pretrained GAN that refines or completes a palette.

An example is available under the Examples directory.

Compared to blurring and resizing, clustering does not seem to be a good idea. Its optimization focuses too much on including all colors present, while we only try to find the important ones.


## Features in Progress
Extracting colors by performing a clustering analysis seems to work well on real world pictures. But for the pictures created artificially, where a lot of pixels can have the exactly same RGB values, clustering does not seem to be a good idea.

Might write a port of colorgram.py (https://github.com/obskyr/colorgram.py). Which exploits the idea of sampling.

Might add a function to sort the palette in some way.

Is outlier detection a good idea here?

Ongoing.
