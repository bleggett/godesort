# gODESort

## Purpose

The GDEMU Phoebe/Rhea optical disc emulators (ODEs) are fantastic hardware mods that
allow you to replace the physical CD drive in a Sega Saturn with an SD card and
load disc images on original hardware - unfortunately the GDEMU ODE is [extremely
picky about how the disc images are ordered, named and otherwise laid out on the
SD card filesystem.](https://gdemu.wordpress.com/details/phoebe-details/)

Grouping disc images into categories requires "dummy" folders with specific text
files and further complicates ordering, especially with a large number of images.

This tool is designed to be pointed at your SD card root and executed - it will
automatically alphabetically order your images, fix numbering gaps, and automate
image (re)ordering after any additions or removals

> Note: Currently gODESort only works with CCD image types and the Phoebe/Rhea
> ODEs - support for the Dreamcast GDEMU and other image types is pending

## Installation

Currently the only way to install is to build from source.

> Supports any OS Golang does - Win/Mac/Lin

1. Install the [Go runtime for your OS](https://golang.org/dl/)

2. Clone this repository

3. Build the source

    ``` sh
    cd godesort
    go build
    ```

## Usage

> Note: This tool will not touch the "01" folder on your SD card since that is
> typically reserved for RMENU - all image folders will be renamed starting at "02"

1. This tool sorts your images based on the image filenames so make sure your
   images are named like you want them - e.g.

   ```text
   45/ ->
        Virtua Cop.ccd
        Virtua Cop.img
        Virtua Cop.sub
   ```

2. Run the command and point it at your Phoebe/Rhea root

    ``` sh
    ./godesort -i <path-to-your-sdcard-root> sort
    ```

3. Done. Rerun the above command as needed after adding or removing images to
   your SD card root to keep things ordered.