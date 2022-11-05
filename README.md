# CGAify

A little utility that takes an image and produces a reduced
version of it compatible with old CGA, EGA, and VGA 
graphics modes.

## Example

Original: 
![original](example/example.jpg)

As CGA 4-color: 
![original](example/example.jpg_CGA1.gif)

As EGA 16-color: 
![original](example/example.jpg_EGA.gif)

As VGA 256-color: 
![original](example/example.jpg_VGA.gif)

## Status

The supported modes are:

  * CGA 320x200 4-color
    * Palette 0 (Black, Red, Green, Yellow)
    * Palette 1 (Black, Magenta, Cyan, White)
  * EGA 320x200 16-color default palette
  * EGA 640x350 16-color default palette
  * VGA 320x200 256-color default palette
  * Bonus modes: CLAY, FLAY, ZANE, MAGENTA, TRANS, LESBIAN, BW

When reducing the size, the aspect ratio is preserved.

Note that this is just for fun/nostalgia, and not to make images
for viewing on actual old-school CGA screens... for that, I'd need 
to adjust for the aspect ratio of rectangular 80's pixels versus the 
square pixels of today.

## Get 

    go install github.com/sammiequee/Go.Cgaify/cmd/cgaify@latest

Giving `-h` will give usage information.

