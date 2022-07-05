# 2phacker

make ma webpage multi media ... builds static websites for your photos and videos.

If you wish, a HUGO preprocessor and executor.

https://github.com/schnoddelbotz/lalaKrautHomebrew

## Was hat sich der Erfinder hierbei so gedacht, hm?

What did the inventor think before coming here, to share his code farts with public on github. Well. Hosting some static website is fun thanks to HUGO.
But maitaining pages for LOADS of multimedia content - photos and videos - can become cumbersome, as it's an annoyingly repetetitve task. On top - my fiber internet connection should be able to serve a few views without involvement (and tracking through) the big players. Still, social media hosting of videos is a great, free service to gain publicity - heck, yeah, thus we SUPPORT BOTH! Geilo, oder? So.

## WHAT THE ...

Yes, features, developers, developers, developers! Hear here:

- 2phacker is implemented in go, as hugo is and überhaupt
- 2phacker scans some top-level media folder (#ORIGINALS), recursively

For the media (images .jpg, movies .mp4/.mov, music .mp3) 2phacker finds, it:

- writes corresponding webpages for later HUGO digestion
- creates animated GIF/webm thumbnails / previews for videos
- creates thumbnails/previews for videos
- makes any EXIF/IPTC meta data available to hugo templates
- provides built-in hugo templates to produce galleries from the above
- permits use of own templates
- writes https://gohugo.io/templates/data-templates/ for hugo template digestion
- 
- 2phacker is not here yet, yeah, meh.

## All TODOs tutti


- add micro payment services so visitors CAN donate (flattr et al)
- would be imaginable to have a (cross-platform) UI on top 
  - more "app" than "program" :/ enabling more users ...
  - permit meta-data editing of mp3/mp3/jpg (IPTC!) as it serves as input for hugo data-templates (single source of truth - THE file)
- when sold as online service instead of local app ...
  - support aws s3 / glacier / google drive / gcs as source 
- never replicate hugo functionality - e.g. `2phacker serve` is only shorthand to preprocess and the run hugo serve

## COME SE DICE ...

Si? Yes.

Docker. Ideally, one day (goal):

```bash
docker run --rm \
    -v/home/jan/MultiMediaCollection:/input \
    -v/var/www/media:/output \ 
    schnoddelbotz/2phacker
```
<!-- 
That is, `2phacker` binary will assume to be run within a docker context by default and use `/input` and `/output` directories operations. The `schnoddelbotz/2phacker` DOCKER image contains a recent HUGO version to further lower installation efforts. You only have to trust the Hacker, lol. User the [Dockerfile](./Dockerfile) yourself - better is. -->

# TOP SECRET FEATURE - todo ...

- Micropayment via a secure AWS Lambda or Google Cloudfunction call. Arbitrary amount people can leave. How best done? Flattr once...?

- 