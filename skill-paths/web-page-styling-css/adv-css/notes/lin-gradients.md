# linear gradients
## what is a gradient?
- graduated blend b/w 2 colors
- 2 types - linear and radial

## linear gradients specifically
- gradient that fades into another color over a line through the direction of the fade
- simplest way to set it up is thru bkgrd-img - gradients are background imgs!

```
div {
    bkgrd-img: linear-gradient(yellow, red);
}
```
- this shows up as a fade from yellow on top to red on the bottom

- helpful when you wanna put a faded color over an img or text that has a function - like a modal

## directions and how it works
- top to bottom is the default
- but to change the direction use this shorthand

```div {
    bkgrd-img: linear-gradient(angle/direction, yellow, red)
}```

- you can specify the directions using keywords like ```to top``` 
- you can also specify a middle color and at which percentage of the gradient you want it to show up as:

```div {
    bkgrd-img: linear-gradient(to right, yellow, blue 30%, red);
}