<template>
  <div class="Watch">
    <h1>This is an watch page</h1>
    <video id="video-player" class="video-js vjs-theme-forest" controls>
      <source :src="'/api/videos/streams/'+this.$route.params.slug+'/part.m3u8'" type="application/x-mpegURL">
    </video>
  </div>
</template>

<script>
import videojs from 'video.js';
import 'video.js/dist/video-js.css';
import '@videojs/themes/dist/forest/index.css';

export default {
  name: 'Watch',
  mounted() {
    videojs('video-player');
  },
};
</script>

<style>
 #video-player {
   display: block;
   margin-left: auto;
   margin-right: auto;
 }
</style>
