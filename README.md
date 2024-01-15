# text-to-image-compression

Reproducing code about https://shkspr.mobi/blog/2024/01/compressing-text-into-images/ this interesting project.

## Reproducing

    ./text-to-image-compression
    pngcrush -ow pg1513.png

    echo -n "PNG: "; cat pg1513.png | wc -c
    echo -n "Gzip: "; cat pg1513.txt | gzip -c | wc -c
    echo -n "Gzip (-9): "; cat pg1513.txt | gzip -9c | wc -c

## License

See [LICENSE](LICENSE).
