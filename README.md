# mm-webpage

make ma webpage multi media ... builds static websites for your photos and videos

## Was hat sich der Erfinder hierbei so gedacht, hm?

What did the inventor think before coming here, to share his code farts with public on github. Well. Hosting some static website is fun thanks to HUGO.
But maitaining pages for LOADS of multimedia content - photos and videos - can become cumbersome, as it's an annoyingly repetetitve task. On top - my fiber internet connection should be able to serve a few views without involvement (and tracking through) the big players. Still, social media hosting of videos is a great, free service to gain publicity - heck, yeah, thus we SUPPORT BOTH! Geilo, oder? So.

## WHAT THE ...

Yes, features, developers, developers, developers! Hear here:

- mm-webpage is implemented in go, as hugo is and überhaupt
- mm-webpage scans some top-level media folder (#ORIGINALS), recursively

For the media (images .jpg, movies .mp4/.mov, music .mp3) mm-webpage finds, it:

- writes corresponding webpages for later HUGO digestion
- creates animated GIF/webm thumbnails / previews for videos
- creates thumbnails/previews for videos
- makes any EXIF/IPTC meta data available to hugo templates 
- 
- mm-webpage is not here yet, yeah, meh.

## COME SE DICE ...

Si? Yes.

Docker. Ideally, one day (goal):

```bash
docker run --rm \
    -v/home/jan/MultiMediaCollection:/input \
    -v/var/www/media:/output \ 
    schnoddelbotz/mm-webpage
```

That is, `mm-webpage` binary will assume to be run within a docker context by default and use `/input` and `/output` directories operations. The `schnoddelbotz/mm-webpage` DOCKER image contains a recent HUGO version to further lower installation efforts. You only have to trust the Hacker, lol. User the [Dockerfile](./Dockerfile) yourself - better is.

# TOP SECRET FEATURE 

Micropayment via a secure AWS Lambda or Google Cloudfunction call. Arbitrary amount people can leave. How best done? Flattr once...?
