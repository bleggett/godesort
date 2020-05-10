# gODESort

## Purpose

The GDEMU Phoebe/Rhea optical disc emulators (ODEs) are fantastic hardware mods that
allow you to replace the physical CD drive in a Sega Saturn with an SD card and
load disc images on original hardware - unfortunately the GDEMU ODE is [extremely
picky about how the disc images are ordered, named and otherwise laid out on the
SD card filesystem.](https://gdemu.wordpress.com/details/phoebe-details/)

Grouping disc images into categories requires "dummy" folders with specific text
files and further complicates ordering, especially with a large number of images.

This tool is designed to be pointed at your SD card root and executed and
currently supports the following operations:

- `godesort sort -i /path/to/your/sdcard` will automatically alphabetically
  order your images, fix folder numbering gaps, and automate image (re)ordering
  after any additions or removals
- `godesort generate -i /path/to/your/sdcard` will scan your sorted images,
  regenerate the RMENU INI file, and rebuild the RMENU ISO-based menu that
  Phoebe/Rhea boot into.
  
  
> NOTE: gODEsort builds the RMENU menu using the **image filenames for the menu
> entries**, rather than the disc header titles that the native RMENU tool
> uses - which means you can name your images according to how you want them to
> be displayed and ordered in the RMENU menu.

> Note: Currently gODESort only works with CCD image types and the Phoebe/Rhea
> ODEs - support for the Dreamcast GDEMU and other image types is pending

## Prerequisites 

- [Install cdrtools for your
  platform](http://cdrtools.sourceforge.net/private/cdrecord.html)-
  specifically, `mkisofs` must be in your path
- Install the [Go runtime for your OS](https://golang.org/dl/)
- Have the latest RMENU files extracted to a folder named `01` at the root of
  your SD

## Installation

Currently the only way to install is to build from source.

> Supports any OS Golang does - Win/Mac/Lin

1. Clone this repository

2. Build the source

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
