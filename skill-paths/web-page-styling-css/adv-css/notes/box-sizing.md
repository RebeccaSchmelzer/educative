## a quick note on how box sizing works

- creating a box and giving it a specific width and height and adding padding all around it will increase the size of the box as so:

````.box {
  width: 200px;
  height: 200px;
  background-color: red;
}```

```.box {
  width: 200px;
  height: 200px;
  background-color: red;
	padding: 100px
}```
````

- the size of the padding and border are added to the width and height of the box
- box-sizing stops this behavior
- by setting this to border-box, the box size will not add the padding and borders to the container - when introduced, they will not increase the size of the element.
