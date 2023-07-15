## bkgrd img prop
- required prop for setting a bkgrd image.
- url() to img

### default behavior of bkgrd imgs
1. imgs repeat until they fill up the avail space
2. repeat from left to right, top to bottom
3. if img is too large, it gets cut off.

## bkgrd-repeat prop 
- controls whether bkgrd images repeat or not
- set to no repeat to only show up once
- repeat-x says only repeat on x-axis from left to right
- repeat-y is on the y-axis from top to bottom

## sizing bkgrd img
- using keywords
    - contain will contain the img within the space required, if more space is left then img is repeated
    - cover will cover the entire space required while keeping the aspect ratio of the image 